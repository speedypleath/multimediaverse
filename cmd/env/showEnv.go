package env

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	keys := viper.AllKeys()
	for _, key := range keys {
		fmt.Printf("%s=%s \n", key, viper.GetString(key))
	}
}
