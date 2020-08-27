//https://www.ffmpeg.org/doxygen/4.0/decode__video_8c_source.html
package main

import (
	"fmt"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avutil"
	"io"
	"log"
	"os"
	"unsafe"
)

const INBUF_SIZE = 4096

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

func main() {
	filename := "/Users/jason12360/Desktop/128x128.264/128x128.264"

	// Register all formats and codecs

	avCodec := avcodec.AvcodecFindDecoder((avcodec.CodecId)(avcodec.AV_CODEC_ID_H264))
	if avCodec == nil{
		log.Fatal("Codec not found")
	}
	parser := avcodec.AvParserInit(avCodec.AvCodecGetId())
	if parser==nil{
		log.Fatal("parser not found")
	}

	avCodecContext := avCodec.AvcodecAllocContext3()
	if avCodecContext == nil {
		log.Fatalf("failed to allocated memory for AVCodecContext")
		return
	}
	defer avCodecContext.AvcodecFreeContext()

	if avCodecContext.AvcodecOpen2(avCodec, nil) < 0 {
		log.Fatalf("failed to open avCodec through avcodec_open2")
		return
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

	mediaFile,err := os.Open(filename)
	if err != nil {
		log.Fatal("read fail")
	}
	defer mediaFile.Close()

	buf := make([]byte,INBUF_SIZE)

	for {
		n,err := mediaFile.Read(buf)
		if err != nil && err != io.EOF{
			log.Fatal("read buf fail",err)
		}
		if n ==0{
			break
		}
		data := buf

		for n >0{
			var avPacketData *byte
			var avPacketSize int
			ret := avCodecContext.AvParserParse2(parser,&avPacketData,&avPacketSize,&data[0],n,avutil.AV_NOPTS_VALUE,avutil.AV_NOPTS_VALUE,0)
			if ret<0{
				log.Fatal("Error while parsing")
			}
			data = data[ret:]
			n = n-ret
			if avPacketSize>0{
				avPacket.SetData(avPacketData)
				avPacket.SetSize(avPacketSize)
				decodePacket(avPacket, avCodecContext, (*avcodec.Frame)(unsafe.Pointer(avFrame)))
			}
		}

	}
}
