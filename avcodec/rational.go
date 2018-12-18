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
import "fmt"

func (r Rational) String() string {
	return fmt.Sprintln("%d/%d", int(r.num), int(r.den))
}

func (r *Rational) Assign(o Rational) {
	*r = o
}

func ToRational(r *C.struct_AVRational) Rational {
	return Rational{
		num: r.num,
		den: r.den,
	}
}
