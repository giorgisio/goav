/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
*/
package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"github.com/giorgisio/goav/avcodec"
	"unsafe"
)

//int av_format_get_probe_score (const Context *s)
func (s *Context) AvFormatGetProbeScore() int {
	return int(C.av_format_get_probe_score((*C.struct_AVFormatContext)(s)))
}

//AvCodec * av_format_get_video_codec (const Context *s)
func (s *Context) AvFormatGetVideoCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_video_codec (Context *s, AvCodec *c)
func (s *Context) AvFormatSetVideoCodec(c *AvCodec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

//AvCodec * av_format_get_audio_codec (const Context *s)
func (s *Context) AvFormatGetAudioCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_audio_codec (Context *s, AvCodec *c)
func (s *Context) AvFormatSetAudioCodec(c *AvCodec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

//AvCodec * av_format_get_subtitle_codec (const Context *s)
func (s *Context) AvFormatGetSubtitleCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_subtitle_codec (Context *s, AvCodec *c)
func (s *Context) AvFormatSetSubtitleCodec(c *AvCodec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// //AvCodec * av_format_get_data_codec (const Context *s)
// func (s *Context)AvFormatGetDataCodec() *AvCodec {
// 	return (*AvCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(s)))
// }

// //void av_format_set_data_codec (Context *s, AvCodec *c)
// func (s *Context)AvFormatSetDataCodec( c *AvCodec) {
// 	C.av_format_set_data_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
// }

//int av_format_get_metadata_header_padding (const Context *s)
func (s *Context) AvFormatGetMetadataHeaderPadding() int {
	return int(C.av_format_get_metadata_header_padding((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_metadata_header_padding (Context *s, int c)
func (s *Context) AvFormatSetMetadataHeaderPadding(c int) {
	C.av_format_set_metadata_header_padding((*C.struct_AVFormatContext)(s), C.int(c))
}

//void * av_format_get_opaque (const Context *s)
func (s *Context) AvFormatGetOpaque() {
	C.av_format_get_opaque((*C.struct_AVFormatContext)(s))
}

//void av_format_set_opaque (Context *s, void *opaque)
func (s *Context) AvFormatSetOpaque(o int) {
	C.av_format_set_opaque((*C.struct_AVFormatContext)(s), unsafe.Pointer(&o))
}

// //av_format_control_message av_format_get_control_message_cb (const Context *s)
// func (s *Context) AvFormatControlMessage() C.av_format_get_control_message_cb {
// 	return C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(s))
// }

// //void av_format_set_control_message_cb (Context *s, av_format_control_message callback)
// func (s *Context) AvFormatSetControlMessageCb(c AvFormatControlMessage) C.av_format_get_control_message_cb {
// 	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(s), (C.struct_AVFormatControlMessage)(c))
// }

//void av_format_inject_global_side_data (Context *s)
//This function will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
func (s *Context) AvFormatInjectGlobalSideData() {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(s))
}

//enum AvDurationEstimationMethod av_fmt_ctx_get_duration_estimation_method (const Context *ctx)
//Returns the method used to set ctx->duration.
func (s *Context) AvFmtCtxGetDurationEstimationMethod() AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(s)))
}

//void avformat_free_context (Context *s)
//Free an Context and all its streams.
func (s *Context) AvformatFreeContext() {
	C.avformat_free_context((*C.struct_AVFormatContext)(s))
}

//Stream * avformat_new_stream (Context *s, const AvCodec *c)
//Add a new stream to a media file.
func (s *Context) AvformatNewStream(c *AvCodec) *Stream {
	return (*Stream)(C.avformat_new_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c)))
}

//AvProgram * av_new_program (Context *s, int id)
func (s *Context) AvNewProgram(id int) *AvProgram {
	return (*AvProgram)(C.av_new_program((*C.struct_AVFormatContext)(s), C.int(id)))
}

//int avformat_find_stream_info (Context *ic, Dictionary **options)
//Read packets of a media file to get stream information.
func (s *Context) AvformatFindStreamInfo(d **Dictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//AvProgram * av_find_program_from_stream (Context *ic, AvProgram *last, int s)
//Find the programs which belong to a given stream.
func (s *Context) AvFindProgramFromStream(l *AvProgram, su int) *AvProgram {
	return (*AvProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVProgram)(l), C.int(su)))
}

