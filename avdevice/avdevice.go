/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)

	The input and output devices provided by the libavdevice library
	The libavdevice library provides the same interface as libavformat.
	Namely, an input device is considered like a demuxer, and an output device like a muxer,
	and the interface and generic device options are the same provided by libavformat
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

type (
	AVDeviceRect              C.struct_AVDeviceRect
	AVDeviceCapabilitiesQuery C.struct_AVDeviceCapabilitiesQuery
	AVDeviceInfo              C.struct_AVDeviceInfo
	AVDeviceInfoList          C.struct_AVDeviceInfoList
	AVInputFormat             C.struct_AVInputFormat
	AVOutputFormat            C.struct_AVOutputFormat
	AVFormatContext           C.struct_AVFormatContext
	AVDictionary              C.struct_AVDictionary
)

type (
	AVAppToDevMessageType C.enum_AVAppToDevMessageType
	AVDevToAppMessageType C.enum_AVDevToAppMessageType
)

//unsigned 	avdevice_version (void)
//Return the LIBAVDEVICE_VERSION_INT constant.
func Avdevice_version() uint {
	return uint(C.avdevice_version())
}

//const char * 	avdevice_configuration (void)
//Return the libavdevice build-time configuration.
func Avdevice_configuration() string {
	return C.GoString(C.avdevice_configuration())
}

//const char * 	avdevice_license (void)
//Return the libavdevice license.
func Avdevice_license() string {
	return C.GoString(C.avdevice_license())
}

//void 	avdevice_register_all (void)
//Initialize libavdevice and register all the input and output devices.
func Avdevice_register_all() {
	C.avdevice_register_all()
}

//AVInputFormat * 	av_input_audio_device_next (AVInputFormat *d)
//Audio input devices iterator.
func Av_input_audio_device_next(d *AVInputFormat) *AVInputFormat {
	return (*AVInputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(d)))
}

//AVInputFormat * 	av_input_video_device_next (AVInputFormat *d)
//Video input devices iterator.
func Av_input_video_device_next(d *AVInputFormat) *AVInputFormat {
	return (*AVInputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(d)))
}

//AVOutputFormat * 	av_output_audio_device_next (AVOutputFormat *d)
//Audio output devices iterator.
func Av_output_audio_device_next(d *AVOutputFormat) *AVOutputFormat {
	return (*AVOutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(d)))
}

//AVOutputFormat * 	av_output_video_device_next (AVOutputFormat *d)
//Video output devices iterator.
func Av_output_video_device_next(d *AVOutputFormat) *AVOutputFormat {
	return (*AVOutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(d)))
}

//int 	avdevice_app_to_dev_control_message (struct AVFormatContext *s, enum AVAppToDevMessageType type, void *data, size_t data_size)
//Send control message from application to device.
func Avdevice_app_to_dev_control_message(s *AVFormatContext, m AVAppToDevMessageType, da int, d uintptr) int {
	return int(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(s), (C.enum_AVAppToDevMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//int 	avdevice_dev_to_app_control_message (struct AVFormatContext *s, enum AVDevToAppMessageType type, void *data, size_t data_size)
//Send control message from device to application.
func Avdevice_dev_to_app_control_message(fcxt *AVFormatContext, m AVDevToAppMessageType, da int, d uintptr) int {
	return int(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(fcxt), (C.enum_AVDevToAppMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//int 	avdevice_capabilities_create (AVDeviceCapabilitiesQuery **caps, AVFormatContext *s, AVDictionary **device_options)
//Initialize capabilities probing API based on AVOption API.
func Avdevice_capabilities_create(c **AVDeviceCapabilitiesQuery, s *AVFormatContext, d **AVDictionary) int {
	return int(C.avdevice_capabilities_create((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//void 	avdevice_capabilities_free (AVDeviceCapabilitiesQuery **caps, AVFormatContext *s)
//Free resources created by avdevice_capabilities_create()
func Avdevice_capabilities_free(c **AVDeviceCapabilitiesQuery, s *AVFormatContext) {
	C.avdevice_capabilities_free((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(s))
}

//int 	avdevice_list_devices (struct AVFormatContext *s, AVDeviceInfoList **device_list)
//List devices.
func Avdevice_list_devices(s *AVFormatContext, d **AVDeviceInfoList) int {
	return int(C.avdevice_list_devices((*C.struct_AVFormatContext)(s), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(d))))
}

//void 	avdevice_free_list_devices (AVDeviceInfoList **device_list)
//Convenient function to free result of avdevice_list_devices().
func Avdevice_free_list_devices(d **AVDeviceInfoList) {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(d)))
}

// //int 	avdevice_list_input_sources (struct AVInputFormat *device, const char *device_name, AVDictionary *device_options, AVDeviceInfoList **device_list)
// //List devices.
// func Avdevice_list_input_sources(d *AVInputFormat, dv string, do *AVDictionary, dl **AVDeviceInfoList) int {
// 	return int(C.avdevice_list_input_sources((*C.struct_AVInputFormat)(d), C.CString(dv), (*C.struct_AVDictionary)(do), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }

// //int 	avdevice_list_output_sinks (struct AVOutputFormat *device, const char *device_name, AVDictionary *device_options, AVDeviceInfoList **device_list)
// func Avdevice_list_output_sinks(d *AVOutputFormat, dn string, di *AVDictionary, dl **AVDeviceInfoList) int {
// 	return int(C.avdevice_list_output_sinks((*C.struct_AVOutputFormat)(d), C.CString(dn), (*C.struct_AVDictionary)(di), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }
