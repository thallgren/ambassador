package edgectl

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/sirupsen/logrus"
)

var RunHelp = `The Edge Control Run is a shorthand command for starting the daemon,
connecting to the traffic mangager, add an intercept, run a command, and then disconnect
(unless already connected) and and quit the daemon (unless already started).

Run a command:
    edgectl run hello -n example-url -t 9000 -- <command> arguments...

Examine the Daemon's log output in
    ` + logfile + `
to troubleshoot problems.
`

type RunInfo struct {
	InterceptInfo
	Self       string
	DNS        string
	Fallback   string
	Preview    bool
	PreviewSet bool
}

func (ri *RunInfo) RunCommand(cmd *cobra.Command, args []string) error {
	logrus.SetLevel(logrus.DebugLevel)

	ri.Self = os.Args[0]
	ri.PreviewSet = cmd.Flags().Changed("preview")
	return ri.withIntercept(func() error {
		// TODO: Sensible signal handling such as trapping SIGINT and SIGKILL. Should
		//  propagate signal to subprocess and then terminate gracefully here.
		return run(args[0], args[1:])
	})
}

// channelWriter will receive the lines printed by the CommandViaDaemon function and
// write them on the channel
type channelWriter chan []byte

func (w channelWriter) Write(p []byte) (n int, err error) {
	w <- p
	return os.Stdout.Write(p)
}

func (w channelWriter) waitFor(t time.Duration, f func(data []byte) bool) bool {
	defer func() {
		// ensure writer accepts remaining output without blocking
		go func() {
			for range w {
			}
		}()
	}()

	timeout := time.NewTimer(t) // timeout waiting for ssh tunnel create
	for {
		select {
		case <-timeout.C:
			return false
		case bts, ok := <-w:
			if !ok {
				return false
			}
			if f(bts) {
				timeout.Stop()
				return true
			}
		}
	}
}

var interceptReadyMessage = []byte("starting SSH")

// withIntercept runs the given function after asserting that an intercept is in place.
func (ri *RunInfo) withIntercept(f func() error) error {
	return ri.withConnection(func() error {
		args := []string{ri.Self, "intercept", "add",
			ri.Deployment, "--name", ri.Name, "--target", ri.TargetHost}
		if ri.PreviewSet {
			args = append(args, "--preview", fmt.Sprintf("%t", ri.Preview))
		}
		for h, rx := range ri.Patterns {
			args = append(args, "--match", h+"="+rx)
		}
		if ri.Namespace != "" {
			args = append(args, "--namespace", ri.Namespace)
		}
		if ri.Prefix != "" {
			args = append(args, "--prefix", ri.Prefix)
		}
		if ri.GRPC {
			args = append(args, "--grpc")
		}

		// create an io.Writer that writes a message on a channel when the desired message has been received
		ready := make(chan bool, 1)
		out := channelWriter(make(chan []byte, 5))
		go func() {
			ready <- out.waitFor(30*time.Second, func(bts []byte) bool {
				return bytes.Contains(bts, interceptReadyMessage)
			})
		}()

		if logrus.IsLevelEnabled(logrus.DebugLevel) {
			logrus.Debug(strings.Join(args, " "))
		}
		err, exitCode := CommandViaDaemon(args, out)
		if err == nil && exitCode != 0 {
			err = fmt.Errorf("%s intercept add exited with %d", ri.Self, exitCode)
		}
		if err != nil {
			close(out) // terminates the above go routine
			return err
		}
		defer func() {
			logrus.Debugf("Removing intercept %s", ri.Name)
			_, _ = CommandViaDaemon([]string{ri.Self, "intercept", "remove", ri.Name}, os.Stdout)
		}()
		if <-ready {
			return f()
		}
		return fmt.Errorf("timeout waiting for intercept add")
	})
}

