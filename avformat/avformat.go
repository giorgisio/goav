/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)

	Supported formats (muxers and demuxers) provided by the libavformat library
	The libavformat library provides some generic global options,
	which can be set on all the muxers and demuxers.
	In addition each muxer or demuxer may support so-called private options,
	which are specific for that component.
*/
package avformat

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"
import (
	"github.com/giorgisio/gfg/avcodec"
	"unsafe"
)

type (
	AVProbeData          C.struct_AVProbeData
	AVOutputFormat       C.struct_AVOutputFormat
	AVInputFormat        C.struct_AVInputFormat
	AVIndexEntry         C.struct_AVIndexEntry
	AVStream             C.struct_AVStream
	AVProgram            C.struct_AVProgram
	AVChapter            C.struct_AVChapter
	AVFormatContext      C.struct_AVFormatContext
	AVPacketList         C.struct_AVPacketList
	AVPacket             C.struct_AVPacket
	AVCodecParserContext C.struct_AVCodecParserContext
	AVIOContext          C.struct_AVIOContext
	AVRational           C.struct_AVRational
	AVCodec              C.struct_AVCodec
	AVCodecTag           C.struct_AVCodecTag
	AVDictionary         C.struct_AVDictionary
	AVFrame              C.struct_AVFrame
	AVClass              C.struct_AVClass
	AVCodecContext       C.struct_AVCodecContext
)
type (
	AVMediaType                C.enum_AVMediaType
	AVDurationEstimationMethod C.enum_AVDurationEstimationMethod
	AVPacketSideDataType       C.enum_AVPacketSideDataType
	AVCodecID                  C.enum_AVCodecID
)

type File C.FILE

//int av_get_packet (AVIOContext *s, AVPacket *pkt, int size)
//Allocate and read the payload of a packet and initialize its fields with default values.
func Av_get_packet(ctxt *AVIOContext, pkt *AVPacket, s int) int {
	return int(C.av_get_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(s)))
}

//int av_append_packet (AVIOContext *s, AVPacket *pkt, int size)
//Read data and append it to the current content of the AVPacket.
func Av_append_packet(ctxt *AVIOContext, pkt *AVPacket, s int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(s)))
}

//AVRational av_stream_get_r_frame_rate (const AVStream *s)
func Av_stream_get_r_frame_rate(s *AVStream) AVRational {
	return (AVRational)(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

//void av_stream_set_r_frame_rate (AVStream *s, AVRational r)
func Av_stream_set_r_frame_rate(s *AVStream, r AVRational) {
	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), (C.struct_AVRational)(r))
}

//struct AVCodecParserContext * av_stream_get_parser (const AVStream *s)
func Av_stream_get_parser(s *AVStream) *AVCodecParserContext {
	return (*AVCodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

//char * av_stream_get_recommended_encoder_configuration (const AVStream *s)
func Av_stream_get_recommended_encoder_configuration(s *AVStream) string {
	return C.GoString(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(s)))
}

//void av_stream_set_recommended_encoder_configuration (AVStream *s, char *configuration)
func Av_stream_set_recommended_encoder_configuration(s *AVStream, c string) {
	C.av_stream_set_recommended_encoder_configuration((*C.struct_AVStream)(s), C.CString(c))
}

//int64_t av_stream_get_end_pts (const AVStream *st)
//Returns the pts of the last muxed packet + its duration.
func Av_stream_get_end_pts(s *AVStream) int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}

//int av_format_get_probe_score (const AVFormatContext *s)
func Av_format_get_probe_score(s *AVFormatContext) int {
	return int(C.av_format_get_probe_score((*C.struct_AVFormatContext)(s)))
}

