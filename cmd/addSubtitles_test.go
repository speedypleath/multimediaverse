package cmd

import "testing"

func TestAddSubtitles(t *testing.T) {
	video_path := "/Users/speedypleath/Projects/multimediaverse/processing/sample_data/video.mp4"
	subtitles_path := "/Users/speedypleath/Projects/multimediaverse/processing/sample_data/subtitles.srt"
	output_path := "/Users/speedypleath/Projects/multimediaverse/processing/sample_data/output.mp4"
	AddSubtitles(video_path, subtitles_path, output_path)
}
