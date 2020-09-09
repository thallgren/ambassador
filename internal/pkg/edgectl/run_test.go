package edgectl

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestRunCommand(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)
	ri := &RunInfo{
		InterceptInfo: InterceptInfo{
			Name:       "example",
			Deployment: "hello",
			TargetPort: 9000,
		},
		Self: "edgectl",
	}
	if err := ri.withIntercept(func() error { return run("python3", []string{"-m", "http.server", "9000"}) }); err != nil {
		t.Error(err)
	}
}