//AVCodec * av_format_get_video_codec (const AVFormatContext *s)
func Av_format_get_video_codec(s *AVFormatContext) *AVCodec {
	return (*AVCodec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_video_codec (AVFormatContext *s, AVCodec *c)
func Av_format_set_video_codec(s *AVFormatContext, c *AVCodec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

//AVCodec * av_format_get_audio_codec (const AVFormatContext *s)
func Av_format_get_audio_codec(s *AVFormatContext) *AVCodec {
	return (*AVCodec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_audio_codec (AVFormatContext *s, AVCodec *c)
func Av_format_set_audio_codec(s *AVFormatContext, c *AVCodec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

//AVCodec * av_format_get_subtitle_codec (const AVFormatContext *s)
func Av_format_get_subtitle_codec(s *AVFormatContext) *AVCodec {
	return (*AVCodec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_subtitle_codec (AVFormatContext *s, AVCodec *c)
func Av_format_set_subtitle_codec(s *AVFormatContext, c *AVCodec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

//AVCodec * av_format_get_data_codec (const AVFormatContext *s)
func Av_format_get_data_codec(s *AVFormatContext) *AVCodec {
	return (*AVCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_data_codec (AVFormatContext *s, AVCodec *c)
func Av_format_set_data_codec(s *AVFormatContext, c *AVCodec) {
	C.av_format_set_data_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

//int av_format_get_metadata_header_padding (const AVFormatContext *s)
func Av_format_get_metadata_header_padding(s *AVFormatContext) int {
	return int(C.av_format_get_metadata_header_padding((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_metadata_header_padding (AVFormatContext *s, int c)
func Av_format_set_metadata_header_padding(s *AVFormatContext, c int) {
	C.av_format_set_metadata_header_padding((*C.struct_AVFormatContext)(s), C.int(c))
}

//void * av_format_get_opaque (const AVFormatContext *s)
func Av_format_get_opaque(s *AVFormatContext) {
	C.av_format_get_opaque((*C.struct_AVFormatContext)(s))
}

//void av_format_set_opaque (AVFormatContext *s, void *opaque)
func Av_format_set_opaque(s *AVFormatContext, o int) {
	C.av_format_set_opaque((*C.struct_AVFormatContext)(s), unsafe.Pointer(&o))
}

//av_format_control_message av_format_get_control_message_cb (const AVFormatContext *s)
func Av_format_control_message(s *AVFormatContext) C.av_format_control_message {
	return C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(s))
}

//void av_format_set_control_message_cb (AVFormatContext *s, av_format_control_message callback)
func Av_format_set_control_message_cb(s *AVFormatContext, c C.av_format_control_message) {
	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(s), c)
}

//void av_format_inject_global_side_data (AVFormatContext *s)
//This function will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
func Av_format_inject_global_side_data(s *AVFormatContext) {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(s))
}

//enum AVDurationEstimationMethod av_fmt_ctx_get_duration_estimation_method (const AVFormatContext *ctx)
//Returns the method used to set ctx->duration.
func Av_fmt_ctx_get_duration_estimation_method(ctx *AVFormatContext) AVDurationEstimationMethod {
	return (AVDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(ctx)))
}

//unsigned avformat_version (void)
//Return the LIBAVFORMAT_VERSION_INT constant.
func Avformat_version() uint {
	return uint(C.avformat_version())
}

//const char * avformat_configuration (void)
//Return the libavformat build-time configuration.
func Avformat_configuration() string {
	return C.GoString(C.avformat_configuration())
}

//const char * avformat_license (void)
//Return the libavformat license.
func Avformat_license() string {
	return C.GoString(C.avformat_license())
}

//void av_register_all (void)
//Initialize libavformat and register all the muxers, demuxers and protocols.
func Av_register_all() {
	C.av_register_all()
}

//void av_register_input_format (AVInputFormat *format)
func Av_register_input_format(f *AVInputFormat) {
	C.av_register_input_format((*C.struct_AVInputFormat)(f))
}

//void av_register_output_format (AVOutputFormat *format)
func Av_register_output_format(f *AVOutputFormat) {
	C.av_register_output_format((*C.struct_AVOutputFormat)(f))
}

//int avformat_network_init (void)
//Do global initialization of network components.
func Avformat_network_init() int {
	return int(C.avformat_network_init())
}

//int avformat_network_deinit (void)
//Undo the initialization done by avformat_network_init.
func Avformat_network_deinit() int {
	return int(C.avformat_network_deinit())
}

//AVInputFormat * av_iformat_next (const AVInputFormat *f)
//If f is NULL, returns the first registered input format, if f is non-NULL, returns the next registered input format after f or NULL if f is the last one.
func Av_iformat_next(f *AVInputFormat) *AVInputFormat {
	return (*AVInputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

//AVOutputFormat * av_oformat_next (const AVOutputFormat *f)
//If f is NULL, returns the first registered output format, if f is non-NULL, returns the next registered output format after f or NULL if f is the last one.
func Av_oformat_next(f *AVOutputFormat) *AVOutputFormat {
	return (*AVOutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

//AVFormatContext * avformat_alloc_context (void)
//Allocate an AVFormatContext.
func Avformat_alloc_context() *AVFormatContext {
	return (*AVFormatContext)(C.avformat_alloc_context())
}

//void avformat_free_context (AVFormatContext *s)
//Free an AVFormatContext and all its streams.
func Avformat_free_context(s *AVFormatContext) {
	C.avformat_free_context((*C.struct_AVFormatContext)(s))
}

//const AVClass * avformat_get_class (void)
//Get the AVClass for AVFormatContext.
func Avformat_get_class() *AVClass {
	return (*AVClass)(C.avformat_get_class())
}

//AVStream * avformat_new_stream (AVFormatContext *s, const AVCodec *c)
//Add a new stream to a media file.
func Avformat_new_stream(s *AVFormatContext, c *AVCodec) *AVStream {
	return (*AVStream)(C.avformat_new_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c)))
}

//uint8_t * av_stream_get_side_data (AVStream *stream, enum AVPacketSideDataType type, int *size)
//Get side information from stream.
func Av_stream_get_side_data(s *AVStream, t AVPacketSideDataType, z int) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(s), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&z))))
}

//AVProgram * av_new_program (AVFormatContext *s, int id)
func Av_new_program(s *AVFormatContext, id int) *AVProgram {
	return (*AVProgram)(C.av_new_program((*C.struct_AVFormatContext)(s), C.int(id)))
}

//int avformat_alloc_output_context2 (AVFormatContext **ctx, AVOutputFormat *oformat, const char *format_name, const char *filename)
//Allocate an AVFormatContext for an output format.
func Avformat_alloc_output_context2(ctx **AVFormatContext, o *AVOutputFormat, fo, fi string) int {
	return int(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(ctx)), (*C.struct_AVOutputFormat)(o), C.CString(fo), C.CString(fi)))
}

//AVInputFormat * av_find_input_format (const char *short_name)
//Find AVInputFormat based on the short name of the input format.
func Av_find_input_format(s string) *AVInputFormat {
	return (*AVInputFormat)(C.av_find_input_format(C.CString(s)))
}

//AVInputFormat * av_probe_input_format (AVProbeData *pd, int is_opened)
//Guess the file format.
func Av_probe_input_format(pd *AVProbeData, i int) *AVInputFormat {
	return (*AVInputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), C.int(i)))
}

//AVInputFormat * av_probe_input_format2 (AVProbeData *pd, int is_opened, int *score_max)
//Guess the file format.
func Av_probe_input_format2(pd *AVProbeData, o int, sm *int) *AVInputFormat {
	return (*AVInputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sm))))
}

//AVInputFormat * av_probe_input_format3 (AVProbeData *pd, int is_opened, int *score_ret)
//Guess the file format.
func Av_probe_input_format3(pd *AVProbeData, o int, sl *int) *AVInputFormat {
	return (*AVInputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sl))))
}

