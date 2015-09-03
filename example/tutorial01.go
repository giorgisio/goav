package main

import (
	"fmt"
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"github.com/giorgisio/goav/swscale"
	"log"
	"os"
	"unsafe"
)

func main() {

	filename := "sample.mp4"

	var (
		pFormatCtx    *avformat.AVFormatContext
		pCodecCtxOrig *avcodec.AVCodecContext
		pCodecCtx     *avcodec.AVCodecContext
		pCodec        *avcodec.AVCodec
		pFrame        *avutil.AVFrame
		pFrameRGB     *avutil.AVFrame
		packet        *avcodec.AVPacket
		sws_ctx       *swscale.SwsContext
		videoStream   int
		frameFinished int
		numBytes      int
		url           string
	)
	//media_type    *avutil.AVMediaType

	// Register all formats and codecs
	avformat.Av_register_all()

	// Open video file
	if avformat.Avformat_open_input(&pFormatCtx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if avformat.Avformat_find_stream_info(pFormatCtx, nil) < 0 {
		log.Println("Error: Couldn't find stream information.")
		return
	}

	// Dump information about file onto standard error
	avformat.Av_dump_format(pFormatCtx, 0, url, 0)

	// Find the first video stream
	videoStream = -1

	//pFormatCtx->nb_streams
	n := pFormatCtx.Nb_streams()

	//pFormatCtx->streams[]
	s := pFormatCtx.Streams()
	//s2 := avformat.StreamsOne(pFormatCtx, 1)

	log.Print("Number of Streams:", n)

	for i := 0; i < int(n); i++ {
		// pFormatCtx->streams[i]->codec->codec_type
		log.Println("Stream Number:", i)

		//FIX: AVMEDIA_TYPE_VIDEO
		if (*avformat.AVCodecContext)(s.Codec()) != nil {
			videoStream = i
			break
		}
	}

	if videoStream == -1 {
		log.Println("Couldn't find a video stream")
		return
	}

	codec := s.Codec()

	// Get a pointer to the codec context for the video stream
	//pCodecCtxOrig = pFormatCtx.streams[videoStream].codec
	pCodecCtxOrig = (*avcodec.AVCodecContext)(unsafe.Pointer(&codec))
	log.Println("Bit Rate:", pCodecCtxOrig.Bit_rate())
	log.Println("Channels:", pCodecCtxOrig.Channels())
	log.Println("Coded_height:", pCodecCtxOrig.Coded_height())
	log.Println("Coded_width:", pCodecCtxOrig.Coded_width())
	log.Println("Coder_type:", pCodecCtxOrig.Coder_type())
	log.Println("Height:", pCodecCtxOrig.Height())
	log.Println("Profile:", pCodecCtxOrig.Profile())
	log.Println("Width:", pCodecCtxOrig.Width())
	log.Println("Codec ID:", pCodecCtxOrig.Codec_id())

	//C.enum_AVCodecID
	codec_id := pCodecCtxOrig.Codec_id()

	// Find the decoder for the video stream
	pCodec = avcodec.Avcodec_find_decoder(codec_id)
	if pCodec == nil {
		log.Println("Error: Unsupported codec!")
		return // Codec not found
	}

	// Copy context
	pCodecCtx = avcodec.Avcodec_alloc_context3(pCodec)

	if avcodec.Avcodec_copy_context(pCodecCtx, pCodecCtxOrig) != 0 {
		log.Println("Error: Couldn't copy codec context")
		return // Error copying codec context
	}

	// Open codec
	if avcodec.Avcodec_open2(pCodecCtx, pCodec, nil) < 0 {
		return // Could not open codec
	}

	// Allocate video frame
	pFrame = avutil.Av_frame_alloc()

	// Allocate an AVFrame structure
	if pFrameRGB = avutil.Av_frame_alloc(); pFrameRGB == nil {
		return
	}

	//##TODO
	var a swscale.AVPixelFormat
	var b int
	//avcodec.AVPixelFormat
	//avcodec.PIX_FMT_RGB24
	//avcodec.SWS_BILINEAR

	w := pCodecCtx.Width()
	h := pCodecCtx.Height()
	pix_fmt := pCodecCtx.Pix_fmt()

	// Determine required buffer size and allocate buffer
	numBytes = avcodec.Avpicture_get_size((avcodec.AVPixelFormat)(a), w, h)

	buffer := avutil.Av_malloc(uintptr(numBytes))

	// Assign appropriate parts of buffer to image planes in pFrameRGB
	// Note that pFrameRGB is an AVFrame, but AVFrame is a superset
	// of AVPicture
	avcodec.Avpicture_fill((*avcodec.AVPicture)(unsafe.Pointer(pFrameRGB)), (*uint8)(buffer), (avcodec.AVPixelFormat)(a), w, h)

	// initialize SWS context for software scaling
	sws_ctx = swscale.Sws_getContext(w,
		h,
		(swscale.AVPixelFormat)(pix_fmt),
		w,
		h,
		a,
		b,
		nil,
		nil,
		nil,
	)

	// Read frames and save first five frames to disk
	i := 0

	for avformat.Av_read_frame(pFormatCtx, packet) >= 0 {
		// Is this a packet from the video stream?
		s := packet.Stream_index()
		if s == videoStream {
			// Decode video frame
			avcodec.Avcodec_decode_video2(pCodecCtx, (*avcodec.AVFrame)(unsafe.Pointer(pFrame)), &frameFinished, packet)

			// Did we get a video frame?
			if frameFinished > 0 {
				// Convert the image from its native format to RGB
				d := avutil.Data(pFrame)
				l := avutil.Linesize(pFrame)
				dr := avutil.Data(pFrameRGB)
				lr := avutil.Linesize(pFrameRGB)
				swscale.Sws_scale(sws_ctx,
					d,
					l,
					0,
					h,
					dr,
					lr,
				)

				// Save the frame to disk
				if i <= 5 {
					saveFrame(pFrameRGB, w, h, i)
				}
				i++
			}
		}

		// Free the packet that was allocated by av_read_frame
		avcodec.Av_free_packet(packet)
	}

	// Free the RGB image
	avutil.Av_free(buffer)
	avutil.Av_frame_free(pFrameRGB)

	// Free the YUV frame
	avutil.Av_frame_free(pFrame)

	// Close the codecs
	avcodec.Avcodec_close(pCodecCtx)
	avcodec.Avcodec_close(pCodecCtxOrig)

	// Close the video file
	avformat.Avformat_close_input(pFormatCtx)

}

func saveFrame(pFrame *avutil.AVFrame, width int, height int, iFrame int) {

	var szFilename string
	var y int
	var file *os.File
	var err error

	szFilename = ""

	// Open file
	szFilename = fmt.Sprintf("frame%d.ppm", iFrame)

	if file, err = os.Open(szFilename); err != nil {
		log.Println("Error Reading")
	}

	// Write header
	fh := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
	log.Println(fh)

	// Write pixel data
	for y = 0; y < height; y++ {
		// d := avutil.Data(pFrame)
		// l := avutil.Linesize(pFrame)
		//##TODO
		f := make([]byte, 100)
		file.Write(f)
	}

	file.Close()

}
