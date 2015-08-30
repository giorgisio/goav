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

func (avs *AVStream) Codec() *AVCodecContext {
	return (*AVCodecContext)(unsafe.Pointer(avs.codec))
}

func (avs *AVStream) Metadata() *AVDictionary {
	return (*AVDictionary)(unsafe.Pointer(avs.metadata))
}

func (avs *AVStream) Index_entries() *AVIndexEntry {
	return (*AVIndexEntry)(unsafe.Pointer(avs.index_entries))
}

func (avs *AVStream) Attached_pic() AVPacket {
	return AVPacket(avs.attached_pic)
}

func (avs *AVStream) Side_data() *AVPacketSideData {
	return (*AVPacketSideData)(unsafe.Pointer(avs.side_data))
}

func (avs *AVStream) Probe_data() AVProbeData {
	return AVProbeData(avs.probe_data)
}

func (avs *AVStream) Avg_frame_rate() AVRational {
	return AVRational(avs.avg_frame_rate)
}

// func (avs *AVStream) Display_aspect_ratio() *AVRational {
// 	return (*AVRational)(unsafe.Pointer(avs.display_aspect_ratio))
// }

func (avs *AVStream) R_frame_rate() AVRational {
	return AVRational(avs.r_frame_rate)
}

func (avs *AVStream) Sample_aspect_ratio() AVRational {
	return AVRational(avs.sample_aspect_ratio)
}

func (avs *AVStream) Time_base() AVRational {
	return AVRational(avs.time_base)
}

// func (avs *AVStream) Recommended_encoder_configuration() string {
// 	return C.GoString(avs.recommended_encoder_configuration)
// }

func (avs *AVStream) Discard() AVDiscard {
	return AVDiscard(avs.discard)
}

func (avs *AVStream) Need_parsing() AVStreamParseType {
	return AVStreamParseType(avs.need_parsing)
}

func (avs *AVStream) Codec_info_nb_frames() int {
	return int(avs.codec_info_nb_frames)
}

func (avs *AVStream) Disposition() int {
	return int(avs.disposition)
}

func (avs *AVStream) Event_flags() int {
	return int(avs.event_flags)
}

func (avs *AVStream) Id() int {
	return int(avs.id)
}

func (avs *AVStream) Index() int {
	return int(avs.index)
}

func (avs *AVStream) Inject_global_side_data() int {
	return int(avs.inject_global_side_data)
}

func (avs *AVStream) Last_ip_duration() int {
	return int(avs.last_IP_duration)
}

func (avs *AVStream) Nb_decoded_frames() int {
	return int(avs.nb_decoded_frames)
}

func (avs *AVStream) Nb_index_entries() int {
	return int(avs.nb_index_entries)
}

func (avs *AVStream) Nb_side_data() int {
	return int(avs.nb_side_data)
}

func (avs *AVStream) Probe_packets() int {
	return int(avs.probe_packets)
}

func (avs *AVStream) Pts_wrap_behavior() int {
	return int(avs.pts_wrap_behavior)
}

func (avs *AVStream) Request_probe() int {
	return int(avs.request_probe)
}

func (avs *AVStream) Skip_samples() int {
	return int(avs.skip_samples)
}

func (avs *AVStream) Skip_to_keyframe() int {
	return int(avs.skip_to_keyframe)
}

func (avs *AVStream) Stream_identifier() int {
	return int(avs.stream_identifier)
}

func (avs *AVStream) Update_initial_durations_done() int {
	return int(avs.update_initial_durations_done)
}

func (avs *AVStream) Cur_dts() int64 {
	return int64(avs.cur_dts)
}

func (avs *AVStream) Duration() int64 {
	return int64(avs.duration)
}

// func (avs *AVStream) First_discard_sample() int64 {
// 	return int64(avs.first_discard_sample)
// }

func (avs *AVStream) First_dts() int64 {
	return int64(avs.first_dts)
}

func (avs *AVStream) Interleaver_chunk_duration() int64 {
	return int64(avs.interleaver_chunk_duration)
}

func (avs *AVStream) Interleaver_chunk_size() int64 {
	return int64(avs.interleaver_chunk_size)
}

// func (avs *AVStream) Last_discard_sample() int64 {
// 	return int64(avs.last_discard_sample)
// }

func (avs *AVStream) Last_dts_for_order_check() int64 {
	return int64(avs.last_dts_for_order_check)
}

func (avs *AVStream) Last_ip_pts() int64 {
	return int64(avs.last_IP_pts)
}

func (avs *AVStream) Mux_ts_offset() int64 {
	return int64(avs.mux_ts_offset)
}

func (avs *AVStream) Nb_frames() int64 {
	return int64(avs.nb_frames)
}

func (avs *AVStream) Pts_buffer() int64 {
	return int64(avs.pts_buffer[0])
}

func (avs *AVStream) Pts_reorder_error() int64 {
	return int64(avs.pts_reorder_error[0])
}

func (avs *AVStream) Pts_wrap_reference() int64 {
	return int64(avs.pts_wrap_reference)
}

// func (avs *AVStream) Start_skip_samples() int64 {
// 	return int64(avs.start_skip_samples)
// }

func (avs *AVStream) Start_time() int64 {
	return int64(avs.start_time)
}

func (avs *AVStream) Parser() *AVCodecParserContext {
	return (*AVCodecParserContext)(unsafe.Pointer(avs.parser))
}

func (avs *AVStream) Last_in_packet_buffer() *AVPacketList {
	return (*AVPacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
}

// func (avs *AVStream) Priv_pts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.priv_pts))
// }

func (avs *AVStream) Dts_misordered() uint8 {
	return uint8(avs.dts_misordered)
}

func (avs *AVStream) Dts_ordered() uint8 {
	return uint8(avs.dts_ordered)
}

func (avs *AVStream) Pts_reorder_error_count() uint8 {
	return uint8(avs.pts_reorder_error_count[0])
}

func (avs *AVStream) Index_entries_allocated_size() uint {
	return uint(avs.index_entries_allocated_size)
}
