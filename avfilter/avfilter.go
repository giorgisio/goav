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
	AVFilter        C.struct_AVFilter
	AVFilterContext C.struct_AVFilterContext
	AVFilterLink    C.struct_AVFilterLink
	AVFilterGraph   C.struct_AVFilterGraph
	AVFilterInOut   C.struct_AVFilterInOut
	AVFilterPad     C.struct_AVFilterPad
	AVDictionary    C.struct_AVDictionary
	AVClass         C.struct_AVClass
)
type (
	AVMediaType C.enum_AVMediaType
)

//unsigned 	avfilter_version (void)
//Return the LIBAVFILTER_VERSION_INT constant.
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

//int 	avfilter_pad_count (const AVFilterPad *pads)
//Get the number of elements in a NULL-terminated array of AVFilterPads (e.g.
func Avfilter_pad_count(p *AVFilterPad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

//const char * 	avfilter_pad_get_name (const AVFilterPad *pads, int pad_idx)
//Get the name of an AVFilterPad.
func Avfilter_pad_get_name(p *AVFilterPad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//enum AVMediaType 	avfilter_pad_get_type (const AVFilterPad *pads, int pad_idx)
//Get the type of an AVFilterPad.
func Avfilter_pad_get_type(p *AVFilterPad, pi int) AVMediaType {
	return (AVMediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//int 	avfilter_link (AVFilterContext *src, unsigned srcpad, AVFilterContext *dst, unsigned dstpad)
//Link two filters together.
func Avfilter_link(s *AVFilterContext, sp uint, d *AVFilterContext, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp)))
}

//void 	avfilter_link_free (AVFilterLink **link)
//Free the link in *link, and set its pointer to NULL.
func Avfilter_link_free(l **AVFilterLink) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

//int 	avfilter_link_get_channels (AVFilterLink *link)
//Get the number of channels of a link.
func Avfilter_link_get_channels(l *AVFilterLink) int {
	return int(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(l)))
}

//void 	avfilter_link_set_closed (AVFilterLink *link, int closed)
//Set the closed field of a link.
func Avfilter_link_set_closed(l *AVFilterLink, c int) {
	C.avfilter_link_set_closed((*C.struct_AVFilterLink)(l), C.int(c))
}

//int 	avfilter_config_links (AVFilterContext *filter)
//Negotiate the media format, dimensions, etc of all inputs to a filter.
func Avfilter_config_links(f *AVFilterContext) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

//int 	avfilter_process_command (AVFilterContext *filter, const char *cmd, const char *arg, char *res, int res_len, int flags)
//Make the filter instance process a command.
func Avfilter_process_command(f *AVFilterContext, cmd, arg, res string, l, fl int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), C.CString(res), C.int(l), C.int(fl)))
}

//void 	avfilter_register_all (void)
//Initialize the filter system.
func Avfilter_register_all() {
	C.avfilter_register_all()
}

//int 	avfilter_register (AVFilter *filter)
//Register a filter.
func Avfilter_register(f *AVFilter) int {
	return int(C.avfilter_register((*C.struct_AVFilter)(f)))
}

//const AVFilter * 	avfilter_get_by_name (const char *name)
//Get a filter definition matching the given name.
func Avfilter_get_by_name(n string) *AVFilter {
	return (*AVFilter)(C.avfilter_get_by_name(C.CString(n)))
}

//const AVFilter * 	avfilter_next (const AVFilter *prev)
//Iterate over all registered filters.
func Avfilter_next(f *AVFilter) *AVFilter {
	return (*AVFilter)(C.avfilter_next((*C.struct_AVFilter)(f)))
}

//int 	avfilter_init_str (AVFilterContext *ctx, const char *args)
//Initialize a filter with the supplied parameters.
func Avfilter_init_str(ctx *AVFilterContext, args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), C.CString(args)))
}

