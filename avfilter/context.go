// Use of this source code is governed by a MIT license that can be found in the LICENSE file.


//Package avfilter contains methods that deal with ffmpeg filters
//filters in the same linear chain are separated by commas, and distinct linear chains of filters are separated by semicolons.
//FFmpeg is enabled through the "C" libavfilter library
package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/buffersrc.h>
*/
import "C"
import (
	"unsafe"
	"github.com/giorgisio/goav/avutil"
)


//Create a new filter instance in a filter graph.
func (c *Context) AvBuffersrcAddFrame(frame * avutil.Frame) int {
	var c_frame = (*C.struct_AVFrame)(unsafe.Pointer(frame))
	return int(C.av_buffersrc_add_frame((*C.struct_AVFilterContext)(c),c_frame))
}