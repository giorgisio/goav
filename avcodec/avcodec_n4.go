// +build !n4
// require ffmpeg tag n4.x

package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libswresample
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"
import "unsafe"

func AvCodecIterate(p *unsafe.Pointer) *Codec {
	return (*Codec)(C.av_codec_iterate(p))
}
