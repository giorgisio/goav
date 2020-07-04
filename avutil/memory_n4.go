// +build !n4
// require ffmpeg tag n4.x

package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/avutil.h>
	#include <libavutil/frame.h>
	#include <stdlib.h>
*/
import "C"
import "unsafe"

func AvMallocArray(n, s uintptr) unsafe.Pointer {
	return C.av_malloc_array(C.size_t(n), C.size_t(s))
}

func AvMalloczArray(n, s uintptr) unsafe.Pointer {
	return C.av_mallocz_array(C.size_t(n), C.size_t(s))
}
