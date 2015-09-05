// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/avutil.h>
	#include <libavutil/frame.h>
	#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

//Allocate a block of size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU).
func AvMalloc(s uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.av_malloc(C.size_t(s)))
}

func AvMallocArray(n, s uintptr) unsafe.Pointer {
	return C.av_malloc_array(C.size_t(n), C.size_t(s))
}

//Allocate or reallocate a block of memory.
func AvRealloc(p *int, s uintptr) unsafe.Pointer {
	return C.av_realloc(unsafe.Pointer(&p), C.size_t(s))
}

//Allocate or reallocate a block of memory.
func AvReallocF(p int, n, e uintptr) unsafe.Pointer {
	return C.av_realloc_f(unsafe.Pointer(&p), C.size_t(n), C.size_t(e))
}

//Allocate or reallocate a block of memory.
func AvReallocp(p int, s uintptr) int {
	return int(C.av_reallocp(unsafe.Pointer(&p), C.size_t(s)))
}

func AvReallocArray(p int, n, s uintptr) unsafe.Pointer {
	return C.av_realloc_array(unsafe.Pointer(&p), C.size_t(n), C.size_t(s))
}

func AvReallocpArray(p int, n, s uintptr) int {
	return int(C.av_reallocp_array(unsafe.Pointer(&p), C.size_t(n), C.size_t(s)))
}

//Free a memory block which has been allocated with av_malloc(z)() or av_realloc().
func AvFree(p unsafe.Pointer) {
	C.av_free(p)
}

//Allocate a block of size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU) and zero all the bytes of the block.
func AvMallocz(s uintptr) unsafe.Pointer {
	return C.av_mallocz(C.size_t(s))
}

//Allocate a block of nmemb * size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU) and zero all the bytes of the block.
func AvCalloc(n, s uintptr) unsafe.Pointer {
	return C.av_calloc(C.size_t(n), C.size_t(s))
}

func AvMalloczArray(n, s uintptr) unsafe.Pointer {
	return C.av_mallocz_array(C.size_t(n), C.size_t(s))
}

//Duplicate the string s.
func AvStrdup(s string) string {
	return C.GoString(C.av_strdup(C.CString(s)))
}

//char * 	av_strndup (const char *s, size_t len) av_malloc_attrib
//Duplicate a substring of the string s.
func AvStrndup(s string, l uintptr) string {
	return C.GoString(C.av_strndup(C.CString(s), C.size_t(l)))
}

//Duplicate the buffer p.
func AvMemdup(p *int, s uintptr) unsafe.Pointer {
	return C.av_memdup(unsafe.Pointer(p), C.size_t(s))
}

//Free a memory block which has been allocated with av_malloc(z)() or av_realloc() and set the pointer pointing to it to NULL.
func AvFreep(p unsafe.Pointer) {
	C.av_freep(p)
}

//Add an element to a dynamic array.
func AvDynarrayAdd(t unsafe.Pointer, n *int, e unsafe.Pointer) {
	C.av_dynarray_add(t, (*C.int)(unsafe.Pointer(n)), e)
}

//Add an element to a dynamic array.
func AvDynarrayAddNofree(p unsafe.Pointer, n *int, e unsafe.Pointer) int {
	return int(C.av_dynarray_add_nofree(p, (*C.int)(unsafe.Pointer(&n)), e))
}

//Add an element of size elem_size to a dynamic array.
func AvDynarray2Add(t *unsafe.Pointer, n *int, e uintptr, d uint8) unsafe.Pointer {
	return C.av_dynarray2_add(t, (*C.int)(unsafe.Pointer(&n)), (C.size_t)(e), (*C.uint8_t)(unsafe.Pointer(&d)))
}

//Multiply two size_t values checking for overflow.
func AvSizeMult(a, b uintptr, r *uintptr) int {
	return int(C.av_size_mult((C.size_t)(a), (C.size_t)(b), (*C.size_t)(unsafe.Pointer(r))))
}

//Set the maximum size that may me allocated in one block.
func AvMaxAlloc(m uintptr) {
	C.av_max_alloc(C.size_t(m))
}

//deliberately overlapping memcpy implementation
func AvMemcpyBackptr(d *uintptr, b, c int) {
	C.av_memcpy_backptr((*C.uint8_t)(unsafe.Pointer(d)), C.int(b), C.int(c))
}

//Reallocate the given block if it is not large enough, otherwise do nothing.
func AvFastRealloc(p unsafe.Pointer, s *uint, m uintptr) unsafe.Pointer {
	return C.av_fast_realloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(m))
}

//Allocate a buffer, reusing the given one if large enough.
func AvFastMalloc(p unsafe.Pointer, s *uint, m uintptr) {
	C.av_fast_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(m))
}
