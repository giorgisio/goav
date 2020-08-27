package main

import (
	"fmt"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"log"
	"unsafe"
)

const OUTPUT_VIDEO_CODEC = "libx265"
const OUTPUT_CODEC_PRIV_KEY = "x265-params"
const OUTPUT_CODEC_PRIV_VALUE = "keyint=60:min-keyint=60:scenecut=0"

func CreateStream(formatContext *avformat.Context, inputCodecContext *avcodec.Context, framerate avcodec.Rational,
	codecName string, codecPrivKey string, codecPrivValue string) *Stream {
	avStream := formatContext.AvformatNewStream(nil)
	avCodec := avcodec.AvcodecFindEncoderByName(codecName)
	log.Printf("\tCodec %s ID %d ", avCodec.AvCodecGetName(), avCodec.AvCodecGetId())
	if avCodec == nil {
		log.Fatal("could not find the proper codec")
	}
	avCodecContext := avCodec.AvcodecAllocContext3()
	if avCodecContext == nil {
		log.Fatal("could not allocated memory for codec context")
	}
	avutil.AvOptSet(formatContext.AvFormatGetPrivData(), "present", "fast", 0)
	if codecPrivKey != "" && codecPrivValue != "" {
		avutil.AvOptSet(formatContext.AvFormatGetPrivData(), codecPrivKey, codecPrivValue, 0)
	}
	avCodecContext.SetHeight(inputCodecContext.Height())
	avCodecContext.SetWidth(inputCodecContext.Width())
	avCodecContext.SetSampleAspectRatio(inputCodecContext.SampleAspectRatio())
	if len(avCodec.GetPixFmts()) != 0 {
		avCodecContext.SetPixFmt(avCodec.GetPixFmts()[0])
	} else {
		avCodecContext.SetPixFmt(inputCodecContext.GetPixFmt())
	}
	avCodecContext.SetBitRate(2 * 1000 * 1000)
	avCodecContext.SetRcBufferSize(4 * 1000 * 1000)
	avCodecContext.SetRcMaxRate(2 * 1000 * 1000)
	avCodecContext.SetRcMinRate(2.5 * 1000 * 1000)

	timeBase := avcodec.AVInvQ(framerate)
	avCodecContext.SetTimebase(timeBase.Num(), timeBase.Den())
	avStream.SetTimeBase(timeBase)

	if avCodecContext.AvcodecOpen2(avCodec, nil) < 0 {
		log.Fatal("could not open the codec")
	}
	ret := avStream.CodecParameters().AvCodecParametersFromContext(avCodecContext)
	log.Print(ret)
	return &Stream{
		Codec:        avCodec,
		Stream:       avStream,
		CodecContext: avCodecContext,
	}
}

type Stream struct {
	Codec        *avcodec.Codec
	Stream       *avformat.Stream
	CodecContext *avcodec.Context
	CodecType    interface{}
	Index        int
}

func (s *Stream) FillStreamInfo(stream *avformat.Stream) {
	s.Stream = stream
	avCodecParameters := stream.CodecParameters()
	//defer avCodecParameters.AvCodecParametersFree()
	avCodec := avcodec.AvcodecFindDecoder(avCodecParameters.AvCodecGetId())
	if avCodec == nil {
		log.Fatalf("Error unsupported codec")
	}
	avCodecContext := avCodec.AvcodecAllocContext3()
	if avCodecContext == nil {
		log.Fatalf("failed to allocated memory for AVCodecContext")
	}

	if avCodecContext.AvcodecParametersToContext(avCodecParameters) < 0 {
		log.Fatalf("failed to copy codec params to codec context")
	}
	if avCodecContext.AvcodecOpen2(avCodec, nil) < 0 {
		log.Fatalf("failed to open codec through avcodec_open2")
	}
	s.Codec = avCodec
	s.CodecContext = avCodecContext
	log.Printf("\tCodec %s ID %d ", avCodec.AvCodecGetName(), avCodec.AvCodecGetId())
}

func (s *Stream) CloseStream() {
	s.CodecContext.AvcodecFreeContext()
}

