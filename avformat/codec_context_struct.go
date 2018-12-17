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
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/giorgisio/goav/avcodec"
)

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

func (cctxt *CodecContext) GetTimeBase() avcodec.Rational {
	return newRational(cctxt.time_base)
}

func (cctxt *CodecContext) SetTimeBase(timeBase avcodec.Rational) {
	cctxt.time_base.num = C.int(timeBase.Num())
	cctxt.time_base.den = C.int(timeBase.Den())
}

func (cctx *CodecContext) GetWidth() int {
	return int(cctx.width)
}

func (cctx *CodecContext) SetWidth(w int) {
	cctx.width = C.int(w)
}

func (cctx *CodecContext) GetHeight() int {
	return int(cctx.height)
}

func (cctx *CodecContext) SetHeight(h int) {
	cctx.height = C.int(h)
}

func (cctx *CodecContext) GetPixelFormat() avcodec.PixelFormat {
	return avcodec.PixelFormat(C.int(cctx.pix_fmt))
}

func (cctx *CodecContext) SetPixelFormat(fmt avcodec.PixelFormat) {
	cctx.pix_fmt = C.enum_AVPixelFormat(C.int(fmt))
}

func (cctx *CodecContext) GetFlags() int {
	return int(cctx.flags)
}

func (cctx *CodecContext) SetFlags(flags int) {
	cctx.flags = C.int(flags)
}

func (cctx *CodecContext) GetMeRange() int {
	return int(cctx.me_range)
}

func (cctx *CodecContext) SetMeRange(r int) {
	cctx.me_range = C.int(r)
}

func (cctx *CodecContext) GetMaxQDiff() int {
	return int(cctx.max_qdiff)
}

func (cctx *CodecContext) SetMaxQDiff(v int) {
	cctx.max_qdiff = C.int(v)
}

func (cctx *CodecContext) GetQMin() int {
	return int(cctx.qmin)
}

func (cctx *CodecContext) SetQMin(v int) {
	cctx.qmin = C.int(v)
}

func (cctx *CodecContext) GetQMax() int {
	return int(cctx.qmax)
}

func (cctx *CodecContext) SetQMax(v int) {
	cctx.qmax = C.int(v)
}

func (cctx *CodecContext) GetQCompress() float32 {
	return float32(cctx.qcompress)
}

func (cctx *CodecContext) SetQCompress(v float32) {
	cctx.qcompress = C.float(v)
}

func (cctx *CodecContext) GetExtraData() []byte {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cctx.extradata)),
		Len:  int(cctx.extradata_size),
		Cap:  int(cctx.extradata_size),
	}

	return *((*[]byte)(unsafe.Pointer(&header)))
}

func (cctx *CodecContext) SetExtraData(data []byte) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&data))

	cctx.extradata = (*C.uint8_t)(unsafe.Pointer(header.Data))
	cctx.extradata_size = C.int(header.Len)
}

func (cctx *CodecContext) Release() {
	C.avcodec_close((*C.struct_AVCodecContext)(unsafe.Pointer(cctx)))
	C.av_freep(unsafe.Pointer(cctx))
}
