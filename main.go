package main

import "multimediaverse/processing"

func main() {
	video_path := "/Users/speedypleath/Projects/multimediaverse/processing/sample_data/video.mp4"
	subtitles_path := "/Users/speedypleath/Projects/multimediaverse/processing/sample_data/subtitles.srt"
	output_path := "/Users/speedypleath/Projects/multimediaverse/processing/sample_data/output.mp4"
	processing.AddSubtitles(video_path, subtitles_path, output_path)
}
