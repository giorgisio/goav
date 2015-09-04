/*
	Format
*/
package avdevice

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"
import (
	"unsafe"
)

//Audio input devices iterator.
func (d *InputFormat) Av_input_audio_device_next() *InputFormat {
	return (*InputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(d)))
}

//Video input devices iterator.
func (d *InputFormat) Av_input_video_device_next() *InputFormat {
	return (*InputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(d)))
}

//Audio output devices iterator.
func (d *OutputFormat) Av_output_audio_device_next() *OutputFormat {
	return (*OutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(d)))
}

//Video output devices iterator.
func (d *OutputFormat) Av_output_video_device_next() *OutputFormat {
	return (*OutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(d)))
}
