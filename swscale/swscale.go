/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)

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
	SwsContext    C.struct_SwsContext
	SwsFilter     C.struct_SwsFilter
	SwsVector     C.struct_SwsVector
	AvClass       C.struct_AVClass
	AvPixelFormat C.enum_AVPixelFormat
)

//unsigned swscale_version (void)
//Return the LIBSWSCALE_VERSION_INT constant.
func Swscale_version() uint {
	return uint(C.swscale_version())
}

//const char * swscale_configuration (void)
//Return the libswscale build-time configuration.
func Swscale_configuration() string {
	return C.GoString(C.swscale_configuration())
}

//const char * swscale_license (void)
//Return the libswscale license.
func Swscale_license() string {
	return C.GoString(C.swscale_license())
}

//const int * sws_getCoefficients (int colorspace)
//Return a pointer to yuv<->rgb coefficients for the given colorspace suitable for sws_setColorspaceDetails().
func Sws_getCoefficients(c int) *int {
	return (*int)(unsafe.Pointer(C.sws_getCoefficients(C.int(c))))
}

//int sws_isSupportedInput (enum AvPixelFormat pix_fmt)
//Return a positive value if pix_fmt is a supported input format, 0 otherwise.
func Sws_isSupportedInput(p AvPixelFormat) int {
	return int(C.sws_isSupportedInput((C.enum_AVPixelFormat)(p)))
}

//int sws_isSupportedOutput (enum AvPixelFormat pix_fmt)
//Return a positive value if pix_fmt is a supported output format, 0 otherwise.
func Sws_isSupportedOutput(p AvPixelFormat) int {
	return int(C.sws_isSupportedOutput((C.enum_AVPixelFormat)(p)))
}

//int sws_isSupportedEndiannessConversion (enum AvPixelFormat pix_fmt)
func Sws_isSupportedEndiannessConversion(p AvPixelFormat) int {
	return int(C.sws_isSupportedEndiannessConversion((C.enum_AVPixelFormat)(p)))
}

//struct SwsContext * sws_alloc_context (void)
//Allocate an empty SwsContext.
func Sws_alloc_context() *SwsContext {
	return (*SwsContext)(C.sws_alloc_context())
}

//int sws_init_context (struct SwsContext *sws_context, SwsFilter *srcFilter, SwsFilter *dstFilter)
//Initialize the swscaler context sws_context.
func Sws_init_context(ctxt *SwsContext, sf, df *SwsFilter) int {
	return int(C.sws_init_context((*C.struct_SwsContext)(ctxt), (*C.struct_SwsFilter)(sf), (*C.struct_SwsFilter)(df)))
}

//void sws_freeContext (struct SwsContext *swsContext)
//Free the swscaler context swsContext.
func Sws_freeContext(ctxt *SwsContext) {
	C.sws_freeContext((*C.struct_SwsContext)(ctxt))
}

//struct SwsContext * sws_getContext (int srcW, int srcH, enum AvPixelFormat srcFormat, int dstW, int dstH, enum AvPixelFormat dstFormat, int flags, SwsFilter *srcFilter, SwsFilter *dstFilter, const double *param)
//Allocate and return an SwsContext.
func Sws_getContext(sw, sh int, sf AvPixelFormat, dw, dh int, df AvPixelFormat, f int, sfl, dfl *SwsFilter, p *int) *SwsContext {
	return (*SwsContext)(C.sws_getContext(C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(unsafe.Pointer(p))))
}

////int sws_scale (struct SwsContext *c, const uint8_t *const srcSlice[], const int srcStride[], int srcSliceY, int srcSliceH, uint8_t *const dst[], const int dstStride[])
////Scale the image slice in srcSlice and put the resulting scaled slice in the image in dst.
func Sws_scale(ctxt *SwsContext, src *uint8, str int, y, h int, d *uint8, ds int) int {
	cctxt := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	csrc := (*C.uint8_t)(unsafe.Pointer(src))
	cstr := (*C.int)(unsafe.Pointer(&str))
	cd := (*C.uint8_t)(unsafe.Pointer(d))
	cds := (*C.int)(unsafe.Pointer(&ds))
	return int(C.sws_scale(cctxt, &csrc, cstr, C.int(y), C.int(h), &cd, cds))
}

//int sws_setColorspaceDetails (struct SwsContext *c, const int inv_table[4], int srcRange, const int table[4], int dstRange, int brightness, int contrast, int saturation)
func Sws_setColorspaceDetails(ctxt *SwsContext, it *int, sr int, t *int, dr, b, c, s int) int {
	cit := (*C.int)(unsafe.Pointer(it))
	ct := (*C.int)(unsafe.Pointer(t))
	return int(C.sws_setColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, C.int(sr), ct, C.int(dr), C.int(b), C.int(c), C.int(s)))
}

//int sws_getColorspaceDetails (struct SwsContext *c, int **inv_table, int *srcRange, int **table, int *dstRange, int *brightness, int *contrast, int *saturation)
func Sws_getColorspaceDetails(ctxt *SwsContext, it, sr, t, dr, b, c, s *int) int {
	cit := (**C.int)(unsafe.Pointer(it))
	csr := (*C.int)(unsafe.Pointer(sr))
	ct := (**C.int)(unsafe.Pointer(t))
	cdr := (*C.int)(unsafe.Pointer(dr))
	cb := (*C.int)(unsafe.Pointer(b))
	cc := (*C.int)(unsafe.Pointer(c))
	cs := (*C.int)(unsafe.Pointer(s))
	return int(C.sws_getColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, csr, ct, cdr, cb, cc, cs))
}

