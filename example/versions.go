package main

import (
	"fmt"
	//"github.com/giorgisio/goav/avcodec"
	"github.com/giorgisio/goav/avdevice" //Just for test
	"github.com/giorgisio/goav/avfilter" //Just for test
	"github.com/giorgisio/goav/avformat"
	//"github.com/giorgisio/goav/avutil"
	"github.com/giorgisio/goav/swresample" //Just for test
	//"github.com/giorgisio/goav/swscale"
)

func main() {
	fmt.Println("Testing:")

	// Register all formats and codecs
	avformat.Av_register_all()
	avfilter.Avfilter_version()
	avdevice.Avdevice_version()
	swresample.Swresample_license()
}
