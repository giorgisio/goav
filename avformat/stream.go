// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"github.com/giorgisio/goav/avcodec"
)

//Rational av_stream_get_r_frame_rate (const Stream *s)
func (s *Stream) AvStreamGetRFrameRate() avcodec.Rational {
	return newRational(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

//void av_stream_set_r_frame_rate (Stream *s, Rational r)
func (s *Stream) AvStreamSetRFrameRate(r avcodec.Rational) {
	rat := C.struct_AVRational{
		num: C.int(r.Num()),
		den: C.int(r.Den()),
	}

	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), rat)
}

//struct CodecParserContext * av_stream_get_parser (const Stream *s)
func (s *Stream) AvStreamGetParser() *CodecParserContext {
	return (*CodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

// //char * av_stream_get_recommended_encoder_configuration (const Stream *s)
// func (s *Stream) AvStreamGetRecommendedEncoderConfiguration() string {
// 	return C.GoString(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(s)))
// }

// //void av_stream_set_recommended_encoder_configuration (Stream *s, char *configuration)
// func (s *Stream) AvStreamSetRecommendedEncoderConfiguration( c string) {
// 	C.av_stream_set_recommended_encoder_configuration((*C.struct_AVStream)(s), C.CString(c))
// }

//int64_t av_stream_get_end_pts (const Stream *st)
//Returns the pts of the last muxed packet + its duration.
func (s *Stream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}
