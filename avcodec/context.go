/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
	by Giorgis (habtom@giorgis.io)
*/
package avcodec

func (ctxt *AVCodecContext) Active_thread_type() int {
	return int(ctxt.active_thread_type)
}

func (ctxt *AVCodecContext) B_frame_strategy() int {
	return int(ctxt.b_frame_strategy)
}

func (ctxt *AVCodecContext) B_quant_factor() float64 {
	return float64(ctxt.b_quant_factor)
}

func (ctxt *AVCodecContext) B_quant_offset() float64 {
	return float64(ctxt.b_quant_offset)
}

func (ctxt *AVCodecContext) B_sensitivity() int {
	return int(ctxt.b_sensitivity)
}

func (ctxt *AVCodecContext) Bidir_refine() int {
	return int(ctxt.bidir_refine)
}

func (ctxt *AVCodecContext) Bit_rate() int {
	return int(ctxt.bit_rate)
}

func (ctxt *AVCodecContext) Bit_rate_tolerance() int {
	return int(ctxt.bit_rate_tolerance)
}

func (ctxt *AVCodecContext) Bits_per_coded_sample() int {
	return int(ctxt.bits_per_coded_sample)
}

func (ctxt *AVCodecContext) Bits_per_raw_sample() int {
	return int(ctxt.bits_per_raw_sample)
}

func (ctxt *AVCodecContext) Block_align() int {
	return int(ctxt.block_align)
}

func (ctxt *AVCodecContext) Brd_scale() int {
	return int(ctxt.brd_scale)
}

func (ctxt *AVCodecContext) Channels() int {
	return int(ctxt.channels)
}

func (ctxt *AVCodecContext) Chromaoffset() int {
	return int(ctxt.chromaoffset)
}

func (ctxt *AVCodecContext) Coded_height() int {
	return int(ctxt.coded_height)
}

func (ctxt *AVCodecContext) Coded_width() int {
	return int(ctxt.coded_width)
}

func (ctxt *AVCodecContext) Coder_type() int {
	return int(ctxt.coder_type)
}

func (ctxt *AVCodecContext) Compression_level() int {
	return int(ctxt.compression_level)
}

func (ctxt *AVCodecContext) Context_model() int {
	return int(ctxt.context_model)
}

func (ctxt *AVCodecContext) Cutoff() int {
	return int(ctxt.cutoff)
}

func (ctxt *AVCodecContext) Dark_masking() float64 {
	return float64(ctxt.dark_masking)
}

func (ctxt *AVCodecContext) Dct_algo() int {
	return int(ctxt.dct_algo)
}

func (ctxt *AVCodecContext) Debug() int {
	return int(ctxt.debug)
}

func (ctxt *AVCodecContext) Debug_mv() int {
	return int(ctxt.debug_mv)
}

func (ctxt *AVCodecContext) Delay() int {
	return int(ctxt.delay)
}

func (ctxt *AVCodecContext) Dia_size() int {
	return int(ctxt.dia_size)
}

func (ctxt *AVCodecContext) Err_recognition() int {
	return int(ctxt.err_recognition)
}

func (ctxt *AVCodecContext) Error_concealment() int {
	return int(ctxt.error_concealment)
}

func (ctxt *AVCodecContext) Extradata_size() int {
	return int(ctxt.extradata_size)
}

func (ctxt *AVCodecContext) Flags() int {
	return int(ctxt.flags)
}

func (ctxt *AVCodecContext) Flags2() int {
	return int(ctxt.flags2)
}

func (ctxt *AVCodecContext) Frame_bits() int {
	return int(ctxt.frame_bits)
}

func (ctxt *AVCodecContext) Frame_number() int {
	return int(ctxt.frame_number)
}

func (ctxt *AVCodecContext) Frame_size() int {
	return int(ctxt.frame_size)
}

func (ctxt *AVCodecContext) Frame_skip_cmp() int {
	return int(ctxt.frame_skip_cmp)
}

func (ctxt *AVCodecContext) Frame_skip_exp() int {
	return int(ctxt.frame_skip_exp)
}

