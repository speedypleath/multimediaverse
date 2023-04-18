package env

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var showEnvCmd = &cobra.Command{
	Use:   "showEnv",
	Short: "Show the environment variables",
	Long:  ``,
	RunE:  showEnvRunE,
}

func init() {
	EnvCmd.AddCommand(showEnvCmd)
}

func showEnvRunE(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return errors.New("showEnv requires no arguments")
	}

	showEnv()
	return nil
}

// TODO: Add a function to show the environment variables
func showEnv() {
	fmt.Println("showEnv called")
}
