/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package video

import (
	"testing"

	"github.com/spf13/cobra"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// addSubtitlesCmd represents the addSubtitles command
var addSubtitlesCmd = &cobra.Command{
	Use:   "addSubtitles",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		video_path := "/Users/speedypleath/Projects/multimediaverse/cmd/video/sample_data/video.mp4"
		subtitles_path := "/Users/speedypleath/Projects/multimediaverse/cmd/video/sample_data/subtitles.srt"
		addSubtitles(video_path, subtitles_path, "output.mp4")
	},
}

func init() {
	VideoCmd.AddCommand(addSubtitlesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addSubtitlesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addSubtitlesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addSubtitles(videoPath string, subtitlesPath string, outputPath string) {
	// Add subtitles to a video
	input := ffmpeg.Input(videoPath)
	subtitles := ffmpeg.Input(subtitlesPath)

	err := ffmpeg.
		Output([]*ffmpeg.Stream{input.Video(), input.Audio(), subtitles},
			outputPath,
			ffmpeg.KwArgs{"c:v": "copy", "c:a": "copy", "c:s": "mov_text"}).
		OverWriteOutput().
		Run()

	if err != nil {
		panic(err)
	}
}

// TODO: Test addSubtitles function
func TestAddSubtitles(t *testing.T) {
	video_path := "/Users/speedypleath/Projects/multimediaverse/cmd/video/sample_data/video.mp4"
	subtitles_path := "/Users/speedypleath/Projects/multimediaverse/cmd/video/sample_data/subtitles.srt"
	output_path := "/Users/speedypleath/Projects/multimediaverse/processing/sample_data/output.mp4"
	addSubtitles(video_path, subtitles_path, output_path)
}
