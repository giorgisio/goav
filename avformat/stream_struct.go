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

func (avs *AvStream) IndexEntries() *AvIndexEntry {
	return (*AvIndexEntry)(unsafe.Pointer(avs.index_entries))
}

func (avs *AvStream) AttachedPic() AvPacket {
	return AvPacket(avs.attached_pic)
}

func (avs *AvStream) SideData() *AvPacketSideData {
	return (*AvPacketSideData)(unsafe.Pointer(avs.side_data))
}

func (avs *AvStream) ProbeData() AvProbeData {
	return AvProbeData(avs.probe_data)
}

func (avs *AvStream) AvgFrameRate() AvRational {
	return AvRational(avs.avg_frame_rate)
}

// func (avs *AvStream) DisplayAspectRatio() *AvRational {
// 	return (*AvRational)(unsafe.Pointer(avs.display_aspect_ratio))
// }

func (avs *AvStream) RFrameRate() AvRational {
	return AvRational(avs.r_frame_rate)
}

func (avs *AvStream) SampleAspectRatio() AvRational {
	return AvRational(avs.sample_aspect_ratio)
}

func (avs *AvStream) TimeBase() AvRational {
	return AvRational(avs.time_base)
}

// func (avs *AvStream) RecommendedEncoderConfiguration() string {
// 	return C.GoString(avs.recommended_encoder_configuration)
// }

func (avs *AvStream) Discard() AvDiscard {
	return AvDiscard(avs.discard)
}

func (avs *AvStream) NeedParsing() AvStreamParseType {
	return AvStreamParseType(avs.need_parsing)
}

func (avs *AvStream) CodecInfoNbFrames() int {
	return int(avs.codec_info_nb_frames)
}

func (avs *AvStream) Disposition() int {
	return int(avs.disposition)
}

func (avs *AvStream) EventFlags() int {
	return int(avs.event_flags)
}

func (avs *AvStream) Id() int {
	return int(avs.id)
}

func (avs *AvStream) Index() int {
	return int(avs.index)
}

func (avs *AvStream) InjectGlobalSideData() int {
	return int(avs.inject_global_side_data)
}

func (avs *AvStream) LastIpDuration() int {
	return int(avs.last_IP_duration)
}

func (avs *AvStream) NbDecodedFrames() int {
	return int(avs.nb_decoded_frames)
}

func (avs *AvStream) NbIndexEntries() int {
	return int(avs.nb_index_entries)
}

func (avs *AvStream) NbSideData() int {
	return int(avs.nb_side_data)
}

func (avs *AvStream) ProbePackets() int {
	return int(avs.probe_packets)
}

func (avs *AvStream) PtsWrapBehavior() int {
	return int(avs.pts_wrap_behavior)
}

func (avs *AvStream) RequestProbe() int {
	return int(avs.request_probe)
}

func (avs *AvStream) SkipSamples() int {
	return int(avs.skip_samples)
}

func (avs *AvStream) SkipToKeyframe() int {
	return int(avs.skip_to_keyframe)
}

func (avs *AvStream) StreamIdentifier() int {
	return int(avs.stream_identifier)
}

func (avs *AvStream) UpdateInitialDurationsDone() int {
	return int(avs.update_initial_durations_done)
}

func (avs *AvStream) CurDts() int64 {
	return int64(avs.cur_dts)
}

func (avs *AvStream) Duration() int64 {
	return int64(avs.duration)
}

// func (avs *AvStream) FirstDiscardSample() int64 {
// 	return int64(avs.first_discard_sample)
// }

func (avs *AvStream) FirstDts() int64 {
	return int64(avs.first_dts)
}

func (avs *AvStream) InterleaverChunkDuration() int64 {
	return int64(avs.interleaver_chunk_duration)
}

func (avs *AvStream) InterleaverChunkSize() int64 {
	return int64(avs.interleaver_chunk_size)
}

// func (avs *AvStream) LastDiscardSample() int64 {
// 	return int64(avs.last_discard_sample)
// }

func (avs *AvStream) LastDtsForOrderCheck() int64 {
	return int64(avs.last_dts_for_order_check)
}

func (avs *AvStream) LastIpPts() int64 {
	return int64(avs.last_IP_pts)
}

func (avs *AvStream) MuxTsOffset() int64 {
	return int64(avs.mux_ts_offset)
}

func (avs *AvStream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

func (avs *AvStream) PtsBuffer() int64 {
	return int64(avs.pts_buffer[0])
}

func (avs *AvStream) PtsReorderError() int64 {
	return int64(avs.pts_reorder_error[0])
}

func (avs *AvStream) PtsWrapReference() int64 {
	return int64(avs.pts_wrap_reference)
}

// func (avs *AvStream) StartSkipSamples() int64 {
// 	return int64(avs.start_skip_samples)
// }

func (avs *AvStream) StartTime() int64 {
	return int64(avs.start_time)
}

func (avs *AvStream) Parser() *AvCodecParserContext {
	return (*AvCodecParserContext)(unsafe.Pointer(avs.parser))
}

func (avs *AvStream) LastInPacketBuffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
}

// func (avs *AvStream) PrivPts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.priv_pts))
// }

func (avs *AvStream) DtsMisordered() uint8 {
	return uint8(avs.dts_misordered)
}

func (avs *AvStream) DtsOrdered() uint8 {
	return uint8(avs.dts_ordered)
}

func (avs *AvStream) PtsReorderErrorCount() uint8 {
	return uint8(avs.pts_reorder_error_count[0])
}

func (avs *AvStream) IndexEntriesAllocatedSize() uint {
	return uint(avs.index_entries_allocated_size)
}
