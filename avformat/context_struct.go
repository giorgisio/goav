// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/giorgisio/goav/avutil"
)

func (ctxt *Context) Chapters() **AvChapter {
	return (**AvChapter)(unsafe.Pointer(ctxt.chapters))
}

func (ctxt *Context) AudioCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctxt.audio_codec))
}

func (ctxt *Context) SubtitleCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctxt.subtitle_codec))
}

func (ctxt *Context) VideoCodec() *AvCodec {
	return (*AvCodec)(unsafe.Pointer(ctxt.video_codec))
}

func (ctxt *Context) Metadata() *avutil.Dictionary {
	return (*avutil.Dictionary)(unsafe.Pointer(ctxt.metadata))
}

func (ctxt *Context) Internal() *AvFormatInternal {
	return (*AvFormatInternal)(unsafe.Pointer(ctxt.internal))
}

func (ctxt *Context) Pb() *AvIOContext {
	return (*AvIOContext)(unsafe.Pointer(ctxt.pb))
}

func (ctxt *Context) InterruptCallback() AvIOInterruptCB {
	return AvIOInterruptCB(ctxt.interrupt_callback)
}

func (ctxt *Context) Programs() []*AvProgram {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctxt.programs)),
		Len:  int(ctxt.NbPrograms()),
		Cap:  int(ctxt.NbPrograms()),
	}

	return *((*[]*AvProgram)(unsafe.Pointer(&header)))
}

func (ctxt *Context) Streams() []*Stream {
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(ctxt.streams)),
		Len:  int(ctxt.NbStreams()),
		Cap:  int(ctxt.NbStreams()),
	}

	return *((*[]*Stream)(unsafe.Pointer(&header)))
}

func (ctxt *Context) Filename() string {
	return C.GoString((*C.char)(unsafe.Pointer(&ctxt.filename[0])))
}

// func (ctxt *Context) CodecWhitelist() string {
// 	return C.GoString(ctxt.codec_whitelist)
// }

// func (ctxt *Context) FormatWhitelist() string {
// 	return C.GoString(ctxt.format_whitelist)
// }

func (ctxt *Context) AudioCodecId() CodecId {
	return CodecId(ctxt.audio_codec_id)
}

func (ctxt *Context) SubtitleCodecId() CodecId {
	return CodecId(ctxt.subtitle_codec_id)
}

func (ctxt *Context) VideoCodecId() CodecId {
	return CodecId(ctxt.video_codec_id)
}

func (ctxt *Context) DurationEstimationMethod() AvDurationEstimationMethod {
	return AvDurationEstimationMethod(ctxt.duration_estimation_method)
}

func (ctxt *Context) AudioPreload() int {
	return int(ctxt.audio_preload)
}

func (ctxt *Context) AvioFlags() int {
	return int(ctxt.avio_flags)
}

func (ctxt *Context) AvoidNegativeTs() int {
	return int(ctxt.avoid_negative_ts)
}

func (ctxt *Context) BitRate() int {
	return int(ctxt.bit_rate)
}

func (ctxt *Context) CtxFlags() int {
	return int(ctxt.ctx_flags)
}

func (ctxt *Context) Debug() int {
	return int(ctxt.debug)
}

func (ctxt *Context) ErrorRecognition() int {
	return int(ctxt.error_recognition)
}

func (ctxt *Context) EventFlags() int {
	return int(ctxt.event_flags)
}

func (ctxt *Context) Flags() int {
	return int(ctxt.flags)
}

func (ctxt *Context) FlushPackets() int {
	return int(ctxt.flush_packets)
}

func (ctxt *Context) FormatProbesize() int {
	return int(ctxt.format_probesize)
}

func (ctxt *Context) FpsProbeSize() int {
	return int(ctxt.fps_probe_size)
}

func (ctxt *Context) IoRepositioned() int {
	return int(ctxt.io_repositioned)
}

func (ctxt *Context) Keylen() int {
	return int(ctxt.keylen)
}

func (ctxt *Context) MaxChunkDuration() int {
	return int(ctxt.max_chunk_duration)
}

func (ctxt *Context) MaxChunkSize() int {
	return int(ctxt.max_chunk_size)
}

func (ctxt *Context) MaxDelay() int {
	return int(ctxt.max_delay)
}

func (ctxt *Context) MaxTsProbe() int {
	return int(ctxt.max_ts_probe)
}

func (ctxt *Context) MetadataHeaderPadding() int {
	return int(ctxt.metadata_header_padding)
}

func (ctxt *Context) ProbeScore() int {
	return int(ctxt.probe_score)
}

func (ctxt *Context) Seek2any() int {
	return int(ctxt.seek2any)
}

func (ctxt *Context) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

func (ctxt *Context) TsId() int {
	return int(ctxt.ts_id)
}

func (ctxt *Context) UseWallclockAsTimestamps() int {
	return int(ctxt.use_wallclock_as_timestamps)
}

func (ctxt *Context) Duration() int64 {
	return int64(ctxt.duration)
}

func (ctxt *Context) MaxAnalyzeDuration2() int64 {
	return int64(ctxt.max_analyze_duration)
}

func (ctxt *Context) MaxInterleaveDelta() int64 {
	return int64(ctxt.max_interleave_delta)
}

func (ctxt *Context) OutputTsOffset() int64 {
	return int64(ctxt.output_ts_offset)
}

func (ctxt *Context) Probesize2() int64 {
	return int64(ctxt.probesize)
}

func (ctxt *Context) SkipInitialBytes() int64 {
	return int64(ctxt.skip_initial_bytes)
}

func (ctxt *Context) StartTime() int64 {
	return int64(ctxt.start_time)
}

func (ctxt *Context) StartTimeRealtime() int64 {
	return int64(ctxt.start_time_realtime)
}

func (ctxt *Context) Iformat() *InputFormat {
	return (*InputFormat)(unsafe.Pointer(ctxt.iformat))
}

func (ctxt *Context) Oformat() *OutputFormat {
	return (*OutputFormat)(unsafe.Pointer(ctxt.oformat))
}

// func (ctxt *Context) DumpSeparator() uint8 {
// 	return uint8(ctxt.dump_separator)
// }

func (ctxt *Context) CorrectTsOverflow() int {
	return int(ctxt.correct_ts_overflow)
}

func (ctxt *Context) MaxIndexSize() uint {
	return uint(ctxt.max_index_size)
}

func (ctxt *Context) MaxPictureBuffer() uint {
	return uint(ctxt.max_picture_buffer)
}

func (ctxt *Context) NbChapters() uint {
	return uint(ctxt.nb_chapters)
}

func (ctxt *Context) NbPrograms() uint {
	return uint(ctxt.nb_programs)
}

func (ctxt *Context) NbStreams() uint {
	return uint(ctxt.nb_streams)
}

func (ctxt *Context) PacketSize() uint {
	return uint(ctxt.packet_size)
}

func (ctxt *Context) Probesize() uint {
	return uint(ctxt.probesize)
}

func (ctxt *Context) SetPb(pb *AvIOContext) {
	ctxt.pb = (*C.struct_AVIOContext)(unsafe.Pointer(pb))
}

func (ctxt *Context) Pb2() **AvIOContext {
	return (**AvIOContext)(unsafe.Pointer(&ctxt.pb))
}