//int av_probe_input_buffer2 (AVIOContext *pb, AVInputFormat **fmt, const char *filename, void *logctx, unsigned int offset, unsigned int max_probe_size)
//Probe a bytestream to determine the input format.
func Av_probe_input_buffer2(pb *AVIOContext, f **AVInputFormat, fi string, l int, o, m uint) int {
	return int(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), C.CString(fi), unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//int av_probe_input_buffer (AVIOContext *pb, AVInputFormat **fmt, const char *filename, void *logctx, unsigned int offset, unsigned int max_probe_size)
//Like av_probe_input_buffer2() but returns 0 on success.
func Av_probe_input_buffer(pb *AVIOContext, f **AVInputFormat, fi string, l int, o, m uint) int {
	return int(C.av_probe_input_buffer((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), C.CString(fi), unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//int avformat_open_input (AVFormatContext **ps, const char *filename, AVInputFormat *fmt, AVDictionary **options)
//Open an input stream and read the header.
func Avformat_open_input(ps **AVFormatContext, fi string, fmt *AVInputFormat, d **AVDictionary) int {
	return int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)), C.CString(fi), (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//int avformat_find_stream_info (AVFormatContext *ic, AVDictionary **options)
//Read packets of a media file to get stream information.
func Avformat_find_stream_info(ic *AVFormatContext, d **AVDictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(ic), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//AVProgram * av_find_program_from_stream (AVFormatContext *ic, AVProgram *last, int s)
//Find the programs which belong to a given stream.
func Av_find_program_from_stream(ic *AVFormatContext, l *AVProgram, s int) *AVProgram {
	return (*AVProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(ic), (*C.struct_AVProgram)(l), C.int(s)))
}

//int av_find_best_stream (AVFormatContext *ic, enum AVMediaType type, int wanted_stream_nb, int related_stream, AVCodec **decoder_ret, int flags)
//Find the "best" stream in the file.
func Av_find_best_stream(ic *AVFormatContext, t AVMediaType, ws, rs int, c **AVCodec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

//int av_read_frame (AVFormatContext *s, AVPacket *pkt)
//Return the next frame of a stream.
func Av_read_frame(s *AVFormatContext, pkt *avcodec.AVPacket) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(s)), (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}

//int av_seek_frame (AVFormatContext *s, int stream_index, int64_t timestamp, int flags)
//Seek to the keyframe at timestamp.
func Av_seek_frame(s *AVFormatContext, st int, t int64, f int) int {
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(s), C.int(st), C.int64_t(t), C.int(f)))
}

//int avformat_seek_file (AVFormatContext *s, int stream_index, int64_t min_ts, int64_t ts, int64_t max_ts, int flags)
//Seek to timestamp ts.
func Avformat_seek_file(s *AVFormatContext, si int, mit, ts, mat int64, f int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(s), C.int(si), C.int64_t(mit), C.int64_t(ts), C.int64_t(mat), C.int(f)))
}

