package main

import (
	"log"

	"github.com/tetsu-koba/goav/avcodec"
	"github.com/tetsu-koba/goav/avdevice"
	"github.com/tetsu-koba/goav/avfilter"
	"github.com/tetsu-koba/goav/avformat"
	"github.com/tetsu-koba/goav/avutil"
	"github.com/tetsu-koba/goav/swresample"
	"github.com/tetsu-koba/goav/swscale"
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
