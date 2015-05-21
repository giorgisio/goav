/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)

	The process of changing the sampling rate of a discrete signal to obtain
	a new discrete representation of the underlying continuous signal.
	resampler provides a high-level interface to the libswresample library audio resampling utilities
*/
package swresample

/*
	#cgo pkg-config: libswresample
	#include <libswresample/swresample.h>
*/
import "C"
import (
	"unsafe"
)

type (
	SwrContext     C.struct_SwrContext
	AVFrame        C.struct_AVFrame
	AVClass        C.struct_AVClass
	AVSampleFormat C.enum_AVSampleFormat
)

//const AVClass * swr_get_class (void)
//Get the AVClass for SwrContext.
func Swr_get_class() *AVClass {
	return (*AVClass)(C.swr_get_class())
}

//SwrContext constructor functions
//struct SwrContext * swr_alloc (void)
//Allocate SwrContext.
func Swr_alloc() *SwrContext {
	return (*SwrContext)(C.swr_alloc())
}

//int swr_init (struct SwrContext *s)
//Initialize context after user parameters have been set.
func Swr_init(s *SwrContext) int {
	return int(C.swr_init((*C.struct_SwrContext)(s)))
}

//int swr_is_initialized (struct SwrContext *s)
//Check whether an swr context has been initialized or not.
func Swr_is_initialized(s *SwrContext) int {
	return int(C.swr_is_initialized((*C.struct_SwrContext)(s)))
}

//struct SwrContext * swr_alloc_set_opts (struct SwrContext *s, int64_t out_ch_layout, enum AVSampleFormat out_sample_fmt, int out_sample_rate, int64_t in_ch_layout, enum AVSampleFormat in_sample_fmt, int in_sample_rate, int log_offset, void *log_ctx)
//Allocate SwrContext if needed and set/reset common parameters.
func Swr_alloc_set_opts(s *SwrContext, ocl int64, osf AVSampleFormat, osr int, icl int64, isf AVSampleFormat, isr, lo, lc int) *SwrContext {
	return (*SwrContext)(C.swr_alloc_set_opts((*C.struct_SwrContext)(s), C.int64_t(ocl), (C.enum_AVSampleFormat)(osf), C.int(osr), C.int64_t(icl), (C.enum_AVSampleFormat)(isf), C.int(isr), C.int(lo), unsafe.Pointer(&lc)))
}

//SwrContext destructor functions
//void swr_free (struct SwrContext **s)
//Free the given SwrContext and set the pointer to NULL.
func Swr_free(s **SwrContext) {
	C.swr_free((**C.struct_SwrContext)(unsafe.Pointer(s)))
}

//void swr_close (struct SwrContext *s)
//Closes the context so that swr_is_initialized() returns 0.
func Swr_close(s *SwrContext) {
	C.swr_close((*C.struct_SwrContext)(s))
}

//Core conversion functions
//int swr_convert (struct SwrContext *s, uint8_t **out, int out_count, const uint8_t **in, int in_count)
//Convert audio.
func Swr_convert(s *SwrContext, out **uint8, oc int, in **uint8, ic int) int {
	return int(C.swr_convert((*C.struct_SwrContext)(s), (**C.uint8_t)(unsafe.Pointer(out)), C.int(oc), (**C.uint8_t)(unsafe.Pointer(in)), C.int(ic)))
}

//int64_t swr_next_pts (struct SwrContext *s, int64_t pts)
//Convert the next timestamp from input to output timestamps are in 1/(in_sample_rate * out_sample_rate) units.
func Swr_next_pts(s *SwrContext, pts int64) int64 {
	return int64(C.swr_next_pts((*C.struct_SwrContext)(s), C.int64_t(pts)))
}

//Low-level option setting functions
//These functons provide a means to set low-level options that is not possible with the AVOption API.
//int swr_set_compensation (struct SwrContext *s, int sample_delta, int compensation_distance)
//Activate resampling compensation ("soft" compensation).
func Swr_set_compensation(s *SwrContext, sd, cd int) int {
	return int(C.swr_set_compensation((*C.struct_SwrContext)(s), C.int(sd), C.int(cd)))
}

//int swr_set_channel_mapping (struct SwrContext *s, const int *channel_map)
//Set a customized input channel mapping.
func Swr_set_channel_mapping(s *SwrContext, cm *int) int {
	return int(C.swr_set_channel_mapping((*C.struct_SwrContext)(s), (*C.int)(unsafe.Pointer(cm))))
}

//int swr_set_matrix (struct SwrContext *s, const double *matrix, int stride)
//Set a customized remix matrix.
func Swr_set_matrix(s *SwrContext, m *int, t int) int {
	return int(C.swr_set_matrix((*C.struct_SwrContext)(s), (*C.double)(unsafe.Pointer(m)), C.int(t)))
}

//Sample handling functions
//int swr_drop_output (struct SwrContext *s, int count)
//Drops the specified number of output samples.
func Swr_drop_output(s *SwrContext, c int) int {
	return int(C.swr_drop_output((*C.struct_SwrContext)(s), C.int(c)))
}

//int swr_inject_silence (struct SwrContext *s, int count)
//Injects the specified number of silence samples.
func Swr_inject_silence(s *SwrContext, c int) int {
	return int(C.swr_inject_silence((*C.struct_SwrContext)(s), C.int(c)))
}

//int64_t swr_get_delay (struct SwrContext *s, int64_t base)
//Gets the delay the next input sample will experience relative to the next output sample.
func Swr_get_delay(s *SwrContext, b int64) int64 {
	return int64(C.swr_get_delay((*C.struct_SwrContext)(s), C.int64_t(b)))
}

//Configuration accessors
//unsigned swresample_version (void)
//Return the LIBSWRESAMPLE_VERSION_INT constant.
func Swresample_version() uint {
	return uint(C.swresample_version())
}

//const char * swresample_configuration (void)
//Return the swr build-time configuration.
func Swresample_configuration() string {
	return C.GoString(C.swresample_configuration())
}

//const char * swresample_license (void)
//Return the swr license.
func Swresample_license() string {
	return C.GoString(C.swresample_license())
}

//AVFrame based API
//int swr_convert_frame (SwrContext *swr, AVFrame *output, const AVFrame *input)
//Convert the samples in the input AVFrame and write them to the output AVFrame.
func Swr_convert_frame(s *SwrContext, o, i *AVFrame) int {
	return int(C.swr_convert_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}

//int swr_config_frame (SwrContext *swr, const AVFrame *out, const AVFrame *in)
//Configure or reconfigure the SwrContext using the information provided by the AVFrames.
func Swr_config_frame(s *SwrContext, o, i *AVFrame) int {
	return int(C.swr_config_frame((*C.struct_SwrContext)(s), (*C.struct_AVFrame)(o), (*C.struct_AVFrame)(i)))
}
