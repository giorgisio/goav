package avcodec

import (
	"testing"
)

func TestAllocAndFree(t *testing.T) {
	codecId := CodecId(AV_CODEC_ID_OPUS)
	inputCodec := AvcodecFindDecoder(codecId)
	if inputCodec == nil {
		t.Errorf("Could not find input codec\n")
	}
	avCtx := inputCodec.AvcodecAllocContext3()
	if avCtx == nil {
		t.Error("Could not allocate a decoding context\n")
	}
	// Check this does not cause SEGV
	avCtx.AvcodecFreeContext()
}
