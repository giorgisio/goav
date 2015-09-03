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
	"github.com/giorgisio/goav/avcodec"
	"unsafe"
)

type (
	AvProbeData                C.struct_AVProbeData
	AvOutputFormat             C.struct_AVOutputFormat
	AvInputFormat              C.struct_AVInputFormat
	AvIndexEntry               C.struct_AVIndexEntry
	AvStream                   C.struct_AVStream
	AvProgram                  C.struct_AVProgram
	AvChapter                  C.struct_AVChapter
	AvFormatContext            C.struct_AVFormatContext
	AvPacketList               C.struct_AVPacketList
	AvPacket                   C.struct_AVPacket
	AvCodecParserContext       C.struct_AVCodecParserContext
	AvIOContext                C.struct_AVIOContext
	AvRational                 C.struct_AVRational
	AvCodec                    C.struct_AVCodec
	AvCodecTag                 C.struct_AVCodecTag
	AvDictionary               C.struct_AVDictionary
	AvFrame                    C.struct_AVFrame
	AvClass                    C.struct_AVClass
	AvCodecContext             C.struct_AVCodecContext
	AvFormatInternal           C.struct_AVFormatInternal
	AvIOInterruptCB            C.struct_AVIOInterruptCB
	AvPacketSideData           C.struct_AVPacketSideData
	FFFrac                     C.struct_FFFrac
	AvStreamParseType          C.enum_AVStreamParseType
	AvDiscard                  C.enum_AVDiscard
	AvMediaType                C.enum_AVMediaType
	AvDurationEstimationMethod C.enum_AVDurationEstimationMethod
	AvPacketSideDataType       C.enum_AVPacketSideDataType
	AvCodecId                  C.enum_AVCodecID
)

type File C.FILE

//int av_get_packet (AvIOContext *s, AvPacket *pkt, int size)
//Allocate and read the payload of a packet and initialize its fields with default values.
func Av_get_packet(ctxt *AvIOContext, pkt *AvPacket, s int) int {
	return int(C.av_get_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(s)))
}

//int av_append_packet (AvIOContext *s, AvPacket *pkt, int size)
//Read data and append it to the current content of the AvPacket.
func Av_append_packet(ctxt *AvIOContext, pkt *AvPacket, s int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(s)))
}

//AvRational av_stream_get_r_frame_rate (const AvStream *s)
func Av_stream_get_r_frame_rate(s *AvStream) AvRational {
	return (AvRational)(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(s)))
}

//void av_stream_set_r_frame_rate (AvStream *s, AvRational r)
func Av_stream_set_r_frame_rate(s *AvStream, r AvRational) {
	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(s), (C.struct_AVRational)(r))
}

//struct AvCodecParserContext * av_stream_get_parser (const AvStream *s)
func Av_stream_get_parser(s *AvStream) *AvCodecParserContext {
	return (*AvCodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(s)))
}

// //char * av_stream_get_recommended_encoder_configuration (const AvStream *s)
// func Av_stream_get_recommended_encoder_configuration(s *AvStream) string {
// 	return C.GoString(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(s)))
// }

// //void av_stream_set_recommended_encoder_configuration (AvStream *s, char *configuration)
// func Av_stream_set_recommended_encoder_configuration(s *AvStream, c string) {
// 	C.av_stream_set_recommended_encoder_configuration((*C.struct_AVStream)(s), C.CString(c))
// }

//int64_t av_stream_get_end_pts (const AvStream *st)
//Returns the pts of the last muxed packet + its duration.
func Av_stream_get_end_pts(s *AvStream) int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(s)))
}

//int av_format_get_probe_score (const AvFormatContext *s)
func Av_format_get_probe_score(s *AvFormatContext) int {
	return int(C.av_format_get_probe_score((*C.struct_AVFormatContext)(s)))
}

