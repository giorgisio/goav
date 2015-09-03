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
	Context                    C.struct_AVFormatContext
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
func (ctxt *AvIOContext) AvGetPacket(pkt *AvPacket, s int) int {
	return int(C.av_get_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(s)))
}

//int av_append_packet (AvIOContext *s, AvPacket *pkt, int size)
//Read data and append it to the current content of the AvPacket.
func (ctxt *AvIOContext) AvAppendPacket(pkt *AvPacket, s int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(ctxt), (*C.struct_AVPacket)(pkt), C.int(s)))
}

//unsigned avformat_version (void)
//Return the LIBAvFORMAT_VERSION_INT constant.
func AvformatVersion() uint {
	return uint(C.avformat_version())
}

//const char * avformat_configuration (void)
//Return the libavformat build-time configuration.
func AvformatConfiguration() string {
	return C.GoString(C.avformat_configuration())
}

//const char * avformat_license (void)
//Return the libavformat license.
func AvformatLicense() string {
	return C.GoString(C.avformat_license())
}

//void av_register_all (void)
//Initialize libavformat and register all the muxers, demuxers and protocols.
func AvRegisterAll() {
	C.av_register_all()
}

//void av_register_input_format (AvInputFormat *format)
func AvRegisterInputFormat(f *AvInputFormat) {
	C.av_register_input_format((*C.struct_AVInputFormat)(f))
}

//void av_register_output_format (AvOutputFormat *format)
func (f *AvOutputFormat) AvRegisterOutputFormat() {
	C.av_register_output_format((*C.struct_AVOutputFormat)(f))
}

//int avformat_network_init (void)
//Do global initialization of network components.
func AvformatNetworkInit() int {
	return int(C.avformat_network_init())
}

//int avformat_network_deinit (void)
//Undo the initialization done by avformat_network_init.
func AvformatNetworkDeinit() int {
	return int(C.avformat_network_deinit())
}

//AvInputFormat * av_iformat_next (const AvInputFormat *f)
//If f is NULL, returns the first registered input format, if f is non-NULL, returns the next registered input format after f or NULL if f is the last one.
func (f *AvInputFormat) AvIformatNext() *AvInputFormat {
	return (*AvInputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

//AvOutputFormat * av_oformat_next (const AvOutputFormat *f)
//If f is NULL, returns the first registered output format, if f is non-NULL, returns the next registered output format after f or NULL if f is the last one.
func (f *AvOutputFormat) AvOformatNext() *AvOutputFormat {
	return (*AvOutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

//Context * avformat_alloc_context (void)
//Allocate an Context.
func AvformatAllocContext() *Context {
	return (*Context)(C.avformat_alloc_context())
}

//const AvClass * avformat_get_class (void)
//Get the AvClass for Context.
func AvformatGetClass() *AvClass {
	return (*AvClass)(C.avformat_get_class())
}

//uint8_t * av_stream_get_side_data (AvStream *stream, enum AvPacketSideDataType type, int *size)
//Get side information from stream.
func (s *AvStream) AvStreamGetSideData(t AvPacketSideDataType, z int) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(s), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&z))))
}

//int avformat_alloc_output_context2 (Context **ctx, AvOutputFormat *oformat, const char *format_name, const char *filename)
//Allocate an Context for an output format.
func AvformatAllocOutputContext2(ctx **Context, o *AvOutputFormat, fo, fi string) int {
	return int(C.avformat_alloc_output_context2((**C.struct_AVFormatContext)(unsafe.Pointer(ctx)), (*C.struct_AVOutputFormat)(o), C.CString(fo), C.CString(fi)))
}

//AvInputFormat * av_find_input_format (const char *short_name)
//Find AvInputFormat based on the short name of the input format.
func AvFindInputFormat(s string) *AvInputFormat {
	return (*AvInputFormat)(C.av_find_input_format(C.CString(s)))
}

//AvInputFormat * av_probe_input_format (AvProbeData *pd, int is_opened)
//Guess the file format.
func AvProbeInputFormat(pd *AvProbeData, i int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format((*C.struct_AVProbeData)(pd), C.int(i)))
}

//AvInputFormat * av_probe_input_format2 (AvProbeData *pd, int is_opened, int *score_max)
//Guess the file format.
func AvProbeInputFormat2(pd *AvProbeData, o int, sm *int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format2((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sm))))
}

