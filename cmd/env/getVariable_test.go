package env

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

// TODO: Test the getVariable function
func Test_GetVariableCommandOneArgument(t *testing.T) {
	root := &cobra.Command{Use: "getVariable", RunE: getVariableRunE}

	os.Args = []string{"getVariable", "test"}
	err := root.Execute()

	if err != nil {
		t.Errorf("getVariableRunE returned an error: %s", err)
	}
}

func Test_GetVariableCommandNoArguments(t *testing.T) {
	root := &cobra.Command{Use: "getVariable", RunE: getVariableRunE}

	os.Args = []string{"getVariable"}
	err := root.Execute()

	if err == nil {
		t.Errorf("getVariableRunE should have returned an error")
	}
}

func Test_GetVariableCommandTwoArguments(t *testing.T) {
	root := &cobra.Command{Use: "getVariable", RunE: getVariableRunE}

	os.Args = []string{"getVariable", "test", "test"}
	err := root.Execute()

	if err == nil {
		t.Errorf("getVariableRunE should have returned an error")
	}
}
