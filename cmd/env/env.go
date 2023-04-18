/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package env

import (
	"github.com/spf13/cobra"
)

// envCmd represents the env command
var EnvCmd = &cobra.Command{
	Use:   "env",
	Short: "Configure envirmonment variables",
	Long:  ``,
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// envCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// envCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
