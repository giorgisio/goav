package swresample

import (
	"testing"
)

func TestAllocAndFree(t *testing.T) {
	resampleCtx := SwrAllocSetOpts(
		3, //AV_CH_LAYOUT_STEREO
		1, //AV_SAMPLE_FMT_S16
		48000,
		3, //AV_CH_LAYOUT_STEREO
		6, //AV_SAMPLE_FMT_S16P
		48000)
	if resampleCtx == nil {
		t.Error("Could not allocate resample context\n")
	}
	ret := resampleCtx.SwrInit()
	if ret < 0 {
		t.Error("Could not open resample context\n")

	}
	// Check this does not cause SEGV
	resampleCtx.SwrFree()
}
