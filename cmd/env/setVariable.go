/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package env

import (
	"fmt"

	"github.com/spf13/cobra"
)

// setVariableCmd represents the setVariable command
var setVariableCmd = &cobra.Command{
	Use:   "setVariable",
	Short: "Set a variable in the environment",
	Long:  ``,
	// TODO: Add a function to set the variable
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("setVariable called")
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
