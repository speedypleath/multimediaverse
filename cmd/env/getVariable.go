package env

import (
	"testing"

	"github.com/spf13/cobra"
)

// getVariableCmd represents the getVariable command
var getVariableCmd = &cobra.Command{
	Use:   "getVariable",
	Short: "Get a variable from the environment",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		getVariable()
	},
}

func init() {
	EnvCmd.AddCommand(getVariableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getVariableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getVariableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// TODO: Add a function to get a variable from the environment
func getVariable() {

}

// TODO: Test the getVariable function
func TestGetVariable(t *testing.T) {
}
