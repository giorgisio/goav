/*
	Use of this source code is governed by a MIT license that can be found in the LICENSE file.
*/
package avcodec

func (ctxt *AvCodecContext) ActiveThreadType() int {
	return int(ctxt.active_thread_type)
}

func (ctxt *AvCodecContext) BFrameStrategy() int {
	return int(ctxt.b_frame_strategy)
}

func (ctxt *AvCodecContext) BQuantFactor() float64 {
	return float64(ctxt.b_quant_factor)
}

func (ctxt *AvCodecContext) BQuantOffset() float64 {
	return float64(ctxt.b_quant_offset)
}

func (ctxt *AvCodecContext) BSensitivity() int {
	return int(ctxt.b_sensitivity)
}

func (ctxt *AvCodecContext) BidirRefine() int {
	return int(ctxt.bidir_refine)
}

func (ctxt *AvCodecContext) BitRate() int {
	return int(ctxt.bit_rate)
}

func (ctxt *AvCodecContext) BitRateTolerance() int {
	return int(ctxt.bit_rate_tolerance)
}

func (ctxt *AvCodecContext) BitsPerCodedSample() int {
	return int(ctxt.bits_per_coded_sample)
}

func (ctxt *AvCodecContext) BitsPerRawSample() int {
	return int(ctxt.bits_per_raw_sample)
}

func (ctxt *AvCodecContext) BlockAlign() int {
	return int(ctxt.block_align)
}

func (ctxt *AvCodecContext) BrdScale() int {
	return int(ctxt.brd_scale)
}

func (ctxt *AvCodecContext) Channels() int {
	return int(ctxt.channels)
}

func (ctxt *AvCodecContext) Chromaoffset() int {
	return int(ctxt.chromaoffset)
}

func (ctxt *AvCodecContext) CodedHeight() int {
	return int(ctxt.coded_height)
}

func (ctxt *AvCodecContext) CodedWidth() int {
	return int(ctxt.coded_width)
}

func (ctxt *AvCodecContext) CoderType() int {
	return int(ctxt.coder_type)
}

func (ctxt *AvCodecContext) CompressionLevel() int {
	return int(ctxt.compression_level)
}

func (ctxt *AvCodecContext) ContextModel() int {
	return int(ctxt.context_model)
}

func (ctxt *AvCodecContext) Cutoff() int {
	return int(ctxt.cutoff)
}

func (ctxt *AvCodecContext) DarkMasking() float64 {
	return float64(ctxt.dark_masking)
}

func (ctxt *AvCodecContext) DctAlgo() int {
	return int(ctxt.dct_algo)
}

func (ctxt *AvCodecContext) Debug() int {
	return int(ctxt.debug)
}

func (ctxt *AvCodecContext) DebugMv() int {
	return int(ctxt.debug_mv)
}

func (ctxt *AvCodecContext) Delay() int {
	return int(ctxt.delay)
}

func (ctxt *AvCodecContext) DiaSize() int {
	return int(ctxt.dia_size)
}

func (ctxt *AvCodecContext) ErrRecognition() int {
	return int(ctxt.err_recognition)
}

func (ctxt *AvCodecContext) ErrorConcealment() int {
	return int(ctxt.error_concealment)
}

func (ctxt *AvCodecContext) ExtradataSize() int {
	return int(ctxt.extradata_size)
}

func (ctxt *AvCodecContext) Flags() int {
	return int(ctxt.flags)
}

func (ctxt *AvCodecContext) Flags2() int {
	return int(ctxt.flags2)
}

func (ctxt *AvCodecContext) FrameBits() int {
	return int(ctxt.frame_bits)
}

func (ctxt *AvCodecContext) FrameNumber() int {
	return int(ctxt.frame_number)
}

func (ctxt *AvCodecContext) FrameSize() int {
	return int(ctxt.frame_size)
}

