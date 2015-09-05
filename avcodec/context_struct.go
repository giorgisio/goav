// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

func (ctxt *Context) ActiveThreadType() int {
	return int(ctxt.active_thread_type)
}

func (ctxt *Context) BFrameStrategy() int {
	return int(ctxt.b_frame_strategy)
}

func (ctxt *Context) BQuantFactor() float64 {
	return float64(ctxt.b_quant_factor)
}

func (ctxt *Context) BQuantOffset() float64 {
	return float64(ctxt.b_quant_offset)
}

func (ctxt *Context) BSensitivity() int {
	return int(ctxt.b_sensitivity)
}

func (ctxt *Context) BidirRefine() int {
	return int(ctxt.bidir_refine)
}

func (ctxt *Context) BitRate() int {
	return int(ctxt.bit_rate)
}

func (ctxt *Context) BitRateTolerance() int {
	return int(ctxt.bit_rate_tolerance)
}

func (ctxt *Context) BitsPerCodedSample() int {
	return int(ctxt.bits_per_coded_sample)
}

func (ctxt *Context) BitsPerRawSample() int {
	return int(ctxt.bits_per_raw_sample)
}

func (ctxt *Context) BlockAlign() int {
	return int(ctxt.block_align)
}

func (ctxt *Context) BrdScale() int {
	return int(ctxt.brd_scale)
}

func (ctxt *Context) Channels() int {
	return int(ctxt.channels)
}

func (ctxt *Context) Chromaoffset() int {
	return int(ctxt.chromaoffset)
}

func (ctxt *Context) CodedHeight() int {
	return int(ctxt.coded_height)
}

func (ctxt *Context) CodedWidth() int {
	return int(ctxt.coded_width)
}

func (ctxt *Context) CoderType() int {
	return int(ctxt.coder_type)
}

func (ctxt *Context) CompressionLevel() int {
	return int(ctxt.compression_level)
}

func (ctxt *Context) ContextModel() int {
	return int(ctxt.context_model)
}

func (ctxt *Context) Cutoff() int {
	return int(ctxt.cutoff)
}

func (ctxt *Context) DarkMasking() float64 {
	return float64(ctxt.dark_masking)
}

func (ctxt *Context) DctAlgo() int {
	return int(ctxt.dct_algo)
}

func (ctxt *Context) Debug() int {
	return int(ctxt.debug)
}

func (ctxt *Context) DebugMv() int {
	return int(ctxt.debug_mv)
}

func (ctxt *Context) Delay() int {
	return int(ctxt.delay)
}

func (ctxt *Context) DiaSize() int {
	return int(ctxt.dia_size)
}

func (ctxt *Context) ErrRecognition() int {
	return int(ctxt.err_recognition)
}

func (ctxt *Context) ErrorConcealment() int {
	return int(ctxt.error_concealment)
}

func (ctxt *Context) ExtradataSize() int {
	return int(ctxt.extradata_size)
}

func (ctxt *Context) Flags() int {
	return int(ctxt.flags)
}

func (ctxt *Context) Flags2() int {
	return int(ctxt.flags2)
}

func (ctxt *Context) FrameBits() int {
	return int(ctxt.frame_bits)
}

func (ctxt *Context) FrameNumber() int {
	return int(ctxt.frame_number)
}

func (ctxt *Context) FrameSize() int {
	return int(ctxt.frame_size)
}

func (ctxt *Context) FrameSkipCmp() int {
	return int(ctxt.frame_skip_cmp)
}

func (ctxt *Context) FrameSkipExp() int {
	return int(ctxt.frame_skip_exp)
}

func (ctxt *Context) FrameSkipFactor() int {
	return int(ctxt.frame_skip_factor)
}

func (ctxt *Context) FrameSkipThreshold() int {
	return int(ctxt.frame_skip_threshold)
}

func (ctxt *Context) GlobalQuality() int {
	return int(ctxt.global_quality)
}

func (ctxt *Context) GopSize() int {
	return int(ctxt.gop_size)
}

func (ctxt *Context) HasBFrames() int {
	return int(ctxt.has_b_frames)
}

func (ctxt *Context) HeaderBits() int {
	return int(ctxt.header_bits)
}

func (ctxt *Context) Height() int {
	return int(ctxt.height)
}

func (ctxt *Context) ICount() int {
	return int(ctxt.i_count)
}

func (ctxt *Context) IQuantFactor() float64 {
	return float64(ctxt.i_quant_factor)
}

func (ctxt *Context) IQuantOffset() float64 {
	return float64(ctxt.i_quant_offset)
}

