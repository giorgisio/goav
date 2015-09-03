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

// cctx := (*C.struct_AVFormatContext)(ctx)
// return (*AvStream)(unsafe.Pointer(cctx.streams))

func (ctxt *AvFormatContext) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(ctxt.chapters))
}

func (ctxt *AvFormatContext) Audio_codec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctxt.audio_codec))
}

func (ctxt *AvFormatContext) Subtitle_codec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctxt.subtitle_codec))
}

func (ctxt *AvFormatContext) Video_codec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctxt.video_codec))
}

func (ctxt *AvFormatContext) Metadata() *AvDictionary {
	return (*AvDictionary)(unsafe.Pointer(ctxt.metadata))
}

func (ctxt *AvFormatContext) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(ctxt.internal))
}

func (ctxt *AvFormatContext) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(ctxt.pb))
}

func (ctxt *AvFormatContext) Interrupt_callback() AvIOInterruptCB {
	return AvIOInterruptCB(ctxt.interrupt_callback)
}

func (ctxt *AvFormatContext) Programs() **AvProgram {
	return (**AvProgram)(unsafe.Pointer(ctxt.programs))
}

func (ctxt *AvFormatContext) Offset_timebase() AvRational {
	return AvRational(ctxt.offset_timebase)
}

func (ctxt *AvFormatContext) Streams() *AvStream {
	return (*AvStream)(unsafe.Pointer(ctxt.streams))
}

func (ctxt *AvFormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&ctxt.filename[0])))
}

// func (ctxt *AvFormatContext) Codec_whitelist() string {
// 	return C.GoString(ctxt.codec_whitelist)
// }

// func (ctxt *AvFormatContext) Format_whitelist() string {
// 	return C.GoString(ctxt.format_whitelist)
// }

func (ctxt *AvFormatContext) Audio_codec_id() AvCodecId {
	return AvCodecId(ctxt.audio_codec_id)
}

func (ctxt *AvFormatContext) Subtitle_codec_id() AvCodecId {
	return AvCodecId(ctxt.subtitle_codec_id)
}

func (ctxt *AvFormatContext) Video_codec_id() AvCodecId {
	return AvCodecId(ctxt.video_codec_id)
}

func (ctxt *AvFormatContext) Duration_estimation_method() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(ctxt.duration_estimation_method)
}

func (ctxt *AvFormatContext) Audio_preload() int {
	return int(ctxt.audio_preload)
}

func (ctxt *AvFormatContext) Avio_flags() int {
	return int(ctxt.avio_flags)
}

func (ctxt *AvFormatContext) Avoid_negative_ts() int {
	return int(ctxt.avoid_negative_ts)
}

func (ctxt *AvFormatContext) Bit_rate() int {
	return int(ctxt.bit_rate)
}

func (ctxt *AvFormatContext) Ctx_flags() int {
	return int(ctxt.ctx_flags)
}

func (ctxt *AvFormatContext) Debug() int {
	return int(ctxt.debug)
}

func (ctxt *AvFormatContext) Error_recognition() int {
	return int(ctxt.error_recognition)
}

func (ctxt *AvFormatContext) Event_flags() int {
	return int(ctxt.event_flags)
}

func (ctxt *AvFormatContext) Flags() int {
	return int(ctxt.flags)
}

func (ctxt *AvFormatContext) Flush_packets() int {
	return int(ctxt.flush_packets)
}

func (ctxt *AvFormatContext) Format_probesize() int {
	return int(ctxt.format_probesize)
}

func (ctxt *AvFormatContext) Fps_probe_size() int {
	return int(ctxt.fps_probe_size)
}

func (ctxt *AvFormatContext) Io_repositioned() int {
	return int(ctxt.io_repositioned)
}

func (ctxt *AvFormatContext) Keylen() int {
	return int(ctxt.keylen)
}

func (ctxt *AvFormatContext) Max_chunk_duration() int {
	return int(ctxt.max_chunk_duration)
}

func (ctxt *AvFormatContext) Max_chunk_size() int {
	return int(ctxt.max_chunk_size)
}

func (ctxt *AvFormatContext) Max_delay() int {
	return int(ctxt.max_delay)
}

