package main

import (
	"github.com/selfmodify/goav/avcodec"
	"github.com/selfmodify/goav/avdevice"
	"github.com/selfmodify/goav/avfilter"
	"github.com/selfmodify/goav/avformat"
	"github.com/selfmodify/goav/avutil"
	"github.com/selfmodify/goav/swresample"
	"github.com/selfmodify/goav/swscale"
	"log"
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