func encodeVideo(outputContext *avformat.Context, inputStream *Stream, outputStream *Stream, inputFrame *avutil.Frame) {
	inputFrame.SetPicType(avutil.AV_PICTURE_TYPE_NONE)
	outputPacket := avcodec.AvPacketAlloc()
	if outputPacket == nil {
		log.Fatalf("failed to allocated memory for AVPacket")
		return
	}
	defer outputPacket.AVPacketFree()
	defer outputPacket.AvPacketUnref()

	response := outputStream.CodecContext.AvcodecSendFrame((*avcodec.Frame)(unsafe.Pointer(inputFrame)))

	for response >= 0 {
		response = outputStream.CodecContext.AvcodecReceivePacket(outputPacket)
		if response == avutil.AvErrorEAGAIN || response == avutil.AvErrorEOF {
			break
		} else if response < 0 {
			log.Fatalf("Error while receiving packet from encoder: %s", avutil.ErrorFromCode(response))
		}

		outputPacket.SetStreamIndex(inputStream.Index)
		outputPacket.SetDuration(outputStream.Stream.TimeBase().Den() / outputStream.Stream.TimeBase().Num() / inputStream.Stream.AvgFrameRate().Num() * inputStream.Stream.AvgFrameRate().Den())
		outputPacket.AvPacketRescaleTs(inputStream.Stream.TimeBase(), outputStream.Stream.TimeBase())
		if outputContext.AvInterleavedWriteFrame(outputPacket) < 0 {
			log.Fatal("Error muxing packet")
			return
		}
	}
	return
}

func transcodeVideo(outputContext *avformat.Context, inputStream *Stream, outputStream *Stream, inputPacket *avcodec.Packet, inputFrame *avutil.Frame) {
	response := inputStream.CodecContext.AvcodecSendPacket(inputPacket)
	if response < 0 {
		log.Fatalf("Error while sending packet to decoder: %s", avutil.ErrorFromCode(response))
	}
	for response >= 0 {
		response = inputStream.CodecContext.AvcodecReceiveFrame((*avcodec.Frame)(unsafe.Pointer(inputFrame)))
		if response == avutil.AvErrorEAGAIN || response == avutil.AvErrorEOF {
			break
		} else if response < 0 {
			fmt.Printf("Error while receiving a frame from the decoder: %s\n", avutil.ErrorFromCode(response))
			return
		}
		encodeVideo(outputContext, inputStream, outputStream, inputFrame)
	}
	avutil.AvFrameUnref(inputFrame)
}

