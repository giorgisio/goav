/*
	libswscale library performs highly optimized image scaling and colorspace and pixel format conversion operations.
	Rescaling: is the process of changing the video size. Several rescaling options and algorithms are available.
	Pixel format conversion: is the process of converting the image format and colorspace of the image.
*/
package swscale

//#cgo pkg-config: libswscale libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <string.h>
//#include <stdint.h>
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

type (
	Context     C.struct_SwsContext
	Filter      C.struct_SwsFilter
	Vector      C.struct_SwsVector
	Class       C.struct_AVClass
	PixelFormat C.enum_AVPixelFormat
)

//Return the LIBSWSCALE_VERSION_INT constant.
func Swscale_version() uint {
	return uint(C.swscale_version())
}

//Return the libswscale build-time configuration.
func Swscale_configuration() string {
	return C.GoString(C.swscale_configuration())
}

//Return the libswscale license.
func Swscale_license() string {
	return C.GoString(C.swscale_license())
}

//Return a pointer to yuv<->rgb coefficients for the given colorspace suitable for sws_setColorspaceDetails().
func Sws_getCoefficients(c int) *int {
	return (*int)(unsafe.Pointer(C.sws_getCoefficients(C.int(c))))
}

//Return a positive value if pix_fmt is a supported input format, 0 otherwise.
func Sws_isSupportedInput(p PixelFormat) int {
	return int(C.sws_isSupportedInput((C.enum_AVPixelFormat)(p)))
}

//Return a positive value if pix_fmt is a supported output format, 0 otherwise.
func Sws_isSupportedOutput(p PixelFormat) int {
	return int(C.sws_isSupportedOutput((C.enum_AVPixelFormat)(p)))
}

func Sws_isSupportedEndiannessConversion(p PixelFormat) int {
	return int(C.sws_isSupportedEndiannessConversion((C.enum_AVPixelFormat)(p)))
}

//Allocate an empty Context.
func Sws_alloc_context() *Context {
	return (*Context)(C.sws_alloc_context())
}

//Initialize the swscaler context sws_context.
func Sws_init_context(ctxt *Context, sf, df *Filter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(ctxt), (*C.struct_SwsFilter)(sf), (*C.struct_SwsFilter)(df)))
}

//Free the swscaler context swsContext.
func Sws_freeContext(ctxt *Context) {
	C.sws_freeContext((*C.struct_SwsContext)(ctxt))
}

//Allocate and return an Context.
func Sws_getContext(sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *Filter, p *int) *Context {
	return (*Context)(C.sws_getContext(C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(unsafe.Pointer(p))))
}

////Scale the image slice in srcSlice and put the resulting scaled slice in the image in dst.
func Sws_scale(ctxt *Context, src *uint8, str int, y, h int, d *uint8, ds int) int {
	cctxt := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	csrc := (*C.uint8_t)(unsafe.Pointer(src))
	cstr := (*C.int)(unsafe.Pointer(&str))
	cd := (*C.uint8_t)(unsafe.Pointer(d))
	cds := (*C.int)(unsafe.Pointer(&ds))
	return int(C.sws_scale(cctxt, &csrc, cstr, C.int(y), C.int(h), &cd, cds))
}

func Sws_setColorspaceDetails(ctxt *Context, it *int, sr int, t *int, dr, b, c, s int) int {
	cit := (*C.int)(unsafe.Pointer(it))
	ct := (*C.int)(unsafe.Pointer(t))
	return int(C.sws_setColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, C.int(sr), ct, C.int(dr), C.int(b), C.int(c), C.int(s)))
}

func Sws_getColorspaceDetails(ctxt *Context, it, sr, t, dr, b, c, s *int) int {
	cit := (**C.int)(unsafe.Pointer(it))
	csr := (*C.int)(unsafe.Pointer(sr))
	ct := (**C.int)(unsafe.Pointer(t))
	cdr := (*C.int)(unsafe.Pointer(dr))
	cb := (*C.int)(unsafe.Pointer(b))
	cc := (*C.int)(unsafe.Pointer(c))
	cs := (*C.int)(unsafe.Pointer(s))
	return int(C.sws_getColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, csr, ct, cdr, cb, cc, cs))
}

//Allocate and return an uninitialized vector with length coefficients.
func Sws_allocVec(l int) *C.struct_SwsVector {
	return C.sws_allocVec(C.int(l))
}

//Return a normalized Gaussian curve used to filter stuff quality = 3 is high quality, lower is lower quality.
func Sws_getGaussianVec(v, q float64) *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(v), C.double(q))))
}

//Allocate and return a vector with length coefficients, all with the same value c.
func Sws_getConstVec(c float64, l int) *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getConstVec(C.double(c), C.int(l))))
}

//Allocate and return a vector with just one coefficient, with value 1.0.
func Sws_getIdentityVec() *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getIdentityVec()))
}

//Scale all the coefficients of a by the scalar value.
func Sws_scaleVec(a *Vector, s float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(unsafe.Pointer(a)), C.double(s))
}

//Scale all the coefficients of a so that their sum equals height.
func Sws_normalizeVec(a *Vector, h float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(a), C.double(h))
}

func Sws_convVec(a, b *Vector) {
	C.sws_convVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

func Sws_addVec(a, b *Vector) {
	C.sws_addVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

func Sws_subVec(a, b *Vector) {
	C.sws_subVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

func Sws_shiftVec(a *Vector, s int) {
	C.sws_shiftVec((*C.struct_SwsVector)(a), C.int(s))
}

//Allocate and return a clone of the vector a, that is a vector with the same coefficients as a.
func Sws_cloneVec(a *Vector) *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_cloneVec((*C.struct_SwsVector)(a))))
}

//Print with av_log() a textual representation of the vector a if log_level <= av_log_level.
func Sws_printVec2(a *Vector, lctx *Class, l int) {
	C.sws_printVec2((*C.struct_SwsVector)(a), (*C.struct_AVClass)(lctx), C.int(l))
}

func Sws_freeVec(a *Vector) {
	C.sws_freeVec((*C.struct_SwsVector)(a))
}

func Sws_getDefaultFilter(lb, cb, ls, cs, chs, cvs float32, v int) *Filter {
	return (*Filter)(unsafe.Pointer(C.sws_getDefaultFilter(C.float(lb), C.float(cb), C.float(ls), C.float(cs), C.float(chs), C.float(cvs), C.int(v))))
}

func Sws_freeFilter(f *Filter) {
	C.sws_freeFilter((*C.struct_SwsFilter)(f))
}

//Check if context can be reused, otherwise reallocate a new one.
func Sws_getCachedContext(ctxt *Context, sw, sh int, sf PixelFormat, dw, dh int, df PixelFormat, f int, sfl, dfl *Filter, p *float64) *Context {
	return (*Context)(C.sws_getCachedContext((*C.struct_SwsContext)(ctxt), C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(p)))
}

//Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
func Sws_convertPalette8ToPacked32(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
func Sws_convertPalette8ToPacked24(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//Get the Class for swsContext.
func Sws_get_class() *Class {
	return (*Class)(C.sws_get_class())
}
