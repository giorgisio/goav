// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/audio_fifo.h>
import "C"
import (
	"unsafe"
)

type (
	AvAudioFifo    C.struct_AVAudioFifo
	AvSampleFormat C.enum_AVSampleFormat
)

func AvAudioFifoAlloc(sampleFmt AvSampleFormat, channels, nbSamples int) *AvAudioFifo {
	return (*AvAudioFifo)(C.av_audio_fifo_alloc(C.enum_AVSampleFormat(sampleFmt), C.int(channels), C.int(nbSamples)))
}

func (af *AvAudioFifo) AvAudioFifoRealloc(nbSamples int) int {
	return int(C.av_audio_fifo_realloc((*C.struct_AVAudioFifo)(af), C.int(nbSamples)))
}

func (af *AvAudioFifo) AvAudioFifoWrite(data **uint8, nbSamples int) int {
	return int(C.av_audio_fifo_write((*C.struct_AVAudioFifo)(af), (*unsafe.Pointer)(unsafe.Pointer(data)), C.int(nbSamples)))
}

func (af *AvAudioFifo) AvAudioFifoRead(data **uint8, nbSamples int) int {
	return int(C.av_audio_fifo_read((*C.struct_AVAudioFifo)(af), (*unsafe.Pointer)(unsafe.Pointer(data)), C.int(nbSamples)))
}

func (af *AvAudioFifo) AvAudioFifoSize() int {
	return int(C.av_audio_fifo_size((*C.struct_AVAudioFifo)(af)))
}

func (af *AvAudioFifo) AvAudioFifoFree() {
	C.av_audio_fifo_free((*C.struct_AVAudioFifo)(af))
}
