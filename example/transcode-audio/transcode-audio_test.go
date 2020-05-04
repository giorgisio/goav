package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"os"
)

func TestTranscodeAudio(t *testing.T) {

	inputFile := "test001.opus"
	outputFile := "out.mp4"
	assert.Equal(t, TranscodeAudio(inputFile, outputFile), 0)
	f, err := os.Stat(outputFile)
	if err != nil {
		t.Errorf("failed to read %s, err=%v", outputFile, err)
	}
	if f.Size() < 100_000 {
		t.Errorf("%s seems too short", outputFile)
	}
	os.Remove(outputFile)
}
