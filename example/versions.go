package main

import (
	"log"

	"github.com/sigmaseven/goav/avcodec"
	"github.com/sigmaseven/goav/avdevice"
	"github.com/sigmaseven/goav/avfilter"
	"github.com/sigmaseven/goav/avformat"
	"github.com/sigmaseven/goav/avutil"
	"github.com/sigmaseven/goav/swresample"
	"github.com/sigmaseven/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
