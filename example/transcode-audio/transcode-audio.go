package main

// derived from transcode_aac.c from ffmpeg example
// https://git.ffmpeg.org/gitweb/ffmpeg.git/blob_plain/HEAD:/doc/examples/transcode_aac.c

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/tetsu-koba/goav/avcodec"
	"github.com/tetsu-koba/goav/avformat"
	"github.com/tetsu-koba/goav/avutil"
	"github.com/tetsu-koba/goav/swresample"
)

const (
	/* The output bit rate in bit/s */
	OUTPUT_BIT_RATE = 96000
	/* The number of output channels */
	OUTPUT_CHANNELS = 2
)

func openInputFile(filename string) (int, *avformat.Context, *avcodec.Context) {
	ret := 0
	ifmtCtx := avformat.AvformatAllocContext()
	if avformat.AvformatOpenInput(&ifmtCtx, filename, nil, nil) != 0 {
		fmt.Fprintf(os.Stderr, "Unable to open file %s\n", filename)
		return avutil.AvErrorEXIT, nil, nil
	}
	defer func() {
		if ret < 0 {
			ifmtCtx.AvformatCloseInput()
		}
	}()

	if ifmtCtx.AvformatFindStreamInfo(nil) < 0 {
		fmt.Fprintf(os.Stderr, "Couldn't find stream information")
		ret = avutil.AvErrorEXIT
		return ret, nil, nil
	}
	ifmtCtx.AvDumpFormat(0, filename, 0)

	if ifmtCtx.NbStreams() != 1 {
		fmt.Fprintf(os.Stderr, "Expected one audio input stream, but found %d\n", ifmtCtx.NbStreams())
		ret = avutil.AvErrorEXIT
		return ret, nil, nil
	}

	codecId := ifmtCtx.Streams()[0].CodecParameters().AvCodecGetId()
	inputCodec := avcodec.AvcodecFindDecoder(codecId)
	if inputCodec == nil {
		fmt.Fprintf(os.Stderr, "Could not find input codec\n")
		ret = avutil.AvErrorEXIT
		return ret, nil, nil
	}
	avCtx := inputCodec.AvcodecAllocContext3()
	if avCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate a decoding context\n")
		ret = avutil.AvErrorENOMEM
		return ret, nil, nil
	}
	defer func() {
		if ret < 0 {
			avCtx.AvcodecFreeContext()
		}
	}()
	ret = ifmtCtx.Streams()[0].CodecParameters().AvCodecParametersToContext(avCtx)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not get codec parameters\n")
		return ret, nil, nil
	}
	if ret = avCtx.AvcodecOpen2(inputCodec, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open input codec (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret, nil, nil
	}

	return 0, ifmtCtx, avCtx
}

