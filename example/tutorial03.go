package main

import (
	"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/avutil"
	"log"
)

// tutorial02.go
// Code based on a tutorial at https://ffmpeg.org/doxygen/trunk/remuxing_8c-example.html

func main() {
	//argc := os.Args
	var fragmentedMp4Options int
	//if len(argc) < 3 {
	//	log.Fatal("You need to pass at least two parameters.")
	//}
	//if len(argc) == 4{
	//	//fragmentedMp4Options = 1
	//}
	//inputFile := argc[1]
	//outFile := argc[2]
	inputFile := "/Users/jason12360/Desktop/study/ffmpeg/ffmpeg-libav-tutorial/small_bunny_1080p_60fps.mp4"
	outputFile := "/Users/jason12360/Desktop/goav/small_bunny_1080p_60fps.mov"
	//initialize input file with Context
	var inputFmtCtx *avformat.Context
	if avformat.AvformatOpenInput(&inputFmtCtx, inputFile, nil, nil) < 0 {
		log.Fatalf("Could not open input file '%s", inputFile)
		return
	}
	defer inputFmtCtx.AvformatFreeContext()
	//read stream information

	if inputFmtCtx.AvformatFindStreamInfo(nil) < 0 {
		log.Fatalf("Failed to retrieve input stream information")
		return
	}

	// initialize output file with Context
	var outputFmtCtx *avformat.Context
	avformat.AvAllocOutputContext2(&outputFmtCtx, nil, nil, &outputFile)
	if outputFmtCtx == nil {
		log.Fatalf("Could not create output context")
		return
	}
	defer outputFmtCtx.AvformatFreeContext()

	//initialize streamMapping
	streamMappingSize := int(inputFmtCtx.NbStreams())
	streamMapping := make([]int, streamMappingSize)
	var streamIndex int

	validTypeMap := map[avcodec.MediaType]int{
		avformat.AVMEDIA_TYPE_VIDEO:    1,
		avformat.AVMEDIA_TYPE_AUDIO:    1,
		avformat.AVMEDIA_TYPE_SUBTITLE: 1,
	}
	var inCodecParam *avcodec.AvCodecParameters
	defer inCodecParam.AvCodecParametersFree()
	for index, inStream := range inputFmtCtx.Streams() {
		inCodecParam = inStream.CodecParameters()
		inCodecType := inCodecParam.AvCodecGetType()

		if validTypeMap[inCodecType] == 0 {
			streamMapping[index] = -1
			continue
		}
		streamMapping[index] = streamIndex
		streamIndex++
		outStream := outputFmtCtx.AvformatNewStream(nil)
		if outStream == nil {
			log.Fatalf("Failed allocating output stream")
			return
		}
		if inCodecParam.AvCodecParametersCopyTo(outStream.CodecParameters()) < 0 {
			log.Fatalf("Failed to copy codec parameters")
			return
		}
	}
	outputFmtCtx.AvDumpFormat(0, outputFile, 1)
	if outputFmtCtx.Oformat().GetFlags()&avformat.AVFMT_NOFILE == 0 {
		avIOContext, err := avformat.AvIOOpen(outputFile, avformat.AVIO_FLAG_WRITE)
		if err != nil {
			log.Fatalf("Could not open output file '%s'", outputFile)
			return
		}
		outputFmtCtx.SetPb(avIOContext)
	}
	// initialize opts
	var opts *avutil.Dictionary
	defer opts.AvDictFree()
	if fragmentedMp4Options != 0 {
		opts.AvDictSet("movflags", "frag_keyframe+empty_moov+default_base_moof", 0)
	}

	if outputFmtCtx.AvformatWriteHeader(&opts) < 0 {
		log.Fatal("Error occurred when opening output file")
		return
	}

	var packet avcodec.Packet
	defer packet.AvPacketUnref()
	for true {
		if inputFmtCtx.AvReadFrame(&packet) < 0 {
			break
		}
		index := packet.StreamIndex()
		inputStream := inputFmtCtx.Streams()[index]
		if index >= streamMappingSize || streamMapping[index] < 0 {
			continue
		}
		packet.SetStreamIndex(streamMapping[index])
		outputStream := outputFmtCtx.Streams()[index]
		packet.AvPacketRescaleTs(inputStream.TimeBase(), outputStream.TimeBase())
		packet.SetPos(-1)
		if outputFmtCtx.AvInterleavedWriteFrame(&packet) < 0 {
			log.Fatal("Error muxing packet")
			return
		}
	}
	outputFmtCtx.AvWriteTrailer()
	if outputFmtCtx.Oformat().GetFlags()&avformat.AVFMT_NOFILE == 0 {
		outputFmtCtx.Pb().Close()
	}
	return
}
