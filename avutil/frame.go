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
	AVBuffer            C.struct_AVBuffer
	AVBufferRef         C.struct_AVBufferRef
	AVBufferPool        C.struct_AVBufferPool
	AVFrame             C.struct_AVFrame
	AVFrameSideData     C.struct_AVFrameSideData
	AVFrameSideDataType C.enum_AVFrameSideDataType
)

//AVDictionary ** avpriv_frame_get_metadatap (AVFrame *frame)
func Avpriv_frame_get_metadatap(f *AVFrame) **AVDictionary {
	return (**AVDictionary)(unsafe.Pointer(C.avpriv_frame_get_metadatap((*C.struct_AVFrame)(unsafe.Pointer(f)))))
}

//int av_frame_set_qp_table (AVFrame *f, AVBufferRef *buf, int stride, int qp_type)
func Av_frame_set_qp_table(f *AVFrame, b *AVBufferRef, s, q int) int {
	return int(C.av_frame_set_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.struct_AVBufferRef)(unsafe.Pointer(b)), C.int(s), C.int(q)))
}

//int8_t * av_frame_get_qp_table (AVFrame *f, int *stride, int *type)
func Av_frame_get_qp_table(f *AVFrame, s, t *int) int8 {
	return int8(*C.av_frame_get_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.int)(unsafe.Pointer(s)), (*C.int)(unsafe.Pointer(t))))
}

// //static void get_frame_defaults (AVFrame *frame)
// func Get_frame_defaults(f *AVFrame) {
// 	C.get_frame_defaults(*C.struct_AVFrame(f))
// }

//AVFrame * av_frame_alloc (void)
//Allocate an AVFrame and set its fields to default values.
func Av_frame_alloc() *AVFrame {
	return (*AVFrame)(unsafe.Pointer(C.av_frame_alloc()))
}

//void av_frame_free (AVFrame **frame)
//Free the frame and any dynamically allocated objects in it, e.g.
func Av_frame_free(f *AVFrame) {
	C.av_frame_free((**C.struct_AVFrame)(unsafe.Pointer(f)))
}

// //static int get_video_buffer (AVFrame *frame, int align)
// func Get_video_buffer(f *AVFrame, a int) int {
// 	return int(C.get_video_buffer(f, C.int(a)))
// }

// //static int get_audio_buffer (AVFrame *frame, int align)
// func Get_audio_buffer(f *AVFrame, a int) int {
// 	return C.get_audio_buffer(f, C.int(a))
// }

//int av_frame_get_buffer (AVFrame *frame, int align)
//Allocate new buffer(s) for audio or video data.
func Av_frame_get_buffer(f *AVFrame, a int) int {
	return int(C.av_frame_get_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(a)))
}

//int av_frame_ref (AVFrame *dst, AVFrame *src)
//Setup a new reference to the data described by an given frame.
func Av_frame_ref(d, s *AVFrame) int {
	return int(C.av_frame_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//AVFrame * av_frame_clone (AVFrame *src)
//Create a new frame that references the same data as src.
func Av_frame_clone(f *AVFrame) *AVFrame {
	return (*AVFrame)(C.av_frame_clone((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//void av_frame_unref (AVFrame *frame)
//Unreference all the buffers referenced by frame and reset the frame fields.
func Av_frame_unref(f *AVFrame) {
	cf := (*C.struct_AVFrame)(unsafe.Pointer(f))
	C.av_frame_unref(cf)
}

//void av_frame_move_ref (AVFrame *dst, AVFrame *src)
//Move everythnig contained in src to dst and reset src.
func Av_frame_move_ref(d, s *AVFrame) {
	C.av_frame_move_ref((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s)))
}

//int av_frame_is_writable (AVFrame *frame)
//Check if the frame data is writable.
func Av_frame_is_writable(f *AVFrame) int {
	return int(C.av_frame_is_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//int av_frame_make_writable (AVFrame *frame)
//Ensure that the frame data is writable, avoiding data copy if possible.
func Av_frame_make_writable(f *AVFrame) int {
	return int(C.av_frame_make_writable((*C.struct_AVFrame)(unsafe.Pointer(f))))
}

//int av_frame_copy_props (AVFrame *dst, const AVFrame *src)
//Copy only "metadata" fields from src to dst.
func Av_frame_copy_props(d, s *AVFrame) int {
	return int(C.av_frame_copy_props((*C.struct_AVFrame)(unsafe.Pointer(d)), (*C.struct_AVFrame)(unsafe.Pointer(s))))
}

//AVBufferRef * av_frame_get_plane_buffer (AVFrame *frame, int plane)
//Get the buffer reference a given data plane is stored in.
func Av_frame_get_plane_buffer(f *AVFrame, p int) *AVBufferRef {
	return (*AVBufferRef)(C.av_frame_get_plane_buffer((*C.struct_AVFrame)(unsafe.Pointer(f)), C.int(p)))
}

//AVFrameSideData * av_frame_new_side_data (AVFrame *frame, enum AVFrameSideDataType type, int size)
//Add a new side data to a frame.
func Av_frame_new_side_data(f *AVFrame, d AVFrameSideDataType, s int) *AVFrameSideData {
	return (*AVFrameSideData)(C.av_frame_new_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(d), C.int(s)))
}

//AVFrameSideData * av_frame_get_side_data (AVFrame *frame, enum AVFrameSideDataType type)
func Av_frame_get_side_data(f *AVFrame, t AVFrameSideDataType) *AVFrameSideData {
	return (*AVFrameSideData)(C.av_frame_get_side_data((*C.struct_AVFrame)(unsafe.Pointer(f)), (C.enum_AVFrameSideDataType)(t)))
}

func Data(f *AVFrame) *uint8 {
	return (*uint8)(unsafe.Pointer((*C.uint8_t)(unsafe.Pointer(&f.data))))
}
func Linesize(f *AVFrame) int {
	return int(*(*C.int)(unsafe.Pointer(&f.linesize)))
}