func (ctxt *Context) ITexBits() int {
	return int(ctxt.i_tex_bits)
}

func (ctxt *Context) IdctAlgo() int {
	return int(ctxt.idct_algo)
}

func (ctxt *Context) IldctCmp() int {
	return int(ctxt.ildct_cmp)
}

func (ctxt *Context) IntraDcPrecision() int {
	return int(ctxt.intra_dc_precision)
}

func (ctxt *Context) KeyintMin() int {
	return int(ctxt.keyint_min)
}

func (ctxt *Context) LastPredictorCount() int {
	return int(ctxt.last_predictor_count)
}

func (ctxt *Context) Level() int {
	return int(ctxt.level)
}

func (ctxt *Context) LogLevelOffset() int {
	return int(ctxt.log_level_offset)
}

func (ctxt *Context) Lowres() int {
	return int(ctxt.lowres)
}

func (ctxt *Context) LumiMasking() float64 {
	return float64(ctxt.lumi_masking)
}

func (ctxt *Context) MaxBFrames() int {
	return int(ctxt.max_b_frames)
}

func (ctxt *Context) MaxPredictionOrder() int {
	return int(ctxt.max_prediction_order)
}

func (ctxt *Context) MaxQdiff() int {
	return int(ctxt.max_qdiff)
}

func (ctxt *Context) MbCmp() int {
	return int(ctxt.mb_cmp)
}

func (ctxt *Context) MbDecision() int {
	return int(ctxt.mb_decision)
}

func (ctxt *Context) MbLmax() int {
	return int(ctxt.mb_lmax)
}

func (ctxt *Context) MbLmin() int {
	return int(ctxt.mb_lmin)
}

func (ctxt *Context) MeCmp() int {
	return int(ctxt.me_cmp)
}

func (ctxt *Context) MePenaltyCompensation() int {
	return int(ctxt.me_penalty_compensation)
}

func (ctxt *Context) MePreCmp() int {
	return int(ctxt.me_pre_cmp)
}

func (ctxt *Context) MeRange() int {
	return int(ctxt.me_range)
}

func (ctxt *Context) MeSubCmp() int {
	return int(ctxt.me_sub_cmp)
}

func (ctxt *Context) MeSubpelQuality() int {
	return int(ctxt.me_subpel_quality)
}

func (ctxt *Context) MinPredictionOrder() int {
	return int(ctxt.min_prediction_order)
}

func (ctxt *Context) MiscBits() int {
	return int(ctxt.misc_bits)
}

func (ctxt *Context) MpegQuant() int {
	return int(ctxt.mpeg_quant)
}

func (ctxt *Context) Mv0Threshold() int {
	return int(ctxt.mv0_threshold)
}

func (ctxt *Context) MvBits() int {
	return int(ctxt.mv_bits)
}

func (ctxt *Context) NoiseReduction() int {
	return int(ctxt.noise_reduction)
}

func (ctxt *Context) NsseWeight() int {
	return int(ctxt.nsse_weight)
}

func (ctxt *Context) PCount() int {
	return int(ctxt.p_count)
}

func (ctxt *Context) PMasking() float64 {
	return float64(ctxt.p_masking)
}

func (ctxt *Context) PTexBits() int {
	return int(ctxt.p_tex_bits)
}

func (ctxt *Context) PreDiaSize() int {
	return int(ctxt.pre_dia_size)
}

func (ctxt *Context) PreMe() int {
	return int(ctxt.pre_me)
}

func (ctxt *Context) PredictionMethod() int {
	return int(ctxt.prediction_method)
}

func (ctxt *Context) Profile() int {
	return int(ctxt.profile)
}

func (ctxt *Context) Qblur() float64 {
	return float64(ctxt.qblur)
}

func (ctxt *Context) Qcompress() float64 {
	return float64(ctxt.qcompress)
}

func (ctxt *Context) Qmax() int {
	return int(ctxt.qmax)
}

func (ctxt *Context) Qmin() int {
	return int(ctxt.qmin)
}

func (ctxt *Context) RcBufferSize() int {
	return int(ctxt.rc_buffer_size)
}

func (ctxt *Context) RcInitialBufferOccupancy() int {
	return int(ctxt.rc_initial_buffer_occupancy)
}

func (ctxt *Context) RcMaxAvailableVbvUse() float64 {
	return float64(ctxt.rc_max_available_vbv_use)
}

func (ctxt *Context) RcMaxRate() int {
	return int(ctxt.rc_max_rate)
}