func (ctxt *AvFormatContext) Max_ts_probe() int {
	return int(ctxt.max_ts_probe)
}

func (ctxt *AvFormatContext) Metadata_header_padding() int {
	return int(ctxt.metadata_header_padding)
}

func (ctxt *AvFormatContext) Probe_score() int {
	return int(ctxt.probe_score)
}

func (ctxt *AvFormatContext) Raw_packet_buffer_remaining_size() int {
	return int(ctxt.raw_packet_buffer_remaining_size)
}

func (ctxt *AvFormatContext) Seek2any() int {
	return int(ctxt.seek2any)
}

func (ctxt *AvFormatContext) Strict_std_compliance() int {
	return int(ctxt.strict_std_compliance)
}

func (ctxt *AvFormatContext) Ts_id() int {
	return int(ctxt.ts_id)
}

func (ctxt *AvFormatContext) Use_wallclock_as_timestamps() int {
	return int(ctxt.use_wallclock_as_timestamps)
}

func (ctxt *AvFormatContext) Data_offset() int64 {
	return int64(ctxt.data_offset)
}

func (ctxt *AvFormatContext) Duration() int64 {
	return int64(ctxt.duration)
}

func (ctxt *AvFormatContext) Max_analyze_duration2() int64 {
	return int64(ctxt.max_analyze_duration2)
}

func (ctxt *AvFormatContext) Max_interleave_delta() int64 {
	return int64(ctxt.max_interleave_delta)
}

func (ctxt *AvFormatContext) Offset() int64 {
	return int64(ctxt.offset)
}

func (ctxt *AvFormatContext) Output_ts_offset() int64 {
	return int64(ctxt.output_ts_offset)
}

func (ctxt *AvFormatContext) Probesize2() int64 {
	return int64(ctxt.probesize2)
}

func (ctxt *AvFormatContext) Skip_initial_bytes() int64 {
	return int64(ctxt.skip_initial_bytes)
}

func (ctxt *AvFormatContext) Start_time() int64 {
	return int64(ctxt.start_time)
}

func (ctxt *AvFormatContext) Start_time_realtime() int64 {
	return int64(ctxt.start_time_realtime)
}

func (ctxt *AvFormatContext) Iformat() *AvInputFormat {
	return (*AvInputFormat)(unsafe.Pointer(ctxt.iformat))
}

func (ctxt *AvFormatContext) Oformat() *AvOutputFormat {
	return (*AvOutputFormat)(unsafe.Pointer(ctxt.oformat))
}

func (ctxt *AvFormatContext) Packet_buffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(ctxt.packet_buffer))
}

func (ctxt *AvFormatContext) Packet_buffer_end() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(ctxt.packet_buffer_end))
}

func (ctxt *AvFormatContext) Parse_queue() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(ctxt.parse_queue))
}

func (ctxt *AvFormatContext) Parse_queue_end() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(ctxt.parse_queue_end))
}

func (ctxt *AvFormatContext) Raw_packet_buffer() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(ctxt.raw_packet_buffer))
}

func (ctxt *AvFormatContext) Raw_packet_buffer_end() *AvPacketList {
	return (*AvPacketList)(unsafe.Pointer(ctxt.raw_packet_buffer_end))
}

// func (ctxt *AvFormatContext) Dump_separator() uint8 {
// 	return uint8(ctxt.dump_separator)
// }

func (ctxt *AvFormatContext) Correct_ts_overflow() int {
	return int(ctxt.correct_ts_overflow)
}

func (ctxt *AvFormatContext) Max_index_size() uint {
	return uint(ctxt.max_index_size)
}

func (ctxt *AvFormatContext) Max_picture_buffer() uint {
	return uint(ctxt.max_picture_buffer)
}

func (ctxt *AvFormatContext) Nb_chapters() uint {
	return uint(ctxt.nb_chapters)
}

func (ctxt *AvFormatContext) Nb_programs() uint {
	return uint(ctxt.nb_programs)
}

func (ctxt *AvFormatContext) Nb_streams() uint {
	return uint(ctxt.nb_streams)
}

func (ctxt *AvFormatContext) Packet_size() uint {
	return uint(ctxt.packet_size)
}

func (ctxt *AvFormatContext) Probesize() uint {
	return uint(ctxt.probesize)
}
