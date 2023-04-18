package env

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func Test_SetVariableLessThanTwoArguments(t *testing.T) {
	cmd := &cobra.Command{Use: "setVariable", RunE: setVariableRunE}

	os.Args = []string{"setVariable", "test"}
	err := cmd.Execute()

	if err == nil {
		t.Errorf("getVariableRunE should have returned an error")
	}
}

func Test_SetVariableMoreThanTwoArguments(t *testing.T) {
	cmd := &cobra.Command{Use: "setVariable", RunE: setVariableRunE}

	os.Args = []string{"setVariable", "test", "test", "test"}
	err := cmd.Execute()

	if err == nil {
		t.Errorf("getVariableRunE should have returned an error")
	}
}

func Test_SetVariableTwoArguments(t *testing.T) {
	cmd := &cobra.Command{Use: "setVariable", RunE: setVariableRunE}

	os.Args = []string{"setVariable", "test", "test"}
	err := cmd.Execute()

	if err != nil {
		t.Errorf("getVariableRunE returned an error: %s", err)
	}
}
