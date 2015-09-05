// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avdevice deals with the input and output devices provided by the libavdevice library
// The libavdevice library provides the same interface as libavformat.
// Namely, an input device is considered like a demuxer, and an output device like a muxer,
// and the interface and generic device options are the same provided by libavformat
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
func AvdeviceVersion() uint {
	return uint(C.avdevice_version())
}

//Return the libavdevice build-time configuration.
func AvdeviceConfiguration() string {
	return C.GoString(C.avdevice_configuration())
}

//Return the libavdevice license.
func AvdeviceLicense() string {
	return C.GoString(C.avdevice_license())
}

//Initialize libavdevice and register all the input and output devices.
func AvdeviceRegisterAll() {
	C.avdevice_register_all()
}

//Send control message from application to device.
func AvdeviceAppToDevControlMessage(s *AvFormatContext, m AvAppToDevMessageType, da int, d uintptr) int {
	return int(C.avdevice_app_to_dev_control_message((*C.struct_AVFormatContext)(s), (C.enum_AVAppToDevMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//Send control message from device to application.
func AvdeviceDevToAppControlMessage(fcxt *AvFormatContext, m AvDevToAppMessageType, da int, d uintptr) int {
	return int(C.avdevice_dev_to_app_control_message((*C.struct_AVFormatContext)(fcxt), (C.enum_AVDevToAppMessageType)(m), unsafe.Pointer(&da), C.size_t(d)))
}

//Initialize capabilities probing API based on AvOption API.
func AvdeviceCapabilitiesCreate(c **AvDeviceCapabilitiesQuery, s *AvFormatContext, d **Dictionary) int {
	return int(C.avdevice_capabilities_create((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Free resources created by avdevice_capabilities_create()
func AvdeviceCapabilitiesFree(c **AvDeviceCapabilitiesQuery, s *AvFormatContext) {
	C.avdevice_capabilities_free((**C.struct_AVDeviceCapabilitiesQuery)(unsafe.Pointer(c)), (*C.struct_AVFormatContext)(s))
}

//List devices.
func AvdeviceListDevices(s *AvFormatContext, d **AvDeviceInfoList) int {
	return int(C.avdevice_list_devices((*C.struct_AVFormatContext)(s), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(d))))
}

//Convenient function to free result of avdeviceListDevices().
func AvdeviceFreeListDevices(d **AvDeviceInfoList) {
	C.avdevice_free_list_devices((**C.struct_AVDeviceInfoList)(unsafe.Pointer(d)))
}

// //int 	avdevice_list_input_sources (struct InputFormat *device, const char *device_name, Dictionary *device_options, AvDeviceInfoList **device_list)
// //List devices.
// func AvdeviceListInputSources(d *InputFormat, dv string, do *Dictionary, dl **AvDeviceInfoList) int {
// 	return int(C.avdevice_list_input_sources((*C.struct_AVInputFormat)(d), C.CString(dv), (*C.struct_AVDictionary)(do), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }

// //int 	avdevice_list_output_sinks (struct OutputFormat *device, const char *device_name, Dictionary *device_options, AvDeviceInfoList **device_list)
// func AvdeviceListOutputSinks(d *OutputFormat, dn string, di *Dictionary, dl **AvDeviceInfoList) int {
// 	return int(C.avdevice_list_output_sinks((*C.struct_AVOutputFormat)(d), C.CString(dn), (*C.struct_AVDictionary)(di), (**C.struct_AVDeviceInfoList)(unsafe.Pointer(dl))))
// }
