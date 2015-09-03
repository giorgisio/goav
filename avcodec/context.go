/*
	AvCodecContext
*/
package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

func (ctxt *AvCodecContext) AvCodecGetPktTimebase() AvRational {
	return (AvRational)(C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *AvCodecContext) AvCodecSetPktTimebase(r AvRational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(ctxt), (C.struct_AVRational)(r))
}

func (ctxt *AvCodecContext) AvCodecGetCodecDescriptor() *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *AvCodecContext) AvCodecSetCodecDescriptor(d *AvCodecDescriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecDescriptor)(d))
}

func (ctxt *AvCodecContext) AvCodecGetLowres() int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *AvCodecContext) AvCodecSetLowres(i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

func (ctxt *AvCodecContext) AvCodecGetSeekPreroll() int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *AvCodecContext) AvCodecSetSeekPreroll(i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(ctxt), C.int(i))
}

func (ctxt *AvCodecContext) AvCodecGetChromaIntraMatrix() *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt)))
}

func (ctxt *AvCodecContext) AvCodecSetChromaIntraMatrix(t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(ctxt), (*C.uint16_t)(t))
}

//Free the codec context and everything associated with it and write NULL to the provided pointer.
func (ctxt *AvCodecContext) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(ctxt)))
}

//Set the fields of the given AvCodecContext to default values corresponding to the given codec (defaults may be codec-dependent).
func (ctxt *AvCodecContext) AvcodecGetContextDefaults3(c *AvCodec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c)))
}

//Copy the settings of the source AvCodecContext into the destination AvCodecContext.
func (ctxt *AvCodecContext) AvcodecCopyContext(ctxt2 *AvCodecContext) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodecContext)(ctxt2)))
}

//Initialize the AvCodecContext to use the given AvCodec
func (ctxt *AvCodecContext) AvcodecOpen2(c *AvCodec, d **AvDictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Close a given AvCodecContext and free all the data associated with it (but not the AvCodecContext itself).
func (ctxt *AvCodecContext) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctxt)))
}

//The default callback for AvCodecContext.get_buffer2().
func (s *AvCodecContext) AvcodecDefaultGetBuffer2(f *AvFrame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(s), (*C.struct_AVFrame)(f), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctxt *AvCodecContext) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctxt *AvCodecContext) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctxt), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Decode the audio frame of size avpkt->size from avpkt->data into frame.
func (ctxt *AvCodecContext) AvcodecDecodeAudio4(f *AvFrame, g *int, a *AvPacket) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode the video frame of size avpkt->size from avpkt->data into picture.
func (ctxt *AvCodecContext) AvcodecDecodeVideo2(p *AvFrame, g *int, a *AvPacket) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVFrame)(p), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Decode a subtitle message.
func (ctxt *AvCodecContext) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *AvPacket) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

//Encode a frame of audio.
func (ctxt *AvCodecContext) AvcodecEncodeAudio2(p *AvPacket, f *AvFrame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

//Encode a frame of video
func (ctxt *AvCodecContext) AvcodecEncodeVideo2(p *AvPacket, f *AvFrame, gp *int) int {
	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(ctxt), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

func (ctxt *AvCodecContext) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctxt), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (ctxt *AvCodecContext) AvcodecDefaultGetFormat(f *AvPixelFormat) AvPixelFormat {
	return (AvPixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctxt), (*C.enum_AVPixelFormat)(f)))
}

//Reset the internal decoder state / flush internal buffers.
func (ctxt *AvCodecContext) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctxt))
}

//Return audio frame duration.
func (ctxt *AvCodecContext) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctxt), C.int(f)))
}

func (ctxt *AvCodecContext) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctxt)))
}

//Parse a packet.
func (ctxt *AvCodecContext) AvParserParse2(ctxtp *AvCodecParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func (ctxt *AvCodecContext) AvParserChange(ctxtp *AvCodecParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctxt), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

func AvParserInit(c int) *AvCodecParserContext {
	return (*AvCodecParserContext)(C.av_parser_init(C.int(c)))
}

func AvParserClose(ctxtp *AvCodecParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

func (p *AvCodecParser) AvParserNext() *AvCodecParser {
	return (*AvCodecParser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

func (p *AvCodecParser) AvRegisterCodecParser() {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}
