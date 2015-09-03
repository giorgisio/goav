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
	AvBitStreamFilter             C.struct_AVBitStreamFilter
	AvBitStreamFilterContext      C.struct_AVBitStreamFilterContext
	AvClass                       C.struct_AVClass
	AvCodec                       C.struct_AVCodec
	AvCodecContext                C.struct_AVCodecContext
	AvCodecDescriptor             C.struct_AVCodecDescriptor
	AvCodecParser                 C.struct_AVCodecParser
	AvCodecParserContext          C.struct_AVCodecParserContext
	AvDictionary                  C.struct_AVDictionary
	AvFrame                       C.struct_AVFrame
	AvHWAccel                     C.struct_AVHWAccel
	AvPacket                      C.struct_AVPacket
	AvPacketSideData              C.struct_AVPacketSideData
	AvPanScan                     C.struct_AVPanScan
	AvPicture                     C.struct_AVPicture
	AvProfile                     C.struct_AVProfile
	AvRational                    C.struct_AVRational
	AvSubtitle                    C.struct_AVSubtitle
	AvSubtitleRect                C.struct_AVSubtitleRect
	RcOverride                    C.struct_RcOverride
	AvBufferRef                   C.struct_AVBufferRef
	AvAudioServiceType            C.enum_AVAudioServiceType
	AvChromaLocation              C.enum_AVChromaLocation
	AvCodecId                     C.enum_AVCodecID
	AvColorPrimaries              C.enum_AVColorPrimaries
	AvColorRange                  C.enum_AVColorRange
	AvColorSpace                  C.enum_AVColorSpace
	AvColorTransferCharacteristic C.enum_AVColorTransferCharacteristic
	AvDiscard                     C.enum_AVDiscard
	AvFieldOrder                  C.enum_AVFieldOrder
	AvMediaType                   C.enum_AVMediaType
	AvPacketSideDataType          C.enum_AVPacketSideDataType
	AvPixelFormat                 C.enum_AVPixelFormat
	AvSampleFormat                C.enum_AVSampleFormat
)

func (c *AvCodec) AvCodecGetMaxLowres() int {
	return int(C.av_codec_get_max_lowres((*C.struct_AVCodec)(c)))
}

//If c is NULL, returns the first registered codec, if c is non-NULL,
func (c *AvCodec) AvCodecNext() *AvCodec {
	return (*AvCodec)(C.av_codec_next((*C.struct_AVCodec)(c)))
}

//Return the LIBAvCODEC_VERSION_INT constant.
func AvcodecVersion() uint {
	return uint(C.avcodec_version())
}

//Return the libavcodec build-time configuration.
func AvcodecConfiguration() string {
	return C.GoString(C.avcodec_configuration())

}

//Return the libavcodec license.
func AvcodecLicense() string {
	return C.GoString(C.avcodec_license())
}

//Register the codec codec and initialize libavcodec.
func (c *AvCodec) AvcodecRegister() {
	C.avcodec_register((*C.struct_AVCodec)(c))
}

//Register all the codecs, parsers and bitstream filters which were enabled at configuration time.
func AvcodecRegisterAll() {
	C.avcodec_register_all()
}

//Allocate an AvCodecContext and set its fields to default values.
func (c *AvCodec) AvcodecAllocContext3() *AvCodecContext {
	return (*AvCodecContext)(C.avcodec_alloc_context3((*C.struct_AVCodec)(c)))
}

//Get the AvClass for AvCodecContext.
func AvcodecGetClass() *AvClass {
	return (*AvClass)(C.avcodec_get_class())
}

//Get the AvClass for AvFrame.
func AvcodecGetFrameClass() *AvClass {
	return (*AvClass)(C.avcodec_get_frame_class())
}

//Get the AvClass for AvSubtitleRect.
func AvcodecGetSubtitleRectClass() *AvClass {
	return (*AvClass)(C.avcodec_get_subtitle_rect_class())
}

//Free all allocated data in the given subtitle struct.
func AvsubtitleFree(s *AvSubtitle) {
	C.avsubtitle_free((*C.struct_AVSubtitle)(s))
}

//Pack a dictionary for use in side_data.
func AvPacketPackDictionary(d *AvDictionary, s *int) *uint8 {
	return (*uint8)(C.av_packet_pack_dictionary((*C.struct_AVDictionary)(d), (*C.int)(unsafe.Pointer(s))))
}

//Unpack a dictionary from side_data.
func AvPacketUnpackDictionary(d *uint8, s int, dt **AvDictionary) int {
	return int(C.av_packet_unpack_dictionary((*C.uint8_t)(d), C.int(s), (**C.struct_AVDictionary)(unsafe.Pointer(dt))))
}

//Find a registered decoder with a matching codec ID.
func AvcodecFindDecoder(id AvCodecId) *AvCodec {
	return (*AvCodec)(C.avcodec_find_decoder((C.enum_AVCodecID)(id)))
}

//Find a registered decoder with the specified name.
func AvcodecFindDecoderByName(n string) *AvCodec {
	return (*AvCodec)(C.avcodec_find_decoder_by_name(C.CString(n)))
}

//Converts AvChromaLocation to swscale x/y chroma position.
func AvcodecEnumToChromaPos(x, y *int, l AvChromaLocation) int {
	return int(C.avcodec_enum_to_chroma_pos((*C.int)(unsafe.Pointer(x)), (*C.int)(unsafe.Pointer(y)), (C.enum_AVChromaLocation)(l)))
}

//Converts swscale x/y chroma position to AvChromaLocation.
func AvcodecChromaPosToEnum(x, y int) AvChromaLocation {
	return (AvChromaLocation)(C.avcodec_chroma_pos_to_enum(C.int(x), C.int(y)))
}