//AvCodec * av_format_get_video_codec (const AvFormatContext *s)
func Av_format_get_video_codec(s *AvFormatContext) *AvCodec {
	return (*AvCodec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_video_codec (AvFormatContext *s, AvCodec *c)
func Av_format_set_video_codec(s *AvFormatContext, c *AvCodec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

//AvCodec * av_format_get_audio_codec (const AvFormatContext *s)
func Av_format_get_audio_codec(s *AvFormatContext) *AvCodec {
	return (*AvCodec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_audio_codec (AvFormatContext *s, AvCodec *c)
func Av_format_set_audio_codec(s *AvFormatContext, c *AvCodec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

//AvCodec * av_format_get_subtitle_codec (const AvFormatContext *s)
func Av_format_get_subtitle_codec(s *AvFormatContext) *AvCodec {
	return (*AvCodec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_subtitle_codec (AvFormatContext *s, AvCodec *c)
func Av_format_set_subtitle_codec(s *AvFormatContext, c *AvCodec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// //AvCodec * av_format_get_data_codec (const AvFormatContext *s)
// func Av_format_get_data_codec(s *AvFormatContext) *AvCodec {
// 	return (*AvCodec)(C.av_format_get_data_codec((*C.struct_AVFormatContext)(s)))
// }

// //void av_format_set_data_codec (AvFormatContext *s, AvCodec *c)
// func Av_format_set_data_codec(s *AvFormatContext, c *AvCodec) {
// 	C.av_format_set_data_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
// }

//int av_format_get_metadata_header_padding (const AvFormatContext *s)
func Av_format_get_metadata_header_padding(s *AvFormatContext) int {
	return int(C.av_format_get_metadata_header_padding((*C.struct_AVFormatContext)(s)))
}

//void av_format_set_metadata_header_padding (AvFormatContext *s, int c)
func Av_format_set_metadata_header_padding(s *AvFormatContext, c int) {
	C.av_format_set_metadata_header_padding((*C.struct_AVFormatContext)(s), C.int(c))
}

//void * av_format_get_opaque (const AvFormatContext *s)
func Av_format_get_opaque(s *AvFormatContext) {
	C.av_format_get_opaque((*C.struct_AVFormatContext)(s))
}

//void av_format_set_opaque (AvFormatContext *s, void *opaque)
func Av_format_set_opaque(s *AvFormatContext, o int) {
	C.av_format_set_opaque((*C.struct_AVFormatContext)(s), unsafe.Pointer(&o))
}

//av_format_control_message av_format_get_control_message_cb (const AvFormatContext *s)
func Av_format_control_message(s *AvFormatContext) C.av_format_control_message {
	return C.av_format_get_control_message_cb((*C.struct_AVFormatContext)(s))
}

//void av_format_set_control_message_cb (AvFormatContext *s, av_format_control_message callback)
func Av_format_set_control_message_cb(s *AvFormatContext, c C.av_format_control_message) {
	C.av_format_set_control_message_cb((*C.struct_AVFormatContext)(s), c)
}

//void av_format_inject_global_side_data (AvFormatContext *s)
//This function will cause global side data to be injected in the next packet of each stream as well as after any subsequent seek.
func Av_format_inject_global_side_data(s *AvFormatContext) {
	C.av_format_inject_global_side_data((*C.struct_AVFormatContext)(s))
}

//enum AvDurationEstimationMethod av_fmt_ctx_get_duration_estimation_method (const AvFormatContext *ctx)
//Returns the method used to set ctx->duration.
func Av_fmt_ctx_get_duration_estimation_method(ctx *AvFormatContext) AvDurationEstimationMethod {
	return (AvDurationEstimationMethod)(C.av_fmt_ctx_get_duration_estimation_method((*C.struct_AVFormatContext)(ctx)))
}

//unsigned avformat_version (void)
//Return the LIBAvFORMAT_VERSION_INT constant.
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

//void av_register_input_format (AvInputFormat *format)
func Av_register_input_format(f *AvInputFormat) {
	C.av_register_input_format((*C.struct_AVInputFormat)(f))
}

//void av_register_output_format (AvOutputFormat *format)
func Av_register_output_format(f *AvOutputFormat) {
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

//AvInputFormat * av_iformat_next (const AvInputFormat *f)
//If f is NULL, returns the first registered input format, if f is non-NULL, returns the next registered input format after f or NULL if f is the last one.
func Av_iformat_next(f *AvInputFormat) *AvInputFormat {
	return (*AvInputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

//AvOutputFormat * av_oformat_next (const AvOutputFormat *f)
//If f is NULL, returns the first registered output format, if f is non-NULL, returns the next registered output format after f or NULL if f is the last one.
func Av_oformat_next(f *AvOutputFormat) *AvOutputFormat {
	return (*AvOutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

//AvFormatContext * avformat_alloc_context (void)
//Allocate an AvFormatContext.
func Avformat_alloc_context() *AvFormatContext {
	return (*AvFormatContext)(C.avformat_alloc_context())
}

//void avformat_free_context (AvFormatContext *s)
//Free an AvFormatContext and all its streams.
func Avformat_free_context(s *AvFormatContext) {
	C.avformat_free_context((*C.struct_AVFormatContext)(s))
}

//const AvClass * avformat_get_class (void)
//Get the AvClass for AvFormatContext.
func Avformat_get_class() *AvClass {
	return (*AvClass)(C.avformat_get_class())
}

//AvStream * avformat_new_stream (AvFormatContext *s, const AvCodec *c)
//Add a new stream to a media file.
func Avformat_new_stream(s *AvFormatContext, c *AvCodec) *AvStream {
	return (*AvStream)(C.avformat_new_stream((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c)))
}

//uint8_t * av_stream_get_side_data (AvStream *stream, enum AvPacketSideDataType type, int *size)
//Get side information from stream.
func Av_stream_get_side_data(s *AvStream, t AvPacketSideDataType, z int) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(s), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&z))))
}

//AvProgram * av_new_program (AvFormatContext *s, int id)
func Av_new_program(s *AvFormatContext, id int) *AvProgram {
	return (*AvProgram)(C.av_new_program((*C.struct_AVFormatContext)(s), C.int(id)))
}

//int avformat_alloc_output_context2 (AvFormatContext **ctx, AvOutputFormat *oformat, const char *format_name, const char *filename)
//Allocate an AvFormatContext for an output format.
func Avformat_alloc_output_context2(ctx **AvFormatContext, o *AvOutputFormat, fo, fi string) int {
	return int(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(ctx)), (*C.struct_AVOutputFormat)(o), C.CString(fo), C.CString(fi)))
}

//AvInputFormat * av_find_input_format (const char *short_name)
//Find AvInputFormat based on the short name of the input format.
func Av_find_input_format(s string) *AvInputFormat {
	return (*AvInputFormat)(C.av_find_input_format(C.CString(s)))
}

//AvInputFormat * av_probe_input_format (AvProbeData *pd, int is_opened)
//Guess the file format.
func Av_probe_input_format(pd *AvProbeData, i int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), C.int(i)))
}

//AvInputFormat * av_probe_input_format2 (AvProbeData *pd, int is_opened, int *score_max)
//Guess the file format.
func Av_probe_input_format2(pd *AvProbeData, o int, sm *int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sm))))
}

