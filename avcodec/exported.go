/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
package avcodec

func Codec_id(c *AVCodecContext) AVCodecID {
	return (AVCodecID)(c.codec_id)
}

func Width(c *AVCodecContext) int {
	return int(c.width)
}

func Height(c *AVCodecContext) int {
	return int(c.width)
}

func Pix_fmt(c *AVCodecContext) AVPixelFormat {
	return (AVPixelFormat)(c.pix_fmt)
}

func Stream_index(p *AVPacket) int {
	return int(p.stream_index)
}