func (ctxt *Context) RcMinRate() int {
	return int(ctxt.rc_min_rate)
}

func (ctxt *Context) RcMinVbvOverflowUse() float64 {
	return float64(ctxt.rc_min_vbv_overflow_use)
}

func (ctxt *Context) RcOverrideCount() int {
	return int(ctxt.rc_override_count)
}

func (ctxt *Context) RefcountedFrames() int {
	return int(ctxt.refcounted_frames)
}

func (ctxt *Context) Refs() int {
	return int(ctxt.refs)
}

func (ctxt *Context) RtpPayloadSize() int {
	return int(ctxt.rtp_payload_size)
}

func (ctxt *Context) SampleRate() int {
	return int(ctxt.sample_rate)
}

func (ctxt *Context) ScenechangeThreshold() int {
	return int(ctxt.scenechange_threshold)
}

func (ctxt *Context) SeekPreroll() int {
	return int(ctxt.seek_preroll)
}

func (ctxt *Context) SideDataOnlyPackets() int {
	return int(ctxt.side_data_only_packets)
}

func (ctxt *Context) SkipAlpha() int {
	return int(ctxt.skip_alpha)
}

func (ctxt *Context) SkipBottom() int {
	return int(ctxt.skip_bottom)
}

func (ctxt *Context) SkipCount() int {
	return int(ctxt.skip_count)
}

func (ctxt *Context) SkipTop() int {
	return int(ctxt.skip_top)
}

func (ctxt *Context) SliceCount() int {
	return int(ctxt.slice_count)
}

func (ctxt *Context) SliceFlags() int {
	return int(ctxt.slice_flags)
}

func (ctxt *Context) Slices() int {
	return int(ctxt.slices)
}

func (ctxt *Context) SpatialCplxMasking() float64 {
	return float64(ctxt.spatial_cplx_masking)
}

func (ctxt *Context) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

func (ctxt *Context) SubCharencMode() int {
	return int(ctxt.sub_charenc_mode)
}

func (ctxt *Context) SubtitleHeaderSize() int {
	return int(ctxt.subtitle_header_size)
}

func (ctxt *Context) TemporalCplxMasking() float64 {
	return float64(ctxt.temporal_cplx_masking)
}

func (ctxt *Context) ThreadCount() int {
	return int(ctxt.thread_count)
}

func (ctxt *Context) ThreadSafeCallbacks() int {
	return int(ctxt.thread_safe_callbacks)
}

func (ctxt *Context) ThreadType() int {
	return int(ctxt.thread_type)
}

func (ctxt *Context) TicksPerFrame() int {
	return int(ctxt.ticks_per_frame)
}

func (ctxt *Context) Trellis() int {
	return int(ctxt.trellis)
}

func (ctxt *Context) Width() int {
	return int(ctxt.width)
}

func (ctxt *Context) WorkaroundBugs() int {
	return int(ctxt.workaround_bugs)
}

func (ctxt *Context) AudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(ctxt.audio_service_type)
}

func (ctxt *Context) ChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(ctxt.chroma_sample_location)
}

func (ctxt *Context) CodecDescriptor() *Descriptor {
	return (*Descriptor)(ctxt.codec_descriptor)
}

func (ctxt *Context) CodecId() CodecId {
	return (CodecId)(ctxt.codec_id)
}

func (ctxt *Context) CodecType() MediaType {
	return (MediaType)(ctxt.codec_type)
}

func (ctxt *Context) ColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(ctxt.color_primaries)
}

func (ctxt *Context) ColorRange() AvColorRange {
	return (AvColorRange)(ctxt.color_range)
}

func (ctxt *Context) ColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(ctxt.color_trc)
}

func (ctxt *Context) Colorspace() AvColorSpace {
	return (AvColorSpace)(ctxt.colorspace)
}

func (ctxt *Context) FieldOrder() AvFieldOrder {
	return (AvFieldOrder)(ctxt.field_order)
}

func (ctxt *Context) PixFmt() PixelFormat {
	return (PixelFormat)(ctxt.pix_fmt)
}

func (ctxt *Context) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctxt.request_sample_fmt)
}

func (ctxt *Context) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctxt.sample_fmt)
}

func (ctxt *Context) SkipFrame() AvDiscard {
	return (AvDiscard)(ctxt.skip_frame)
}

func (ctxt *Context) SkipIdct() AvDiscard {
	return (AvDiscard)(ctxt.skip_idct)
}

func (ctxt *Context) SkipLoopFilter() AvDiscard {
	return (AvDiscard)(ctxt.skip_loop_filter)
}