func (ctxt *AVCodecContext) Frame_skip_factor() int {
	return int(ctxt.frame_skip_factor)
}

func (ctxt *AVCodecContext) Frame_skip_threshold() int {
	return int(ctxt.frame_skip_threshold)
}

func (ctxt *AVCodecContext) Global_quality() int {
	return int(ctxt.global_quality)
}

func (ctxt *AVCodecContext) Gop_size() int {
	return int(ctxt.gop_size)
}

func (ctxt *AVCodecContext) Has_b_frames() int {
	return int(ctxt.has_b_frames)
}

func (ctxt *AVCodecContext) Header_bits() int {
	return int(ctxt.header_bits)
}

func (ctxt *AVCodecContext) Height() int {
	return int(ctxt.height)
}

func (ctxt *AVCodecContext) I_count() int {
	return int(ctxt.i_count)
}

func (ctxt *AVCodecContext) I_quant_factor() float64 {
	return float64(ctxt.i_quant_factor)
}

func (ctxt *AVCodecContext) I_quant_offset() float64 {
	return float64(ctxt.i_quant_offset)
}

func (ctxt *AVCodecContext) I_tex_bits() int {
	return int(ctxt.i_tex_bits)
}

func (ctxt *AVCodecContext) Idct_algo() int {
	return int(ctxt.idct_algo)
}

func (ctxt *AVCodecContext) Ildct_cmp() int {
	return int(ctxt.ildct_cmp)
}

// func (ctxt *AVCodecContext) Initial_padding() int {
// 	return int(ctxt.initial_padding)
// }

func (ctxt *AVCodecContext) Intra_dc_precision() int {
	return int(ctxt.intra_dc_precision)
}

func (ctxt *AVCodecContext) Keyint_min() int {
	return int(ctxt.keyint_min)
}

func (ctxt *AVCodecContext) Last_predictor_count() int {
	return int(ctxt.last_predictor_count)
}

func (ctxt *AVCodecContext) Level() int {
	return int(ctxt.level)
}

func (ctxt *AVCodecContext) Log_level_offset() int {
	return int(ctxt.log_level_offset)
}

func (ctxt *AVCodecContext) Lowres() int {
	return int(ctxt.lowres)
}

func (ctxt *AVCodecContext) Lumi_masking() float64 {
	return float64(ctxt.lumi_masking)
}

func (ctxt *AVCodecContext) Max_b_frames() int {
	return int(ctxt.max_b_frames)
}

func (ctxt *AVCodecContext) Max_prediction_order() int {
	return int(ctxt.max_prediction_order)
}

func (ctxt *AVCodecContext) Max_qdiff() int {
	return int(ctxt.max_qdiff)
}

func (ctxt *AVCodecContext) Mb_cmp() int {
	return int(ctxt.mb_cmp)
}

func (ctxt *AVCodecContext) Mb_decision() int {
	return int(ctxt.mb_decision)
}

func (ctxt *AVCodecContext) Mb_lmax() int {
	return int(ctxt.mb_lmax)
}

func (ctxt *AVCodecContext) Mb_lmin() int {
	return int(ctxt.mb_lmin)
}

func (ctxt *AVCodecContext) Me_cmp() int {
	return int(ctxt.me_cmp)
}

func (ctxt *AVCodecContext) Me_penalty_compensation() int {
	return int(ctxt.me_penalty_compensation)
}

func (ctxt *AVCodecContext) Me_pre_cmp() int {
	return int(ctxt.me_pre_cmp)
}

func (ctxt *AVCodecContext) Me_range() int {
	return int(ctxt.me_range)
}

func (ctxt *AVCodecContext) Me_sub_cmp() int {
	return int(ctxt.me_sub_cmp)
}

func (ctxt *AVCodecContext) Me_subpel_quality() int {
	return int(ctxt.me_subpel_quality)
}

func (ctxt *AVCodecContext) Min_prediction_order() int {
	return int(ctxt.min_prediction_order)
}

func (ctxt *AVCodecContext) Misc_bits() int {
	return int(ctxt.misc_bits)
}

func (ctxt *AVCodecContext) Mpeg_quant() int {
	return int(ctxt.mpeg_quant)
}

