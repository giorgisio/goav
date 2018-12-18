// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avformat provides some generic global options, which can be set on all the muxers and demuxers.
//In addition each muxer or demuxer may support so-called private options, which are specific for that component.
//Supported formats (muxers and demuxers) provided by the libavformat library
package avformat

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavutil/avutil.h>
import "C"
import (
	"github.com/giorgisio/goav/avutil"


func (cctxt *CodecContext) Type() MediaType {
	return MediaType(cctxt.codec_type)
}

func (cctxt *CodecContext) SetBitRate(br int64) {
	cctxt.bit_rate = C.int64_t(br)
}

func (cctxt *CodecContext) GetCodecId() CodecId {
	return CodecId(cctxt.codec_id)
}

func (cctxt *CodecContext) SetCodecId(codecId CodecId) {
	cctxt.codec_id = C.enum_AVCodecID(codecId)
}

func (cctxt *CodecContext) GetCodecType() MediaType {
	return MediaType(cctxt.codec_type)
}

func (cctxt *CodecContext) SetCodecType(ctype MediaType) {
	cctxt.codec_type = C.enum_AVMediaType(ctype)
}

func (cctxt *CodecContext) GetTimeBase() avutil.Rational {
	return avutil.Rational(cctxt.time_base)
}

func (cctxt *CodecContext) SetTimeBase(timeBase avutil.Rational) {
	cctxt.time_base = C.struct_AVRational(timeBase)
}
