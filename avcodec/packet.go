/*
	AvPackets
*/
package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

//Initialize optional fields of a packet with default values.
func (p *AvPacket) AvInitPacket() {
	C.av_init_packet((*C.struct_AVPacket)(p))
}

//Allocate the payload of a packet and initialize its fields with default values.
func (p *AvPacket) AvNewPacket(s int) int {
	return int(C.av_new_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//Reduce packet size, correctly zeroing padding.
func (p *AvPacket) AvShrinkPacket(s int) {
	C.av_shrink_packet((*C.struct_AVPacket)(p), C.int(s))
}

//Increase packet size, correctly zeroing padding.
func (p *AvPacket) AvGrowPacket(s int) int {
	return int(C.av_grow_packet((*C.struct_AVPacket)(p), C.int(s)))
}

//Initialize a reference-counted packet from av_malloc()ed data.
func (p *AvPacket) AvPacketFromData(d *uint8, s int) int {
	return int(C.av_packet_from_data((*C.struct_AVPacket)(p), (*C.uint8_t)(d), C.int(s)))

}

func (p *AvPacket) AvDupPacket() int {
	return int(C.av_dup_packet((*C.struct_AVPacket)(p)))

}

//Copy packet, including contents.
func (p *AvPacket) AvCopyPacket(r *AvPacket) int {
	return int(C.av_copy_packet((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//Copy packet side data.
func (p *AvPacket) AvCopyPacketSideData(r *AvPacket) int {
	return int(C.av_copy_packet_side_data((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))

}

//Free a packet.
func (p *AvPacket) AvFreePacket() {
	C.av_free_packet((*C.struct_AVPacket)(p))

}

//Allocate new information of a packet.
func (p *AvPacket) AvPacketNewSideData(t AvPacketSideDataType, s int) *uint8 {
	return (*uint8)(C.av_packet_new_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//Shrink the already allocated side data buffer.
func (p *AvPacket) AvPacketShrinkSideData(t AvPacketSideDataType, s int) int {
	return int(C.av_packet_shrink_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), C.int(s)))
}

//Get side information from packet.
func (p *AvPacket) AvPacketGetSideData(t AvPacketSideDataType, s *int) *uint8 {
	return (*uint8)(C.av_packet_get_side_data((*C.struct_AVPacket)(p), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(s))))
}

//int 	av_packet_merge_side_data (AvPacket *pkt)
func (p *AvPacket) AvPacketMergeSideData() int {
	return int(C.av_packet_merge_side_data((*C.struct_AVPacket)(p)))
}

//int 	av_packet_split_side_data (AvPacket *pkt)
func (p *AvPacket) AvPacketSplitSideData() int {
	return int(C.av_packet_split_side_data((*C.struct_AVPacket)(p)))
}

//Convenience function to free all the side data stored.
func (p *AvPacket) AvPacketFreeSideData() {
	C.av_packet_free_side_data((*C.struct_AVPacket)(p))
}

//Setup a new reference to the data described by a given packet.
func (p *AvPacket) AvPacketRef(s *AvPacket) int {
	return int(C.av_packet_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//Wipe the packet.
func (p *AvPacket) AvPacketUnref() {
	C.av_packet_unref((*C.struct_AVPacket)(p))
}

//Move every field in src to dst and reset src.
func (p *AvPacket) AvPacketMoveRef(s *AvPacket) {
	C.av_packet_move_ref((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s))
}

//Copy only "properties" fields from src to dst.
func (p *AvPacket) AvPacketCopyProps(s *AvPacket) int {
	return int(C.av_packet_copy_props((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(s)))
}

//Convert valid timing fields (timestamps / durations) in a packet from one timebase to another.
func (p *AvPacket) AvPacketRescaleTs(r, r2 AvRational) {
	C.av_packet_rescale_ts((*C.struct_AVPacket)(p), (C.struct_AVRational)(r), (C.struct_AVRational)(r2))
}
