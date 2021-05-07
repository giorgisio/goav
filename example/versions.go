package main

import (
	"log"

	"github.com/Turing-Chu/goav/avcodec"
	"github.com/Turing-Chu/goav/avdevice"
	"github.com/Turing-Chu/goav/avfilter"
	"github.com/Turing-Chu/goav/avformat"
	"github.com/Turing-Chu/goav/avutil"
	"github.com/Turing-Chu/goav/swresample"
	"github.com/Turing-Chu/goav/swscale"
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