//Find a registered encoder with a matching codec ID.
func AvcodecFindEncoder(id AvCodecId) *AvCodec {
	return (*AvCodec)(C.avcodec_find_encoder((C.enum_AVCodecID)(id)))
}

//Find a registered encoder with the specified name.
func AvcodecFindEncoderByName(c string) *AvCodec {
	return (*AvCodec)(C.avcodec_find_encoder_by_name(C.CString(c)))
}

//Put a string representing the codec tag codec_tag in buf.
func AvGetCodecTagString(b string, bf uintptr, c uint) uintptr {
	return uintptr(C.av_get_codec_tag_string(C.CString(b), C.size_t(bf), C.uint(c)))
}

func AvcodecString(b string, bs int, ctxt *AvCodecContext, e int) {
	C.avcodec_string(C.CString(b), C.int(bs), (*C.struct_AVCodecContext)(ctxt), C.int(e))
}

//Return a name for the specified profile, if available.
func (c *AvCodec) AvGetProfileName(p int) string {
	return C.GoString(C.av_get_profile_name((*C.struct_AVCodec)(c), C.int(p)))
}

//Fill AvFrame audio data and linesize pointers.
func AvcodecFillAudioFrame(f *AvFrame, c int, s AvSampleFormat, b *uint8, bs, a int) int {
	return int(C.avcodec_fill_audio_frame((*C.struct_AVFrame)(f), C.int(c), (C.enum_AVSampleFormat)(s), (*C.uint8_t)(b), C.int(bs), C.int(a)))
}

//Return codec bits per sample.
func AvGetBitsPerSample(c AvCodecId) int {
	return int(C.av_get_bits_per_sample((C.enum_AVCodecID)(c)))
}

//Return the PCM codec associated with a sample format.
func AvGetPcmCodec(f AvSampleFormat, b int) AvCodecId {
	return (AvCodecId)(C.av_get_pcm_codec((C.enum_AVSampleFormat)(f), C.int(b)))
}

//Return codec bits per sample.
func AvGetExactBitsPerSample(c AvCodecId) int {
	return int(C.av_get_exact_bits_per_sample((C.enum_AVCodecID)(c)))
}

//Register a bitstream filter.
func (b *AvBitStreamFilter) AvRegisterBitstreamFilter() {
	C.av_register_bitstream_filter((*C.struct_AVBitStreamFilter)(b))
}

//Create and initialize a bitstream filter context given a bitstream filter name.
func AvBitstreamFilterInit(n string) *AvBitStreamFilterContext {
	return (*AvBitStreamFilterContext)(C.av_bitstream_filter_init(C.CString(n)))
}

//Filter bitstream.
func (bfx *AvBitStreamFilterContext) AvBitstreamFilterFilter(ctxt *AvCodecContext, a string, p **uint8, ps *int, b *uint8, bs, k int) int {
	return int(C.av_bitstream_filter_filter((*C.struct_AVBitStreamFilterContext)(bfx), (*C.struct_AVCodecContext)(ctxt), C.CString(a), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

//Release bitstream filter context.
func (bfx *AvBitStreamFilterContext) AvBitstreamFilterClose() {
	C.av_bitstream_filter_close((*C.struct_AVBitStreamFilterContext)(bfx))
}

//AvBitStreamFilter *av_bitstream_filter_next (const AvBitStreamFilter *f)
func (f *AvBitStreamFilter) AvBitstreamFilterNext() *AvBitStreamFilter {
	return (*AvBitStreamFilter)(C.av_bitstream_filter_next((*C.struct_AVBitStreamFilter)(f)))
}

//Same behaviour av_fast_malloc but the buffer has additional FF_INPUT_BUFFER_PADDING_SIZE at the end which will always be 0.
func AvFastPaddedMalloc(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

//Same behaviour av_fast_padded_malloc except that buffer will always be 0-initialized after call.
func AvFastPaddedMallocz(p unsafe.Pointer, s *uint, t uintptr) {
	C.av_fast_padded_mallocz(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(t))
}

//Encode extradata length to a buffer.
func AvXiphlacing(s *string, v uint) uint {
	return uint(C.av_xiphlacing((*C.uchar)(unsafe.Pointer(s)), (C.uint)(v)))
}

//If hwaccel is NULL, returns the first registered hardware accelerator, if hwaccel is non-NULL,
//returns the next registered hardware accelerator after hwaccel, or NULL if hwaccel is the last one.
func (a *AvHWAccel) AvHwaccelNext() *AvHWAccel {
	return (*AvHWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(a)))
}

//Get the type of the given codec.
func AvcodecGetType(c AvCodecId) AvMediaType {
	return (AvMediaType)(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

//Get the name of a codec.
func AvcodecGetName(d AvCodecId) string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(d)))
}

func (c *AvCodec) AvCodecIsEncoder() int {
	return int(C.av_codec_is_encoder((*C.struct_AVCodec)(c)))
}

func (c *AvCodec) AvCodecIsDecoder() int {
	return int(C.av_codec_is_decoder((*C.struct_AVCodec)(c)))
}

//const AvCodecDescriptor *avcodec_descriptor_get (enum AvCodecId id)
func AvcodecDescriptorGet(id AvCodecId) *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.avcodec_descriptor_get((C.enum_AVCodecID)(id)))
}

//Iterate over all codec descriptors known to libavcodec.
func AvcodecDescriptorNext(d *AvCodecDescriptor) *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.avcodec_descriptor_next((*C.struct_AVCodecDescriptor)(d)))
}

func AvcodecDescriptorGetByName(n string) *AvCodecDescriptor {
	return (*AvCodecDescriptor)(C.avcodec_descriptor_get_by_name(C.CString(n)))
}
