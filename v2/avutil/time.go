// Author: Turing Zhu
// Date: 5/10/21 3:22 PM
// File: time.go

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/time.h>
//#include <stdlib.h>
import "C"

func GetTime() int64 {
	return C.longlong(C.av_gettime())
}