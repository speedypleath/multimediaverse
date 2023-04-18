package video

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func Test_NoArguments(t *testing.T) {
	root := &cobra.Command{Use: "addSubtitles", RunE: addSubtitlesRunE}

	os.Args = []string{"addSubtitles"}
	err := root.Execute()
	errCode := "addSubtitles requires a video path"

	if err.Error() != errCode {
		t.Errorf("getVariableRunE should return error '%s', Actual error: %s", errCode, err)
	}
}

func Test_NoAutogenerateAndNoSubtitlesPath(t *testing.T) {
	root := &cobra.Command{Use: "addSubtitles", RunE: addSubtitlesRunE}

	os.Args = []string{"addSubtitles", "test"}
	err := root.Execute()
	errCode := "addSubtitles requires either a subtitles file path or the --autogenerate flag"

	if err.Error() != errCode {
		t.Errorf("getVariableRunE should return error '%s', Actual error: %s", errCode, err)
	}
}

func Test_Autogenerate(t *testing.T) {
	cmd := &cobra.Command{Use: "addSubtitles", RunE: addSubtitlesRunE}

	os.Args = []string{"addSubtitles", "test"}
	cmd.Flags().Bool("autogenerate", true, "autogenerate subtitles")
	err := cmd.Execute()
	errCode := "addSubtitles requires either a subtitles file path or the --autogenerate flag"

	if err == nil {
		t.Errorf("getVariableRunE should return error '%s', Actual error: %s", errCode, err)
	}
}
