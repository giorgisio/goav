// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Dmitry Patsura <talk@dmtry.me> https://github.com/ovr

package avformat

//#cgo pkg-config: libavformat
//#cgo pkg-config: libavutil
//#include <libavformat/avformat.h>
//#include <libavutil/display.h>
import "C"
import (
	"strconv"
	"unsafe"
	"strings"
)

//Rational av_stream_get_r_frame_rate (const Stream *s)
func (s *Stream) AvStreamGetRFrameRate() Rational {
	return (Rational)(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

//void av_stream_set_r_frame_rate (Stream *s, Rational r)
func (s *Stream) AvStreamSetRFrameRate(r Rational) {
	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), (C.struct_AVRational)(r))
}

/**
 * Get rotation of the Stream by
 * 1. MetaData
 * 2. Metrix
 */
func (s *Stream) GetRotation() int64 {
	dictionaryEntry := s.Metadata().Get("rotate")
	if (dictionaryEntry != nil) {
		value := dictionaryEntry.Value();
		strings.TrimRight(value, "%")

		rotation, err := strconv.ParseInt(value, 10, 64)
		if (err == nil) {
			return rotation
		}
	}

	var displaymatrix (*C.uint8_t) = C.av_stream_get_side_data((*C.struct_AVStream)(s), C.AV_PKT_DATA_DISPLAYMATRIX, nil);
	if (displaymatrix != nil) {
		return -int64(C.av_display_rotation_get((*C.int32_t)(unsafe.Pointer(displaymatrix))))
	}

	return 0
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