//int av_find_best_stream (Context *ic, enum MediaType type, int wanted_stream_nb, int related_stream, AvCodec **decoder_ret, int flags)
//Find the "best" stream in the file.
func AvFindBestStream(ic *Context, t MediaType, ws, rs int, c **AvCodec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

//int av_read_frame (Context *s, Packet *pkt)
//Return the next frame of a stream.
func (s *Context) AvReadFrame(pkt *avcodec.Packet) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(s)), (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}

//int av_seek_frame (Context *s, int stream_index, int64_t timestamp, int flags)
//Seek to the keyframe at timestamp.
func (s *Context) AvSeekFrame(st int, t int64, f int) int {
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(s), C.int(st), C.int64_t(t), C.int(f)))
}

//int avformat_seek_file (Context *s, int stream_index, int64_t min_ts, int64_t ts, int64_t max_ts, int flags)
//Seek to timestamp ts.
func (s *Context) AvformatSeekFile(si int, mit, ts, mat int64, f int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(s), C.int(si), C.int64_t(mit), C.int64_t(ts), C.int64_t(mat), C.int(f)))
}

//int av_read_play (Context *s)
//Start playing a network-based stream (e.g.
func (s *Context) AvReadPlay() int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(s)))
}

//int av_read_pause (Context *s)
//Pause a network-based stream (e.g.
func (s *Context) AvReadPause() int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(s)))
}

//void avformat_close_input (Context **s)
//Close an opened input Context.
func (s *Context) AvformatCloseInput() {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(s)))
}

//int avformat_write_header (Context *s, Dictionary **options)
//Allocate the stream private data and write the stream header to an output media file.
func (s *Context) AvformatWriteHeader(o **Dictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//int av_write_frame (Context *s, Packet *pkt)
//Write a packet to an output media file.
func (s *Context) AvWriteFrame(pkt *Packet) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(s), (*C.struct_AVPacket)(pkt)))
}

//int av_interleaved_write_frame (Context *s, Packet *pkt)
//Write a packet to an output media file ensuring correct interleaving.
func (s *Context) AvInterleavedWriteFrame(pkt *Packet) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(s), (*C.struct_AVPacket)(pkt)))
}

//int av_write_uncoded_frame (Context *s, int stream_index, Frame *frame)
//Write a uncoded frame to an output media file.
func (s *Context) AvWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(s), C.int(si), (*C.struct_AVFrame)(f)))
}

//int av_interleaved_write_uncoded_frame (Context *s, int stream_index, Frame *frame)
//Write a uncoded frame to an output media file.
func (s *Context) AvInterleavedWriteUncodedFrame(si int, f *Frame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(s), C.int(si), (*C.struct_AVFrame)(f)))
}

//int av_write_uncoded_frame_query (Context *s, int stream_index)
//Test whether a muxer supports uncoded frame.
func (s *Context) AvWriteUncodedFrameQuery(si int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(s), C.int(si)))
}

//int av_write_trailer (Context *s)
//Write the stream trailer to an output media file and free the file private data.
func (s *Context) AvWriteTrailer() int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(s)))
}

//int av_get_output_timestamp (struct Context *s, int stream, int64_t *dts, int64_t *wall)
//Get timing information for the data currently output.
func (s *Context) AvGetOutputTimestamp(st int, dts, wall *int) int {
	return int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(s), C.int(st), (*C.int64_t)(unsafe.Pointer(&dts)), (*C.int64_t)(unsafe.Pointer(&wall))))
}

//int av_find_default_stream_index (Context *s)
func (s *Context) AvFindDefaultStreamIndex() int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(s)))
}

//void av_dump_format (Context *ic, int index, const char *url, int is_output)
//Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func (s *Context) AvDumpFormat(i int, url string, io int) {
	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(s)), C.int(i), C.CString(url), C.int(io))
}

//Rational av_guess_sample_aspect_ratio (Context *format, Stream *stream, Frame *frame)
//Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
func (s *Context) AvGuessSampleAspectRatio(st *Stream, fr *Frame) Rational {
	return (Rational)(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

//Rational av_guess_frame_rate (Context *ctx, Stream *stream, Frame *frame)
//Guess the frame rate, based on both the container and codec information.
func (s *Context) AvGuessFrameRate(st *Stream, fr *Frame) Rational {
	return (Rational)(C.av_guess_frame_rate((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

//int avformat_match_stream_specifier (Context *s, Stream *st, const char *spec)
//Check if the stream st contained in s is matched by the stream specifier spec.
func (s *Context) AvformatMatchStreamSpecifier(st *Stream, spec string) int {
	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st), C.CString(spec)))
}

//int avformat_queue_attached_pictures (Context *s)
func (s *Context) AvformatQueueAttachedPictures() int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(s)))
}