func (ctxt *AvCodecContext) FrameSkipCmp() int {
	return int(ctxt.frame_skip_cmp)
}

func (ctxt *AvCodecContext) FrameSkipExp() int {
	return int(ctxt.frame_skip_exp)
}

func (ctxt *AvCodecContext) FrameSkipFactor() int {
	return int(ctxt.frame_skip_factor)
}

func (ctxt *AvCodecContext) FrameSkipThreshold() int {
	return int(ctxt.frame_skip_threshold)
}

func (ctxt *AvCodecContext) GlobalQuality() int {
	return int(ctxt.global_quality)
}

func (ctxt *AvCodecContext) GopSize() int {
	return int(ctxt.gop_size)
}

func (ctxt *AvCodecContext) HasBFrames() int {
	return int(ctxt.has_b_frames)
}

func (ctxt *AvCodecContext) HeaderBits() int {
	return int(ctxt.header_bits)
}

func (ctxt *AvCodecContext) Height() int {
	return int(ctxt.height)
}

func (ctxt *AvCodecContext) ICount() int {
	return int(ctxt.i_count)
}

func (ctxt *AvCodecContext) IQuantFactor() float64 {
	return float64(ctxt.i_quant_factor)
}

func (ctxt *AvCodecContext) IQuantOffset() float64 {
	return float64(ctxt.i_quant_offset)
}

func (ctxt *AvCodecContext) ITexBits() int {
	return int(ctxt.i_tex_bits)
}

func (ctxt *AvCodecContext) IdctAlgo() int {
	return int(ctxt.idct_algo)
}

func (ctxt *AvCodecContext) IldctCmp() int {
	return int(ctxt.ildct_cmp)
}

func (ctxt *AvCodecContext) IntraDcPrecision() int {
	return int(ctxt.intra_dc_precision)
}

func (ctxt *AvCodecContext) KeyintMin() int {
	return int(ctxt.keyint_min)
}

func (ctxt *AvCodecContext) LastPredictorCount() int {
	return int(ctxt.last_predictor_count)
}

func (ctxt *AvCodecContext) Level() int {
	return int(ctxt.level)
}

func (ctxt *AvCodecContext) LogLevelOffset() int {
	return int(ctxt.log_level_offset)
}

func (ctxt *AvCodecContext) Lowres() int {
	return int(ctxt.lowres)
}

func (ctxt *AvCodecContext) LumiMasking() float64 {
	return float64(ctxt.lumi_masking)
}

func (ctxt *AvCodecContext) MaxBFrames() int {
	return int(ctxt.max_b_frames)
}

func (ctxt *AvCodecContext) MaxPredictionOrder() int {
	return int(ctxt.max_prediction_order)
}

func (ctxt *AvCodecContext) MaxQdiff() int {
	return int(ctxt.max_qdiff)
}

func (ctxt *AvCodecContext) MbCmp() int {
	return int(ctxt.mb_cmp)
}

func (ctxt *AvCodecContext) MbDecision() int {
	return int(ctxt.mb_decision)
}

func (ctxt *AvCodecContext) MbLmax() int {
	return int(ctxt.mb_lmax)
}

func (ctxt *AvCodecContext) MbLmin() int {
	return int(ctxt.mb_lmin)
}

func (ctxt *AvCodecContext) MeCmp() int {
	return int(ctxt.me_cmp)
}

func (ctxt *AvCodecContext) MePenaltyCompensation() int {
	return int(ctxt.me_penalty_compensation)
}

func (ctxt *AvCodecContext) MePreCmp() int {
	return int(ctxt.me_pre_cmp)
}

func (ctxt *AvCodecContext) MeRange() int {
	return int(ctxt.me_range)
}

func (ctxt *AvCodecContext) MeSubCmp() int {
	return int(ctxt.me_sub_cmp)
}

func (ctxt *AvCodecContext) MeSubpelQuality() int {
	return int(ctxt.me_subpel_quality)
}

