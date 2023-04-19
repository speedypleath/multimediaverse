package env

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// setVariableCmd represents the setVariable command
var setVariableCmd = &cobra.Command{
	Use:   "setVariable [variable] [value]",
	Short: "Set a variable in the environment",
	Long:  ``,
	RunE:  setVariableRunE,
}

func init() {
	EnvCmd.AddCommand(setVariableCmd)
}

func setVariableRunE(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("setVariable requires two arguments")
	}

	setVariable(args[0], args[1])
	return nil
}

// TODO: Add a function to set the variable
func setVariable(variable string, value string) {
	viper.Set(variable, value)
	showEnv()
	viper.WriteConfig()
}