//AvInputFormat * av_probe_input_format3 (AvProbeData *pd, int is_opened, int *score_ret)
//Guess the file format.
func Av_probe_input_format3(pd *AvProbeData, o int, sl *int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sl))))
}

//int av_probe_input_buffer2 (AvIOContext *pb, AvInputFormat **fmt, const char *filename, void *logctx, unsigned int offset, unsigned int max_probe_size)
//Probe a bytestream to determine the input format.
func Av_probe_input_buffer2(pb *AvIOContext, f **AvInputFormat, fi string, l int, o, m uint) int {
	return int(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), C.CString(fi), unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//int av_probe_input_buffer (AvIOContext *pb, AvInputFormat **fmt, const char *filename, void *logctx, unsigned int offset, unsigned int max_probe_size)
//Like av_probe_input_buffer2() but returns 0 on success.
func Av_probe_input_buffer(pb *AvIOContext, f **AvInputFormat, fi string, l int, o, m uint) int {
	return int(C.av_probe_input_buffer((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), C.CString(fi), unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//int avformat_open_input (AvFormatContext **ps, const char *filename, AvInputFormat *fmt, AvDictionary **options)
//Open an input stream and read the header.
func Avformat_open_input(ps **AvFormatContext, fi string, fmt *AvInputFormat, d **AvDictionary) int {
	return int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)), C.CString(fi), (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//int avformat_find_stream_info (AvFormatContext *ic, AvDictionary **options)
//Read packets of a media file to get stream information.
func Avformat_find_stream_info(ic *AvFormatContext, d **AvDictionary) int {
	return int(C.avformat_find_stream_info((*C.struct_AVFormatContext)(ic), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//AvProgram * av_find_program_from_stream (AvFormatContext *ic, AvProgram *last, int s)
//Find the programs which belong to a given stream.
func Av_find_program_from_stream(ic *AvFormatContext, l *AvProgram, s int) *AvProgram {
	return (*AvProgram)(C.av_find_program_from_stream((*C.struct_AVFormatContext)(ic), (*C.struct_AVProgram)(l), C.int(s)))
}

//int av_find_best_stream (AvFormatContext *ic, enum AvMediaType type, int wanted_stream_nb, int related_stream, AvCodec **decoder_ret, int flags)
//Find the "best" stream in the file.
func Av_find_best_stream(ic *AvFormatContext, t AvMediaType, ws, rs int, c **AvCodec, f int) int {
	return int(C.av_find_best_stream((*C.struct_AVFormatContext)(ic), (C.enum_AVMediaType)(t), C.int(ws), C.int(rs), (**C.struct_AVCodec)(unsafe.Pointer(c)), C.int(f)))
}

//int av_read_frame (AvFormatContext *s, AvPacket *pkt)
//Return the next frame of a stream.
func Av_read_frame(s *AvFormatContext, pkt *avcodec.AvPacket) int {
	return int(C.av_read_frame((*C.struct_AVFormatContext)(unsafe.Pointer(s)), (*C.struct_AVPacket)(unsafe.Pointer(pkt))))
}

//int av_seek_frame (AvFormatContext *s, int stream_index, int64_t timestamp, int flags)
//Seek to the keyframe at timestamp.
func Av_seek_frame(s *AvFormatContext, st int, t int64, f int) int {
	return int(C.av_seek_frame((*C.struct_AVFormatContext)(s), C.int(st), C.int64_t(t), C.int(f)))
}

//int avformat_seek_file (AvFormatContext *s, int stream_index, int64_t min_ts, int64_t ts, int64_t max_ts, int flags)
//Seek to timestamp ts.
func Avformat_seek_file(s *AvFormatContext, si int, mit, ts, mat int64, f int) int {
	return int(C.avformat_seek_file((*C.struct_AVFormatContext)(s), C.int(si), C.int64_t(mit), C.int64_t(ts), C.int64_t(mat), C.int(f)))
}

// //int avformat_flush (AvFormatContext *s)
// //Discard all internally buffered data.
// func Avformat_flush(s *AvFormatContext) int {
// 	return int(C.avformat_flush((*C.struct_AVFormatContext)(s)))
// }

//int av_read_play (AvFormatContext *s)
//Start playing a network-based stream (e.g.
func Av_read_play(s *AvFormatContext) int {
	return int(C.av_read_play((*C.struct_AVFormatContext)(s)))
}

//int av_read_pause (AvFormatContext *s)
//Pause a network-based stream (e.g.
func Av_read_pause(s *AvFormatContext) int {
	return int(C.av_read_pause((*C.struct_AVFormatContext)(s)))
}

//void avformat_close_input (AvFormatContext **s)
//Close an opened input AvFormatContext.
func Avformat_close_input(s *AvFormatContext) {
	C.avformat_close_input((**C.struct_AVFormatContext)(unsafe.Pointer(s)))
}

//int avformat_write_header (AvFormatContext *s, AvDictionary **options)
//Allocate the stream private data and write the stream header to an output media file.
func Avformat_write_header(s *AvFormatContext, o **AvDictionary) int {
	return int(C.avformat_write_header((*C.struct_AVFormatContext)(s), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//int av_write_frame (AvFormatContext *s, AvPacket *pkt)
//Write a packet to an output media file.
func Av_write_frame(s *AvFormatContext, pkt *AvPacket) int {
	return int(C.av_write_frame((*C.struct_AVFormatContext)(s), (*C.struct_AVPacket)(pkt)))
}

//int av_interleaved_write_frame (AvFormatContext *s, AvPacket *pkt)
//Write a packet to an output media file ensuring correct interleaving.
func Av_interleaved_write_frame(s *AvFormatContext, pkt *AvPacket) int {
	return int(C.av_interleaved_write_frame((*C.struct_AVFormatContext)(s), (*C.struct_AVPacket)(pkt)))
}

//int av_write_uncoded_frame (AvFormatContext *s, int stream_index, AvFrame *frame)
//Write a uncoded frame to an output media file.
func Av_write_uncoded_frame(s *AvFormatContext, si int, f *AvFrame) int {
	return int(C.av_write_uncoded_frame((*C.struct_AVFormatContext)(s), C.int(si), (*C.struct_AVFrame)(f)))
}

//int av_interleaved_write_uncoded_frame (AvFormatContext *s, int stream_index, AvFrame *frame)
//Write a uncoded frame to an output media file.
func Av_interleaved_write_uncoded_frame(s *AvFormatContext, si int, f *AvFrame) int {
	return int(C.av_interleaved_write_uncoded_frame((*C.struct_AVFormatContext)(s), C.int(si), (*C.struct_AVFrame)(f)))
}

//int av_write_uncoded_frame_query (AvFormatContext *s, int stream_index)
//Test whether a muxer supports uncoded frame.
func Av_write_uncoded_frame_query(s *AvFormatContext, si int) int {
	return int(C.av_write_uncoded_frame_query((*C.struct_AVFormatContext)(s), C.int(si)))
}

//int av_write_trailer (AvFormatContext *s)
//Write the stream trailer to an output media file and free the file private data.
func Av_write_trailer(s *AvFormatContext) int {
	return int(C.av_write_trailer((*C.struct_AVFormatContext)(s)))
}

//AvOutputFormat * av_guess_format (const char *short_name, const char *filename, const char *mime_type)
//Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
func Av_guess_format(sn, f, mt string) *AvOutputFormat {
	return (*AvOutputFormat)(C.av_guess_format(C.CString(sn), C.CString(f), C.CString(mt)))
}

//enum AvCodecId av_guess_codec (AvOutputFormat *fmt, const char *short_name, const char *filename, const char *mime_type, enum AvMediaType type)
//Guess the codec ID based upon muxer and filename.
func Av_guess_codec(fmt *AvOutputFormat, sn, f, mt string, t AvMediaType) AvCodecId {
	return (AvCodecId)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt), C.CString(sn), C.CString(f), C.CString(mt), (C.enum_AVMediaType)(t)))
}

//int av_get_output_timestamp (struct AvFormatContext *s, int stream, int64_t *dts, int64_t *wall)
//Get timing information for the data currently output.
func Av_get_output_timestamp(s *AvFormatContext, st int, dts, wall *C.int64_t) int {
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

//void av_pkt_dump2 (FILE *f, const AvPacket *pkt, int dump_payload, const AvStream *st)
//Send a nice dump of a packet to the specified file stream.
func Av_pkt_dump2(f *File, pkt *AvPacket, dp int, st *AvStream) {
	C.av_pkt_dump2((*C.FILE)(f), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//void av_pkt_dump_log2 (void *avcl, int level, const AvPacket *pkt, int dump_payload, const AvStream *st)
//Send a nice dump of a packet to the log.
func Av_pkt_dump_log2(a int, l int, pkt *AvPacket, dp int, st *AvStream) {
	C.av_pkt_dump_log2(unsafe.Pointer(&a), C.int(l), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//enum AvCodecId av_codec_get_id (const struct AvCodecTag *const *tags, unsigned int tag)
//Get the AvCodecId for the given codec tag tag.
func Av_codec_get_id(t **AvCodecTag, tag uint) AvCodecId {
	return (AvCodecId)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(t)), C.uint(tag)))
}

//unsigned int av_codec_get_tag (const struct AvCodecTag *const *tags, enum AvCodecId id)
//Get the codec tag for the given codec id id.
func Av_codec_get_tag(t **AvCodecTag, id AvCodecId) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id)))
}

//int av_codec_get_tag2 (const struct AvCodecTag *const *tags, enum AvCodecId id, unsigned int *tag)
//Get the codec tag for the given codec id.
func Av_codec_get_tag2(t **AvCodecTag, id AvCodecId, tag *uint) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

//int av_find_default_stream_index (AvFormatContext *s)
func Av_find_default_stream_index(s *AvFormatContext) int {
	return int(C.av_find_default_stream_index((*C.struct_AVFormatContext)(s)))
}

//int av_index_search_timestamp (AvStream *st, int64_t timestamp, int flags)
//Get the index for a specific timestamp.
func Av_index_search_timestamp(st *AvStream, t int64, f int) int {
	return int(C.av_index_search_timestamp((*C.struct_AVStream)(st), C.int64_t(t), C.int(f)))
}

//int av_add_index_entry (AvStream *st, int64_t pos, int64_t timestamp, int size, int distance, int flags)
//Add an index entry into a sorted list.
func Av_add_index_entry(st *AvStream, pos, t, int64, s, d, f int) int {
	return int(C.av_add_index_entry((*C.struct_AVStream)(st), C.int64_t(pos), C.int64_t(t), C.int(s), C.int(d), C.int(f)))
}

//void av_url_split (char *proto, int proto_size, char *authorization, int authorization_size, char *hostname, int hostname_size, int *port_ptr, char *path, int path_size, const char *url)
//Split a URL string into components.
func Av_url_split(p string, ps int, a string, as int, h string, hs int, pp *int, path string, psize int, url string) {
	C.av_url_split(C.CString(p), C.int(ps), C.CString(a), C.int(as), C.CString(h), C.int(hs), (*C.int)(unsafe.Pointer(pp)), C.CString(path), C.int(psize), C.CString(url))
}

//void av_dump_format (AvFormatContext *ic, int index, const char *url, int is_output)
//Print detailed information about the input or output format, such as duration, bitrate, streams, container, programs, metadata, side data, codec and time base.
func Av_dump_format(ic *AvFormatContext, i int, url string, io int) {
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

//int av_sdp_create (AvFormatContext *ac[], int n_files, char *buf, int size)
//Generate an SDP for an RTP session.
func Av_sdp_create(ac **AvFormatContext, nf int, b string, s int) int {
	return int(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(ac)), C.int(nf), C.CString(b), C.int(s)))
}

//int av_match_ext (const char *filename, const char *extensions)
//Return a positive value if the given filename has one of the given extensions, 0 otherwise.
func Av_match_ext(f, e string) int {
	return int(C.av_match_ext(C.CString(f), C.CString(e)))
}

//int avformat_query_codec (const AvOutputFormat *ofmt, enum AvCodecId codec_id, int std_compliance)
//Test if the given container can store a codec.
func Avformat_query_codec(o *AvOutputFormat, cd AvCodecId, sc int) int {
	return int(C.avformat_query_codec((*C.struct_AVOutputFormat)(o), (C.enum_AVCodecID)(cd), C.int(sc)))
}

//struct AvCodecTag * avformat_get_riff_video_tags (void)
func Avformat_get_riff_video_tags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_video_tags())
}

//struct AvCodecTag * avformat_get_riff_audio_tags (void)
func Avformat_get_riff_audio_tags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_audio_tags())
}

