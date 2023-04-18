package env

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

// setVariableCmd represents the setVariable command
var setVariableCmd = &cobra.Command{
	Use:   "setVariable [variable] [value]",
	Short: "Set a variable in the environment",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		setVariable(args[0], args[1])
	},
}

func init() {
	EnvCmd.AddCommand(setVariableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setVariableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setVariableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// TODO: Add a function to set the variable
func setVariable(variable string, value string) {
	fmt.Println("setVariable called")
}

// TODO: Test the setVariable function
func TestSetVariable(t *testing.T) {
}
