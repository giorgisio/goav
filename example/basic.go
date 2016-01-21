package main

import "C"
import "github.com/giorgisio/goav/avutil"
import "github.com/giorgisio/goav/avformat"
import "github.com/giorgisio/goav/avfilter"
import "log"

/**
  The purpose of this file is to scale and change the bitrate of a given file.

  Useful links that help understand what is going on:

  1. http://dranger.com/ffmpeg/tutorial01.html
**/
func main() {

	filename := "sample.mp4"

	var (
		ctxtFormat *avformat.Context
	)

	// Register all formats and codecs
	avformat.AvRegisterAll()

	avfilter.AvfilterRegisterAll()

	// Open video file
	if avformat.AvformatOpenInput(&ctxtFormat, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	//This function populates pFormatCtx->streams with the proper information.
	if ctxtFormat.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")
		return
	}

	//We introduce a handy debugging function to show us what's inside
	// Dump information about file onto standard error
	ctxtFormat.AvDumpFormat(0, filename, 0)

	for i := uint(0); i < ctxtFormat.NbStreams(); i++ {
		log.Println(ctxtFormat.Streams(i).Codec().CodecType() == avutil.AVMEDIA_TYPE_VIDEO)
	}

}
