package main

import (
	"fmt"
	//"github.com/giorgisio/gfg/avcodec"
	"github.com/giorgisio/gfg/avdevice" //Just for test
	"github.com/giorgisio/gfg/avfilter" //Just for test
	"github.com/giorgisio/gfg/avformat"
	//"github.com/giorgisio/gfg/avutil"
	"github.com/giorgisio/gfg/swresample" //Just for test
	//"github.com/giorgisio/gfg/swscale"
)

func main() {
	fmt.Println("Testing:")

	// Register all formats and codecs
	avformat.Av_register_all()
	avfilter.Avfilter_version()
	avdevice.Avdevice_version()
	swresample.Swresample_license()
}
