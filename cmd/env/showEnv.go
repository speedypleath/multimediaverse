package env

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

var showEnvCmd = &cobra.Command{
	Use:   "showEnv",
	Short: "Show the environment variables",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		showEnv()
	},
}

func init() {
	EnvCmd.AddCommand(showEnvCmd)
}

// TODO: Add a function to show the environment variables
func showEnv() {
	fmt.Println("showEnv called")
}

// TODO:
func TestShowEnv(t *testing.T) {
}
