package video

import (
	"github.com/spf13/cobra"
)

// videoCmd represents the video command
var VideoCmd = &cobra.Command{
	Use:   "video",
	Short: "Video editing tools",
	Long:  ``,
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// videoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// videoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
