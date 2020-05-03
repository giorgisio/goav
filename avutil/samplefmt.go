// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/samplefmt.h>
import "C"
import (
	"unsafe"
)

const (
	AV_SAMPLE_FMT_NONE = AvSampleFormat(C.AV_SAMPLE_FMT_NONE)
	AV_SAMPLE_FMT_U8   = AvSampleFormat(C.AV_SAMPLE_FMT_U8)
	AV_SAMPLE_FMT_S16  = AvSampleFormat(C.AV_SAMPLE_FMT_S16)
	AV_SAMPLE_FMT_S32  = AvSampleFormat(C.AV_SAMPLE_FMT_S32)
	AV_SAMPLE_FMT_FLT  = AvSampleFormat(C.AV_SAMPLE_FMT_FLT)
	AV_SAMPLE_FMT_DBL  = AvSampleFormat(C.AV_SAMPLE_FMT_DBL)
	AV_SAMPLE_FMT_U8P  = AvSampleFormat(C.AV_SAMPLE_FMT_U8P)
	AV_SAMPLE_FMT_S16P = AvSampleFormat(C.AV_SAMPLE_FMT_S16P)
	AV_SAMPLE_FMT_S32P = AvSampleFormat(C.AV_SAMPLE_FMT_S32P)
	AV_SAMPLE_FMT_FLTP = AvSampleFormat(C.AV_SAMPLE_FMT_FLTP)
	AV_SAMPLE_FMT_DBLP = AvSampleFormat(C.AV_SAMPLE_FMT_DBLP)
	AV_SAMPLE_FMT_NB   = AvSampleFormat(C.AV_SAMPLE_FMT_NB)
)

func AvSamplesAllocArrayAndSamples(nbChannels, nbSamples int, sampleFmt AvSampleFormat, align int) (ret int, audioData **uint8, linesize int) {
	var linesize32 int32
	ret = (int)(C.av_samples_alloc_array_and_samples(
		(***C.uint8_t)(unsafe.Pointer(&audioData)),
		(*C.int)(&linesize32),
		C.int(nbChannels), C.int(nbSamples),
		C.enum_AVSampleFormat(sampleFmt), C.int(align)))
	return ret, audioData, int(linesize32)
}

func AvSamplesFreeArrayAndSamples(audioData **uint8) {
	C.av_freep(unsafe.Pointer(audioData))
	C.av_free(unsafe.Pointer(audioData))
}
