/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)

	The codecs (decoders and encoders) provided by the libavcodec library
	libavcodec provides some generic global options, which can be set on all the encoders and decoders.
*/
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
import (
	"unsafe"
)

type (
	AVCodecDescriptor        C.struct_AVCodecDescriptor
	RcOverride               C.struct_RcOverride
	AVPanScan                C.struct_AVPanScan
	AVPacketSideData         C.struct_AVPacketSideData
	AVPacket                 C.struct_AVPacket
	AVCodecContext           C.struct_AVCodecContext
	AVProfile                C.struct_AVProfile
	AVCodec                  C.struct_AVCodec
	AVHWAccel                C.struct_AVHWAccel
	AVPicture                C.struct_AVPicture
	AVSubtitleRect           C.struct_AVSubtitleRect
	AVSubtitle               C.struct_AVSubtitle
	AVCodecParserContext     C.struct_AVCodecParserContext
	AVCodecParser            C.struct_AVCodecParser
	AVBitStreamFilterContext C.struct_AVBitStreamFilterContext
	AVBitStreamFilter        C.struct_AVBitStreamFilter
	AVRational               C.struct_AVRational
	AVDictionary             C.struct_AVDictionary
	AVClass                  C.struct_AVClass
	AVFrame                  C.struct_AVFrame
)

type (
	AVPixelFormat        C.enum_AVPixelFormat
	AVChromaLocation     C.enum_AVChromaLocation
	AVPacketSideDataType C.enum_AVPacketSideDataType
	AVSampleFormat       C.enum_AVSampleFormat
	AVMediaType          C.enum_AVMediaType
	AVCodecID            C.enum_AVCodecID
)

//AVRational av_codec_get_pkt_timebase (const AVCodecContext *avctx)
func Av_codec_get_pkt_timebase(ctxt *AVCodecContext) AVRational {
	return (AVRational)(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(ctxt)))
}

//void av_codec_set_pkt_timebase (AVCodecContext *avctx, AVRational val)
func Av_codec_set_pkt_timebase(ctxt *AVCodecContext, r AVRational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(ctxt), (C.struct_AVRational)(r))
}

//const AVCodecDescriptor * av_codec_get_codec_descriptor (const AVCodecContext *avctx)
func Av_codec_get_codec_descriptor(ctxt *AVCodecContext) *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(ctxt)))
}

//void 	av_codec_set_codec_descriptor (AVCodecContext *avctx, const AVCodecDescriptor *desc)
func Av_codec_set_codec_descriptor(ctxt *AVCodecContext, d *AVCodecDescriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecDescriptor)(d))
}

//int 	av_codec_get_lowres (const AVCodecContext *avctx)
func Av_codec_get_lowres(ctxt *AVCodecContext) int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(ctxt)))
}

//void 	av_codec_set_lowres (AVCodecContext *avctx, int val)
func Av_codec_set_lowres(ctxt *AVCodecContext, i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

//int 	av_codec_get_seek_preroll (const AVCodecContext *avctx)
func Av_codec_get_seek_preroll(ctxt *AVCodecContext) int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(ctxt)))
}

//void 	av_codec_set_seek_preroll (AVCodecContext *avctx, int val)
func Av_codec_set_seek_preroll(ctxt *AVCodecContext, i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

//uint16_t *av_codec_get_chroma_intra_matrix (const AVCodecContext *avctx)
func Av_codec_get_chroma_intra_matrix(ctxt *AVCodecContext) *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt)))
}

//void 	av_codec_set_chroma_intra_matrix (AVCodecContext *avctx, uint16_t *val)
func Av_codec_set_chroma_intra_matrix(ctxt *AVCodecContext, t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt), (*C.uint16_t)(t))
}

//int 	av_codec_get_max_lowres (const AVCodec *codec)
func Av_codec_get_max_lowres(c *AVCodec) int {
	return int(C.av_codec_get_max_lowres((*C.struct_AVCodec)(c)))
}

//If c is NULL, returns the first registered codec, if c is non-NULL,
//returns the next registered codec after c, or NULL if c is the last one.
//AVCodec *av_codec_next (const AVCodec *c)
func Av_codec_next(c *AVCodec) *AVCodec {
	return (*AVCodec)(C.av_codec_next((*C.struct_AVCodec)(c)))
}

