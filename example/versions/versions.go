package main

import (
	"fmt"

	"github.com/tetsu-koba/goav/avcodec"
	"github.com/tetsu-koba/goav/avdevice"
	"github.com/tetsu-koba/goav/avfilter"
	"github.com/tetsu-koba/goav/avformat"
	"github.com/tetsu-koba/goav/avutil"
	"github.com/tetsu-koba/goav/swresample"
	"github.com/tetsu-koba/goav/swscale"
)

func PrintVersions() {

	// Register all formats and codecs
	//avformat.AvRegisterAll()
	//avcodec.AvcodecRegisterAll()

	fmt.Printf("AvFormat Version:\t%v\n", avformat.AvformatVersion())
	fmt.Printf("AvFilter Version:\t%v\n", avfilter.AvfilterVersion())
	fmt.Printf("AvDevice Version:\t%v\n", avdevice.AvdeviceVersion())
	fmt.Printf("SWScale Version:\t%v\n", swscale.SwscaleVersion())
	fmt.Printf("AvUtil Version:\t%v\n", avutil.AvutilVersion())
	fmt.Printf("AvCodec Version:\t%v\n", avcodec.AvcodecVersion())
	fmt.Printf("Resample Version:\t%v\n", swresample.SwresampleLicense())
}

func main() {
	PrintVersions()
}