func openOutputFile(filename string, inputCodecCtx *avcodec.Context) (int, *avformat.Context, *avcodec.Context) {
	ret := 0

	ofmtCtx := avformat.AvformatAllocContext()
	if avformat.AvformatAllocOutputContext2(&ofmtCtx, nil, "", filename) != 0 {
		fmt.Fprintf(os.Stderr, "Unable to alloc output context for %s\n", filename)
		ret = avutil.AvErrorENOMEM
		return ret, nil, nil
	}
	defer func() {
		if ret < 0 {
			ofmtCtx.AvformatFreeContext()
		}
	}()

	res, err := avformat.AvIOOpen(filename, avformat.AVIO_FLAG_WRITE)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open output file '%s'\n", filename)
		ret = avutil.AvErrorEXIT
		return ret, nil, nil
	}
	ofmtCtx.SetPb(res)
	defer func() {
		if ret < 0 {
			ofmtCtx.Pb().Close()
		}
	}()

	outputCodec := avcodec.AvcodecFindEncoder(avcodec.CodecId(avcodec.AV_CODEC_ID_AAC))
	if outputCodec == nil {
		fmt.Fprintf(os.Stderr, "Could not find an AAC encoder\n")
		ret = avutil.AvErrorEXIT
		return ret, nil, nil
	}
	outStream := ofmtCtx.AvformatNewStream(nil)
	if outStream == nil {
		fmt.Fprintf(os.Stderr, "Could not reate new stream\n")
		ret = avutil.AvErrorENOMEM
		return ret, nil, nil
	}
	avCtx := outputCodec.AvcodecAllocContext3()
	if avCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate an encoding context\n")
		ret = avutil.AvErrorENOMEM
		return ret, nil, nil
	}
	defer func() {
		if ret < 0 {
			avCtx.AvcodecFreeContext()
		}
	}()

	/* Set the basic encoder parameters.
	 * The input file's sample rate is used to avoid a sample rate conversion. */
	avCtx.SetChannels(OUTPUT_CHANNELS)
	avCtx.SetChannelLayout(avutil.AvGetDefaultChannelLayout(OUTPUT_CHANNELS))
	avCtx.SetSampleRate(inputCodecCtx.SampleRate())
	avCtx.SetSampleFmt(outputCodec.SampleFmts()[0])
	avCtx.SetBitRate(OUTPUT_BIT_RATE)

	/* Allow the use of the experimental AAC encoder. */
	avCtx.SetStrictStdCompliance(avcodec.FF_COMPLIANCE_EXPERIMENTAL)

	/* Set the sample rate for the container. */
	outStream.SetTimebase(1, inputCodecCtx.SampleRate())

	/* Some container formats (like MP4) require global headers to be present.
	 * Mark the encoder so that it behaves accordingly. */
	if (ofmtCtx.Oformat().Flags() & avformat.AVFMT_GLOBALHEADER) != 0 {
		avCtx.SetFlags(avCtx.Flags() | avcodec.AV_CODEC_FLAG_GLOBAL_HEADER)
	}

	if ret = avCtx.AvcodecOpen2(outputCodec, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open output codec (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret, nil, nil
	}
	ret = outStream.CodecParameters().AvCodecParametersFromContext(avCtx)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not initialize stream parameters\n")
		return ret, nil, nil
	}

	return 0, ofmtCtx, avCtx
}

func initInputFrame() (int, *avutil.Frame) {
	frame := avutil.AvFrameAlloc()
	if frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate input frame\n")
		return avutil.AvErrorENOMEM, nil
	}
	return 0, frame
}

func initResampler(inputCodecCtx, outputCodecCtx *avcodec.Context) (int, *swresample.Context) {
	resampleCtx := swresample.SwrAllocSetOpts(
		int64(outputCodecCtx.ChannelLayout()),
		swresample.AvSampleFormat(outputCodecCtx.SampleFmt()),
		outputCodecCtx.SampleRate(),
		int64(inputCodecCtx.ChannelLayout()),
		swresample.AvSampleFormat(inputCodecCtx.SampleFmt()),
		inputCodecCtx.SampleRate())
	if resampleCtx == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate resample context\n")
		return avutil.AvErrorENOMEM, nil
	}
	if outputCodecCtx.SampleRate() != inputCodecCtx.SampleRate() {
		fmt.Fprintf(os.Stderr, "sample rate is not the same\n")
		return avutil.AvErrorEXIT, nil
	}
	ret := resampleCtx.SwrInit()
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open resample context\n")
		resampleCtx.SwrFree()
		return ret, nil
	}
	return 0, resampleCtx
}

func initFifo(outputCodecCtx *avcodec.Context) (int, *avutil.AvAudioFifo) {
	fifo := avutil.AvAudioFifoAlloc(
		avutil.AvSampleFormat(outputCodecCtx.SampleFmt()),
		outputCodecCtx.Channels(), 1)
	if fifo == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate FIFO\n")
		return avutil.AvErrorENOMEM, nil
	}
	return 0, fifo
}

func writeOutputFileHeader(outputFormatCtx *avformat.Context) int {
	ret := outputFormatCtx.AvformatWriteHeader(nil)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not write output file header (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret
	}
	return 0
}

