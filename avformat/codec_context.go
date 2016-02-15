// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

func (s *CodecContext) Codec() *Codec {
	codec := C.avcodec_find_decoder(s.codec_id);
	if (codec == nil) {
		panic("Codec not found")
	}

	C.avcodec_open2((*C.struct_AVCodecContext)(s), codec, nil)
	return (*Codec)(s.codec)
}