func (ctxt *AVCodecContext) Mv0_threshold() int {
	return int(ctxt.mv0_threshold)
}

func (ctxt *AVCodecContext) Mv_bits() int {
	return int(ctxt.mv_bits)
}

func (ctxt *AVCodecContext) Noise_reduction() int {
	return int(ctxt.noise_reduction)
}

func (ctxt *AVCodecContext) Nsse_weight() int {
	return int(ctxt.nsse_weight)
}

func (ctxt *AVCodecContext) P_count() int {
	return int(ctxt.p_count)
}

func (ctxt *AVCodecContext) P_masking() float64 {
	return float64(ctxt.p_masking)
}

func (ctxt *AVCodecContext) P_tex_bits() int {
	return int(ctxt.p_tex_bits)
}

func (ctxt *AVCodecContext) Pre_dia_size() int {
	return int(ctxt.pre_dia_size)
}

func (ctxt *AVCodecContext) Pre_me() int {
	return int(ctxt.pre_me)
}

func (ctxt *AVCodecContext) Prediction_method() int {
	return int(ctxt.prediction_method)
}

func (ctxt *AVCodecContext) Profile() int {
	return int(ctxt.profile)
}

func (ctxt *AVCodecContext) Qblur() float64 {
	return float64(ctxt.qblur)
}

func (ctxt *AVCodecContext) Qcompress() float64 {
	return float64(ctxt.qcompress)
}

func (ctxt *AVCodecContext) Qmax() int {
	return int(ctxt.qmax)
}

func (ctxt *AVCodecContext) Qmin() int {
	return int(ctxt.qmin)
}

func (ctxt *AVCodecContext) Rc_buffer_size() int {
	return int(ctxt.rc_buffer_size)
}

func (ctxt *AVCodecContext) Rc_initial_buffer_occupancy() int {
	return int(ctxt.rc_initial_buffer_occupancy)
}

func (ctxt *AVCodecContext) Rc_max_available_vbv_use() float64 {
	return float64(ctxt.rc_max_available_vbv_use)
}

func (ctxt *AVCodecContext) Rc_max_rate() int {
	return int(ctxt.rc_max_rate)
}

func (ctxt *AVCodecContext) Rc_min_rate() int {
	return int(ctxt.rc_min_rate)
}

func (ctxt *AVCodecContext) Rc_min_vbv_overflow_use() float64 {
	return float64(ctxt.rc_min_vbv_overflow_use)
}

func (ctxt *AVCodecContext) Rc_override_count() int {
	return int(ctxt.rc_override_count)
}

func (ctxt *AVCodecContext) Refcounted_frames() int {
	return int(ctxt.refcounted_frames)
}

func (ctxt *AVCodecContext) Refs() int {
	return int(ctxt.refs)
}

func (ctxt *AVCodecContext) Rtp_payload_size() int {
	return int(ctxt.rtp_payload_size)
}

func (ctxt *AVCodecContext) Sample_rate() int {
	return int(ctxt.sample_rate)
}

func (ctxt *AVCodecContext) Scenechange_threshold() int {
	return int(ctxt.scenechange_threshold)
}

func (ctxt *AVCodecContext) Seek_preroll() int {
	return int(ctxt.seek_preroll)
}

func (ctxt *AVCodecContext) Side_data_only_packets() int {
	return int(ctxt.side_data_only_packets)
}

func (ctxt *AVCodecContext) Skip_alpha() int {
	return int(ctxt.skip_alpha)
}

func (ctxt *AVCodecContext) Skip_bottom() int {
	return int(ctxt.skip_bottom)
}

func (ctxt *AVCodecContext) Skip_count() int {
	return int(ctxt.skip_count)
}

func (ctxt *AVCodecContext) Skip_top() int {
	return int(ctxt.skip_top)
}

func (ctxt *AVCodecContext) Slice_count() int {
	return int(ctxt.slice_count)
}

func (ctxt *AVCodecContext) Slice_flags() int {
	return int(ctxt.slice_flags)
}

func (ctxt *AVCodecContext) Slices() int {
	return int(ctxt.slices)
}