func decodeAudioFrame(frame *avcodec.Frame, inputFormatCtx *avformat.Context, inputCodecCtx *avcodec.Context) (ret int, dataPresent, finished bool) {
	packet := avcodec.AvPacketAlloc()
	defer packet.AvFreePacket()

	ret = inputFormatCtx.AvReadFrame(packet)
	if ret < 0 {
		if ret == avutil.AvErrorEOF {
			finished = true
		} else {
			fmt.Fprintf(os.Stderr, "Could not read frame (error '%s')\n", avutil.ErrorFromCode(ret))
			return ret, false, false
		}
	}
	ret = inputCodecCtx.AvcodecSendPacket(packet)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not send packet for decoding (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret, false, false
	}
	ret = inputCodecCtx.AvcodecReceiveFrame(frame)
	if ret == avutil.AvErrorEAGAIN {
		return 0, false, false
	} else if ret == avutil.AvErrorEOF {
		return 0, false, true
	} else if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not decode frame (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret, false, false
	} else {
		return 0, true, false
	}
}

func initConvertedSamples(outputCodecCtx *avcodec.Context, frameSize int) (int, **uint8) {
	ret, samples, _ := avutil.AvSamplesAllocArrayAndSamples(outputCodecCtx.Channels(), frameSize, avutil.AvSampleFormat(outputCodecCtx.SampleFmt()), 0)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate converted input samples (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret, nil
	}
	return 0, samples
}

func convertSamples(inputData, convertedData **uint8, frameSize int, resampleCtx *swresample.Context) int {
	ret := resampleCtx.SwrConvert(convertedData, frameSize, inputData, frameSize)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not convert input samples (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret
	}
	return 0
}

func addSamplesToFifo(fifo *avutil.AvAudioFifo, convertedInputSamples **uint8, frameSize int) int {
	ret := fifo.AvAudioFifoRealloc(fifo.AvAudioFifoSize() + frameSize)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not reallocate FIFO\n")
		return ret
	}
	if fifo.AvAudioFifoWrite(convertedInputSamples, frameSize) < frameSize {
		fmt.Fprintf(os.Stderr, "Could not write data to FIFO\n")
		return avutil.AvErrorEXIT
	}
	return 0
}

func readDecodeConvertAndStore(fifo *avutil.AvAudioFifo, inputFormatCtx *avformat.Context,
	inputCodecCtx *avcodec.Context, outputCodecCtx *avcodec.Context, resampleCtx *swresample.Context) (int, bool) {
	ret, inputFrame := initInputFrame()
	if ret < 0 {
		return ret, false
	}
	defer avutil.AvFrameFree(inputFrame)

	ret, dataPresent, finished := decodeAudioFrame((*avcodec.Frame)(unsafe.Pointer(inputFrame)), inputFormatCtx, inputCodecCtx)
	if ret < 0 {
		return ret, false
	}
	if finished {
		return 0, true
	}
	if !dataPresent {
		return 0, false
	}
	frameSize := inputFrame.NbSamples()
	ret, convertedInputSamples := initConvertedSamples(outputCodecCtx, frameSize)
	if ret < 0 {
		return ret, false
	}
	defer avutil.AvSamplesFreeArrayAndSamples(convertedInputSamples)

	ret = convertSamples((**uint8)(inputFrame.ExtendedData()), convertedInputSamples, frameSize, resampleCtx)
	if ret < 0 {
		return ret, false
	}
	ret = addSamplesToFifo(fifo, convertedInputSamples, frameSize)
	if ret < 0 {
		return ret, false
	}
	return 0, false
}

func initOutputFrame(outputCodecCtx *avcodec.Context, frameSize int) (int, *avutil.Frame) {
	frame := avutil.AvFrameAlloc()
	if frame == nil {
		fmt.Fprintf(os.Stderr, "Could not allocate output frame\n")
		return avutil.AvErrorEXIT, nil
	}
	frame.SetNbSamples(frameSize)
	frame.SetChannelLayout(outputCodecCtx.ChannelLayout())
	frame.SetSampleFmt(avutil.AvSampleFormat(outputCodecCtx.SampleFmt()))
	frame.SetSampleRate(outputCodecCtx.SampleRate())

	ret := avutil.AvFrameGetBuffer(frame, 0)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not allocate output frame samples (error '%s')\n", avutil.ErrorFromCode(ret))
		avutil.AvFrameFree(frame)
		return ret, nil
	}
	return 0, frame
}

/* Global timestamp for the audio frames. */
var pts int64

