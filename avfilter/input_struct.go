package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import "unsafe"

func (i *Input) Name() string {
	return C.GoString(i.name)
}

func (i *Input) SetName(name string) {
	i.name = C.CString(name)
}

func (i *Input) Context() *Context {
	return (*Context)(unsafe.Pointer(i.filter_ctx))
}

func (i *Input) SetContext(ctxt *Context)  {
	i.filter_ctx = (*C.struct_AVFilterContext)(unsafe.Pointer(ctxt));
}

func (i *Input) Next() *Input {
	return (*Input)(unsafe.Pointer(i.next))
}

func (i *Input) SetNext(next *Input) {
	i.next = (*C.struct_AVFilterInOut)(unsafe.Pointer(next));
}

func (i *Input) PadIndex() int {
	return int(i.pad_idx)
}

func (i *Input) SetPadIndex(idx int) {
	i.pad_idx = (C.int)(idx)
}