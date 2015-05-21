/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func Nb_streams(n *AVFormatContext) int {
	cn := (*C.struct_AVFormatContext)(n)
	return int(cn.nb_streams)
}

func Streams(ctx *AVFormatContext) *AVStream {
	cctx := (*C.struct_AVFormatContext)(ctx)
	return (*AVStream)(unsafe.Pointer(cctx.streams))
}

func Codec_type(s *AVStream, n int) *AVMediaType {
	c := Codec(s)
	fmt.Println("Codec Type: Fix Method")
	if c != nil {
		return (*AVMediaType)(unsafe.Pointer(&c.codec_type))
	} else {
		return nil
	}
}

func Codec(s *AVStream) *AVCodecContext {
	return (*AVCodecContext)(s.codec)
}
