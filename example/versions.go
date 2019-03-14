package main

import (
	"log"

	"github.com/ampsight/goav/avcodec"
	"github.com/ampsight/goav/avdevice"
	"github.com/ampsight/goav/avfilter"
	"github.com/ampsight/goav/avformat"
	"github.com/ampsight/goav/avutil"
	"github.com/ampsight/goav/swresample"
	"github.com/ampsight/goav/swscale"
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
