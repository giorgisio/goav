// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package avformat

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"

const (
	AVMEDIA_TYPE_UNKNOWN    = C.AVMEDIA_TYPE_UNKNOWN
	AVMEDIA_TYPE_VIDEO      = C.AVMEDIA_TYPE_VIDEO
	AVMEDIA_TYPE_AUDIO      = C.AVMEDIA_TYPE_AUDIO
	AVMEDIA_TYPE_DATA       = C.AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_SUBTITLE   = C.AVMEDIA_TYPE_SUBTITLE
	AVMEDIA_TYPE_ATTACHMENT = C.AVMEDIA_TYPE_ATTACHMENT
	AVMEDIA_TYPE_NB         = C.AVMEDIA_TYPE_NB
)