//AvInputFormat * av_probe_input_format3 (AvProbeData *pd, int is_opened, int *score_ret)
//Guess the file format.
func AvProbeInputFormat3(pd *AvProbeData, o int, sl *int) *AvInputFormat {
	return (*AvInputFormat)(C.av_probe_input_format3((*C.struct_AVProbeData)(pd), C.int(o), (*C.int)(unsafe.Pointer(sl))))
}

//int av_probe_input_buffer2 (AvIOContext *pb, AvInputFormat **fmt, const char *filename, void *logctx, unsigned int offset, unsigned int max_probe_size)
//Probe a bytestream to determine the input format.
func AvProbeInputBuffer2(pb *AvIOContext, f **AvInputFormat, fi string, l int, o, m uint) int {
	return int(C.av_probe_input_buffer2((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), C.CString(fi), unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//int av_probe_input_buffer (AvIOContext *pb, AvInputFormat **fmt, const char *filename, void *logctx, unsigned int offset, unsigned int max_probe_size)
//Like av_probe_input_buffer2() but returns 0 on success.
func AvProbeInputBuffer(pb *AvIOContext, f **AvInputFormat, fi string, l int, o, m uint) int {
	return int(C.av_probe_input_buffer((*C.struct_AVIOContext)(pb), (**C.struct_AVInputFormat)(unsafe.Pointer(f)), C.CString(fi), unsafe.Pointer(&l), C.uint(o), C.uint(m)))
}

//int avformat_open_input (Context **ps, const char *filename, AvInputFormat *fmt, AvDictionary **options)
//Open an input stream and read the header.
func AvformatOpenInput(ps **Context, fi string, fmt *AvInputFormat, d **AvDictionary) int {
	return int(C.avformat_open_input((**C.struct_AVFormatContext)(unsafe.Pointer(ps)), C.CString(fi), (*C.struct_AVInputFormat)(fmt), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//AvOutputFormat * av_guess_format (const char *short_name, const char *filename, const char *mime_type)
//Return the output format in the list of registered output formats which best matches the provided parameters, or return NULL if there is no match.
func AvGuessFormat(sn, f, mt string) *AvOutputFormat {
	return (*AvOutputFormat)(C.av_guess_format(C.CString(sn), C.CString(f), C.CString(mt)))
}

//enum AvCodecId av_guess_codec (AvOutputFormat *fmt, const char *short_name, const char *filename, const char *mime_type, enum AvMediaType type)
//Guess the codec ID based upon muxer and filename.
func AvGuessCodec(fmt *AvOutputFormat, sn, f, mt string, t AvMediaType) AvCodecId {
	return (AvCodecId)(C.av_guess_codec((*C.struct_AVOutputFormat)(fmt), C.CString(sn), C.CString(f), C.CString(mt), (C.enum_AVMediaType)(t)))
}

//void av_hex_dump (FILE *f, const uint8_t *buf, int size)
//Send a nice hexadecimal dump of a buffer to the specified file stream.
func AvHexDump(f *File, b *uint8, s int) {
	C.av_hex_dump((*C.FILE)(f), (*C.uint8_t)(b), C.int(s))
}

//void av_hex_dump_log (void *avcl, int level, const uint8_t *buf, int size)
//Send a nice hexadecimal dump of a buffer to the log.
func AvHexDumpLog(a, l int, b *uint8, s int) {
	C.av_hex_dump_log(unsafe.Pointer(&a), C.int(l), (*C.uint8_t)(b), C.int(s))
}

//void av_pkt_dump2 (FILE *f, const AvPacket *pkt, int dump_payload, const AvStream *st)
//Send a nice dump of a packet to the specified file stream.
func AvPktDump2(f *File, pkt *AvPacket, dp int, st *AvStream) {
	C.av_pkt_dump2((*C.FILE)(f), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//void av_pkt_dump_log2 (void *avcl, int level, const AvPacket *pkt, int dump_payload, const AvStream *st)
//Send a nice dump of a packet to the log.
func AvPktDumpLog2(a int, l int, pkt *AvPacket, dp int, st *AvStream) {
	C.av_pkt_dump_log2(unsafe.Pointer(&a), C.int(l), (*C.struct_AVPacket)(pkt), C.int(dp), (*C.struct_AVStream)(st))
}

//enum AvCodecId av_codec_get_id (const struct AvCodecTag *const *tags, unsigned int tag)
//Get the AvCodecId for the given codec tag tag.
func AvCodecGetId(t **AvCodecTag, tag uint) AvCodecId {
	return (AvCodecId)(C.av_codec_get_id((**C.struct_AVCodecTag)(unsafe.Pointer(t)), C.uint(tag)))
}

//unsigned int av_codec_get_tag (const struct AvCodecTag *const *tags, enum AvCodecId id)
//Get the codec tag for the given codec id id.
func AvCodecGetTag(t **AvCodecTag, id AvCodecId) uint {
	return uint(C.av_codec_get_tag((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id)))
}

//int av_codec_get_tag2 (const struct AvCodecTag *const *tags, enum AvCodecId id, unsigned int *tag)
//Get the codec tag for the given codec id.
func AvCodecGetTag2(t **AvCodecTag, id AvCodecId, tag *uint) int {
	return int(C.av_codec_get_tag2((**C.struct_AVCodecTag)(unsafe.Pointer(t)), (C.enum_AVCodecID)(id), (*C.uint)(unsafe.Pointer(tag))))
}

//int av_index_search_timestamp (AvStream *st, int64_t timestamp, int flags)
//Get the index for a specific timestamp.
func AvIndexSearchTimestamp(st *AvStream, t int64, f int) int {
	return int(C.av_index_search_timestamp((*C.struct_AVStream)(st), C.int64_t(t), C.int(f)))
}

//int av_add_index_entry (AvStream *st, int64_t pos, int64_t timestamp, int size, int distance, int flags)
//Add an index entry into a sorted list.
func AvAddIndexEntry(st *AvStream, pos, t, int64, s, d, f int) int {
	return int(C.av_add_index_entry((*C.struct_AVStream)(st), C.int64_t(pos), C.int64_t(t), C.int(s), C.int(d), C.int(f)))
}

//void av_url_split (char *proto, int proto_size, char *authorization, int authorization_size, char *hostname, int hostname_size, int *port_ptr, char *path, int path_size, const char *url)
//Split a URL string into components.
func AvUrlSplit(p string, ps int, a string, as int, h string, hs int, pp *int, path string, psize int, url string) {
	C.av_url_split(C.CString(p), C.int(ps), C.CString(a), C.int(as), C.CString(h), C.int(hs), (*C.int)(unsafe.Pointer(pp)), C.CString(path), C.int(psize), C.CString(url))
}

//int av_get_frame_filename (char *buf, int buf_size, const char *path, int number)
//Return in 'buf' the path with 'd' replaced by a number.
func AvGetFrameFilename(b string, bs int, pa string, n int) int {
	return int(C.av_get_frame_filename(C.CString(b), C.int(bs), C.CString(pa), C.int(n)))
}

//int av_filename_number_test (const char *filename)
//Check whether filename actually is a numbered sequence generator.
func AvFilenameNumberTest(f string) int {
	return int(C.av_filename_number_test(C.CString(f)))
}

//int av_sdp_create (Context *ac[], int n_files, char *buf, int size)
//Generate an SDP for an RTP session.
func AvSdpCreate(ac **Context, nf int, b string, s int) int {
	return int(C.av_sdp_create((**C.struct_AVFormatContext)(unsafe.Pointer(ac)), C.int(nf), C.CString(b), C.int(s)))
}

//int av_match_ext (const char *filename, const char *extensions)
//Return a positive value if the given filename has one of the given extensions, 0 otherwise.
func AvMatchExt(f, e string) int {
	return int(C.av_match_ext(C.CString(f), C.CString(e)))
}

//int avformat_query_codec (const AvOutputFormat *ofmt, enum AvCodecId codec_id, int std_compliance)
//Test if the given container can store a codec.
func AvformatQueryCodec(o *AvOutputFormat, cd AvCodecId, sc int) int {
	return int(C.avformat_query_codec((*C.struct_AVOutputFormat)(o), (C.enum_AVCodecID)(cd), C.int(sc)))
}

//struct AvCodecTag * avformat_get_riff_video_tags (void)
func AvformatGetRiffVideoTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_video_tags())
}

//struct AvCodecTag * avformat_get_riff_audio_tags (void)
func AvformatGetRiffAudioTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_riff_audio_tags())
}

//struct AvCodecTag * avformat_get_mov_video_tags (void)
func AvformatGetMovVideoTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_video_tags())
}

//struct AvCodecTag * avformat_get_mov_audio_tags (void)
func AvformatGetMovAudioTags() *AvCodecTag {
	return (*AvCodecTag)(C.avformat_get_mov_audio_tags())
}