// withConnection runs the given function after asserting that a connection is active.
func (ri *RunInfo) withConnection(f func() error) error {
	return ri.withDaemonRunning(func() error {
		logrus.Debug("Connecting to daemon")
		wasConnected := true
		connected := false
		var err error

		var exitCode int
		for i := 0; i < 10; i++ {
			ready := make(chan bool, 1)
			out := channelWriter(make(chan []byte, 5))
			go func() {
				ready <- out.waitFor(2*time.Second, func(bts []byte) bool {
					line := string(bts)
					switch {
					case strings.HasPrefix(line, "Already connected"):
						return true
					case strings.HasPrefix(line, "Connected"):
						wasConnected = false
						return true
					}
					// Not ready
					wasConnected = false
					return false
				})
			}()

			err, exitCode = CommandViaDaemon([]string{ri.Self, "connect"}, out)
			if err == nil && exitCode != 0 {
				err = fmt.Errorf("%s connect exited with %d", ri.Self, exitCode)
				break
			}
			if <-ready {
				connected = true
				break
			}
			logrus.Debug("Connection not ready. Retrying...")
		}
		if !connected {
			return fmt.Errorf("timeout trying to connect")
		}
		if !wasConnected {
			defer func() {
				logrus.Debug("Disconnecting from daemon")
				_, _ = CommandViaDaemon([]string{ri.Self, "disconnect"}, os.Stdout)
			}()
		}
		// Allow time to connect to traffic manager
		time.Sleep(3 * time.Second)
		return f()
	})
}

func (ri *RunInfo) withDaemonRunning(f func() error) error {
	if IsServerRunning() {
		return f()
	}

	daemonError := atomic.Value{}
	go func() {
		logrus.Debug("Starting daemon")
		if err := ri.startDaemon(); err != nil {
			daemonError.Store(err)
		}
	}()

	defer func() {
		logrus.Debug("Quitting daemon")
		if err := ri.quitDaemon(); err != nil {
			logrus.Error(err.Error())
		}
	}()
	for {
		time.Sleep(50 * time.Millisecond)
		if err, ok := daemonError.Load().(error); ok {
			return err
		}
		if IsServerRunning() {
			return f()
		}
	}
}

func (ri *RunInfo) startDaemon() error {
	/* #nosec */
	exe := ri.Self
	args := []string{"daemon"}
	if ri.DNS != "" {
		args = append(args, "--dns", ri.DNS)
	}
	if ri.Fallback != "" {
		args = append(args, "--fallback", ri.Fallback)
	}
	return runAsRoot(exe, args)
}

func (ri *RunInfo) quitDaemon() error {
	return run(ri.Self, []string{"quit"})
}

func runAsRoot(exe string, args []string) error {
	if os.Geteuid() != 0 {
		args = append([]string{"-E", exe}, args...)
		exe = "sudo"
	}
	return run(exe, args)
}

func run(exe string, args []string) error {
	cmd := exec.Command(exe, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logrus.Debugf("executing %s %s\n", exe, strings.Join(args, " "))
	if err := cmd.Start(); err != nil {
		logrus.Debugf("starting %s %s returned error: %s\n", exe, strings.Join(args, " "), err)
		return fmt.Errorf("%s %s: %v\n", exe, strings.Join(args, " "), err)
	}
	proc := cmd.Process
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		if sig == syscall.SIGUSR1 {
			return
		}
		_ = proc.Signal(sig)
	}()
	s, err := proc.Wait()
	sigCh <- syscall.SIGUSR1
	if err != nil {
		logrus.Debugf("running %s %s returned error: %s\n", exe, strings.Join(args, " "), err)
		return fmt.Errorf("%s %s: %v\n", exe, strings.Join(args, " "), err)
	}
	exitCode := s.ExitCode()
	if exitCode != 0 {
		logrus.Debugf("executing %s %s returned exit code: %d\n", exe, strings.Join(args, " "), exitCode)
		return fmt.Errorf("%s %s: exited with %d\n", exe, strings.Join(args, " "), exitCode)
	}
	return nil
}
