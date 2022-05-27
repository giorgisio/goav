// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

//Register a bitstream filter.
func (bsf *BitStreamFilter) AvRegisterBitstreamFilter() {
	C.av_register_bitstream_filter((*C.struct_AVBitStreamFilter)(bsf))
}

//BitStreamFilter *av_bitstream_filter_next (const BitStreamFilter *f)
func (bsf *BitStreamFilter) AvBitstreamFilterNext() *BitStreamFilter {
	return (*BitStreamFilter)(C.av_bitstream_filter_next((*C.struct_AVBitStreamFilter)(bsf)))
}

//Filter bitstream.
func (bfx *BitStreamFilterContext) AvBitstreamFilterFilter(ctxt *Context, a string, p **uint8, ps *int, b *uint8, bs, k int) int {
	return int(C.av_bitstream_filter_filter((*C.struct_AVBitStreamFilterContext)(bfx), (*C.struct_AVCodecContext)(ctxt), C.CString(a), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

//Release bitstream filter context.
func (bfx *BitStreamFilterContext) AvBitstreamFilterClose() {
	C.av_bitstream_filter_close((*C.struct_AVBitStreamFilterContext)(bfx))
}

//Create and initialize a bitstream filter context given a bitstream filter name.
func AvBitstreamFilterInit(n string) *BitStreamFilterContext {
	return (*BitStreamFilterContext)(C.av_bitstream_filter_init(C.CString(n)))
}

func AvBsfGetByName(name string) *BitStreamFilter{
	return (*BitStreamFilter)(C.av_bsf_get_by_name(C.CString(name)))
}

func (bsf *BitStreamFilter) AvBsfAlloc(bsfContext **BSFContext) int{
	return int(C.av_bsf_alloc((*C.struct_AVBitStreamFilter)(bsf),(**C.struct_AVBSFContext)(unsafe.Pointer(bsfContext))))
}

func (bfx *BSFContext) AvBsfInit() int {
	return int(C.av_bsf_init((*C.struct_AVBSFContext)(bfx)))
}

func (bfx *BSFContext) CodecParameters() *AvCodecParameters {
	return (*AvCodecParameters)(unsafe.Pointer(bfx.par_in))
}

func (bfx *BSFContext) AvBfsSendPacket(packet *Packet) int {
	return (int)(C.av_bsf_send_packet((*C.struct_AVBSFContext)(bfx), (*C.struct_AVPacket)(packet)))
}

func (bfx *BSFContext) AvBfsReceivePacket(packet *Packet) int {
	return (int)(C.av_bsf_receive_packet((*C.struct_AVBSFContext)(bfx), (*C.struct_AVPacket)(packet)))
}