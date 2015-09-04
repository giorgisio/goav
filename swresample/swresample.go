// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// The process of changing the sampling rate of a discrete signal to obtain a new discrete representation of the underlying continuous signal.
// resampler provides a high-level interface to the libswresample library audio resampling utilities
package swresample

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"

type (
	Context        C.struct_SwrContext
	Frame          C.struct_AVFrame
	Class          C.struct_AVClass
	AvSampleFormat C.enum_AVSampleFormat
)

//Get the Class for Context.
func Swr_get_class() *Class {
	return (*Class)(C.swr_get_class())
}

//Context constructor functions.Allocate Context.
func Swr_alloc() *Context {
	return (*Context)(C.swr_alloc())
}

//Configuration accessors
func Swresample_version() uint {
	return uint(C.swresample_version())
}

func Swresample_configuration() string {
	return C.GoString(C.swresample_configuration())
}

func Swresample_license() string {
	return C.GoString(C.swresample_license())
}