//int 	avfilter_init_dict (AVFilterContext *ctx, AVDictionary **options)
//Initialize a filter with the supplied dictionary of options.
func Avfilter_init_dict(f *AVFilterContext, o **AVDictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(f), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//void 	avfilter_free (AVFilterContext *filter)
//Free a filter context.
func Avfilter_free(f *AVFilterContext) {
	C.avfilter_free((*C.struct_AVFilterContext)(f))
}

//int 	avfilter_insert_filter (AVFilterLink *link, AVFilterContext *filt, unsigned filt_srcpad_idx, unsigned filt_dstpad_idx)
//Insert a filter in the middle of an existing link.
func Avfilter_insert_filter(l *AVFilterLink, f *AVFilterContext, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
}

//const AVClass * 	avfilter_get_class (void)
//avfilter_get_class
func Avfilter_get_class() *AVClass {
	return (*AVClass)(C.avfilter_get_class())
}

//Allocate a filter graph.
//AVFilterGraph * 	avfilter_graph_alloc (void)
func Avfilter_graph_alloc() *AVFilterGraph {
	return (*AVFilterGraph)(C.avfilter_graph_alloc())
}

//AVFilterContext * 	avfilter_graph_alloc_filter (AVFilterGraph *graph, const AVFilter *filter, const char *name)
//Create a new filter instance in a filter graph.
func Avfilter_graph_alloc_filter(g *AVFilterGraph, f *AVFilter, n string) *AVFilterContext {
	return (*AVFilterContext)(C.avfilter_graph_alloc_filter((*C.struct_AVFilterGraph)(g), (*C.struct_AVFilter)(f), C.CString(n)))
}

//AVFilterContext * 	avfilter_graph_get_filter (AVFilterGraph *graph, const char *name)
//Get a filter instance identified by instance name from graph.
func Avfilter_graph_get_filter(g *AVFilterGraph, n string) *AVFilterContext {
	return (*AVFilterContext)(C.avfilter_graph_get_filter((*C.struct_AVFilterGraph)(g), C.CString(n)))
}

//int 	avfilter_graph_create_filter (AVFilterContext **filt_ctx, const AVFilter *filt, const char *name, const char *args, void *opaque, AVFilterGraph *graph_ctx)
//Create and add a filter instance into an existing graph.
func Avfilter_graph_create_filter(cx **AVFilterContext, f *AVFilter, n, a string, o int, g *AVFilterGraph) int {
	return int(C.avfilter_graph_create_filter((**C.struct_AVFilterContext)(unsafe.Pointer(cx)), (*C.struct_AVFilter)(f), C.CString(n), C.CString(a), unsafe.Pointer(&o), (*C.struct_AVFilterGraph)(g)))
}

//void 	avfilter_graph_set_auto_convert (AVFilterGraph *graph, unsigned flags)
//Enable or disable automatic format conversion inside the graph.
func Avfilter_graph_set_auto_convert(g *AVFilterGraph, f uint) {
	C.avfilter_graph_set_auto_convert((*C.struct_AVFilterGraph)(g), C.uint(f))
}

//int 	avfilter_graph_config (AVFilterGraph *graphctx, void *log_ctx)
//Check validity and configure all the links and formats in the graph.
func Avfilter_graph_config(g *AVFilterGraph, l int) int {
	return int(C.avfilter_graph_config((*C.struct_AVFilterGraph)(g), unsafe.Pointer(&l)))
}

//void 	avfilter_graph_free (AVFilterGraph **graph)
//Free a graph, destroy its links, and set *graph to NULL.
func Avfilter_graph_free(g **AVFilterGraph) {
	C.avfilter_graph_free((**C.struct_AVFilterGraph)(unsafe.Pointer(g)))
}

//AVFilterInOut * 	avfilter_inout_alloc (void)
//Allocate a single AVFilterInOut entry.
func Avfilter_inout_alloc() *AVFilterInOut {
	return (*AVFilterInOut)(C.avfilter_inout_alloc())
}

//void 	avfilter_inout_free (AVFilterInOut **inout)
//Free the supplied list of AVFilterInOut and set *inout to NULL.
func Avfilter_inout_free(i **AVFilterInOut) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}

//int 	avfilter_graph_parse (AVFilterGraph *graph, const char *filters, AVFilterInOut *inputs, AVFilterInOut *outputs, void *log_ctx)
//Add a graph described by a string to a graph.
func Avfilter_graph_parse(g *AVFilterGraph, f string, i, o *AVFilterInOut, l int) int {
	return int(C.avfilter_graph_parse((*C.struct_AVFilterGraph)(g), C.CString(f), (*C.struct_AVFilterInOut)(i), (*C.struct_AVFilterInOut)(o), unsafe.Pointer(&l)))
}

//int 	avfilter_graph_parse_ptr (AVFilterGraph *graph, const char *filters, AVFilterInOut **inputs, AVFilterInOut **outputs, void *log_ctx)
//Add a graph described by a string to a graph.
func Avfilter_graph_parse_ptr(g *AVFilterGraph, f string, i, o **AVFilterInOut, l int) int {
	return int(C.avfilter_graph_parse_ptr((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o)), unsafe.Pointer(&l)))
}

//int 	avfilter_graph_parse2 (AVFilterGraph *graph, const char *filters, AVFilterInOut **inputs, AVFilterInOut **outputs)
//Add a graph described by a string to a graph.
func Avfilter_graph_parse2(g *AVFilterGraph, f string, i, o **AVFilterInOut) int {
	return int(C.avfilter_graph_parse2((*C.struct_AVFilterGraph)(g), C.CString(f), (**C.struct_AVFilterInOut)(unsafe.Pointer(i)), (**C.struct_AVFilterInOut)(unsafe.Pointer(o))))
}

//int 	avfilter_graph_send_command (AVFilterGraph *graph, const char *target, const char *cmd, const char *arg, char *res, int res_len, int flags)
//Send a command to one or more filter instances.
func Avfilter_graph_send_command(g *AVFilterGraph, t, cmd, arg, res string, resl, f int) int {
	return int(C.avfilter_graph_send_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.CString(res), C.int(resl), C.int(f)))
}

//int 	avfilter_graph_queue_command (AVFilterGraph *graph, const char *target, const char *cmd, const char *arg, int flags, double ts)
//Queue a command for one or more filter instances.
func Avfilter_graph_queue_command(g *AVFilterGraph, t, cmd, arg string, f int, ts C.double) int {
	return int(C.avfilter_graph_queue_command((*C.struct_AVFilterGraph)(g), C.CString(t), C.CString(cmd), C.CString(arg), C.int(f), ts))
}

//char * 	avfilter_graph_dump (AVFilterGraph *graph, const char *options)
//Dump a graph into a human-readable string representation.
func Avfilter_graph_dump(g *AVFilterGraph, o string) string {
	return C.GoString(C.avfilter_graph_dump((*C.struct_AVFilterGraph)(g), C.CString(o)))
}

//int 	avfilter_graph_request_oldest (AVFilterGraph *graph)
//Request a frame on the oldest sink
func Avfilter_graph_request_oldestlink(g *AVFilterGraph) int {
	return int(C.avfilter_graph_request_oldest((*C.struct_AVFilterGraph)(g)))
}
