// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Turing Zhu (qishiwenjun@163.com)

// Package avutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type (
	AVMediaType   C.enum_AVMediaType
	AVPictureType C.enum_AVPictureType
	AVRational    C.AVRational
	File          C.FILE
)

// Return the LIBAVUTIL_VERSION_INT constant.
func AvUtilVersion() uint {
	return uint(C.avutil_version())
}

// Return an informative version string. This usually is the actual release
// version number or a git commit description. This string has no fixed format
// and can change any time. It should never be parsed by code.
func AvVersionInfo() string {
	return C.GoString(C.av_version_info())
}

// Return the libavutil build-time configuration.
func AvUtilConfiguration() string {
	return C.GoString(C.avutil_configuration())
}

// Return the libavutil license.
func AvUtilLicense() string {
	return C.GoString(C.avutil_license())
}

//Return a string describing the media_type enum, NULL if media_type is unknown.
func AvGetMediaTypeString(mt AVMediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

/**
 * Return a single letter to describe the given picture type
 * pict_type.
 *
 * @param[in] pict_type the picture type @return a single character
 * representing the picture type, '?' if pict_type is unknown
 */
func AvGetPictureTypeChar(pt AVPictureType) string {
	return string(C.av_get_picture_type_char((C.enum_AVPictureType)(pt)))
}

// Return x default pointer in case p is NULL.
func AvXIfNull(p, x int) {
	C.av_x_if_null(unsafe.Pointer(&p), unsafe.Pointer(&x))
}

/**
 * Compute the length of an integer list.
 *
 * @param elsize  size in bytes of each list element (only 1, 2, 4 or 8)
 * @param term    list terminator (usually 0 or -1)
 * @param list    pointer to the list
 * @return  length of the list, in elements, not counting the terminator
 */
func AvIntListLengthForSize(eachListElement uint, list int, term uint64) uint {
	return uint(C.av_int_list_length_for_size(C.uint(eachListElement), unsafe.Pointer(&list), (C.uint64_t)(term)))
}

/**
 * Open a file using a UTF-8 filename.
 * The API of this function matches POSIX fopen(), errors are returned through
 * errno.
 */
func AvFopenUtf8(path, mode string) *File {
	f := C.av_fopen_utf8(C.CString(path), C.CString(mode))
	return (*File)(f)
}

// Return the fractional representation of the internal time base.
func AvGetTimeBaseQ() AVRational {
	return (AVRational)(C.av_get_time_base_q())
}

/**
 * Fill the provided buffer with a string containing a FourCC (four-character
 * code) representation.
 *
 * @param buf    a buffer with size in bytes of at least AV_FOURCC_MAX_STRING_SIZE
 * @param fourcc the fourcc to represent
 * @return the buffer in input
 */
func AvFourccMakeString(buf string, fourcc uint32) string {
	// return C.GoString(C.av_version_info())

	return C.GoString(C.av_fourcc_make_string(C.CString(buf), C.uint32_t(fourcc)))
}
