// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avdevice

/*
	#cgo pkg-config: libavdevice
	#include <libavdevice/avdevice.h>
*/
import "C"

//Audio input devices iterator.
func (d *InputFormat) AvInputAudioDeviceNext() *InputFormat {
	return (*InputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(d)))
}

//Video input devices iterator.
func (d *InputFormat) AvInputVideoDeviceNext() *InputFormat {
	return (*InputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(d)))
}

//Audio output devices iterator.
func (d *OutputFormat) AvOutputAudioDeviceNext() *OutputFormat {
	return (*OutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(d)))
}

//Video output devices iterator.
func (d *OutputFormat) AvOutputVideoDeviceNext() *OutputFormat {
	return (*OutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(d)))
}
