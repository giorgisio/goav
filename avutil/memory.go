/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
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

//void * 	av_malloc (size_t size) av_malloc_attrib 1(1)
//Allocate a block of size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU).
func Av_malloc(s uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.av_malloc(C.size_t(s)))
}

//static void * 	av_malloc_array (size_t nmemb, size_t size)
func Av_malloc_array(n, s uintptr) unsafe.Pointer {
	return C.av_malloc_array(C.size_t(n), C.size_t(s))
}

//void * 	av_realloc (void *ptr, size_t size) 1(2)
//Allocate or reallocate a block of memory.
func Av_realloc(p *int, s uintptr) unsafe.Pointer {
	return C.av_realloc(unsafe.Pointer(&p), C.size_t(s))
}

//void * 	av_realloc_f (void *ptr, size_t nelem, size_t elsize)
//Allocate or reallocate a block of memory.
func Av_realloc_f(p int, n, e uintptr) unsafe.Pointer {
	return C.av_realloc_f(unsafe.Pointer(&p), C.size_t(n), C.size_t(e))
}

//int 	av_reallocp (void *ptr, size_t size)
//Allocate or reallocate a block of memory.
func Av_reallocp(p int, s uintptr) int {
	return int(C.av_reallocp(unsafe.Pointer(&p), C.size_t(s)))
}

//void * 	av_realloc_array (void *ptr, size_t nmemb, size_t size)
func Av_realloc_array(p int, n, s uintptr) unsafe.Pointer {
	return C.av_realloc_array(unsafe.Pointer(&p), C.size_t(n), C.size_t(s))
}

//int 	av_reallocp_array (void *ptr, size_t nmemb, size_t size)
func Av_reallocp_array(p int, n, s uintptr) int {
	return int(C.av_reallocp_array(unsafe.Pointer(&p), C.size_t(n), C.size_t(s)))
}

//void 	av_free (void *ptr)
//Free a memory block which has been allocated with av_malloc(z)() or av_realloc().
func Av_free(p unsafe.Pointer) {
	C.av_free(p)
}

//void * 	av_mallocz (size_t size) av_malloc_attrib 1(1)
//Allocate a block of size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU) and zero all the bytes of the block.
func Av_mallocz(s uintptr) unsafe.Pointer {
	return C.av_mallocz(C.size_t(s))
}

//void * 	av_calloc (size_t nmemb, size_t size) av_malloc_attrib
//Allocate a block of nmemb * size bytes with alignment suitable for all memory accesses (including vectors if available on the CPU) and zero all the bytes of the block.
func Av_calloc(n, s uintptr) unsafe.Pointer {
	return C.av_calloc(C.size_t(n), C.size_t(s))
}

//static void * 	av_mallocz_array (size_t nmemb, size_t size)
func Av_mallocz_array(n, s uintptr) unsafe.Pointer {
	return C.av_mallocz_array(C.size_t(n), C.size_t(s))
}

//char * 	av_strdup (const char *s) av_malloc_attrib
//Duplicate the string s.
func Av_strdup(s string) string {
	return C.GoString(C.av_strdup(C.CString(s)))
}

//char * 	av_strndup (const char *s, size_t len) av_malloc_attrib
//Duplicate a substring of the string s.
func Av_strndup(s string, l uintptr) string {
	return C.GoString(C.av_strndup(C.CString(s), C.size_t(l)))
}

//void * 	av_memdup (const void *p, size_t size)
//Duplicate the buffer p.
func Av_memdup(p *int, s uintptr) unsafe.Pointer {
	return C.av_memdup(unsafe.Pointer(p), C.size_t(s))
}

//void 	av_freep (void *ptr)
//Free a memory block which has been allocated with av_malloc(z)() or av_realloc() and set the pointer pointing to it to NULL.
func Av_freep(p unsafe.Pointer) {
	C.av_freep(p)
}

//void 	av_dynarray_add (void *tab_ptr, int *nb_ptr, void *elem)
//Add an element to a dynamic array.
func Av_dynarray_add(t unsafe.Pointer, n *int, e unsafe.Pointer) {
	C.av_dynarray_add(t, (*C.int)(unsafe.Pointer(n)), e)
}

//int 	av_dynarray_add_nofree (void *tab_ptr, int *nb_ptr, void *elem)
//Add an element to a dynamic array.
func Av_dynarray_add_nofree(p unsafe.Pointer, n *int, e unsafe.Pointer) int {
	return int(C.av_dynarray_add_nofree(p, (*C.int)(unsafe.Pointer(&n)), e))
}

//void * 	av_dynarray2_add (void **tab_ptr, int *nb_ptr, size_t elem_size, const uint8_t *elem_data)
//Add an element of size elem_size to a dynamic array.
func Av_dynarray2_add(t *unsafe.Pointer, n *int, e uintptr, d uint8) unsafe.Pointer {
	return C.av_dynarray2_add(t, (*C.int)(unsafe.Pointer(&n)), (C.size_t)(e), (*C.uint8_t)(unsafe.Pointer(&d)))
}

//static int 	av_size_mult (size_t a, size_t b, size_t *r)
//Multiply two size_t values checking for overflow.
func Av_size_mult(a, b uintptr, r *uintptr) int {
	return int(C.av_size_mult((C.size_t)(a), (C.size_t)(b), (*C.size_t)(unsafe.Pointer(r))))
}

//void 	av_max_alloc (size_t max)
//Set the maximum size that may me allocated in one block.
func Av_max_alloc(m uintptr) {
	C.av_max_alloc(C.size_t(m))
}

//void 	av_memcpy_backptr (uint8_t *dst, int back, int cnt)
//deliberately overlapping memcpy implementation
func Av_memcpy_backptr(d *uintptr, b, c int) {
	C.av_memcpy_backptr((*C.uint8_t)(unsafe.Pointer(d)), C.int(b), C.int(c))
}

//void * 	av_fast_realloc (void *ptr, unsigned int *size, size_t min_size)
//Reallocate the given block if it is not large enough, otherwise do nothing.
func Av_fast_realloc(p unsafe.Pointer, s *uint, m uintptr) unsafe.Pointer {
	return C.av_fast_realloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(m))
}

//void 	av_fast_malloc (void *ptr, unsigned int *size, size_t min_size)
//Allocate a buffer, reusing the given one if large enough.
func Av_fast_malloc(p unsafe.Pointer, s *uint, m uintptr) {
	C.av_fast_malloc(p, (*C.uint)(unsafe.Pointer(s)), (C.size_t)(m))
}
