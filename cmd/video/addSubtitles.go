/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package video

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// addSubtitlesCmd represents the addSubtitles command
var addSubtitlesCmd = &cobra.Command{
	Use:   "addSubtitles [flags] [video_path]",
	Short: "Add subtitles to a video",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Handle errors
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		input_path := args[0]
		auto_generate, _ := cmd.Flags().GetBool("autogenerate")

		var subtitles_path string

		if auto_generate {
			subtitles_path = generateSubtitles(input_path)
		} else {
			subtitles_path, _ = cmd.Flags().GetString("subtitles")
		}

		output_path, err := cmd.Flags().GetString("output")

		// TODO: FFmpeg doesn't overwrite the file if it exists
		if err != nil || output_path == "" {
			output_path = input_path
		}

		addSubtitles(input_path, subtitles_path, output_path)
	},
}

func init() {
	VideoCmd.AddCommand(addSubtitlesCmd)

	// Flags
	addSubtitlesCmd.Flags().StringP("subtitles", "s", "", "Path to the subtitles file")
	addSubtitlesCmd.Flags().StringP("output", "o", "", "Path to the output file (by default, the input file is overwritten)")
	addSubtitlesCmd.Flags().BoolP("overwrite", "w", false, "Overwrite the output file if it exists")
	addSubtitlesCmd.Flags().BoolP("autogenerate", "a", false, "Automatically generate subtitles using ASR model specified in the environment variables")
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

// TODO: Create a function to generate subtitles using ASR model
func generateSubtitles(input_path string) string {
	return ""
}

// TODO: Test generateSubtitles function
func TestGenerateSubtitles(t *testing.T) {
}