//int avformat_flush (AVFormatContext *s)
//Discard all internally buffered data.
func Avformat_flush(s *AVFormatContext) int {
	return int(C.avformat_flush((*C.struct_AVFormatContext)(s)))
}

//int av_read_play (AVFormatContext *s)
//Start playing a network-based stream (e.g.
func Av_read_play(s *AVFormatContext) int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(s)))
}

//int av_read_pause (AVFormatContext *s)
//Pause a network-based stream (e.g.
func Av_read_pause(s *AVFormatContext) int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(s)))
}

//void avformat_close_input (AVFormatContext **s)
//Close an opened input AVFormatContext.
func Avformat_close_input(s *AVFormatContext) {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(s)))
}

//int avformat_write_header (AVFormatContext *s, AVDictionary **options)
//Allocate the stream private data and write the stream header to an output media file.
func Avformat_write_header(s *AVFormatContext, o **AVDictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//int av_write_frame (AVFormatContext *s, AVPacket *pkt)
//Write a packet to an output media file.
func Av_write_frame(s *AVFormatContext, pkt *AVPacket) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(s), (*C.struct_AVPacket)(pkt)))
}

//int av_interleaved_write_frame (AVFormatContext *s, AVPacket *pkt)
//Write a packet to an output media file ensuring correct interleaving.
func Av_interleaved_write_frame(s *AVFormatContext, pkt *AVPacket) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(s), (*C.struct_AVPacket)(pkt)))
}

//int av_write_uncoded_frame (AVFormatContext *s, int stream_index, AVFrame *frame)
//Write a uncoded frame to an output media file.
func Av_write_uncoded_frame(s *AVFormatContext, si int, f *AVFrame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(s), C.int(si), (*C.struct_AVFrame)(f)))
}

//int av_interleaved_write_uncoded_frame (AVFormatContext *s, int stream_index, AVFrame *frame)
//Write a uncoded frame to an output media file.
func Av_interleaved_write_uncoded_frame(s *AVFormatContext, si int, f *AVFrame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(s), C.int(si), (*C.struct_AVFrame)(f)))
}

//int av_write_uncoded_frame_query (AVFormatContext *s, int stream_index)
//Test whether a muxer supports uncoded frame.
func Av_write_uncoded_frame_query(s *AVFormatContext, si int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(s), C.int(si)))
}

