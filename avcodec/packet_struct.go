// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

var gi int64

func (p *Packet) Buf() *AvBufferRef {
	return (*AvBufferRef)(p.buf)
}
func (p *Packet) Duration() int {
	return int(p.duration)
}
func (p *Packet) Flags() int {
	return int(p.flags)
}
func (p *Packet) SideDataElems() int {
	return int(p.side_data_elems)
}
func (p *Packet) Size() int {
	return int(p.size)
}
func (p *Packet) StreamIndex() int {
	return int(p.stream_index)
}
func (p *Packet) ConvergenceDuration() int64 {
	return int64(p.convergence_duration)
}
func (p *Packet) Dts() int64 {
	return int64(p.dts)
}
func (p *Packet) Pos() int64 {
	return int64(p.pos)
}
func (p *Packet) Pts() int64 {
	return int64(p.pts)
}
func (p *Packet) Data() *uint8 {
	return (*uint8)(p.data)
}
func (p *Packet) SetPts(pts int64, index int) {
	p.pts = C.int64_t(gi * 25)
	gi = gi + 1
	//p.pts = C.int64_t(pts)
	p.stream_index = C.int(index)
}