func (ctxt *AvCodecContext) MinPredictionOrder() int {
	return int(ctxt.min_prediction_order)
}

func (ctxt *AvCodecContext) MiscBits() int {
	return int(ctxt.misc_bits)
}

func (ctxt *AvCodecContext) MpegQuant() int {
	return int(ctxt.mpeg_quant)
}

func (ctxt *AvCodecContext) Mv0Threshold() int {
	return int(ctxt.mv0_threshold)
}

func (ctxt *AvCodecContext) MvBits() int {
	return int(ctxt.mv_bits)
}

func (ctxt *AvCodecContext) NoiseReduction() int {
	return int(ctxt.noise_reduction)
}

func (ctxt *AvCodecContext) NsseWeight() int {
	return int(ctxt.nsse_weight)
}

func (ctxt *AvCodecContext) PCount() int {
	return int(ctxt.p_count)
}

func (ctxt *AvCodecContext) PMasking() float64 {
	return float64(ctxt.p_masking)
}

func (ctxt *AvCodecContext) PTexBits() int {
	return int(ctxt.p_tex_bits)
}

func (ctxt *AvCodecContext) PreDiaSize() int {
	return int(ctxt.pre_dia_size)
}

func (ctxt *AvCodecContext) PreMe() int {
	return int(ctxt.pre_me)
}

func (ctxt *AvCodecContext) PredictionMethod() int {
	return int(ctxt.prediction_method)
}

func (ctxt *AvCodecContext) Profile() int {
	return int(ctxt.profile)
}

func (ctxt *AvCodecContext) Qblur() float64 {
	return float64(ctxt.qblur)
}

func (ctxt *AvCodecContext) Qcompress() float64 {
	return float64(ctxt.qcompress)
}

func (ctxt *AvCodecContext) Qmax() int {
	return int(ctxt.qmax)
}

func (ctxt *AvCodecContext) Qmin() int {
	return int(ctxt.qmin)
}

func (ctxt *AvCodecContext) RcBufferSize() int {
	return int(ctxt.rc_buffer_size)
}

func (ctxt *AvCodecContext) RcInitialBufferOccupancy() int {
	return int(ctxt.rc_initial_buffer_occupancy)
}

func (ctxt *AvCodecContext) RcMaxAvailableVbvUse() float64 {
	return float64(ctxt.rc_max_available_vbv_use)
}

func (ctxt *AvCodecContext) RcMaxRate() int {
	return int(ctxt.rc_max_rate)
}

func (ctxt *AvCodecContext) RcMinRate() int {
	return int(ctxt.rc_min_rate)
}

func (ctxt *AvCodecContext) RcMinVbvOverflowUse() float64 {
	return float64(ctxt.rc_min_vbv_overflow_use)
}

func (ctxt *AvCodecContext) RcOverrideCount() int {
	return int(ctxt.rc_override_count)
}

func (ctxt *AvCodecContext) RefcountedFrames() int {
	return int(ctxt.refcounted_frames)
}

func (ctxt *AvCodecContext) Refs() int {
	return int(ctxt.refs)
}

func (ctxt *AvCodecContext) RtpPayloadSize() int {
	return int(ctxt.rtp_payload_size)
}

func (ctxt *AvCodecContext) SampleRate() int {
	return int(ctxt.sample_rate)
}

func (ctxt *AvCodecContext) ScenechangeThreshold() int {
	return int(ctxt.scenechange_threshold)
}

func (ctxt *AvCodecContext) SeekPreroll() int {
	return int(ctxt.seek_preroll)
}

func (ctxt *AvCodecContext) SideDataOnlyPackets() int {
	return int(ctxt.side_data_only_packets)
}

func (ctxt *AvCodecContext) SkipAlpha() int {
	return int(ctxt.skip_alpha)
}

func (ctxt *AvCodecContext) SkipBottom() int {
	return int(ctxt.skip_bottom)
}

