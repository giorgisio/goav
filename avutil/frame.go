/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
package avutil

/*
	#cgo pkg-config: libavutil
	#include <libavutil/frame.h>
	#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

type (
	AvBuffer            C.struct_AVBuffer
	AvBufferRef         C.struct_AVBufferRef
	AvBufferPool        C.struct_AVBufferPool
	AvFrame             C.struct_AVFrame
	AvFrameSideData     C.struct_AVFrameSideData
	AvFrameSideDataType C.enum_AVFrameSideDataType
)

//AvDictionary ** avpriv_frame_get_metadatap (AvFrame *frame)
func Avpriv_frame_get_metadatap(f *AvFrame) **AvDictionary {
	return (**AvDictionary)(unsafe.Pointer(C.avpriv_frame_get_metadatap((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

//int av_frame_set_qp_table (AvFrame *f, AvBufferRef *buf, int stride, int qp_type)
func Av_frame_set_qp_table(f *AvFrame, b *AvBufferRef, s, q int) int {
	return int(C.av_frame_set_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.struct_AVBufferRef)(unsafe.Pointer(b)), C.int(s), C.int(q)))
}

//int8_t * av_frame_get_qp_table (AvFrame *f, int *stride, int *type)
func Av_frame_get_qp_table(f *AvFrame, s, t *int) int8 {
	return int8(*C.av_frame_get_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.int)(unsafe.Pointer(s)), (*C.int)(unsafe.Pointer(t))))
}

// //static void get_frame_defaults (AvFrame *frame)
// func Get_frame_defaults(f *AvFrame) {
// 	C.get_frame_defaults(*C.struct_AVFrame(f))
// }

//AvFrame * av_frame_alloc (void)
//Allocate an AvFrame and set its fields to default values.
func Av_frame_alloc() *AvFrame {
	return (*AvFrame)(unsafe.Pointer(C.av_frame_alloc()))
}

//void av_frame_free (AvFrame **frame)
//Free the frame and any dynamically allocated objects in it, e.g.
func Av_frame_free(f *AvFrame) {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(f)))
}

// //static int get_video_buffer (AvFrame *frame, int align)
// func Get_video_buffer(f *AvFrame, a int) int {
// 	return int(C.get_video_buffer(f, C.int(a)))
// }

// //static int get_audio_buffer (AvFrame *frame, int align)
// func Get_audio_buffer(f *AvFrame, a int) int {
// 	return C.get_audio_buffer(f, C.int(a))
// }

//int av_frame_get_buffer (AvFrame *frame, int align)
//Allocate new buffer(s) for audio or video data.
func Av_frame_get_buffer(f *AvFrame, a int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(a)))
}

//int av_frame_ref (AvFrame *dst, AvFrame *src)
//Setup a new reference to the data described by an given frame.
func Av_frame_ref(d, s *AvFrame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//AvFrame * av_frame_clone (AvFrame *src)
//Create a new frame that references the same data as src.
func Av_frame_clone(f *AvFrame) *AvFrame {
	return (*AvFrame)(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//void av_frame_unref (AvFrame *frame)
//Unreference all the buffers referenced by frame and reset the frame fields.
func Av_frame_unref(f *AvFrame) {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

//void av_frame_move_ref (AvFrame *dst, AvFrame *src)
//Move everythnig contained in src to dst and reset src.
func Av_frame_move_ref(d, s *AvFrame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s)))
}

//int av_frame_is_writable (AvFrame *frame)
//Check if the frame data is writable.
func Av_frame_is_writable(f *AvFrame) int {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//int av_frame_make_writable (AvFrame *frame)
//Ensure that the frame data is writable, avoiding data copy if possible.
func Av_frame_make_writable(f *AvFrame) int {
	return int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//int av_frame_copy_props (AvFrame *dst, const AvFrame *src)
//Copy only "metadata" fields from src to dst.
func Av_frame_copy_props(d, s *AvFrame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//AvBufferRef * av_frame_get_plane_buffer (AvFrame *frame, int plane)
//Get the buffer reference a given data plane is stored in.
func Av_frame_get_plane_buffer(f *AvFrame, p int) *AvBufferRef {
	return (*AvBufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p)))
}

//AvFrameSideData * av_frame_new_side_data (AvFrame *frame, enum AvFrameSideDataType type, int size)
//Add a new side data to a frame.
func Av_frame_new_side_data(f *AvFrame, d AvFrameSideDataType, s int) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s)))
}

//AvFrameSideData * av_frame_get_side_data (AvFrame *frame, enum AvFrameSideDataType type)
func Av_frame_get_side_data(f *AvFrame, t AvFrameSideDataType) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t)))
}

func Data(f *AvFrame) *uint8 {
	return (*uint8)(unsafe.Pointer((*C.uint8_t)(unsafe.Pointer(&f.data))))
}
func Linesize(f *AvFrame) int {
	return int(*(*C.int)(unsafe.Pointer(&f.linesize)))
}
