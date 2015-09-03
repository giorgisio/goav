/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)

	Filters in the same linear chain are separated by commas,
	and distinct linear chains of filters are separated by semicolons.
	FFmpeg is enabled through the "C" libavfilter library
*/
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
	AvFilter        C.struct_AVFilter
	AvFilterContext C.struct_AVFilterContext
	AvFilterLink    C.struct_AVFilterLink
	AvFilterGraph   C.struct_AVFilterGraph
	AvFilterInOut   C.struct_AVFilterInOut
	AvFilterPad     C.struct_AVFilterPad
	Dictionary      C.struct_AVDictionary
	AvClass         C.struct_AVClass
)
type (
	MediaType C.enum_AVMediaType
)

//unsigned 	avfilter_version (void)
//Return the LIBAvFILTER_VERSION_INT constant.
func Avfilter_version() uint {
	return uint(C.avfilter_version())
}

//const char * 	avfilter_configuration (void)
//Return the libavfilter build-time configuration.
func Avfilter_configuration() string {
	return C.GoString(C.avfilter_configuration())
}

//const char * 	avfilter_license (void)
//Return the libavfilter license.
func Avfilter_license() string {
	return C.GoString(C.avfilter_license())
}

//int 	avfilter_pad_count (const AvFilterPad *pads)
//Get the number of elements in a NULL-terminated array of AvFilterPads (e.g.
func Avfilter_pad_count(p *AvFilterPad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

//const char * 	avfilter_pad_get_name (const AvFilterPad *pads, int pad_idx)
//Get the name of an AvFilterPad.
func Avfilter_pad_get_name(p *AvFilterPad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//enum MediaType 	avfilter_pad_get_type (const AvFilterPad *pads, int pad_idx)
//Get the type of an AvFilterPad.
func Avfilter_pad_get_type(p *AvFilterPad, pi int) MediaType {
	return (MediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//int 	avfilter_link (AvFilterContext *src, unsigned srcpad, AvFilterContext *dst, unsigned dstpad)
//Link two filters together.
func Avfilter_link(s *AvFilterContext, sp uint, d *AvFilterContext, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp)))
}

//void 	avfilter_link_free (AvFilterLink **link)
//Free the link in *link, and set its pointer to NULL.
func Avfilter_link_free(l **AvFilterLink) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

//int 	avfilter_link_get_channels (AvFilterLink *link)
//Get the number of channels of a link.
func Avfilter_link_get_channels(l *AvFilterLink) int {
	return int(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(l)))
}

//void 	avfilter_link_set_closed (AvFilterLink *link, int closed)
//Set the closed field of a link.
func Avfilter_link_set_closed(l *AvFilterLink, c int) {
	C.avfilter_link_set_closed((*C.struct_AVFilterLink)(l), C.int(c))
}

//int 	avfilter_config_links (AvFilterContext *filter)
//Negotiate the media format, dimensions, etc of all inputs to a filter.
func Avfilter_config_links(f *AvFilterContext) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

//int 	avfilter_process_command (AvFilterContext *filter, const char *cmd, const char *arg, char *res, int res_len, int flags)
//Make the filter instance process a command.
func Avfilter_process_command(f *AvFilterContext, cmd, arg, res string, l, fl int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), C.CString(res), C.int(l), C.int(fl)))
}

//void 	avfilter_register_all (void)
//Initialize the filter system.
func Avfilter_register_all() {
	C.avfilter_register_all()
}

//int 	avfilter_register (AvFilter *filter)
//Register a filter.
func Avfilter_register(f *AvFilter) int {
	return int(C.avfilter_register((*C.struct_AVFilter)(f)))
}

//const AvFilter * 	avfilter_get_by_name (const char *name)
//Get a filter definition matching the given name.
func Avfilter_get_by_name(n string) *AvFilter {
	return (*AvFilter)(C.avfilter_get_by_name(C.CString(n)))
}

//const AvFilter * 	avfilter_next (const AvFilter *prev)
//Iterate over all registered filters.
func Avfilter_next(f *AvFilter) *AvFilter {
	return (*AvFilter)(C.avfilter_next((*C.struct_AVFilter)(f)))
}

//int 	avfilter_init_str (AvFilterContext *ctx, const char *args)
//Initialize a filter with the supplied parameters.
func Avfilter_init_str(ctx *AvFilterContext, args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), C.CString(args)))
}

//int 	avfilter_init_dict (AvFilterContext *ctx, Dictionary **options)
//Initialize a filter with the supplied dictionary of options.
func Avfilter_init_dict(f *AvFilterContext, o **Dictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(f), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//void 	avfilter_free (AvFilterContext *filter)
//Free a filter context.
func Avfilter_free(f *AvFilterContext) {
	C.avfilter_free((*C.struct_AVFilterContext)(f))
}

//int 	avfilter_insert_filter (AvFilterLink *link, AvFilterContext *filt, unsigned filt_srcpad_idx, unsigned filt_dstpad_idx)
//Insert a filter in the middle of an existing link.
func Avfilter_insert_filter(l *AvFilterLink, f *AvFilterContext, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
}

//const AvClass * 	avfilter_get_class (void)
//avfilter_get_class
func Avfilter_get_class() *AvClass {
	return (*AvClass)(C.avfilter_get_class())
}

//Allocate a filter graph.
//AvFilterGraph * 	avfilter_graph_alloc (void)
func Avfilter_graph_alloc() *AvFilterGraph {
	return (*AvFilterGraph)(C.avfilter_graph_alloc())
}

//AvFilterContext * 	avfilter_graph_alloc_filter (AvFilterGraph *graph, const AvFilter *filter, const char *name)
//Create a new filter instance in a filter graph.
func Avfilter_graph_alloc_filter(g *AvFilterGraph, f *AvFilter, n string) *AvFilterContext {
	return (*AvFilterContext)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(g), (*C.struct_AVFilter)(f), C.CString(n)))
}

