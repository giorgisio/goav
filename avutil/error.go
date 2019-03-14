// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/error.h>
//#include <stdlib.h>
//static const char *error2string(int code) { return av_err2str(code); }
import "C"
import "errors"

const (
	AvErrorEOF    = -('E' | ('O' << 8) | ('F' << 16) | (' ' << 24))
	AvErrorEAGAIN = -35
)

func ErrorFromCode(code int) error {
	if code >= 0 {
		return nil
	}

	return errors.New(C.GoString(C.error2string(C.int(code))))
}
