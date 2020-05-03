// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/samplefmt.h>
import "C"
import (
	"unsafe"
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