func (ctxt *AvCodecContext) SkipCount() int {
	return int(ctxt.skip_count)
}

func (ctxt *AvCodecContext) SkipTop() int {
	return int(ctxt.skip_top)
}

func (ctxt *AvCodecContext) SliceCount() int {
	return int(ctxt.slice_count)
}

func (ctxt *AvCodecContext) SliceFlags() int {
	return int(ctxt.slice_flags)
}

func (ctxt *AvCodecContext) Slices() int {
	return int(ctxt.slices)
}

func (ctxt *AvCodecContext) SpatialCplxMasking() float64 {
	return float64(ctxt.spatial_cplx_masking)
}

func (ctxt *AvCodecContext) StrictStdCompliance() int {
	return int(ctxt.strict_std_compliance)
}

func (ctxt *AvCodecContext) SubCharencMode() int {
	return int(ctxt.sub_charenc_mode)
}

func (ctxt *AvCodecContext) SubtitleHeaderSize() int {
	return int(ctxt.subtitle_header_size)
}

func (ctxt *AvCodecContext) TemporalCplxMasking() float64 {
	return float64(ctxt.temporal_cplx_masking)
}

func (ctxt *AvCodecContext) ThreadCount() int {
	return int(ctxt.thread_count)
}

func (ctxt *AvCodecContext) ThreadSafeCallbacks() int {
	return int(ctxt.thread_safe_callbacks)
}

func (ctxt *AvCodecContext) ThreadType() int {
	return int(ctxt.thread_type)
}

func (ctxt *AvCodecContext) TicksPerFrame() int {
	return int(ctxt.ticks_per_frame)
}

func (ctxt *AvCodecContext) Trellis() int {
	return int(ctxt.trellis)
}

func (ctxt *AvCodecContext) Width() int {
	return int(ctxt.width)
}

func (ctxt *AvCodecContext) WorkaroundBugs() int {
	return int(ctxt.workaround_bugs)
}

func (ctxt *AvCodecContext) AudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(ctxt.audio_service_type)
}

func (ctxt *AvCodecContext) ChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(ctxt.chroma_sample_location)
}

func (ctxt *AvCodecContext) CodecDescriptor() *AvCodecDescriptor {
	return (*AvCodecDescriptor)(ctxt.codec_descriptor)
}

func (ctxt *AvCodecContext) CodecId() AvCodecId {
	return (AvCodecId)(ctxt.codec_id)
}

func (ctxt *AvCodecContext) CodecType() AvMediaType {
	return (AvMediaType)(ctxt.codec_type)
}

func (ctxt *AvCodecContext) ColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(ctxt.color_primaries)
}

func (ctxt *AvCodecContext) ColorRange() AvColorRange {
	return (AvColorRange)(ctxt.color_range)
}

func (ctxt *AvCodecContext) ColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(ctxt.color_trc)
}

func (ctxt *AvCodecContext) Colorspace() AvColorSpace {
	return (AvColorSpace)(ctxt.colorspace)
}

func (ctxt *AvCodecContext) FieldOrder() AvFieldOrder {
	return (AvFieldOrder)(ctxt.field_order)
}

func (ctxt *AvCodecContext) PixFmt() AvPixelFormat {
	return (AvPixelFormat)(ctxt.pix_fmt)
}

func (ctxt *AvCodecContext) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctxt.request_sample_fmt)
}

func (ctxt *AvCodecContext) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctxt.sample_fmt)
}

func (ctxt *AvCodecContext) SkipFrame() AvDiscard {
	return (AvDiscard)(ctxt.skip_frame)
}

func (ctxt *AvCodecContext) SkipIdct() AvDiscard {
	return (AvDiscard)(ctxt.skip_idct)
}

func (ctxt *AvCodecContext) SkipLoopFilter() AvDiscard {
	return (AvDiscard)(ctxt.skip_loop_filter)
}