//AvFilterContext * 	avfilter_graph_get_filter (AvFilterGraph *graph, const char *name)
//Get a filter instance identified by instance name from graph.
func Avfilter_graph_get_filter(g *AvFilterGraph, n string) *AvFilterContext {
	return (*AvFilterContext)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(g), C.CString(n)))
}

//int 	avfilter_graph_create_filter (AvFilterContext **filt_ctx, const AvFilter *filt, const char *name, const char *args, void *opaque, AvFilterGraph *graph_ctx)
//Create and add a filter instance into an existing graph.
func Avfilter_graph_create_filter(cx **AvFilterContext, f *AvFilter, n, a string, o int, g *AvFilterGraph) int {
	return int(C.avfilter_graph_create_filter((**C.struct_AVFilterContext)(unsafe.Pointer(cx)), (*C.struct_AVFilter)(f), C.CString(n), C.CString(a), unsafe.Pointer(&o), (*C.struct_AVFilterGraph)(g)))
}

//void 	avfilter_graph_set_auto_convert (AvFilterGraph *graph, unsigned flags)
//Enable or disable automatic format conversion inside the graph.
func Avfilter_graph_set_auto_convert(g *AvFilterGraph, f uint) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(g), C.uint(f))
}

//int 	avfilter_graph_config (AvFilterGraph *graphctx, void *log_ctx)
//Check validity and configure all the links and formats in the graph.
func Avfilter_graph_config(g *AvFilterGraph, l int) int {
	return int(C.avfilter_graph_config((*C.struct_AVFilterGraph)(g), unsafe.Pointer(&l)))
}

//void 	avfilter_graph_free (AvFilterGraph **graph)
//Free a graph, destroy its links, and set *graph to NULL.
func Avfilter_graph_free(g **AvFilterGraph) {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(g)))
}

//AvFilterInOut * 	avfilter_inout_alloc (void)
//Allocate a single AvFilterInOut entry.
func Avfilter_inout_alloc() *AvFilterInOut {
	return (*AvFilterInOut)(C.avfilter_inout_alloc())
}

//void 	avfilter_inout_free (AvFilterInOut **inout)
//Free the supplied list of AvFilterInOut and set *inout to NULL.
func Avfilter_inout_free(i **AvFilterInOut) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}

//int 	avfilter_graph_parse (AvFilterGraph *graph, const char *filters, AvFilterInOut *inputs, AvFilterInOut *outputs, void *log_ctx)
//Add a graph described by a string to a graph.
func Avfilter_graph_parse(g *AvFilterGraph, f string, i, o *AvFilterInOut, l int) int {
	return int(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(g), C.CString(f), (*C.struct_AVFilterInOut)(i), (*C.struct_AVFilterInOut)(o), unsafe.Pointer(&l)))
}

//int 	avfilter_graph_parse_ptr (AvFilterGraph *graph, const char *filters, AvFilterInOut **inputs, AvFilterInOut **outputs, void *log_ctx)
//Add a graph described by a string to a graph.
func Avfilter_graph_parse_ptr(g *AvFilterGraph, f string, i, o **AvFilterInOut, l int) int {
	return int(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o)), unsafe.Pointer(&l)))
}

//int 	avfilter_graph_parse2 (AvFilterGraph *graph, const char *filters, AvFilterInOut **inputs, AvFilterInOut **outputs)
//Add a graph described by a string to a graph.
func Avfilter_graph_parse2(g *AvFilterGraph, f string, i, o **AvFilterInOut) int {
	return int(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o))))
}

//int 	avfilter_graph_send_command (AvFilterGraph *graph, const char *target, const char *cmd, const char *arg, char *res, int res_len, int flags)
//Send a command to one or more filter instances.
func Avfilter_graph_send_command(g *AvFilterGraph, t, cmd, arg, res string, resl, f int) int {
	return int(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.CString(res), C.int(resl), C.int(f)))
}

//int 	avfilter_graph_queue_command (AvFilterGraph *graph, const char *target, const char *cmd, const char *arg, int flags, double ts)
//Queue a command for one or more filter instances.
func Avfilter_graph_queue_command(g *AvFilterGraph, t, cmd, arg string, f int, ts C.double) int {
	return int(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.int(f), ts))
}

//char * 	avfilter_graph_dump (AvFilterGraph *graph, const char *options)
//Dump a graph into a human-readable string representation.
func Avfilter_graph_dump(g *AvFilterGraph, o string) string {
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(g), C.CString(o)))
}

//int 	avfilter_graph_request_oldest (AvFilterGraph *graph)
//Request a frame on the oldest sink
func Avfilter_graph_request_oldestlink(g *AvFilterGraph) int {
	return int(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(g)))
}
