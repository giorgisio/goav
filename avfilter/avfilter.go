// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Filters in the same linear chain are separated by commas, and distinct linear chains of filters are separated by semicolons.
//FFmpeg is enabled through the "C" libavfilter library
package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
*/
import "C"
import (
	"unsafe"
)

type (
	Filter     C.struct_AVFilter
	Context    C.struct_AVFilterContext
	Link       C.struct_AVFilterLink
	Graph      C.struct_AVFilterGraph
	Input      C.struct_AVFilterInOut
	Pad        C.struct_AVFilterPad
	Dictionary C.struct_AVDictionary
	Class      C.struct_AVClass
	MediaType  C.enum_AVMediaType
)

//Return the LIBAvFILTER_VERSION_INT constant.
func AvfilterVersion() uint {
	return uint(C.avfilter_version())
}

//Return the libavfilter build-time configuration.
func AvfilterConfiguration() string {
	return C.GoString(C.avfilter_configuration())
}

//Return the libavfilter license.
func AvfilterLicense() string {
	return C.GoString(C.avfilter_license())
}

//Get the number of elements in a NULL-terminated array of Pads (e.g.
func AvfilterPadCount(p *Pad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

//Get the name of an Pad.
func AvfilterPadGetName(p *Pad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//Get the type of an Pad.
func AvfilterPadGetType(p *Pad, pi int) MediaType {
	return (MediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//Link two filters together.
func AvfilterLink(s *Context, sp uint, d *Context, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp)))
}

//Free the link in *link, and set its pointer to NULL.
func AvfilterLinkFree(l **Link) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

//Get the number of channels of a link.
func AvfilterLinkGetChannels(l *Link) int {
	return int(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(l)))
}

//Set the closed field of a link.
func AvfilterLinkSetClosed(l *Link, c int) {
	C.avfilter_link_set_closed((*C.struct_AVFilterLink)(l), C.int(c))
}

//Negotiate the media format, dimensions, etc of all inputs to a filter.
func AvfilterConfigLinks(f *Context) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

//Make the filter instance process a command.
func AvfilterProcessCommand(f *Context, cmd, arg, res string, l, fl int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), C.CString(res), C.int(l), C.int(fl)))
}

//Initialize the filter system.
func AvfilterRegisterAll() {
	C.avfilter_register_all()
}

//Register a filter.
func AvfilterRegister(f *Filter) int {
	return int(C.avfilter_register((*C.struct_AVFilter)(f)))
}

//Get a filter definition matching the given name.
func AvfilterGetByName(n string) *Filter {
	return (*Filter)(C.avfilter_get_by_name(C.CString(n)))
}

//Iterate over all registered filters.
func AvfilterNext(f *Filter) *Filter {
	return (*Filter)(C.avfilter_next((*C.struct_AVFilter)(f)))
}

//Initialize a filter with the supplied parameters.
func (ctx *Context) AvfilterInitStr(args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), C.CString(args)))
}

//Initialize a filter with the supplied dictionary of options.
func (ctx *Context) AvfilterInitDict(o **Dictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//Free a filter context.
func (ctx *Context) AvfilterFree() {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

//Insert a filter in the middle of an existing link.
func AvfilterInsertFilter(l *Link, f *Context, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
}

//avfilter_get_class
func AvfilterGetClass() *Class {
	return (*Class)(C.avfilter_get_class())
}

//Allocate a filter graph.
func AvfilterGraphAlloc() *Graph {
	return (*Graph)(C.avfilter_graph_alloc())
}

//Create a new filter instance in a filter graph.
func AvfilterGraphAllocFilter(g *Graph, f *Filter, n string) *Context {
	return (*Context)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(g), (*C.struct_AVFilter)(f), C.CString(n)))
}

//Get a filter instance identified by instance name from graph.
func AvfilterGraphGetFilter(g *Graph, n string) *Context {
	return (*Context)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(g), C.CString(n)))
}

//Create and add a filter instance into an existing graph.
func AvfilterGraphCreateFilter(cx **Context, f *Filter, n, a string, o int, g *Graph) int {
	return int(C.avfilter_graph_create_filter((**C.struct_AVFilterContext)(unsafe.Pointer(cx)), (*C.struct_AVFilter)(f), C.CString(n), C.CString(a), unsafe.Pointer(&o), (*C.struct_AVFilterGraph)(g)))
}

//Enable or disable automatic format conversion inside the graph.
func AvfilterGraphSetAutoConvert(g *Graph, f uint) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(g), C.uint(f))
}

//Check validity and configure all the links and formats in the graph.
func AvfilterGraphConfig(g *Graph, l int) int {
	return int(C.avfilter_graph_config((*C.struct_AVFilterGraph)(g), unsafe.Pointer(&l)))
}

//Free a graph, destroy its links, and set *graph to NULL.
func AvfilterGraphFree(g **Graph) {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(g)))
}

//Allocate a single Input entry.
func AvfilterInoutAlloc() *Input {
	return (*Input)(C.avfilter_inout_alloc())
}

//Free the supplied list of Input and set *inout to NULL.
func AvfilterInoutFree(i **Input) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}

//Add a graph described by a string to a graph.
func AvfilterGraphParse(g *Graph, f string, i, o *Input, l int) int {
	return int(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(g), C.CString(f), (*C.struct_AVFilterInOut)(i), (*C.struct_AVFilterInOut)(o), unsafe.Pointer(&l)))
}

//Add a graph described by a string to a graph.
func AvfilterGraphParsePtr(g *Graph, f string, i, o **Input, l int) int {
	return int(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o)), unsafe.Pointer(&l)))
}

//Add a graph described by a string to a graph.
func AvfilterGraphParse2(g *Graph, f string, i, o **Input) int {
	return int(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o))))
}

//Send a command to one or more filter instances.
func AvfilterGraphSendCommand(g *Graph, t, cmd, arg, res string, resl, f int) int {
	return int(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.CString(res), C.int(resl), C.int(f)))
}

//Queue a command for one or more filter instances.
func AvfilterGraphQueueCommand(g *Graph, t, cmd, arg string, f int, ts C.double) int {
	return int(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.int(f), ts))
}

//Dump a graph into a human-readable string representation.
func AvfilterGraphDump(g *Graph, o string) string {
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(g), C.CString(o)))
}

//Request a frame on the oldest sink
func AvfilterGraphRequestOldestlink(g *Graph) int {
	return int(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(g)))
}
