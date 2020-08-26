package main
// tutorial02.go
// Code based on a tutorial at https://github.com/leandromoreira/ffmpeg-libav-tutoria

// A small sample program that shows how to use libavformat and libavcodec to
// read video from a file.
//
// Use
//
//
// tutorial01 myvideofile.mpg
//
// to write the first eight frames from "small_bunny_1080p_60fps.mp4" to disk in pgm
// format.
import (
	"fmt"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"log"
	"os"
	"unsafe"
)

func main() {
	filename := "/Users/jason12360/Desktop/test.mp4"

	// Register all formats and codecs

	ctx := avformat.AvformatAllocContext()
	defer ctx.AvformatCloseInput()
	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Fatalf("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Fatalf("Error: Couldn't find stream information.")

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
			log.Fatalf("Error unsupported codec")
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

	avFrame := avutil.AvFrameAlloc()
	if avFrame == nil {
		log.Fatalf("failed to allocated memory for AVFrame")
		return
	}
	defer avutil.AvFrameFree(avFrame)

	avPacket := avcodec.AvPacketAlloc()
	if avPacket == nil {
		log.Fatalf("failed to allocated memory for AVPacket")
		return
	}
	defer avPacket.AVPacketFree()

	avCodecContext := avCodec.AvcodecAllocContext3()
	if avCodecContext == nil {
		log.Fatalf("failed to allocated memory for AVCodecContext")
		return
	}
	defer avCodecContext.AvcodecFreeContext()

	if avCodecContext.AvcodecParametersToContext(avCodecParameters) < 0 {
		log.Fatalf("failed to copy codec params to codec context")
		return
	}
	if avCodecContext.AvcodecOpen2(avCodec, nil) < 0 {
		log.Fatalf("failed to open codec through avcodec_open2")
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

	return

}

func decodePacket(avPacket *avcodec.Packet, avCodecContext *avcodec.Context, avFrame *avcodec.Frame) int {
	sendResponse := avCodecContext.AvcodecSendPacket(avPacket)
	if sendResponse < 0 {
		log.Fatalf("Error while sending a packet to the decoder: %s", avutil.ErrorFromCode(sendResponse))
		return sendResponse
	}
	for sendResponse >= 0 {
		receiveResponse := avCodecContext.AvcodecReceiveFrame(avFrame)
		if receiveResponse == avutil.AvErrorEAGAIN || receiveResponse == avutil.AvErrorEOF {
			break
		} else if receiveResponse < 0 {
			log.Fatalf("Error while receiving a frame from the decoder: %s", avutil.ErrorFromCode(receiveResponse))
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
			fileName := fmt.Sprintf("frame-%d.pgm", avCodecContext.FrameNumber())
			save_gray_frame(avFrame, fileName)
		}
	}
	return 1
}

func save_gray_frame(avFrame *avcodec.Frame, fileName string) {
	//width := avFrame.GetWidth()
	//height := avFrame.GetHeight()
	//fileName := fmt.Sprintf("frame%d.ppm", frameNumber)
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Error Reading")
		return
	}
	defer file.Close()

	header := fmt.Sprintf("P5\n%d %d\n255\n", avFrame.GetWidth(), avFrame.GetHeight())
	_,err =file.Write([]byte(header))
	if err != nil{
		log.Fatalf("can not write to file,the err is %v",err)
	}

	buffer := avFrame.GetData()
	_,err =file.Write(buffer[0])
	if err != nil{
		log.Fatalf("can not write to file,the err is %v",err)
	}
	return
}
