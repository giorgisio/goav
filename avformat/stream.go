/*
	AVStream
*/
package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

//AvRational av_stream_get_r_frame_rate (const AvStream *s)
func (s *AvStream) AvStreamGetRFrameRate() AvRational {
	return (AvRational)(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

//void av_stream_set_r_frame_rate (AvStream *s, AvRational r)
func (s *AvStream) AvStreamSetRFrameRate(r AvRational) {
	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), (C.struct_AVRational)(r))
}

//struct AvCodecParserContext * av_stream_get_parser (const AvStream *s)
func (s *AvStream) AvStreamGetParser() *AvCodecParserContext {
	return (*AvCodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

// //char * av_stream_get_recommended_encoder_configuration (const AvStream *s)
// func (s *AvStream) AvStreamGetRecommendedEncoderConfiguration() string {
// 	return C.GoString(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(s)))
// }

// //void av_stream_set_recommended_encoder_configuration (AvStream *s, char *configuration)
// func (s *AvStream) AvStreamSetRecommendedEncoderConfiguration( c string) {
// 	C.av_stream_set_recommended_encoder_configuration((*C.struct_AVStream)(s), C.CString(c))
// }

//int64_t av_stream_get_end_pts (const AvStream *st)
//Returns the pts of the last muxed packet + its duration.
func (s *AvStream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}
