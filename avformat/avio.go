// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avio.h>
import "C"
import "unsafe"

func AvioOpen(s **AvIOContext, url string, flags int) int {
	return int(C.avio_open((**C.struct_AVIOContext)(unsafe.Pointer(s)), C.CString(url), C.int(flags)))
}