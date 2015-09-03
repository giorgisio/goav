/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"
)

func (avs *AvStream) Codec() *AvCodecContext {
	return (*AvCodecContext)(unsafe.Pointer(avs.codec))
}

func (avs *AvStream) Metadata() *AvDictionary {
	return (*AvDictionary)(unsafe.Pointer(avs.metadata))
}

func (avs *AvStream) Index_entries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(avs.index_entries))
}

func (avs *AvStream) Attached_pic() AvPacket {
	return AvPacket(avs.attached_pic)
}

func (avs *AvStream) Side_data() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(avs.side_data))
}

func (avs *AvStream) Probe_data() AvProbeData {
	return AvProbeData(avs.probe_data)
}

func (avs *AvStream) Avg_frame_rate() AvRational {
	return AvRational(avs.avg_frame_rate)
}

// func (avs *AvStream) Display_aspect_ratio() *AvRational {
// 	return (*AvRational)(unsafe.Pointer(avs.display_aspect_ratio))
// }

func (avs *AvStream) R_frame_rate() AvRational {
	return AvRational(avs.r_frame_rate)
}

func (avs *AvStream) Sample_aspect_ratio() AvRational {
	return AvRational(avs.sample_aspect_ratio)
}

func (avs *AvStream) Time_base() AvRational {
	return AvRational(avs.time_base)
}

// func (avs *AvStream) Recommended_encoder_configuration() string {
// 	return C.GoString(avs.recommended_encoder_configuration)
// }

func (avs *AvStream) Discard() AvDiscard {
	return AvDiscard(avs.discard)
}

func (avs *AvStream) Need_parsing() AvStreamParseType {
	return AvStreamParseType(avs.need_parsing)
}

func (avs *AvStream) Codec_info_nb_frames() int {
	return int(avs.codec_info_nb_frames)
}

func (avs *AvStream) Disposition() int {
	return int(avs.disposition)
}

func (avs *AvStream) Event_flags() int {
	return int(avs.event_flags)
}

func (avs *AvStream) Id() int {
	return int(avs.id)
}

func (avs *AvStream) Index() int {
	return int(avs.index)
}

func (avs *AvStream) Inject_global_side_data() int {
	return int(avs.inject_global_side_data)
}

func (avs *AvStream) Last_ip_duration() int {
	return int(avs.last_IP_duration)
}

func (avs *AvStream) Nb_decoded_frames() int {
	return int(avs.nb_decoded_frames)
}

func (avs *AvStream) Nb_index_entries() int {
	return int(avs.nb_index_entries)
}

func (avs *AvStream) Nb_side_data() int {
	return int(avs.nb_side_data)
}

func (avs *AvStream) Probe_packets() int {
	return int(avs.probe_packets)
}

func (avs *AvStream) Pts_wrap_behavior() int {
	return int(avs.pts_wrap_behavior)
}

func (avs *AvStream) Request_probe() int {
	return int(avs.request_probe)
}

func (avs *AvStream) Skip_samples() int {
	return int(avs.skip_samples)
}

func (avs *AvStream) Skip_to_keyframe() int {
	return int(avs.skip_to_keyframe)
}

func (avs *AvStream) Stream_identifier() int {
	return int(avs.stream_identifier)
}

func (avs *AvStream) Update_initial_durations_done() int {
	return int(avs.update_initial_durations_done)
}

func (avs *AvStream) Cur_dts() int64 {
	return int64(avs.cur_dts)
}

func (avs *AvStream) Duration() int64 {
	return int64(avs.duration)
}

// func (avs *AvStream) First_discard_sample() int64 {
// 	return int64(avs.first_discard_sample)
// }

func (avs *AvStream) First_dts() int64 {
	return int64(avs.first_dts)
}

func (avs *AvStream) Interleaver_chunk_duration() int64 {
	return int64(avs.interleaver_chunk_duration)
}

func (avs *AvStream) Interleaver_chunk_size() int64 {
	return int64(avs.interleaver_chunk_size)
}

// func (avs *AvStream) Last_discard_sample() int64 {
// 	return int64(avs.last_discard_sample)
// }

func (avs *AvStream) Last_dts_for_order_check() int64 {
	return int64(avs.last_dts_for_order_check)
}

func (avs *AvStream) Last_ip_pts() int64 {
	return int64(avs.last_IP_pts)
}

func (avs *AvStream) Mux_ts_offset() int64 {
	return int64(avs.mux_ts_offset)
}

func (avs *AvStream) Nb_frames() int64 {
	return int64(avs.nb_frames)
}

func (avs *AvStream) Pts_buffer() int64 {
	return int64(avs.pts_buffer[0])
}

func (avs *AvStream) Pts_reorder_error() int64 {
	return int64(avs.pts_reorder_error[0])
}

func (avs *AvStream) Pts_wrap_reference() int64 {
	return int64(avs.pts_wrap_reference)
}

// func (avs *AvStream) Start_skip_samples() int64 {
// 	return int64(avs.start_skip_samples)
// }

func (avs *AvStream) Start_time() int64 {
	return int64(avs.start_time)
}

func (avs *AvStream) Parser() *AvCodecParserContext {
	return (*AvCodecParserContext)(unsafe.Pointer(avs.parser))
}

func (avs *AvStream) Last_in_packet_buffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
}

// func (avs *AvStream) Priv_pts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.priv_pts))
// }

func (avs *AvStream) Dts_misordered() uint8 {
	return uint8(avs.dts_misordered)
}

func (avs *AvStream) Dts_ordered() uint8 {
	return uint8(avs.dts_ordered)
}

func (avs *AvStream) Pts_reorder_error_count() uint8 {
	return uint8(avs.pts_reorder_error_count[0])
}

func (avs *AvStream) Index_entries_allocated_size() uint {
	return uint(avs.index_entries_allocated_size)
}
