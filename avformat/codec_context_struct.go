// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package avformat

func (cctxt *CodecContext) Type() MediaType {
	return MediaType(cctxt.codec_type)
}
