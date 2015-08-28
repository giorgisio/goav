package main

import (
	"github.com/giorgisio/goav/avdevice"
	"github.com/giorgisio/goav/avfilter"
	"github.com/giorgisio/goav/avformat"
	"github.com/giorgisio/goav/swresample"
	"log"
	//"github.com/giorgisio/goav/avcodec"
	//"github.com/giorgisio/goav/avutil"
	//"github.com/giorgisio/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.Av_register_all()

	log.Println(avfilter.Avfilter_version())
	log.Println(avdevice.Avdevice_version())
	log.Println(swresample.Swresample_license())

}
