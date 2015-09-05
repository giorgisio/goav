// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package swscale

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

//Allocate and return an uninitialized vector with length coefficients.
func SwsAllocvec(l int) *Vector {
	return (*Vector)(C.sws_allocVec(C.int(l)))
}

//Return a normalized Gaussian curve used to filter stuff quality = 3 is high quality, lower is lower quality.
func SwsGetgaussianvec(v, q float64) *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getGaussianVec(C.double(v), C.double(q))))
}

//Allocate and return a vector with length coefficients, all with the same value c.
func SwsGetconstvec(c float64, l int) *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getConstVec(C.double(c), C.int(l))))
}

//Allocate and return a vector with just one coefficient, with value 1.0.
func SwsGetidentityvec() *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getIdentityVec()))
}

//Scale all the coefficients of a by the scalar value.
func (a *Vector) SwsScalevec(s float64) {
	C.sws_scaleVec((*C.struct_SwsVector)(unsafe.Pointer(a)), C.double(s))
}

//Scale all the coefficients of a so that their sum equals height.
func (a *Vector) SwsNormalizevec(h float64) {
	C.sws_normalizeVec((*C.struct_SwsVector)(a), C.double(h))
}

func (a *Vector) SwsConvvec(b *Vector) {
	C.sws_convVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

func (a *Vector) SwsAddvec(b *Vector) {
	C.sws_addVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

func (a *Vector) SwsSubvec(b *Vector) {
	C.sws_subVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

func (a *Vector) SwsShiftvec(s int) {
	C.sws_shiftVec((*C.struct_SwsVector)(a), C.int(s))
}

//Allocate and return a clone of the vector a, that is a vector with the same coefficients as a.
func (a *Vector) SwsClonevec() *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_cloneVec((*C.struct_SwsVector)(a))))
}

//Print with av_log() a textual representation of the vector a if log_level <= av_log_level.
func (a *Vector) SwsPrintvec2(lctx *Class, l int) {
	C.sws_printVec2((*C.struct_SwsVector)(a), (*C.struct_AVClass)(lctx), C.int(l))
}

func (a *Vector) SwsFreevec() {
	C.sws_freeVec((*C.struct_SwsVector)(a))
}
