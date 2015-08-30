/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// cctx := (*C.struct_AVFormatContext)(ctx)
// return (*AVStream)(unsafe.Pointer(cctx.streams))

func (ctxt *AVFormatContext) Chapters() **AVChapter {
	return (**AVChapter)(unsafe.Pointer(ctxt.chapters))
}

func (ctxt *AVFormatContext) Audio_codec() *AVCodec {
	return (*AVCodec)(unsafe.Pointer(ctxt.audio_codec))
}

func (ctxt *AVFormatContext) Subtitle_codec() *AVCodec {
	return (*AVCodec)(unsafe.Pointer(ctxt.subtitle_codec))
}

func (ctxt *AVFormatContext) Video_codec() *AVCodec {
	return (*AVCodec)(unsafe.Pointer(ctxt.video_codec))
}

func (ctxt *AVFormatContext) Metadata() *AVDictionary {
	return (*AVDictionary)(unsafe.Pointer(ctxt.metadata))
}

func (ctxt *AVFormatContext) Internal() *AVFormatInternal {
	return (*AVFormatInternal)(unsafe.Pointer(ctxt.internal))
}

func (ctxt *AVFormatContext) Pb() *AVIOContext {
	return (*AVIOContext)(unsafe.Pointer(ctxt.pb))
}

func (ctxt *AVFormatContext) Interrupt_callback() AVIOInterruptCB {
	return AVIOInterruptCB(ctxt.interrupt_callback)
}

func (ctxt *AVFormatContext) Programs() **AVProgram {
	return (**AVProgram)(unsafe.Pointer(ctxt.programs))
}

func (ctxt *AVFormatContext) Offset_timebase() AVRational {
	return AVRational(ctxt.offset_timebase)
}

func (ctxt *AVFormatContext) Streams() *AVStream {
	return (*AVStream)(unsafe.Pointer(ctxt.streams))
}

func (ctxt *AVFormatContext) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&ctxt.filename[0])))
}

// func (ctxt *AVFormatContext) Codec_whitelist() string {
// 	return C.GoString(ctxt.codec_whitelist)
// }

// func (ctxt *AVFormatContext) Format_whitelist() string {
// 	return C.GoString(ctxt.format_whitelist)
// }

func (ctxt *AVFormatContext) Audio_codec_id() AVCodecID {
	return AVCodecID(ctxt.audio_codec_id)
}

func (ctxt *AVFormatContext) Subtitle_codec_id() AVCodecID {
	return AVCodecID(ctxt.subtitle_codec_id)
}

func (ctxt *AVFormatContext) Video_codec_id() AVCodecID {
	return AVCodecID(ctxt.video_codec_id)
}

func (ctxt *AVFormatContext) Duration_estimation_method() AVDurationEstimationMethod {
	return AVDurationEstimationMethod(ctxt.duration_estimation_method)
}

func (ctxt *AVFormatContext) Audio_preload() int {
	return int(ctxt.audio_preload)
}

func (ctxt *AVFormatContext) Avio_flags() int {
	return int(ctxt.avio_flags)
}

func (ctxt *AVFormatContext) Avoid_negative_ts() int {
	return int(ctxt.avoid_negative_ts)
}

func (ctxt *AVFormatContext) Bit_rate() int {
	return int(ctxt.bit_rate)
}

func (ctxt *AVFormatContext) Ctx_flags() int {
	return int(ctxt.ctx_flags)
}

func (ctxt *AVFormatContext) Debug() int {
	return int(ctxt.debug)
}

func (ctxt *AVFormatContext) Error_recognition() int {
	return int(ctxt.error_recognition)
}

func (ctxt *AVFormatContext) Event_flags() int {
	return int(ctxt.event_flags)
}

func (ctxt *AVFormatContext) Flags() int {
	return int(ctxt.flags)
}

func (ctxt *AVFormatContext) Flush_packets() int {
	return int(ctxt.flush_packets)
}

func (ctxt *AVFormatContext) Format_probesize() int {
	return int(ctxt.format_probesize)
}

func (ctxt *AVFormatContext) Fps_probe_size() int {
	return int(ctxt.fps_probe_size)
}

func (ctxt *AVFormatContext) Io_repositioned() int {
	return int(ctxt.io_repositioned)
}

func (ctxt *AVFormatContext) Keylen() int {
	return int(ctxt.keylen)
}

func (ctxt *AVFormatContext) Max_chunk_duration() int {
	return int(ctxt.max_chunk_duration)
}

func (ctxt *AVFormatContext) Max_chunk_size() int {
	return int(ctxt.max_chunk_size)
}

func (ctxt *AVFormatContext) Max_delay() int {
	return int(ctxt.max_delay)
}

func (ctxt *AVFormatContext) Max_ts_probe() int {
	return int(ctxt.max_ts_probe)
}

