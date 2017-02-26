// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)
// Dmitry Patsura <talk@dmtry.me> https://github.com/ovr

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

func (c *Codec) Name() string {
	return C.GoString(c.name)
}

func (c *Codec) LongName() string {
	return C.GoString(c.long_name)
}
