/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
package avcodec

func (p *AvPacket) Buf() *AvBufferRef {
	return (*AvBufferRef)(p.buf)
}
func (p *AvPacket) Duration() int {
	return int(p.duration)
}
func (p *AvPacket) Flags() int {
	return int(p.flags)
}
func (p *AvPacket) SideDataElems() int {
	return int(p.side_data_elems)
}
func (p *AvPacket) Size() int {
	return int(p.size)
}
func (p *AvPacket) StreamIndex() int {
	return int(p.stream_index)
}
func (p *AvPacket) ConvergenceDuration() int64 {
	return int64(p.convergence_duration)
}
func (p *AvPacket) Dts() int64 {
	return int64(p.dts)
}
func (p *AvPacket) Pos() int64 {
	return int64(p.pos)
}
func (p *AvPacket) Pts() int64 {
	return int64(p.pts)
}
func (p *AvPacket) Data() *uint8 {
	return (*uint8)(p.data)
}
