// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Utility Frames
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
	Frame               C.struct_AVFrame
	AvFrameSideData     C.struct_AVFrameSideData
	AvFrameSideDataType C.enum_AVFrameSideDataType
)

//Dictionary ** avpriv_frame_get_metadatap (Frame *frame)
func Avpriv_frame_get_metadatap(f *Frame) **Dictionary {
	return (**Dictionary)(unsafe.Pointer(C.avpriv_frame_get_metadatap((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

//int av_frame_set_qp_table (Frame *f, AvBufferRef *buf, int stride, int qp_type)
func Av_frame_set_qp_table(f *Frame, b *AvBufferRef, s, q int) int {
	return int(C.av_frame_set_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.struct_AVBufferRef)(unsafe.Pointer(b)), C.int(s), C.int(q)))
}

//int8_t * av_frame_get_qp_table (Frame *f, int *stride, int *type)
func Av_frame_get_qp_table(f *Frame, s, t *int) int8 {
	return int8(*C.av_frame_get_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.int)(unsafe.Pointer(s)), (*C.int)(unsafe.Pointer(t))))
}

// //static void get_frame_defaults (Frame *frame)
// func Get_frame_defaults(f *Frame) {
// 	C.get_frame_defaults(*C.struct_AVFrame(f))
// }

//Frame * av_frame_alloc (void)
//Allocate an Frame and set its fields to default values.
func Av_frame_alloc() *Frame {
	return (*Frame)(unsafe.Pointer(C.av_frame_alloc()))
}

//void av_frame_free (Frame **frame)
//Free the frame and any dynamically allocated objects in it, e.g.
func Av_frame_free(f *Frame) {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(f)))
}

// //static int get_video_buffer (Frame *frame, int align)
// func Get_video_buffer(f *Frame, a int) int {
// 	return int(C.get_video_buffer(f, C.int(a)))
// }

// //static int get_audio_buffer (Frame *frame, int align)
// func Get_audio_buffer(f *Frame, a int) int {
// 	return C.get_audio_buffer(f, C.int(a))
// }

//int av_frame_get_buffer (Frame *frame, int align)
//Allocate new buffer(s) for audio or video data.
func Av_frame_get_buffer(f *Frame, a int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(a)))
}

//int av_frame_ref (Frame *dst, Frame *src)
//Setup a new reference to the data described by an given frame.
func Av_frame_ref(d, s *Frame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//Frame * av_frame_clone (Frame *src)
//Create a new frame that references the same data as src.
func Av_frame_clone(f *Frame) *Frame {
	return (*Frame)(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//void av_frame_unref (Frame *frame)
//Unreference all the buffers referenced by frame and reset the frame fields.
func Av_frame_unref(f *Frame) {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

//void av_frame_move_ref (Frame *dst, Frame *src)
//Move everythnig contained in src to dst and reset src.
func Av_frame_move_ref(d, s *Frame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s)))
}

//int av_frame_is_writable (Frame *frame)
//Check if the frame data is writable.
func Av_frame_is_writable(f *Frame) int {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//int av_frame_make_writable (Frame *frame)
//Ensure that the frame data is writable, avoiding data copy if possible.
func Av_frame_make_writable(f *Frame) int {
	return int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//int av_frame_copy_props (Frame *dst, const Frame *src)
//Copy only "metadata" fields from src to dst.
func Av_frame_copy_props(d, s *Frame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//AvBufferRef * av_frame_get_plane_buffer (Frame *frame, int plane)
//Get the buffer reference a given data plane is stored in.
func Av_frame_get_plane_buffer(f *Frame, p int) *AvBufferRef {
	return (*AvBufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p)))
}

//AvFrameSideData * av_frame_new_side_data (Frame *frame, enum AvFrameSideDataType type, int size)
//Add a new side data to a frame.
func Av_frame_new_side_data(f *Frame, d AvFrameSideDataType, s int) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s)))
}

//AvFrameSideData * av_frame_get_side_data (Frame *frame, enum AvFrameSideDataType type)
func Av_frame_get_side_data(f *Frame, t AvFrameSideDataType) *AvFrameSideData {
	return (*AvFrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t)))
}

func Data(f *Frame) *uint8 {
	return (*uint8)(unsafe.Pointer((*C.uint8_t)(unsafe.Pointer(&f.data))))
}
func Linesize(f *Frame) int {
	return int(*(*C.int)(unsafe.Pointer(&f.linesize)))
}
