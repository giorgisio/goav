/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
package avcodec

func (c *AVCodecContext) Codec_id() AVCodecID {
	return (AVCodecID)(c.codec_id)
}

func (c *AVCodecContext) Width() int {
	return int(c.width)
}

func (c *AVCodecContext) Height() int {
	return int(c.width)
}

func (c *AVCodecContext) Pix_fmt() AVPixelFormat {
	return (AVPixelFormat)(c.pix_fmt)
}

func Stream_index(p *AVPacket) int {
	return int(p.stream_index)
}