func (ctxt *AVCodecContext) Spatial_cplx_masking() float64 {
	return float64(ctxt.spatial_cplx_masking)
}

func (ctxt *AVCodecContext) Strict_std_compliance() int {
	return int(ctxt.strict_std_compliance)
}

func (ctxt *AVCodecContext) Sub_charenc_mode() int {
	return int(ctxt.sub_charenc_mode)
}

func (ctxt *AVCodecContext) Subtitle_header_size() int {
	return int(ctxt.subtitle_header_size)
}

func (ctxt *AVCodecContext) Temporal_cplx_masking() float64 {
	return float64(ctxt.temporal_cplx_masking)
}

func (ctxt *AVCodecContext) Thread_count() int {
	return int(ctxt.thread_count)
}

func (ctxt *AVCodecContext) Thread_safe_callbacks() int {
	return int(ctxt.thread_safe_callbacks)
}

func (ctxt *AVCodecContext) Thread_type() int {
	return int(ctxt.thread_type)
}

func (ctxt *AVCodecContext) Ticks_per_frame() int {
	return int(ctxt.ticks_per_frame)
}

func (ctxt *AVCodecContext) Trellis() int {
	return int(ctxt.trellis)
}

func (ctxt *AVCodecContext) Width() int {
	return int(ctxt.width)
}

func (ctxt *AVCodecContext) Workaround_bugs() int {
	return int(ctxt.workaround_bugs)
}

func (ctxt *AVCodecContext) Audio_service_type() AVAudioServiceType { //enum
	return (AVAudioServiceType)(ctxt.audio_service_type)
}

func (ctxt *AVCodecContext) Chroma_sample_location() AVChromaLocation { //enum
	return (AVChromaLocation)(ctxt.chroma_sample_location)
}

func (ctxt *AVCodecContext) Codec_descriptor() *AVCodecDescriptor { //const
	return (*AVCodecDescriptor)(ctxt.codec_descriptor)
}

func (ctxt *AVCodecContext) Codec_id() AVCodecID { //enum
	return (AVCodecID)(ctxt.codec_id)
}

func (ctxt *AVCodecContext) Codec_type() AVMediaType { //enum
	return (AVMediaType)(ctxt.codec_type)
}

func (ctxt *AVCodecContext) Color_primaries() AVColorPrimaries { //enum
	return (AVColorPrimaries)(ctxt.color_primaries)
}

func (ctxt *AVCodecContext) Color_range() AVColorRange { //enum
	return (AVColorRange)(ctxt.color_range)
}

func (ctxt *AVCodecContext) Color_trc() AVColorTransferCharacteristic { //enum
	return (AVColorTransferCharacteristic)(ctxt.color_trc)
}

func (ctxt *AVCodecContext) Colorspace() AVColorSpace { //enum
	return (AVColorSpace)(ctxt.colorspace)
}

func (ctxt *AVCodecContext) Field_order() AVFieldOrder { //enum
	return (AVFieldOrder)(ctxt.field_order)
}

func (ctxt *AVCodecContext) Pix_fmt() AVPixelFormat { //enum
	return (AVPixelFormat)(ctxt.pix_fmt)
}

func (ctxt *AVCodecContext) Request_sample_fmt() AVSampleFormat { //enum
	return (AVSampleFormat)(ctxt.request_sample_fmt)
}

func (ctxt *AVCodecContext) Sample_fmt() AVSampleFormat { //enum
	return (AVSampleFormat)(ctxt.sample_fmt)
}

func (ctxt *AVCodecContext) Skip_frame() AVDiscard { //enum
	return (AVDiscard)(ctxt.skip_frame)
}

func (ctxt *AVCodecContext) Skip_idct() AVDiscard { //enum
	return (AVDiscard)(ctxt.skip_idct)
}

func (ctxt *AVCodecContext) Skip_loop_filter() AVDiscard { //enum
	return (AVDiscard)(ctxt.skip_loop_filter)
}

// func (ctxt *AVCodecContext) Sw_pix_fmt() AVPixelFormat { //enum
// 	return (AVPixelFormat)(ctxt.sw_pix_fmt)
// }