//int av_write_trailer (AVFormatContext *s)
//Write the stream trailer to an output media file and free the file private data.
func Av_write_trailer(s *AVFormatContext) int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(s)))
}

//AVOutputFormat * av_guess_format (const char *short_name, const char *filename, const char *mime_type)
//Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
func Av_guess_format(sn, f, mt string) *AVOutputFormat {
	return (*AVOutputFormat)(C.av_guess_format(C.CString(sn), C.CString(f), C.CString(mt)))
}

//enum AVCodecID av_guess_codec (AVOutputFormat *fmt, const char *short_name, const char *filename, const char *mime_type, enum AVMediaType type)
//Guess the codec ID based upon muxer and filename.
func Av_guess_codec(fmt *AVOutputFormat, sn, f, mt string, t AVMediaType) AVCodecID {
	return (AVCodecID)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt), C.CString(sn), C.CString(f), C.CString(mt), (C.enum_AVMediaType)(t)))
}

//int av_get_output_timestamp (struct AVFormatContext *s, int stream, int64_t *dts, int64_t *wall)
//Get timing information for the data currently output.
func Av_get_output_timestamp(s *AVFormatContext, st int, dts, wall *C.int64_t) int {
	return int(C.av_get_output_timestamp((*C.struct_AVFormatContext)(s), C.int(st), dts, wall))
}

//void av_hex_dump (FILE *f, const uint8_t *buf, int size)
//Send a nice hexadecimal dump of a buffer to the specified file stream.
func Av_hex_dump(f *File, b *uint8, s int) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(b), C.int(s))
}

//void av_hex_dump_log (void *avcl, int level, const uint8_t *buf, int size)
//Send a nice hexadecimal dump of a buffer to the log.
func Av_hex_dump_log(a, l int, b *uint8, s int) {
	C.av_hex_dump_log(unsafe.Pointer(&a), C.int(l), (*C.uint8_t)(b), C.int(s))
}

//void av_pkt_dump2 (FILE *f, const AVPacket *pkt, int dump_payload, const AVStream *st)
//Send a nice dump of a packet to the specified file stream.
func Av_pkt_dump2(f *File, pkt *AVPacket, dp int, st *AVStream) {
	C.av_pkt_dump2((*C.FILE)(f), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//void av_pkt_dump_log2 (void *avcl, int level, const AVPacket *pkt, int dump_payload, const AVStream *st)
//Send a nice dump of a packet to the log.
func Av_pkt_dump_log2(a int, l int, pkt *AVPacket, dp int, st *AVStream) {
	C.av_pkt_dump_log2(unsafe.Pointer(&a), C.int(l), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//enum AVCodecID av_codec_get_id (const struct AVCodecTag *const *tags, unsigned int tag)
//Get the AVCodecID for the given codec tag tag.
func Av_codec_get_id(t **AVCodecTag, tag uint) AVCodecID {
	return (AVCodecID)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(t)), C.uint(tag)))
}

//unsigned int av_codec_get_tag (const struct AVCodecTag *const *tags, enum AVCodecID id)
//Get the codec tag for the given codec id id.
func Av_codec_get_tag(t **AVCodecTag, id AVCodecID) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id)))
}

//int av_codec_get_tag2 (const struct AVCodecTag *const *tags, enum AVCodecID id, unsigned int *tag)
//Get the codec tag for the given codec id.
func Av_codec_get_tag2(t **AVCodecTag, id AVCodecID, tag *uint) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

//int av_find_default_stream_index (AVFormatContext *s)
func Av_find_default_stream_index(s *AVFormatContext) int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(s)))
}

//int av_index_search_timestamp (AVStream *st, int64_t timestamp, int flags)
//Get the index for a specific timestamp.
func Av_index_search_timestamp(st *AVStream, t int64, f int) int {
	return int(C.av_index_search_timestamp((*C.struct_AVStream)(st), C.int64_t(t), C.int(f)))
}

//int av_add_index_entry (AVStream *st, int64_t pos, int64_t timestamp, int size, int distance, int flags)
//Add an index entry into a sorted list.
func Av_add_index_entry(st *AVStream, pos, t, int64, s, d, f int) int {
	return int(C.av_add_index_entry((*C.struct_AVStream)(st), C.int64_t(pos), C.int64_t(t), C.int(s), C.int(d), C.int(f)))
}