func encodeAudioFrame(frame *avutil.Frame, outputFormatCtx *avformat.Context, outputCodecCtx *avcodec.Context) (int, bool) {
	outputPacket := avcodec.AvPacketAlloc()
	defer outputPacket.AvFreePacket()

	if frame != nil {
		frame.SetPts(pts)
		pts += int64(frame.NbSamples())
	}
	ret := outputCodecCtx.AvcodecSendFrame((*avcodec.Frame)(unsafe.Pointer(frame)))
	if ret == avutil.AvErrorEOF {
		return 0, false
	} else if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not send packet for encoding (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret, false
	}
	ret = outputCodecCtx.AvcodecReceivePacket(outputPacket)
	if ret == avutil.AvErrorEAGAIN {
		return 0, false
	} else if ret == avutil.AvErrorEOF {
		return 0, false
	} else if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not encode frame (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret, false
	}
	ret = outputFormatCtx.AvWriteFrame(outputPacket)
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not write frame (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret, false
	}
	return 0, true
}

func loadEncodeAndWrite(fifo *avutil.AvAudioFifo, outputFormatCtx *avformat.Context, outputCodecCtx *avcodec.Context) int {
	frameSize := fifo.AvAudioFifoSize()
	if frameSize > outputCodecCtx.FrameSize() {
		frameSize = outputCodecCtx.FrameSize()
	}
	ret, outputFrame := initOutputFrame(outputCodecCtx, frameSize)
	if ret < 0 {
		return ret
	}
	defer avutil.AvFrameFree(outputFrame)
	if fifo.AvAudioFifoRead(outputFrame.GetDataP(), frameSize) < frameSize {
		fmt.Fprintf(os.Stderr, "Could not read data from FIFO\n")
		return avutil.AvErrorEXIT
	}
	ret, _ = encodeAudioFrame(outputFrame, outputFormatCtx, outputCodecCtx)
	if ret < 0 {
		return ret
	}
	return 0
}

func writeOutputFileTrailer(outputFormatCtx *avformat.Context) int {
	ret := outputFormatCtx.AvWriteTrailer()
	if ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not write output file trailer (error '%s')\n", avutil.ErrorFromCode(ret))
		return ret
	}
	return 0
}

func TranscodeAudio(infile, outfile string) int {
	ret, inputFormatCtx, inputCodecCtx := openInputFile(infile)
	if ret < 0 {
		return ret
	}
	defer inputFormatCtx.AvformatCloseInput()
	defer inputCodecCtx.AvcodecFreeContext()

	ret, outputFormatCtx, outputCodecCtx := openOutputFile(outfile, inputCodecCtx)
	if ret < 0 {
		return ret
	}
	defer outputFormatCtx.AvformatFreeContext()
	defer outputFormatCtx.Pb().Close()
	defer outputCodecCtx.AvcodecFreeContext()

	ret, resampleCtx := initResampler(inputCodecCtx, outputCodecCtx)
	if ret < 0 {
		return ret
	}
	defer resampleCtx.SwrFree()

	ret, fifo := initFifo(outputCodecCtx)
	if ret < 0 {
		return ret
	}
	defer fifo.AvAudioFifoFree()

	ret = writeOutputFileHeader(outputFormatCtx)
	if ret < 0 {
		return ret
	}
	defer writeOutputFileTrailer(outputFormatCtx)

	for {
		finished := false
		outputFrameSize := outputCodecCtx.FrameSize()
		for fifo.AvAudioFifoSize() < outputFrameSize {
			ret, finished = readDecodeConvertAndStore(fifo, inputFormatCtx, inputCodecCtx, outputCodecCtx, resampleCtx)
			if ret < 0 {
				os.Exit(-ret)
			}
			if finished {
				break
			}
		}
		for fifo.AvAudioFifoSize() >= outputFrameSize ||
			(finished && fifo.AvAudioFifoSize() > 0) {
			ret = loadEncodeAndWrite(fifo, outputFormatCtx, outputCodecCtx)
			if ret < 0 {
				return ret
			}
		}
		if finished {
			dataWritten := true
			for dataWritten {
				ret, dataWritten = encodeAudioFrame(nil, outputFormatCtx, outputCodecCtx)
				if ret < 0 {
					return ret
				}
			}
			break
		}
	}
	return 0
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <input file> <output file>\n", os.Args[0])
		os.Exit(1)
	}
	os.Exit(-TranscodeAudio(os.Args[1], os.Args[2]))
}