func main() {
	inputFile := "/Users/jason12360/Desktop/study/ffmpeg/ffmpeg-libav-tutorial/bunny_1080p_60fps.mp4"
	outputFile := "/Users/jason12360/Desktop/goav/small_bunny_1080p_60fps.mp4"

	var inputFmtCtx *avformat.Context
	if avformat.AvformatOpenInput(&inputFmtCtx, inputFile, nil, nil) < 0 {
		log.Fatalf("Could not open input file '%s", inputFile)
		return
	}
	defer inputFmtCtx.AvformatFreeContext()
	if inputFmtCtx.AvformatFindStreamInfo(nil) < 0 {
		log.Fatalf("Failed to retrieve input stream information")
		return
	}

	var inputVideoStream = &Stream{
		CodecType: avformat.AVMEDIA_TYPE_VIDEO,
	}
	defer inputVideoStream.CloseStream()
	var inputAudioStream = &Stream{
		CodecType: avformat.AVMEDIA_TYPE_AUDIO,
	}
	defer inputAudioStream.CloseStream()
	for index, avStream := range inputFmtCtx.Streams() {
		localCodecParameters := avStream.CodecParameters()
		switch localCodecParameters.AvCodecGetType() {
		case avformat.AVMEDIA_TYPE_VIDEO:
			inputVideoStream.FillStreamInfo(avStream)
			inputVideoStream.Index = index
		case avformat.AVMEDIA_TYPE_AUDIO:
			inputAudioStream.FillStreamInfo(avStream)
			inputAudioStream.Index = index
		}
	}

	var outputFmtCtx *avformat.Context
	avformat.AvAllocOutputContext2(&outputFmtCtx, nil, nil, &outputFile)
	if outputFmtCtx == nil {
		log.Fatalf("Could not create output context")
		return
	}
	defer outputFmtCtx.AvformatFreeContext()

	//video
	inputFramerate := inputFmtCtx.AvGuessFrameRate(inputVideoStream.Stream, nil)
	outputVideoStream := CreateStream(outputFmtCtx, inputVideoStream.CodecContext, inputFramerate,
		OUTPUT_VIDEO_CODEC, OUTPUT_CODEC_PRIV_KEY, OUTPUT_CODEC_PRIV_VALUE)
	outputVideoStream.CodecType = avformat.AVMEDIA_TYPE_VIDEO
	if outputVideoStream == nil {
		log.Fatal("sss")
	}

	//audio

	outputAudioStream := outputFmtCtx.AvformatNewStream(nil)
	if outputAudioStream == nil {
		log.Fatalf("Failed allocating output stream")
		return
	}
	if inputAudioStream.Stream.CodecParameters().AvCodecParametersCopyTo(outputAudioStream.CodecParameters()) < 0 {
		log.Fatalf("Failed to copy codec parameters")
		return
	}
	//log.Printf("\tCodec %s ID %d ",outputAudioStream.CodecParameters().AvCodecGetId(),avcodec.AvcodecFindEncoder(outputAudioStream.CodecParameters().AvCodecGetId()).AvCodecGetName() )

	if outputFmtCtx.Oformat().GetFlags()&avformat.AVFMT_GLOBALHEADER == 1 {
		inputFmtCtx.Oformat().SetFlags(inputFmtCtx.Oformat().GetFlags() | avformat.AVFMT_GLOBALHEADER)
	}

	if outputFmtCtx.Oformat().GetFlags()&avformat.AVFMT_NOFILE == 0 {
		_, err := avformat.AvIOOpen(outputFile, avformat.AVIO_FLAG_WRITE)
		if err != nil {
			log.Fatalf("Could not open output file '%s'", outputFile)
			return
		}
	}
	//var opts *avutil.Dictionary = &avutil.Dictionary{}
	//defer opts.AvDictFree()
	//
	//opts.AvDictSet("movflags", "frag_keyframe+empty_moov+default_base_moof", 0)

	outputFmtCtx.AvDumpFormat(0, outputFile, 1)
	if outputFmtCtx.Oformat().GetFlags()&avformat.AVFMT_NOFILE == 0 {
		avIOContext, err := avformat.AvIOOpen(outputFile, avformat.AVIO_FLAG_WRITE)
		if err != nil {
			log.Fatalf("Could not open output file '%s'", outputFile)
			return
		}
		outputFmtCtx.SetPb(avIOContext)
	}



	if outputFmtCtx.AvformatWriteHeader(nil) < 0 {
		log.Fatal("Error occurred when opening output file")
		return
	}

	inputFrame := avutil.AvFrameAlloc()
	if inputFrame == nil {
		log.Fatalf("failed to allocated memory for AVFrame")
		return
	}
	defer avutil.AvFrameFree(inputFrame)

	inputPacket := avcodec.AvPacketAlloc()
	if inputPacket == nil {
		log.Fatalf("failed to allocated memory for AVPacket")
		return
	}
	defer inputPacket.AVPacketFree()

	for true {
		if inputFmtCtx.AvReadFrame(inputPacket) < 0 {
			break
		}
		if inputFmtCtx.Streams()[inputPacket.StreamIndex()].CodecParameters().AvCodecGetType() == avformat.AVMEDIA_TYPE_VIDEO {
			transcodeVideo(outputFmtCtx, inputVideoStream, outputVideoStream, inputPacket, inputFrame)
		} else if inputFmtCtx.Streams()[inputPacket.StreamIndex()].CodecParameters().AvCodecGetType() == avformat.AVMEDIA_TYPE_AUDIO {
			inputPacket.AvPacketRescaleTs(inputAudioStream.Stream.TimeBase(), outputAudioStream.TimeBase())
			if outputFmtCtx.AvInterleavedWriteFrame(inputPacket) < 0 {
				log.Fatal("Error muxing packet")

			}
		} else {
			log.Fatal("ignoring all non video or audio packets")
		}
		//inputPacket.AvPacketUnref()
	}
	outputFmtCtx.AvWriteTrailer()
	if outputFmtCtx.Oformat().GetFlags()&avformat.AVFMT_NOFILE == 0 {
		outputFmtCtx.Pb().Close()
	}
}
