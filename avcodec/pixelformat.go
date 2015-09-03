/*
	AvPixelFormat
*/
package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

//Utility function to access log2_chroma_w log2_chroma_h from the pixel format AvPixFmtDescriptor.
func (p AvPixelFormat) Avcodec_get_chroma_sub_sample(h, v *int) {
	C.avcodec_get_chroma_sub_sample((C.enum_AVPixelFormat)(p), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(v)))
}

//Return a value representing the fourCC code associated to the pixel format pix_fmt, or 0 if no associated fourCC code can be found.
func (p AvPixelFormat) Avcodec_pix_fmt_to_codec_tag() uint {
	return uint(C.avcodec_pix_fmt_to_codec_tag((C.enum_AVPixelFormat)(p)))
}

func (p AvPixelFormat) Avcodec_get_pix_fmt_loss(f AvPixelFormat, a int) int {
	return int(C.avcodec_get_pix_fmt_loss((C.enum_AVPixelFormat)(p), (C.enum_AVPixelFormat)(f), C.int(a)))
}

//Find the best pixel format to convert to given a certain source pixel format.
func (p *AvPixelFormat) Avcodec_find_best_pix_fmt_of_list(s AvPixelFormat, a int, l *int) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_find_best_pix_fmt_of_list((*C.enum_AVPixelFormat)(p), (C.enum_AVPixelFormat)(s), C.int(a), (*C.int)(unsafe.Pointer(l))))
}

func (p AvPixelFormat) Avcodec_find_best_pix_fmt_of_2(f2, s AvPixelFormat, a int, l *int) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_find_best_pix_fmt_of_2((C.enum_AVPixelFormat)(p), (C.enum_AVPixelFormat)(f2), (C.enum_AVPixelFormat)(s), C.int(a), (*C.int)(unsafe.Pointer(l))))
}