//Return the LIBAVCODEC_VERSION_INT constant.
//unsigned 	avcodec_version (void)
func Avcodec_version() uint {
	return uint(C.avcodec_version())
}

//Return the libavcodec build-time configuration.
//const char *avcodec_configuration (void)
func Avcodec_configuration() string {
	return C.GoString(C.avcodec_configuration())

}

//Return the libavcodec license.
//const char *avcodec_license (void)
func Avcodec_license() string {
	return C.GoString(C.avcodec_license())
}

//Register the codec codec and initialize libavcodec.
//void 	avcodec_register (AVCodec *codec)
func Avcodec_register(c *AVCodec) {
	C.avcodec_register((*C.struct_AVCodec)(c))
}

//Register all the codecs, parsers and bitstream filters which were enabled at configuration time.
//void 	avcodec_register_all (void)
func Avcodec_register_all() {
	C.avcodec_register_all()
}

//Allocate an AVCodecContext and set its fields to default values.
//AVCodecContext *avcodec_alloc_context3 (const AVCodec *codec)
func Avcodec_alloc_context3(c *AVCodec) *AVCodecContext {
	return (*AVCodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
}

//Free the codec context and everything associated with it and write NULL to the provided pointer.
//void 	avcodec_free_context (AVCodecContext **avctx)
func Avcodec_free_context(ctxt **AVCodecContext) {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(ctxt)))
}

//Set the fields of the given AVCodecContext to default values corresponding to the given codec (defaults may be codec-dependent).
//int 	avcodec_get_context_defaults3 (AVCodecContext *s, const AVCodec *codec)
func Avcodec_get_context_defaults3(ctxt *AVCodecContext, c *AVCodec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c)))
}

//Get the AVClass for AVCodecContext.
//const AVClass *avcodec_get_class (void)
func Avcodec_get_class() *AVClass {
	return (*AVClass)(C.avcodec_get_class())
}

//Get the AVClass for AVFrame.
//const AVClass *avcodec_get_frame_class (void)
func Avcodec_get_frame_class() *AVClass {
	return (*AVClass)(C.avcodec_get_frame_class())
}

//Get the AVClass for AVSubtitleRect.
//const AVClass *avcodec_get_subtitle_rect_class (void)
func Avcodec_get_subtitle_rect_class() *AVClass {
	return (*AVClass)(C.avcodec_get_subtitle_rect_class())
}

//Copy the settings of the source AVCodecContext into the destination AVCodecContext.
//int 	avcodec_copy_context (AVCodecContext *dest, const AVCodecContext *src)
func Avcodec_copy_context(ctxt, ctxt2 *AVCodecContext) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecContext)(ctxt2)))
}

//Initialize the AVCodecContext to use the given AVCodec
//int 	avcodec_open2 (AVCodecContext *avctx, const AVCodec *codec, AVDictionary **options)
func Avcodec_open2(ctxt *AVCodecContext, c *AVCodec, d **AVDictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Close a given AVCodecContext and free all the data associated with it (but not the AVCodecContext itself).
//int 	avcodec_close (AVCodecContext *avctx)
func Avcodec_close(ctxt *AVCodecContext) int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctxt)))
}

