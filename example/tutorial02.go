package main

// tutorial01.c
// Code based on a tutorial by Martin Bohme (boehme@inb.uni-luebeckREMOVETHIS.de)
// Tested on Gentoo, CVS version 5/01/07 compiled with GCC 4.1.1
// With updates from https://github.com/chelyaev/ffmpeg-tutorial
// Updates tested on:
// LAVC 54.59.100, LAVF 54.29.104, LSWS 2.1.101
// on GCC 4.7.2 in Debian February 2015

// A small sample program that shows how to use libavformat and libavcodec to
// read video from a file.
//
// Use
//
// gcc -o tutorial01 tutorial01.c -lavformat -lavcodec -lswscale -lz
//
// to build (assuming libavformat and libavcodec are correctly installed
// your system).
//
// Run using
//
// tutorial01 myvideofile.mpg
//
// to write the first five frames from "myvideofile.mpg" to disk in PPM
// format.
import (
	"fmt"
	"log"
	"os"
	"unsafe"

	"github.com/ampsight/goav/swscale"

	"github.com/ampsight/goav/avcodec"
	"github.com/ampsight/goav/avformat"
	"github.com/ampsight/goav/avutil"
)

func SaveFrame(pFrame *avutil.Frame, width, height, iFrame int) {
	var szFilename string
	var y int
	var file *os.File
	var err error

	szFilename = ""

	// Open file
	szFilename = fmt.Sprintf("frame%d.ppm", iFrame)

	if file, err = os.Create(szFilename); err != nil {
		log.Println("Error Reading")
	}
	defer file.Close()

	// Write header
	fh := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
	file.Write([]byte(fh))

	// Write pixel data
	for y = 0; y < height; y++ {
		// d := avutil.Data(videoFrame)
		// l := avutil.Linesize(videoFrame)
		//##TODO
		data0 := avutil.Data(pFrame)[0]
		buf := make([]byte, width*3)
		startPos := uintptr(unsafe.Pointer(data0)) + uintptr(y)*uintptr(avutil.Linesize(pFrame)[0])
		//fmt.Printf("%d + %d*%d=%d\n", uintptr(unsafe.Pointer(data0)), uintptr(y), uintptr(avutil.Linesize(videoFrame)[0]), startPos)
		for i := 0; i < width*3; i++ {
			element := *(*uint8)(unsafe.Pointer(startPos + uintptr(i)))
			buf[i] = element
			//fmt.Printf("Set element %d\n", i)
		}
		file.Write(buf)
	}
	// FILE *pFile;
	// char szFilename[32];
	// int  y;

	// // Open file
	// sprintf(szFilename, "frame%d.ppm", iFrame);
	// pFile=fopen(szFilename, "wb");
	// if(pFile==nil)
	//   return;

	// // Write header
	// fprintf(pFile, "P6\n%d %d\n255\n", width, height);

	// // Write pixel data
	// for(y=0; y<height; y++)
	//   fwrite(pFrame->data[0]+y*pFrame->linesize[0], 1, width*3, pFile);

	// // Close file
	// fclose(pFile);
}

