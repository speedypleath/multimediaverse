package env

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// getVariableCmd represents the getVariable command
var getVariableCmd = &cobra.Command{
	Use:   "getVariable",
	Short: "Get a variable from the environment",
	Long:  ``,
	RunE:  getVariableRunE,
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

func getVariableRunE(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("getVariable requires one argument")
	}

	getVariable()
	return nil
}

// TODO: Add a function to get a variable from the environment
func getVariable() {
	fmt.Println("getVariable called")
}
