/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
package avcodec

func (p *AVPacket) Buf() *AVBufferRef {
	return (*AVBufferRef)(p.buf)
}
func (p *AVPacket) Duration() int {
	return int(p.duration)
}
func (p *AVPacket) Flags() int {
	return int(p.flags)
}
func (p *AVPacket) Side_data_elems() int {
	return int(p.side_data_elems)
}
func (p *AVPacket) Size() int {
	return int(p.size)
}
func (p *AVPacket) Stream_index() int {
	return int(p.stream_index)
}
func (p *AVPacket) Convergence_duration() int64 {
	return int64(p.convergence_duration)
}
func (p *AVPacket) Dts() int64 {
	return int64(p.dts)
}
func (p *AVPacket) Pos() int64 {
	return int64(p.pos)
}
func (p *AVPacket) Pts() int64 {
	return int64(p.pts)
}
func (p *AVPacket) Data() *uint8 {
	return (*uint8)(p.data)
}