func main() {
	// Initalizing these to nil prevents segfaults!
	// AVFormatContext   *pFormatCtx = nil;
	// int               i, videoStream;
	// AVCodecContext    *pCodecCtxOrig = nil;
	// AVCodecContext    *pCodecCtx = nil;
	// AVCodec           *pCodec = nil;
	// AVFrame           *pFrame = nil;
	// AVFrame           *pFrameRGB = nil;
	// AVPacket          packet;
	// int               frameFinished;
	// int               numBytes;
	// uint8_t           *buffer = nil;
	// struct SwsContext *sws_ctx = nil;

	// if(argc < 2) {
	//   printf("Please provide a movie file\n");
	//   return -1;
	// }
	// // Register all formats and codecs
	// av_register_all();

	// Open video file
	pFormatContext := avformat.AvformatAllocContext()
	if avformat.AvformatOpenInput(&pFormatContext, "sample.mp4", nil, nil) != 0 {
		return //-1 // Couldn't open file
	}

	// Retrieve stream information
	if pFormatContext.AvformatFindStreamInfo(nil) < 0 {
		return //-1 // Couldn't find stream information
	}

	// Dump information about file onto standard error
	//pFormatCtx.AvDumpFormat, 0, argv[1], 0)

	// Find the first video stream
	videoStream := -1
	for i := 0; i < int(pFormatContext.NbStreams()); i++ {
		if pFormatContext.Streams()[i].CodecParameters().AvCodecGetType() == avformat.AVMEDIA_TYPE_VIDEO {
			videoStream = i
		} else {
			break
		}
		if videoStream == -1 {
			return //-1 // Didn't find a video stream
		}

		// Get a pointer to the codec context for the video stream
		pCodecCtxOrig := pFormatContext.Streams()[videoStream].Codec()
		// Find the decoder for the video stream
		pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(pCodecCtxOrig.GetCodecId()))
		if pCodec == nil {
			fmt.Println("Unsupported codec!")
			return //-1 // Codec not found
		}
		// Copy context
		pCodecCtx := pCodec.AvcodecAllocContext3()
		if pCodecCtx.AvcodecCopyContext((*avcodec.Context)(unsafe.Pointer(pCodecCtxOrig))) != 0 {
			fmt.Println("Couldn't copy codec context")
			return //-1 // Error copying codec context
		}

		// Open codec
		if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
			return //-1; // Could not open codec
		}

		// Allocate video frame
		pFrame := avutil.AvFrameAlloc()

		// Allocate an AVFrame structure
		pFrameRGB := avutil.AvFrameAlloc()
		if pFrameRGB == nil {
			return //-1;
		}

		// Determine required buffer size and allocate buffer
		numBytes := uintptr(avcodec.AvpictureGetSize(avcodec.AV_PIX_FMT_RGB24, pCodecCtx.Width(),
			pCodecCtx.Height()))
		buffer := avutil.AvMalloc(numBytes)

		// Assign appropriate parts of buffer to image planes in pFrameRGB
		// Note that pFrameRGB is an AVFrame, but AVFrame is a superset
		// of AVPicture
		avp := (*avcodec.Picture)(unsafe.Pointer(pFrameRGB))
		avp.AvpictureFill((*uint8)(buffer), avcodec.AV_PIX_FMT_RGB24, pCodecCtx.Width(), pCodecCtx.Height())

		// initialize SWS context for software scaling
		sws_ctx := swscale.SwsGetcontext(pCodecCtx.Width(),
			pCodecCtx.Height(),
			(swscale.PixelFormat)(pCodecCtx.PixFmt()),
			pCodecCtx.Width(),
			pCodecCtx.Height(),
			avcodec.AV_PIX_FMT_RGB24,
			avcodec.SWS_BILINEAR,
			nil,
			nil,
			nil,
		)

		// Read frames and save first five frames to disk
		i = 0
		packet := avcodec.AvPacketAlloc()
		for pFormatContext.AvReadFrame(packet) >= 0 {
			// Is this a packet from the video stream?
			if packet.StreamIndex() == videoStream {
				// Decode video frame
				response := pCodecCtx.AvcodecSendPacket(packet)
				if response < 0 {
					fmt.Printf("Error while sending a packet to the decoder: %s\n", avutil.ErrorFromCode(response))
				}
				for response >= 0 {
					response = pCodecCtx.AvcodecReceiveFrame((*avcodec.Frame)(unsafe.Pointer(pFrame)))
					fmt.Printf("Receive frame response: %d\n", response)
					if response == avutil.AvErrorEAGAIN || response == avutil.AvErrorEOF {
						break
					} else if response < 0 {
						fmt.Printf("Error while receiving a frame from the decoder: %s\n", avutil.ErrorFromCode(response))
						return
					}

					// Convert the image from its native format to RGB
					swscale.SwsScale2(sws_ctx, avutil.Data(pFrame),
						avutil.Linesize(pFrame), 0, pCodecCtx.Height(),
						avutil.Data(pFrameRGB), avutil.Linesize(pFrameRGB))

					// Save the frame to disk
					i++
					if i <= 5 {
						SaveFrame(pFrameRGB, pCodecCtx.Width(), pCodecCtx.Height(),
							i)
					}
				}
			}

			// Free the packet that was allocated by av_read_frame
			packet.AvFreePacket()
		}

		// Free the RGB image
		avutil.AvFree(buffer)
		avutil.AvFrameFree(pFrameRGB)

		// Free the YUV frame
		avutil.AvFrameFree(pFrame)

		// Close the codecs
		pCodecCtx.AvcodecClose()
		(*avcodec.Context)(unsafe.Pointer(pCodecCtxOrig)).AvcodecClose()

		// Close the video file
		pFormatContext.AvformatCloseInput()
	}
}
