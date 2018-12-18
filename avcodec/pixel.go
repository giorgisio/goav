// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"

const (
	AV_PIX_FMT_YUV        = 0
	AV_PIX_FMT_YUV420P9   = int(C.AV_PIX_FMT_YUV420P9)
	AV_PIX_FMT_YUV422P9   = int(C.AV_PIX_FMT_YUV422P9)
	AV_PIX_FMT_YUV444P9   = int(C.AV_PIX_FMT_YUV444P9)
	AV_PIX_FMT_YUV420P10  = int(C.AV_PIX_FMT_YUV420P10)
	AV_PIX_FMT_YUV422P10  = int(C.AV_PIX_FMT_YUV422P10)
	AV_PIX_FMT_YUV440P10  = int(C.AV_PIX_FMT_YUV440P10)
	AV_PIX_FMT_YUV444P10  = int(C.AV_PIX_FMT_YUV444P10)
	AV_PIX_FMT_YUV420P12  = int(C.AV_PIX_FMT_YUV420P12)
	AV_PIX_FMT_YUV422P12  = int(C.AV_PIX_FMT_YUV422P12)
	AV_PIX_FMT_YUV440P12  = int(C.AV_PIX_FMT_YUV440P12)
	AV_PIX_FMT_YUV444P12  = int(C.AV_PIX_FMT_YUV444P12)
	AV_PIX_FMT_YUV420P14  = int(C.AV_PIX_FMT_YUV420P14)
	AV_PIX_FMT_YUV422P14  = int(C.AV_PIX_FMT_YUV422P14)
	AV_PIX_FMT_YUV444P14  = int(C.AV_PIX_FMT_YUV444P14)
	AV_PIX_FMT_YUV420P16  = int(C.AV_PIX_FMT_YUV420P16)
	AV_PIX_FMT_YUV422P16  = int(C.AV_PIX_FMT_YUV422P16)
	AV_PIX_FMT_YUV444P16  = int(C.AV_PIX_FMT_YUV444P16)
	AV_PIX_FMT_YUVA420P9  = int(C.AV_PIX_FMT_YUVA420P9)
	AV_PIX_FMT_YUVA422P9  = int(C.AV_PIX_FMT_YUVA422P9)
	AV_PIX_FMT_YUVA444P9  = int(C.AV_PIX_FMT_YUVA444P9)
	AV_PIX_FMT_YUVA420P10 = int(C.AV_PIX_FMT_YUVA420P10)
	AV_PIX_FMT_YUVA422P10 = int(C.AV_PIX_FMT_YUVA422P10)
	AV_PIX_FMT_YUVA444P10 = int(C.AV_PIX_FMT_YUVA444P10)
	AV_PIX_FMT_YUVA420P16 = int(C.AV_PIX_FMT_YUVA420P16)
	AV_PIX_FMT_YUVA422P16 = int(C.AV_PIX_FMT_YUVA422P16)
	AV_PIX_FMT_YUVA444P16 = int(C.AV_PIX_FMT_YUVA444P16)
	AV_PIX_FMT_RGB24      = int(C.AV_PIX_FMT_RGB24)
	AV_PIX_FMT_RGBA       = int(C.AV_PIX_FMT_RGBA)
)

func (pf PixelFormat) String() string {
	switch int(pf) {
	case AV_PIX_FMT_YUV420P9:
		return "YUV420P9"

	case AV_PIX_FMT_YUV422P9:
		return "YUV422P9"

	case AV_PIX_FMT_YUV444P9:
		return "YUV444P9"

	case AV_PIX_FMT_YUV420P10:
		return "YUV420P10"

	case AV_PIX_FMT_YUV422P10:
		return "YUV422P10"

	case AV_PIX_FMT_YUV440P10:
		return "YUV440P10"

	case AV_PIX_FMT_YUV444P10:
		return "YUV444P10"

	case AV_PIX_FMT_YUV420P12:
		return "YUV420P12"

	case AV_PIX_FMT_YUV422P12:
		return "YUV422P12"

	case AV_PIX_FMT_YUV440P12:
		return "YUV440P12"

	case AV_PIX_FMT_YUV444P12:
		return "YUV444P12"

	case AV_PIX_FMT_YUV420P14:
		return "YUV420P14"

	case AV_PIX_FMT_YUV422P14:
		return "YUV422P14"

	case AV_PIX_FMT_YUV444P14:
		return "YUV444P14"

	case AV_PIX_FMT_YUV420P16:
		return "YUV420P16"

	case AV_PIX_FMT_YUV422P16:
		return "YUV422P16"

	case AV_PIX_FMT_YUV444P16:
		return "YUV444P16"

	case AV_PIX_FMT_YUVA420P9:
		return "YUVA420P9"

	case AV_PIX_FMT_YUVA422P9:
		return "YUVA422P9"

	case AV_PIX_FMT_YUVA444P9:
		return "YUVA444P9"

	case AV_PIX_FMT_YUVA420P10:
		return "YUVA420P10"

	case AV_PIX_FMT_YUVA422P10:
		return "YUVA422P10"

	case AV_PIX_FMT_YUVA444P10:
		return "YUVA444P10"

	case AV_PIX_FMT_YUVA420P16:
		return "YUVA420P16"

	case AV_PIX_FMT_YUVA422P16:
		return "YUVA422P16"

	case AV_PIX_FMT_YUVA444P16:
		return "YUVA444P16"

	case AV_PIX_FMT_RGB24:
		return "RGB24"

	case AV_PIX_FMT_RGBA:
		return "RGBA"
	}

	return "{UNKNOWN}"
}