//Free all allocated data in the given subtitle struct.
//void 	avsubtitle_free (AVSubtitle *sub)
func Avsubtitle_free(s *AVSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

//Initialize optional fields of a packet with default values.
//void 	av_init_packet (AVPacket *pkt)
func Av_init_packet(p *AVPacket) {
	C.av_init_packet((*C.struct_AVPacket)(p))
}

//Allocate the payload of a packet and initialize its fields with default values.
//int 	av_new_packet (AVPacket *pkt, int size)
func Av_new_packet(p *AVPacket, s int) int {
	return int(C.av_new_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//Reduce packet size, correctly zeroing padding.
//void 	av_shrink_packet (AVPacket *pkt, int size)
func Av_shrink_packet(p *AVPacket, s int) {
	C.av_shrink_packet((*C.struct_AVPacket)(p), C.int(s))
}

//Increase packet size, correctly zeroing padding.
//int 	av_grow_packet (AVPacket *pkt, int grow_by)
func Av_grow_packet(p *AVPacket, s int) int {
	return int(C.av_grow_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//Initialize a reference-counted packet from av_malloc()ed data.
//int 	av_packet_from_data (AVPacket *pkt, uint8_t *data, int size)
func Av_packet_from_data(p *AVPacket, d *uint8, s int) int {
	return int(C.av_packet_from_data((*C.struct_AVPacket)(p), (*C.uint8_t)(d), C.int(s)))

}

//int 	av_dup_packet (AVPacket *pkt)
func Av_dup_packet(p *AVPacket) int {
	return int(C.av_dup_packet((*C.struct_AVPacket)(p)))

}

//Copy packet, including contents.
//int 	av_copy_packet (AVPacket *dst, const AVPacket *src)
func Av_copy_packet(p, r *AVPacket) int {
	return int(C.av_copy_packet((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//Copy packet side data.
//int 	av_copy_packet_side_data (AVPacket *dst, const AVPacket *src)
func Av_copy_packet_side_data(p, r *AVPacket) int {
	return int(C.av_copy_packet_side_data((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//Free a packet.
//void 	av_free_packet (AVPacket *pkt)
func Av_free_packet(p *AVPacket) {
	C.av_free_packet((*C.struct_AVPacket)(p))

}

//Allocate new information of a packet.
//uint8_t *av_packet_new_side_data (AVPacket *pkt, enum AVPacketSideDataType type, int size)
func Av_packet_new_side_data(p *AVPacket, t AVPacketSideDataType, s int) *uint8 {
	return (*uint8)(C.av_packet_new_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//Shrink the already allocated side data buffer.
//int 	av_packet_shrink_side_data (AVPacket *pkt, enum AVPacketSideDataType type, int size)
func Av_packet_shrink_side_data(p *AVPacket, t AVPacketSideDataType, s int) int {
	return int(C.av_packet_shrink_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//Get side information from packet.
//uint8_t *av_packet_get_side_data (AVPacket *pkt, enum AVPacketSideDataType type, int *size)
func Av_packet_get_side_data(p *AVPacket, t AVPacketSideDataType, s *int) *uint8 {
	return (*uint8)(C.av_packet_get_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(s))))
}

//int 	av_packet_merge_side_data (AVPacket *pkt)
func Av_packet_merge_side_data(p *AVPacket) int {
	return int(C.av_packet_merge_side_data((*C.struct_AVPacket)(p)))
}

//int 	av_packet_split_side_data (AVPacket *pkt)
func Av_packet_split_side_data(p *AVPacket) int {
	return int(C.av_packet_split_side_data((*C.struct_AVPacket)(p)))
}

//const char *av_packet_side_data_name (enum AVPacketSideDataType type)
func Av_packet_side_data_name(t AVPacketSideDataType) string {
	return C.GoString(C.av_packet_side_data_name((C.enum_AVPacketSideDataType)(t)))
}

//Pack a dictionary for use in side_data.
//uint8_t *av_packet_pack_dictionary (AVDictionary *dict, int *size)
func Av_packet_pack_dictionary(d *AVDictionary, s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

//Unpack a dictionary from side_data.
//int 	av_packet_unpack_dictionary (const uint8_t *data, int size, AVDictionary **dict)
func Av_packet_unpack_dictionary(d *uint8, s int, dt **AVDictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt))))
}

//Convenience function to free all the side data stored.
//void 	av_packet_free_side_data (AVPacket *pkt)
func Av_packet_free_side_data(p *AVPacket) {
	C.av_packet_free_side_data((*C.struct_AVPacket)(p))
}

//Setup a new reference to the data described by a given packet.
//int 	av_packet_ref (AVPacket *dst, const AVPacket *src)
func Av_packet_ref(p, s *AVPacket) int {
	return int(C.av_packet_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//Wipe the packet.
//void 	av_packet_unref (AVPacket *pkt)
func Av_packet_unref(p *AVPacket) {
	C.av_packet_unref((*C.struct_AVPacket)(p))
}

//Move every field in src to dst and reset src.
//void 	av_packet_move_ref (AVPacket *dst, AVPacket *src)
func Av_packet_move_ref(p, s *AVPacket) {
	C.av_packet_move_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s))
}

//Copy only "properties" fields from src to dst.
//int 	av_packet_copy_props (AVPacket *dst, const AVPacket *src)
func Av_packet_copy_props(p, s *AVPacket) int {
	return int(C.av_packet_copy_props((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//Convert valid timing fields (timestamps / durations) in a packet from one timebase to another.
//void 	av_packet_rescale_ts (AVPacket *pkt, AVRational tb_src, AVRational tb_dst)
func Av_packet_rescale_ts(p *AVPacket, r, r2 AVRational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(p), (C.struct_AVRational)(r), (C.struct_AVRational)(r2))
}

//Find a registered decoder with a matching codec ID.
//AVCodec *avcodec_find_decoder (enum AVCodecID id)
func Avcodec_find_decoder(id AVCodecID) *AVCodec {
	return (*AVCodec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

//Find a registered decoder with the specified name.
//AVCodec *avcodec_find_decoder_by_name (const char *name)
func Avcodec_find_decoder_by_name(n string) *AVCodec {
	return (*AVCodec)(C.avcodec_find_decoder_by_name(C.CString(n)))
}

//The default callback for AVCodecContext.get_buffer2().
//int 	avcodec_default_get_buffer2 (AVCodecContext *s, AVFrame *frame, int flags)
func Avcodec_default_get_buffer2(s *AVCodecContext, f *AVFrame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(s), (*C.struct_AVFrame)(f), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
//void 	avcodec_align_dimensions (AVCodecContext *s, int *width, int *height)
func Avcodec_align_dimensions(ctxt *AVCodecContext, w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
//void 	avcodec_align_dimensions2 (AVCodecContext *s, int *width, int *height, int linesize_align[AV_NUM_DATA_POINTERS])
func Avcodec_align_dimensions2(ctxt *AVCodecContext, w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Converts AVChromaLocation to swscale x/y chroma position.
//int 	avcodec_enum_to_chroma_pos (int *xpos, int *ypos, enum AVChromaLocation pos)
func Avcodec_enum_to_chroma_pos(x, y *int, l AVChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)), (C.enum_AVChromaLocation)(l)))
}

//Converts swscale x/y chroma position to AVChromaLocation.
//enum AVChromaLocation 	avcodec_chroma_pos_to_enum (int xpos, int ypos)
func Avcodec_chroma_pos_to_enum(x, y int) AVChromaLocation {
	return (AVChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

//Decode the audio frame of size avpkt->size from avpkt->data into frame.
//int 	avcodec_decode_audio4 (AVCodecContext *avctx, AVFrame *frame, int *got_frame_ptr, const AVPacket *avpkt)
func Avcodec_decode_audio4(ctxt *AVCodecContext, f *AVFrame, g *int, a *AVPacket) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode the video frame of size avpkt->size from avpkt->data into picture.
//int 	avcodec_decode_video2 (AVCodecContext *avctx, AVFrame *picture, int *got_picture_ptr, const AVPacket *avpkt)
func Avcodec_decode_video2(ctxt *AVCodecContext, p *AVFrame, g *int, a *AVPacket) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(p), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode a subtitle message.
//int 	avcodec_decode_subtitle2 (AVCodecContext *avctx, AVSubtitle *sub, int *got_sub_ptr, AVPacket *avpkt)
func Avcodec_decode_subtitle2(ctxt *AVCodecContext, s *AVSubtitle, g *int, a *AVPacket) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//AVCodecParser *av_parser_next (const AVCodecParser *c)
func Av_parser_next(p *AVCodecParser) *AVCodecParser {
	return (*AVCodecParser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

//void 	av_register_codec_parser (AVCodecParser *parser)
func Av_register_codec_parser(p *AVCodecParser) {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}

//AVCodecParserContext *av_parser_init (int codec_id)
func Av_parser_init(c int) *AVCodecParserContext {
	return (*AVCodecParserContext)(C.av_parser_init(C.int(c)))
}

//Parse a packet.
//int 	av_parser_parse2 (AVCodecParserContext *s, AVCodecContext *avctx, uint8_t **poutbuf, int *poutbuf_size, const uint8_t *buf, int buf_size, int64_t pts, int64_t dts, int64_t pos)
func Av_parser_parse2(ctxtp *AVCodecParserContext, ctxt *AVCodecContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

//int 	av_parser_change (AVCodecParserContext *s, AVCodecContext *avctx, uint8_t **poutbuf, int *poutbuf_size, const uint8_t *buf, int buf_size, int keyframe)
func Av_parser_change(ctxtp *AVCodecParserContext, ctxt *AVCodecContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

//void 	av_parser_close (AVCodecParserContext *s)
func Av_parser_close(ctxtp *AVCodecParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

//Find a registered encoder with a matching codec ID.
//AVCodec *avcodec_find_encoder (enum AVCodecID id)
func Avcodec_find_encoder(id AVCodecID) *AVCodec {
	return (*AVCodec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

//Find a registered encoder with the specified name.
//AVCodec *avcodec_find_encoder_by_name (const char *name)
func Avcodec_find_encoder_by_name(c string) *AVCodec {
	return (*AVCodec)(C.avcodec_find_encoder_by_name(C.CString(c)))
}

//Encode a frame of audio.
//int 	avcodec_encode_audio2 (AVCodecContext *avctx, AVPacket *avpkt, const AVFrame *frame, int *got_packet_ptr)
func Avcodec_encode_audio2(ctxt *AVCodecContext, p *AVPacket, f *AVFrame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

//int 	avcodec_encode_subtitle (AVCodecContext *avctx, uint8_t *buf, int buf_size, const AVSubtitle *sub)
func Avcodec_encode_subtitle(ctxt *AVCodecContext, b *uint8, bs int, s *AVSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctxt), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

//Free a picture previously allocated by avpicture_alloc().
//Allocate memory for the pixels of a picture and setup the AVPicture fields for it.
func Avpicture_alloc(p *AVPicture, t AVPixelFormat, w, h int) int {
	return int(C.avpicture_alloc((*C.struct_AVPicture)(p), (C.enum_AVPixelFormat)(t), C.int(w), C.int(h)))
}

//void 	avpicture_free (AVPicture *picture)
func Avpicture_free(p *AVPicture) {
	C.avpicture_free((*C.struct_AVPicture)(p))
}

//Setup the picture fields based on the specified image parameters and the provided image data buffer.
//int 	avpicture_fill (AVPicture *picture, const uint8_t *ptr, enum AVPixelFormat pix_fmt, int width, int height)
func Avpicture_fill(p *AVPicture, pt *uint8, pf AVPixelFormat, w, h int) int {
	return int(C.avpicture_fill((*C.struct_AVPicture)(p), (*C.uint8_t)(pt), (C.enum_AVPixelFormat)(pf), C.int(w), C.int(h)))
}

//Copy pixel data from an AVPicture into a buffer.
//int 	avpicture_layout (const AVPicture *src, enum AVPixelFormat pix_fmt, int width, int height, unsigned char *dest, int dest_size)
func Avpicture_layout(s *AVPicture, pf AVPixelFormat, w, h int, d *string, ds int) int {
	return int(C.avpicture_layout((*C.struct_AVPicture)(s), (C.enum_AVPixelFormat)(pf), C.int(w), C.int(h), (*C.uchar)(unsafe.Pointer(d)), C.int(ds)))
}

//Calculate the size in bytes that a picture of the given width and height would occupy if stored in the given picture format.
//int 	avpicture_get_size (enum AVPixelFormat pix_fmt, int width, int height)
func Avpicture_get_size(pf AVPixelFormat, w, h int) int {
	return int(C.avpicture_get_size((C.enum_AVPixelFormat)(pf), C.int(w), C.int(h)))
}

//Copy image src to dst.
//void 	av_picture_copy (AVPicture *dst, const AVPicture *src, enum AVPixelFormat pix_fmt, int width, int height)
func Av_picture_copy(d, s *AVPicture, pf AVPixelFormat, w, h int) {
	C.av_picture_copy((*C.struct_AVPicture)(d), (*C.struct_AVPicture)(s), (C.enum_AVPixelFormat)(pf), C.int(w), C.int(h))
}

//Crop image top and left side.
//int 	av_picture_crop (AVPicture *dst, const AVPicture *src, enum AVPixelFormat pix_fmt, int top_band, int left_band)
func Av_picture_crop(d, s *AVPicture, pf AVPixelFormat, t, l int) int {
	return int(C.av_picture_crop((*C.struct_AVPicture)(d), (*C.struct_AVPicture)(s), (C.enum_AVPixelFormat)(pf), C.int(t), C.int(l)))
}

//Pad image.
//int 	av_picture_pad (AVPicture *dst, const AVPicture *src, int height, int width, enum AVPixelFormat pix_fmt, int padtop, int padbottom, int padleft, int padright, int *color)
func Av_picture_pad(p, s *AVPicture, h, w int, pf AVPixelFormat, t, b, l, r int, c *int) int {
	return int(C.av_picture_pad((*C.struct_AVPicture)(p), (*C.struct_AVPicture)(s), (C.int)(h), (C.int)(w), (C.enum_AVPixelFormat)(pf), (C.int)(t), (C.int)(b), (C.int)(l), (C.int)(r), (*C.int)(unsafe.Pointer(c))))
}

//Utility function to access log2_chroma_w log2_chroma_h from the pixel format AVPixFmtDescriptor.
//void 	avcodec_get_chroma_sub_sample (enum AVPixelFormat pix_fmt, int *h_shift, int *v_shift)
func Avcodec_get_chroma_sub_sample(pf AVPixelFormat, h, v *int) {
	C.avcodec_get_chroma_sub_sample((C.enum_AVPixelFormat)(pf), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(v)))
}

//Return a value representing the fourCC code associated to the pixel format pix_fmt, or 0 if no associated fourCC code can be found.
//unsigned int 	avcodec_pix_fmt_to_codec_tag (enum AVPixelFormat pix_fmt)
func Avcodec_pix_fmt_to_codec_tag(pf AVPixelFormat) uint {
	return uint(C.avcodec_pix_fmt_to_codec_tag((C.enum_AVPixelFormat)(pf)))
}

//int 	avcodec_get_pix_fmt_loss (enum AVPixelFormat dst_pix_fmt, enum AVPixelFormat src_pix_fmt, int has_alpha)
func Avcodec_get_pix_fmt_loss(p, f AVPixelFormat, a int) int {
	return int(C.avcodec_get_pix_fmt_loss((C.enum_AVPixelFormat)(p), (C.enum_AVPixelFormat)(f), C.int(a)))
}

//Find the best pixel format to convert to given a certain source pixel format.
//enum AVPixelFormat 	avcodec_find_best_pix_fmt_of_list (const enum AVPixelFormat *pix_fmt_list, enum AVPixelFormat src_pix_fmt, int has_alpha, int *loss_ptr)
func Avcodec_find_best_pix_fmt_of_list(f *AVPixelFormat, s AVPixelFormat, a int, p *int) AVPixelFormat {
	return (AVPixelFormat)(C.avcodec_find_best_pix_fmt_of_list((*C.enum_AVPixelFormat)(f), (C.enum_AVPixelFormat)(s), C.int(a), (*C.int)(unsafe.Pointer(p))))
}

//enum AVPixelFormat 	avcodec_find_best_pix_fmt_of_2 (enum AVPixelFormat dst_pix_fmt1, enum AVPixelFormat dst_pix_fmt2, enum AVPixelFormat src_pix_fmt, int has_alpha, int *loss_ptr)
func Avcodec_find_best_pix_fmt_of_2(f1, f2, s AVPixelFormat, a int, l *int) AVPixelFormat {
	return (AVPixelFormat)(C.avcodec_find_best_pix_fmt_of_2((C.enum_AVPixelFormat)(f1), (C.enum_AVPixelFormat)(f2), (C.enum_AVPixelFormat)(s), C.int(a), (*C.int)(unsafe.Pointer(l))))
}

//enum AVPixelFormat 	avcodec_default_get_format (struct AVCodecContext *s, const enum AVPixelFormat *fmt)
func Avcodec_default_get_format(ctxt *AVCodecContext, f *AVPixelFormat) AVPixelFormat {
	return (AVPixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctxt), (*C.enum_AVPixelFormat)(f)))
}

//size_t 	av_get_codec_tag_string (char *buf, size_t buf_size, unsigned int codec_tag)
//Put a string representing the codec tag codec_tag in buf.
func Av_get_codec_tag_string(b string, bf uintptr, c uint) uintptr {
	return uintptr(C.av_get_codec_tag_string(C.CString(b), C.size_t(bf), C.uint(c)))
}

//void 	avcodec_string (char *buf, int buf_size, AVCodecContext *enc, int encode)
func Avcodec_string(b string, bs int, ctxt *AVCodecContext, e int) {
	C.avcodec_string(C.CString(b), C.int(bs), (*C.struct_AVCodecContext)(ctxt), C.int(e))
}

//Return a name for the specified profile, if available.
//const char *av_get_profile_name (const AVCodec *codec, int profile)
func Av_get_profile_name(c *AVCodec, p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

// //int 	avcodec_default_execute (AVCodecContext *c, int(*func)(AVCodecContext *c2, void *arg2), void *arg, int *ret, int count, int size)
// func Avcodec_default_execute(ctxt AVCodecContext, int(*func)(AVCodecContext *c2, void *arg2), void *arg, int *ret, int count, s int){
// _,err := C.avcodec_default_execute(AVCodecContext *c, int(*func)(AVCodecContext *c2, void *arg2), void *arg, int *ret, int count, s int)
// }

// //int 	avcodec_default_execute2 (AVCodecContext *c, int(*func)(AVCodecContext *c2, void *arg2, int, int), void *arg, int *ret, int count)
// func Avcodec_default_execute2(AVCodecContext *c, int(*func)(AVCodecContext *c2, void *arg2, int, int), void *arg, int *ret, int count){
// _,err := C.avcodec_default_execute2(AVCodecContext *c, int(*func)(AVCodecContext *c2, void *arg2, int, int), void *arg, int *ret, int count)
// }

//Fill AVFrame audio data and linesize pointers.
//int 	avcodec_fill_audio_frame (AVFrame *frame, int nb_channels, enum AVSampleFormat sample_fmt, const uint8_t *buf, int buf_size, int align)
func Avcodec_fill_audio_frame(f *AVFrame, c int, s AVSampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

//Reset the internal decoder state / flush internal buffers.
//void 	avcodec_flush_buffers (AVCodecContext *avctx)
func Avcodec_flush_buffers(ctxt *AVCodecContext) {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctxt))
}

//Return codec bits per sample.
//int 	av_get_bits_per_sample (enum AVCodecID codec_id)
func Av_get_bits_per_sample(c AVCodecID) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

//Return the PCM codec associated with a sample format.
//enum AVCodecID 	av_get_pcm_codec (enum AVSampleFormat fmt, int be)
func Av_get_pcm_codec(f AVSampleFormat, b int) AVCodecID {
	return (AVCodecID)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(f), C.int(b)))
}

//Return codec bits per sample.
//int 	av_get_exact_bits_per_sample (enum AVCodecID codec_id)
func Av_get_exact_bits_per_sample(c AVCodecID) int {
	return int(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(c)))
}

//Return audio frame duration.
//int 	av_get_audio_frame_duration (AVCodecContext *avctx, int frame_bytes)
func Av_get_audio_frame_duration(ctxt *AVCodecContext, f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctxt), C.int(f)))
}

//Register a bitstream filter.
//void 	av_register_bitstream_filter (AVBitStreamFilter *bsf)
func Av_register_bitstream_filter(b *AVBitStreamFilter) {
	C.av_register_bitstream_filter((*C.struct_AVBitStreamFilter)(b))
}

//Create and initialize a bitstream filter context given a bitstream filter name.
//AVBitStreamFilterContext *av_bitstream_filter_init (const char *name)
func Av_bitstream_filter_init(n string) *AVBitStreamFilterContext {
	return (*AVBitStreamFilterContext)(C.av_bitstream_filter_init(C.CString(n)))
}

//Filter bitstream.
//int av_bitstream_filter_filter (AVBitStreamFilterContext *bsfc, AVCodecContext *avctx, const char *args, uint8_t **poutbuf, int *poutbuf_size, const uint8_t *buf, int buf_size, int keyframe)
func Av_bitstream_filter_filter(bsctxt *AVBitStreamFilterContext, ctxt *AVCodecContext, a string, p **uint8, ps *int, b *uint8, bs, k int) int {
	return int(C.av_bitstream_filter_filter((*C.struct_AVBitStreamFilterContext)(bsctxt), (*C.struct_AVCodecContext)(ctxt), C.CString(a), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

//Release bitstream filter context.
//void 	av_bitstream_filter_close (AVBitStreamFilterContext *bsf)
func Av_bitstream_filter_close(bsctxt *AVBitStreamFilterContext) {
	C.av_bitstream_filter_close((*C.struct_AVBitStreamFilterContext)(bsctxt))
}

//If f is NULL, return the first registered bitstream filter, if f is non-NULL, return the next registered bitstream filter after f, or NULL if f is the last one.
//AVBitStreamFilter *av_bitstream_filter_next (const AVBitStreamFilter *f)
func Av_bitstream_filter_next(f *AVBitStreamFilter) *AVBitStreamFilter {
	return (*AVBitStreamFilter)(C.av_bitstream_filter_next((*C.struct_AVBitStreamFilter)(f)))
}

//Same behaviour av_fast_malloc but the buffer has additional FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
//void av_fast_padded_malloc (void *ptr, unsigned int *size, size_t min_size)
func Av_fast_padded_malloc(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

//Same behaviour av_fast_padded_malloc except that buffer will always be 0-initialized after call.
//void av_fast_padded_mallocz (void *ptr, unsigned int *size, size_t min_size)
func Av_fast_padded_mallocz(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_mallocz(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

//Encode extradata length to a buffer.
//unsigned int 	av_xiphlacing (unsigned char *s, unsigned int v)
func Av_xiphlacing(s *string, v uint) uint {
	return uint(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), (C.uint)(v)))
}

//If hwaccel is NULL, returns the first registered hardware accelerator, if hwaccel is non-NULL,
//returns the next registered hardware accelerator after hwaccel, or NULL if hwaccel is the last one.
//AVHWAccel *av_hwaccel_next (const AVHWAccel *hwaccel)
func Av_hwaccel_next(a *AVHWAccel) *AVHWAccel {
	return (*AVHWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(a)))
}

// //Register a user provided lock manager supporting the operations specified by AVLockOp.
// //int av_lockmgr_register (int(*cb)(void **mutex, enum AVLockOp op))
// func Av_lockmgr_register(int(*cb)(void **mutex, enum AVLockOp op)) int {
// return C.av_lockmgr_register(int(*cb)(void **mutex, enum AVLockOp op))
// }

//Get the type of the given codec.
//enum AVMediaType 	avcodec_get_type (enum AVCodecID codec_id)
func Avcodec_get_type(c AVCodecID) AVMediaType {
	return (AVMediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

//Get the name of a codec.
//const char *avcodec_get_name (enum AVCodecID id)
func Avcodec_get_name(d AVCodecID) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(d)))
}

//int 	avcodec_is_open (AVCodecContext *s)
func Avcodec_is_open(ctxt *AVCodecContext) int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctxt)))
}

//int 	av_codec_is_encoder (const AVCodec *codec)
func Av_codec_is_encoder(c *AVCodec) int {
	return int(C.av_codec_is_encoder((*C.struct_AVCodec)(c)))
}

//int 	av_codec_is_decoder (const AVCodec *codec)
func Av_codec_is_decoder(c *AVCodec) int {
	return int(C.av_codec_is_decoder((*C.struct_AVCodec)(c)))
}

//const AVCodecDescriptor *avcodec_descriptor_get (enum AVCodecID id)
func Avcodec_descriptor_get(id AVCodecID) *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

//Iterate over all codec descriptors known to libavcodec.
//const AVCodecDescriptor *avcodec_descriptor_next (const AVCodecDescriptor *prev)
func Avcodec_descriptor_next(d *AVCodecDescriptor) *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(d)))
}

//const AVCodecDescriptor *avcodec_descriptor_get_by_name (const char *name)
func Avcodec_descriptor_get_by_name(n string) *AVCodecDescriptor {
	return (*AVCodecDescriptor)(C.avcodec_descriptor_get_by_name(C.CString(n)))
}
