// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/log.h>
	#include <stdlib.h>
*/
import "C"

const (
	AV_LOG_QUIET   = -8
	AV_LOG_PANIC   = 0
	AV_LOG_FATAL   = 8
	AV_LOG_ERROR   = 16
	AV_LOG_WARNING = 24
	AV_LOG_INFO    = 32
	AV_LOG_VERBOSE = 40
	AV_LOG_DEBUG   = 48
	AV_LOG_TRACE   = 56
)

func AvLogSetLevel(level int) {
	C.av_log_set_level(C.int(level))
}

func AvLogGetLevel() int {
	return int(C.av_log_get_level())
}
