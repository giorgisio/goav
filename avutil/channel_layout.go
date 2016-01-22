// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/channel_layout.h>
*/
import "C"


func AvGetDefaultChannelLayout(nb_channels int) int64 {
	return int64(C.av_get_default_channel_layout(C.int(nb_channels)))
}
	