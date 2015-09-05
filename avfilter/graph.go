// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"
)

//Allocate a filter graph.
func AvfilterGraphAlloc() *Graph {
	return (*Graph)(C.avfilter_graph_alloc())
}

//Create a new filter instance in a filter graph.
func (g *Graph) AvfilterGraphAllocFilter(f *Filter, n string) *Context {
	return (*Context)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(g), (*C.struct_AVFilter)(f), C.CString(n)))
}

//Get a filter instance identified by instance name from graph.
func (g *Graph) AvfilterGraphGetFilter(n string) *Context {
	return (*Context)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(g), C.CString(n)))
}

//Enable or disable automatic format conversion inside the graph.
func (g *Graph) AvfilterGraphSetAutoConvert(f uint) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(g), C.uint(f))
}

//Check validity and configure all the links and formats in the graph.
func (g *Graph) AvfilterGraphConfig(l int) int {
	return int(C.avfilter_graph_config((*C.struct_AVFilterGraph)(g), unsafe.Pointer(&l)))
}

//Free a graph, destroy its links, and set *graph to NULL.
func (g *Graph) AvfilterGraphFree() {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(g)))
}

//Add a graph described by a string to a graph.
func (g *Graph) AvfilterGraphParse(f string, i, o *Input, l int) int {
	return int(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(g), C.CString(f), (*C.struct_AVFilterInOut)(i), (*C.struct_AVFilterInOut)(o), unsafe.Pointer(&l)))
}

//Add a graph described by a string to a graph.
func (g *Graph) AvfilterGraphParsePtr(f string, i, o **Input, l int) int {
	return int(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o)), unsafe.Pointer(&l)))
}

//Add a graph described by a string to a graph.
func (g *Graph) AvfilterGraphParse2(f string, i, o **Input) int {
	return int(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o))))
}

//Send a command to one or more filter instances.
func (g *Graph) AvfilterGraphSendCommand(t, cmd, arg, res string, resl, f int) int {
	return int(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.CString(res), C.int(resl), C.int(f)))
}

//Queue a command for one or more filter instances.
func (g *Graph) AvfilterGraphQueueCommand(t, cmd, arg string, f int, ts C.double) int {
	return int(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.int(f), ts))
}

//Dump a graph into a human-readable string representation.
func (g *Graph) AvfilterGraphDump(o string) string {
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(g), C.CString(o)))
}

//Request a frame on the oldest sink
func (g *Graph) AvfilterGraphRequestOldestlink() int {
	return int(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(g)))
}

//Create and add a filter instance into an existing graph.
func AvfilterGraphCreateFilter(cx **Context, f *Filter, n, a string, o int, g *Graph) int {
	return int(C.avfilter_graph_create_filter((**C.struct_AVFilterContext)(unsafe.Pointer(cx)), (*C.struct_AVFilter)(f), C.CString(n), C.CString(a), unsafe.Pointer(&o), (*C.struct_AVFilterGraph)(g)))
}
