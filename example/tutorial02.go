package main

import (
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"log"
	"unsafe"
)

func main() {
	filename := "/Users/jason12360/Desktop/732/mediaCollector/video/bunny_1080p_60fps.mp4"

	// Register all formats and codecs

	ctx := avformat.AvformatAllocContext()

	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

	var avCodec *avcodec.Codec
	var avCodecParameters *avcodec.AvCodecParameters
	videoStreamIndex := -1
	for index, avStream := range ctx.Streams() {
		localCodecParameters := avStream.CodecParameters()
		log.Printf("AVStream->time_base before open coded %d/%d", avStream.TimeBase().Num(), avStream.TimeBase().Den())
		log.Printf("AVStream->r_frame_rate before open coded %d/%d", avStream.RFrameRate().Num(), avStream.RFrameRate().Den())
		log.Printf("AVStream->start_time %d", avStream.StartTime())
		log.Printf("AVStream->duration %d", avStream.Duration())

		log.Println("finding the proper decoder (CODEC)")

		localCodec := avcodec.AvcodecFindDecoder(localCodecParameters.AvCodecGetId())
		if localCodec == nil {
			log.Printf("Error unsupported codec")
			return
		}

		switch localCodecParameters.AvCodecGetType() {
		case avformat.AVMEDIA_TYPE_VIDEO:
			if videoStreamIndex == -1 {
				videoStreamIndex = index
				avCodec = localCodec
				avCodecParameters = localCodecParameters
			}
			log.Printf("Video Codec: resolution %d x %d", localCodecParameters.AvCodecGetWidth(), localCodecParameters.AvCodecGetHeight())
		case avformat.AVMEDIA_TYPE_AUDIO:
			log.Printf("Audio Codec: %d channels, sample rate %d", localCodecParameters.AvCodecGetChannels(), localCodecParameters.AvCodecGetSampleRate())
		}
		log.Printf("\tCodec %s ID %d bit_rate %d", localCodec.AvCodecGetName(), localCodec.AvCodecGetId(), localCodecParameters.AvCodecGetBitRate())
	}
	avCodecContext := avCodec.AvcodecAllocContext3()
	if avCodecContext == nil {
		log.Printf("failed to allocated memory for AVCodecContext")
		return
	}
	if avCodecContext.AvcodecParametersToContext(avCodecParameters) < 0 {
		log.Printf("failed to copy codec params to codec context")
		return
	}
	if avCodecContext.AvcodecOpen2(avCodec, nil) < 0 {
		log.Printf("failed to open codec through avcodec_open2")
		return
	}

	avFrame := avutil.AvFrameAlloc()
	if avFrame == nil {
		log.Printf("failed to allocated memory for AVFrame")
		return
	}

	avPacket := avcodec.AvPacketAlloc()
	if avPacket == nil {
		log.Printf("failed to allocated memory for AVPacket")
		return
	}
	timesLimit := 8
	for ctx.AvReadFrame(avPacket) >= 0 {
		if avPacket.StreamIndex() == videoStreamIndex {
			log.Printf("AVPacket->pts %d", avPacket.Pts())
			response := decodePacket(avPacket, avCodecContext, (*avcodec.Frame)(unsafe.Pointer(avFrame)))
			if response < 0 {
				break
			}
			timesLimit--
			if (timesLimit) < 0 {
				break
			}
		}
		avPacket.AvPacketUnref()
	}
	log.Printf("releasing all the resources")

	ctx.AvformatCloseInput()
	avPacket.AvFreePacket()
	avutil.AvFrameFree(avFrame)
	avCodecContext.AvcodecFreeContext()
	return

}


func decodePacket(avPacket *avcodec.Packet, avCodecContext *avcodec.Context, avFrame *avcodec.Frame) int {
	sendResponse := avCodecContext.AvcodecSendPacket(avPacket)
	if sendResponse < 0 {
		log.Printf("Error while sending a packet to the decoder: %s", avutil.ErrorFromCode(sendResponse))
		return sendResponse
	}
	for sendResponse >= 0 {
		receiveResponse := avCodecContext.AvcodecReceiveFrame(avFrame)
		if receiveResponse == avutil.AvErrorEAGAIN || receiveResponse == avutil.AvErrorEOF {
			break
		} else if receiveResponse < 0 {
			log.Printf("Error while receiving a frame from the decoder: %s", avutil.ErrorFromCode(receiveResponse))
			return receiveResponse
		}
		if receiveResponse >= 0 {
			log.Printf(
				"Frame %d (type=%s, size=%d bytes) pts %d key_frame %d [DTS %d]",
				avCodecContext.FrameNumber(),
				avutil.AvGetPictureTypeChar(avFrame.GetPicType()),
				avFrame.GetPktSize(),
				avFrame.GetPts(),
				avFrame.GetKeyFrame(),
				avFrame.GetCodedPictureNumber(),
			)
		}
	}
	return 1
}