//void av_url_split (char *proto, int proto_size, char *authorization, int authorization_size, char *hostname, int hostname_size, int *port_ptr, char *path, int path_size, const char *url)
//Split a URL string into components.
func Av_url_split(p string, ps int, a string, as int, h string, hs int, pp *int, path string, psize int, url string) {
	C.av_url_split(C.CString(p), C.int(ps), C.CString(a), C.int(as), C.CString(h), C.int(hs), (*C.int)(unsafe.Pointer(pp)), C.CString(path), C.int(psize), C.CString(url))
}

//void av_dump_format (AVFormatContext *ic, int index, const char *url, int is_output)
//Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func Av_dump_format(ic *AVFormatContext, i int, url string, io int) {
	C.av_dump_format((*C.struct_AVFormatContext)(unsafe.Pointer(ic)), C.int(i), C.CString(url), C.int(io))
}

//int av_get_frame_filename (char *buf, int buf_size, const char *path, int number)
//Return in 'buf' the path with 'd' replaced by a number.
func Av_get_frame_filename(b string, bs int, pa string, n int) int {
	return int(C.av_get_frame_filename(C.CString(b), C.int(bs), C.CString(pa), C.int(n)))
}

//int av_filename_number_test (const char *filename)
//Check whether filename actually is a numbered sequence generator.
func Av_filename_number_test(f string) int {
	return int(C.av_filename_number_test(C.CString(f)))
}

//int av_sdp_create (AVFormatContext *ac[], int n_files, char *buf, int size)
//Generate an SDP for an RTP session.
func Av_sdp_create(ac **AVFormatContext, nf int, b string, s int) int {
	return int(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(ac)), C.int(nf), C.CString(b), C.int(s)))
}

//int av_match_ext (const char *filename, const char *extensions)
//Return a positive value if the given filename has one of the given extensions, 0 otherwise.
func Av_match_ext(f, e string) int {
	return int(C.av_match_ext(C.CString(f), C.CString(e)))
}

//int avformat_query_codec (const AVOutputFormat *ofmt, enum AVCodecID codec_id, int std_compliance)
//Test if the given container can store a codec.
func Avformat_query_codec(o *AVOutputFormat, cd AVCodecID, sc int) int {
	return int(C.avformat_query_codec((*C.struct_AVOutputFormat)(o), (C.enum_AVCodecID)(cd), C.int(sc)))
}

//struct AVCodecTag * avformat_get_riff_video_tags (void)
func Avformat_get_riff_video_tags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_riff_video_tags())
}

//struct AVCodecTag * avformat_get_riff_audio_tags (void)
func Avformat_get_riff_audio_tags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_riff_audio_tags())
}

//struct AVCodecTag * avformat_get_mov_video_tags (void)
func Avformat_get_mov_video_tags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_mov_video_tags())
}

//struct AVCodecTag * avformat_get_mov_audio_tags (void)
func Avformat_get_mov_audio_tags() *AVCodecTag {
	return (*AVCodecTag)(C.avformat_get_mov_audio_tags())
}

//AVRational av_guess_sample_aspect_ratio (AVFormatContext *format, AVStream *stream, AVFrame *frame)
//Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
func Av_guess_sample_aspect_ratio(f *AVFormatContext, st *AVStream, fr *AVFrame) AVRational {
	return (AVRational)(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(f), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

//AVRational av_guess_frame_rate (AVFormatContext *ctx, AVStream *stream, AVFrame *frame)
//Guess the frame rate, based on both the container and codec information.
func Av_guess_frame_rate(ctx *AVFormatContext, st *AVStream, fr *AVFrame) AVRational {
	return (AVRational)(C.av_guess_frame_rate((*C.struct_AVFormatContext)(ctx), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

//int avformat_match_stream_specifier (AVFormatContext *s, AVStream *st, const char *spec)
//Check if the stream st contained in s is matched by the stream specifier spec.
func Avformat_match_stream_specifier(s *AVFormatContext, st *AVStream, spec string) int {
	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st), C.CString(spec)))
}

//int avformat_queue_attached_pictures (AVFormatContext *s)
func Avformat_queue_attached_pictures(s *AVFormatContext) int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(s)))
}
