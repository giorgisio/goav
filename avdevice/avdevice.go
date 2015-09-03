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
	AvDeviceRect              C.struct_AVDeviceRect
	AvDeviceCapabilitiesQuery C.struct_AVDeviceCapabilitiesQuery
	AvDeviceInfo              C.struct_AVDeviceInfo
	AvDeviceInfoList          C.struct_AVDeviceInfoList
	InputFormat               C.struct_AVInputFormat
	OutputFormat              C.struct_AVOutputFormat
	AvFormatContext           C.struct_AVFormatContext
	Dictionary                C.struct_AVDictionary
	AvAppToDevMessageType     C.enum_AVAppToDevMessageType
	AvDevToAppMessageType     C.enum_AVDevToAppMessageType
)

//unsigned 	avdevice_version (void)
//Return the LIBAvDEVICE_VERSION_INT constant.
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

//InputFormat * 	av_input_audio_device_next (InputFormat *d)
//Audio input devices iterator.
func Av_input_audio_device_next(d *InputFormat) *InputFormat {
	return (*InputFormat)(C.av_input_audio_device_next((*C.struct_AVInputFormat)(d)))
}

//InputFormat * 	av_input_video_device_next (InputFormat *d)
//Video input devices iterator.
func Av_input_video_device_next(d *InputFormat) *InputFormat {
	return (*InputFormat)(C.av_input_video_device_next((*C.struct_AVInputFormat)(d)))
}

//OutputFormat * 	av_output_audio_device_next (OutputFormat *d)
//Audio output devices iterator.
func Av_output_audio_device_next(d *OutputFormat) *OutputFormat {
	return (*OutputFormat)(C.av_output_audio_device_next((*C.struct_AVOutputFormat)(d)))
}

//OutputFormat * 	av_output_video_device_next (OutputFormat *d)
//Video output devices iterator.
func Av_output_video_device_next(d *OutputFormat) *OutputFormat {
	return (*OutputFormat)(C.av_output_video_device_next((*C.struct_AVOutputFormat)(d)))
}

//int 	avdevice_app_to_dev_control_message (struct AvFormatContext *s, enum AvAppToDevMessageType type, void *data, size_t data_size)
//Send control message from application to device.
func Avdevice_app_to_dev_control_message(s *AvFormatContext, m AvAppToDevMessageType, da int, d uintptr) int {
	return int(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(s), (C.enum_AVAppToDevMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//int 	avdevice_dev_to_app_control_message (struct AvFormatContext *s, enum AvDevToAppMessageType type, void *data, size_t data_size)
//Send control message from device to application.
func Avdevice_dev_to_app_control_message(fcxt *AvFormatContext, m AvDevToAppMessageType, da int, d uintptr) int {
	return int(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(fcxt), (C.enum_AVDevToAppMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//int 	avdevice_capabilities_create (AvDeviceCapabilitiesQuery **caps, AvFormatContext *s, Dictionary **device_options)
//Initialize capabilities probing API based on AvOption API.
func Avdevice_capabilities_create(c **AvDeviceCapabilitiesQuery, s *AvFormatContext, d **Dictionary) int {
	return int(C.avdevice_capabilities_create((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//void 	avdevice_capabilities_free (AvDeviceCapabilitiesQuery **caps, AvFormatContext *s)
//Free resources created by avdevice_capabilities_create()
func Avdevice_capabilities_free(c **AvDeviceCapabilitiesQuery, s *AvFormatContext) {
	C.avdevice_capabilities_free((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(s))
}

//int 	avdevice_list_devices (struct AvFormatContext *s, AvDeviceInfoList **device_list)
//List devices.
func Avdevice_list_devices(s *AvFormatContext, d **AvDeviceInfoList) int {
	return int(C.avdevice_list_devices((*C.struct_AVFormatContext)(s), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(d))))
}

//void 	avdevice_free_list_devices (AvDeviceInfoList **device_list)
//Convenient function to free result of avdevice_list_devices().
func Avdevice_free_list_devices(d **AvDeviceInfoList) {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(d)))
}

// //int 	avdevice_list_input_sources (struct InputFormat *device, const char *device_name, Dictionary *device_options, AvDeviceInfoList **device_list)
// //List devices.
// func Avdevice_list_input_sources(d *InputFormat, dv string, do *Dictionary, dl **AvDeviceInfoList) int {
// 	return int(C.avdevice_list_input_sources((*C.struct_AVInputFormat)(d), C.CString(dv), (*C.struct_AVDictionary)(do), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }

// //int 	avdevice_list_output_sinks (struct OutputFormat *device, const char *device_name, Dictionary *device_options, AvDeviceInfoList **device_list)
// func Avdevice_list_output_sinks(d *OutputFormat, dn string, di *Dictionary, dl **AvDeviceInfoList) int {
// 	return int(C.avdevice_list_output_sinks((*C.struct_AVOutputFormat)(d), C.CString(dn), (*C.struct_AVDictionary)(di), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }
