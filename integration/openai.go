package integrations

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func Transcribe(filepath string) {
	c := openai.NewClient("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")
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: filepath,
	}
	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {

		fmt.Printf("Transcription error: %v\n", err)
		return
	}
	fmt.Println(resp.Text)
}
