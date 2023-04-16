package processing

import (
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func AddSubtitles(videoPath string, subtitlesPath string, outputPath string) {
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