//SwsVector * sws_allocVec (int length)
//Allocate and return an uninitialized vector with length coefficients.
func Sws_allocVec(l int) *C.struct_SwsVector {
	return C.sws_allocVec(C.int(l))
}

//SwsVector * sws_getGaussianVec (double variance, double quality)
//Return a normalized Gaussian curve used to filter stuff quality = 3 is high quality, lower is lower quality.
func Sws_getGaussianVec(v, q float64) *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(v), C.double(q))))
}

//SwsVector * sws_getConstVec (double c, int length)
//Allocate and return a vector with length coefficients, all with the same value c.
func Sws_getConstVec(c float64, l int) *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_getConstVec(C.double(c), C.int(l))))
}

//SwsVector * sws_getIdentityVec (void)
//Allocate and return a vector with just one coefficient, with value 1.0.
func Sws_getIdentityVec() *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_getIdentityVec()))
}

//void sws_scaleVec (SwsVector *a, double scalar)
//Scale all the coefficients of a by the scalar value.
func Sws_scaleVec(a *SwsVector, s float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(unsafe.Pointer(a)), C.double(s))
}

//void sws_normalizeVec (SwsVector *a, double height)
//Scale all the coefficients of a so that their sum equals height.
func Sws_normalizeVec(a *SwsVector, h float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(a), C.double(h))
}

//void sws_convVec (SwsVector *a, SwsVector *b)
func Sws_convVec(a, b *SwsVector) {
	C.sws_convVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

//void sws_addVec (SwsVector *a, SwsVector *b)
func Sws_addVec(a, b *SwsVector) {
	C.sws_addVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

//void sws_subVec (SwsVector *a, SwsVector *b)
func Sws_subVec(a, b *SwsVector) {
	C.sws_subVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

//void sws_shiftVec (SwsVector *a, int shift)
func Sws_shiftVec(a *SwsVector, s int) {
	C.sws_shiftVec((*C.struct_SwsVector)(a), C.int(s))
}

//SwsVector * sws_cloneVec (SwsVector *a)
//Allocate and return a clone of the vector a, that is a vector with the same coefficients as a.
func Sws_cloneVec(a *SwsVector) *SwsVector {
	return (*SwsVector)(unsafe.Pointer(C.sws_cloneVec((*C.struct_SwsVector)(a))))
}

//void sws_printVec2 (SwsVector *a, AvClass *log_ctx, int log_level)
//Print with av_log() a textual representation of the vector a if log_level <= av_log_level.
func Sws_printVec2(a *SwsVector, lctx *AvClass, l int) {
	C.sws_printVec2((*C.struct_SwsVector)(a), (*C.struct_AVClass)(lctx), C.int(l))
}

//void sws_freeVec (SwsVector *a)
func Sws_freeVec(a *SwsVector) {
	C.sws_freeVec((*C.struct_SwsVector)(a))
}

//SwsFilter * sws_getDefaultFilter (float lumaGBlur, float chromaGBlur, float lumaSharpen, float chromaSharpen, float chromaHShift, float chromaVShift, int verbose)
func Sws_getDefaultFilter(lb, cb, ls, cs, chs, cvs float32, v int) *SwsFilter {
	return (*SwsFilter)(unsafe.Pointer(C.sws_getDefaultFilter(C.float(lb), C.float(cb), C.float(ls), C.float(cs), C.float(chs), C.float(cvs), C.int(v))))
}

//void sws_freeFilter (SwsFilter *filter)
func Sws_freeFilter(f *SwsFilter) {
	C.sws_freeFilter((*C.struct_SwsFilter)(f))
}

//struct SwsContext * sws_getCachedContext (struct SwsContext *context, int srcW, int srcH, enum AvPixelFormat srcFormat, int dstW, int dstH, enum AvPixelFormat dstFormat, int flags, SwsFilter *srcFilter, SwsFilter *dstFilter, const double *param)
//Check if context can be reused, otherwise reallocate a new one.
func Sws_getCachedContext(ctxt *SwsContext, sw, sh int, sf AvPixelFormat, dw, dh int, df AvPixelFormat, f int, sfl, dfl *SwsFilter, p *float64) *SwsContext {
	return (*SwsContext)(C.sws_getCachedContext((*C.struct_SwsContext)(ctxt), C.int(sw), C.int(sh), (C.enum_AVPixelFormat)(sf), C.int(dw), C.int(dh), (C.enum_AVPixelFormat)(df), C.int(f), (*C.struct_SwsFilter)(sfl), (*C.struct_SwsFilter)(dfl), (*C.double)(p)))
}

//void sws_convertPalette8ToPacked32 (const uint8_t *src, uint8_t *dst, int num_pixels, const uint8_t *palette)
//Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
func Sws_convertPalette8ToPacked32(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//void sws_convertPalette8ToPacked24 (const uint8_t *src, uint8_t *dst, int num_pixels, const uint8_t *palette)
//Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
func Sws_convertPalette8ToPacked24(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//const AvClass * sws_get_class (void)
//Get the AvClass for swsContext.
func Sws_get_class() *AvClass {
	return (*AvClass)(C.sws_get_class())
}