//struct AvCodecTag * avformat_get_mov_video_tags (void)
func Avformat_get_mov_video_tags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_video_tags())
}

//struct AvCodecTag * avformat_get_mov_audio_tags (void)
func Avformat_get_mov_audio_tags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_audio_tags())
}

//AvRational av_guess_sample_aspect_ratio (AvFormatContext *format, AvStream *stream, AvFrame *frame)
//Guess the sample aspect ratio of a frame, based on both the stream and the frame aspect ratio.
func Av_guess_sample_aspect_ratio(f *AvFormatContext, st *AvStream, fr *AvFrame) AvRational {
	return (AvRational)(C.av_guess_sample_aspect_ratio((*C.struct_AVFormatContext)(f), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

//AvRational av_guess_frame_rate (AvFormatContext *ctx, AvStream *stream, AvFrame *frame)
//Guess the frame rate, based on both the container and codec information.
func Av_guess_frame_rate(ctx *AvFormatContext, st *AvStream, fr *AvFrame) AvRational {
	return (AvRational)(C.av_guess_frame_rate((*C.struct_AVFormatContext)(ctx), (*C.struct_AVStream)(st), (*C.struct_AVFrame)(fr)))
}

//int avformat_match_stream_specifier (AvFormatContext *s, AvStream *st, const char *spec)
//Check if the stream st contained in s is matched by the stream specifier spec.
func Avformat_match_stream_specifier(s *AvFormatContext, st *AvStream, spec string) int {
	return int(C.avformat_match_stream_specifier((*C.struct_AVFormatContext)(s), (*C.struct_AVStream)(st), C.CString(spec)))
}

//int avformat_queue_attached_pictures (AvFormatContext *s)
func Avformat_queue_attached_pictures(s *AvFormatContext) int {
	return int(C.avformat_queue_attached_pictures((*C.struct_AVFormatContext)(s)))
}
