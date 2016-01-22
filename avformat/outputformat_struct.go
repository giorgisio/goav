// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"

//Close an opened input Context.
func (s *OutputFormat) Flags() int {
	return int(s.flags)
}