package env

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func Test_ShowEnvNoArguments(t *testing.T) {
	root := &cobra.Command{Use: "showEnv", RunE: showEnvRunE}

	os.Args = []string{"showEnv"}
	err := root.Execute()

	if err != nil {
		t.Errorf("getVariableRunE returned an error: %s", err)
	}
}

func Test_ShowEnvWithArguments(t *testing.T) {
	root := &cobra.Command{Use: "showEnv", RunE: showEnvRunE}

	os.Args = []string{"showEnv", "test"}
	err := root.Execute()

	if err == nil {
		t.Errorf("getVariableRunE should have returned an error")
	}
}