func (ctxt *AVFormatContext) Metadata_header_padding() int {
	return int(ctxt.metadata_header_padding)
}

func (ctxt *AVFormatContext) Probe_score() int {
	return int(ctxt.probe_score)
}

func (ctxt *AVFormatContext) Raw_packet_buffer_remaining_size() int {
	return int(ctxt.raw_packet_buffer_remaining_size)
}

func (ctxt *AVFormatContext) Seek2any() int {
	return int(ctxt.seek2any)
}

func (ctxt *AVFormatContext) Strict_std_compliance() int {
	return int(ctxt.strict_std_compliance)
}

func (ctxt *AVFormatContext) Ts_id() int {
	return int(ctxt.ts_id)
}

func (ctxt *AVFormatContext) Use_wallclock_as_timestamps() int {
	return int(ctxt.use_wallclock_as_timestamps)
}

func (ctxt *AVFormatContext) Data_offset() int64 {
	return int64(ctxt.data_offset)
}

func (ctxt *AVFormatContext) Duration() int64 {
	return int64(ctxt.duration)
}

func (ctxt *AVFormatContext) Max_analyze_duration2() int64 {
	return int64(ctxt.max_analyze_duration2)
}

func (ctxt *AVFormatContext) Max_interleave_delta() int64 {
	return int64(ctxt.max_interleave_delta)
}

func (ctxt *AVFormatContext) Offset() int64 {
	return int64(ctxt.offset)
}

func (ctxt *AVFormatContext) Output_ts_offset() int64 {
	return int64(ctxt.output_ts_offset)
}

func (ctxt *AVFormatContext) Probesize2() int64 {
	return int64(ctxt.probesize2)
}

func (ctxt *AVFormatContext) Skip_initial_bytes() int64 {
	return int64(ctxt.skip_initial_bytes)
}

func (ctxt *AVFormatContext) Start_time() int64 {
	return int64(ctxt.start_time)
}

func (ctxt *AVFormatContext) Start_time_realtime() int64 {
	return int64(ctxt.start_time_realtime)
}

func (ctxt *AVFormatContext) Iformat() *AVInputFormat {
	return (*AVInputFormat)(unsafe.Pointer(ctxt.iformat))
}

func (ctxt *AVFormatContext) Oformat() *AVOutputFormat {
	return (*AVOutputFormat)(unsafe.Pointer(ctxt.oformat))
}

func (ctxt *AVFormatContext) Packet_buffer() *AVPacketList {
	return (*AVPacketList)(unsafe.Pointer(ctxt.packet_buffer))
}

func (ctxt *AVFormatContext) Packet_buffer_end() *AVPacketList {
	return (*AVPacketList)(unsafe.Pointer(ctxt.packet_buffer_end))
}

func (ctxt *AVFormatContext) Parse_queue() *AVPacketList {
	return (*AVPacketList)(unsafe.Pointer(ctxt.parse_queue))
}

func (ctxt *AVFormatContext) Parse_queue_end() *AVPacketList {
	return (*AVPacketList)(unsafe.Pointer(ctxt.parse_queue_end))
}

func (ctxt *AVFormatContext) Raw_packet_buffer() *AVPacketList {
	return (*AVPacketList)(unsafe.Pointer(ctxt.raw_packet_buffer))
}

func (ctxt *AVFormatContext) Raw_packet_buffer_end() *AVPacketList {
	return (*AVPacketList)(unsafe.Pointer(ctxt.raw_packet_buffer_end))
}

// func (ctxt *AVFormatContext) Dump_separator() uint8 {
// 	return uint8(ctxt.dump_separator)
// }

func (ctxt *AVFormatContext) Correct_ts_overflow() int {
	return int(ctxt.correct_ts_overflow)
}

func (ctxt *AVFormatContext) Max_index_size() uint {
	return uint(ctxt.max_index_size)
}

func (ctxt *AVFormatContext) Max_picture_buffer() uint {
	return uint(ctxt.max_picture_buffer)
}

func (ctxt *AVFormatContext) Nb_chapters() uint {
	return uint(ctxt.nb_chapters)
}

func (ctxt *AVFormatContext) Nb_programs() uint {
	return uint(ctxt.nb_programs)
}

func (ctxt *AVFormatContext) Nb_streams() uint {
	return uint(ctxt.nb_streams)
}

func (ctxt *AVFormatContext) Packet_size() uint {
	return uint(ctxt.packet_size)
}

func (ctxt *AVFormatContext) Probesize() uint {
	return uint(ctxt.probesize)
}

func Codec_type(s *AVStream, n int) *AVMediaType {
	c := s.Codec()
	fmt.Println("Codec Type: Fix Method")
	if c != nil {
		return (*AVMediaType)(unsafe.Pointer(&c.codec_type))
	} else {
		return nil
	}
}
