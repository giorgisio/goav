// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package avformat

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

const (
	AVMEDIA_TYPE_UNKNOWN    = int(C.AVMEDIA_TYPE_UNKNOWN)
	AVMEDIA_TYPE_VIDEO      = int(C.AVMEDIA_TYPE_VIDEO)
	AVMEDIA_TYPE_AUDIO      = int(C.AVMEDIA_TYPE_AUDIO)
	AVMEDIA_TYPE_DATA       = int(C.AVMEDIA_TYPE_DATA)
	AVMEDIA_TYPE_SUBTITLE   = int(C.AVMEDIA_TYPE_SUBTITLE)
	AVMEDIA_TYPE_ATTACHMENT = int(C.AVMEDIA_TYPE_ATTACHMENT)
	AVMEDIA_TYPE_NB         = int(C.AVMEDIA_TYPE_NB)
)
