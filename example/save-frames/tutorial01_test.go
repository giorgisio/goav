package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"fmt"
	"os"
)

func TestSaveFrames(t *testing.T) {

	for i := 1; i <= 5; i++ {
		os.Remove(fmt.Sprintf("frame%d.ppm", i))
	}
	assert.Equal(t, SaveFrames("sample.mp4"), 0)
	for i := 1; i <= 5; i++ {
		fname := fmt.Sprintf("frame%d.ppm", i)
		f, err := os.Stat(fname)
		if err != nil {
			t.Errorf("failed to read %s, err=%v", fname, err)
		}
		if f.Size() < 2_000_000 {
			t.Errorf("%s seems too short", fname)
		}
		os.Remove(fname)
	}
}
