# goav [![Build Status](https://travis-ci.org/giorgisio/goav.svg?branch=master)](https://travis-ci.org/giorgisio/goav)
Golang bindings for FFmpeg

## Notice
goav comes with ABSOLUTELY NO WARRANTY.

## Installation

go get github.com/giorgisio/goav

export FFMPEG_ROOT=$HOME/ffmpeg

export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"

export CGO_CFLAGS="-I$FFMPEG_ROOT/include"

export LD_LIBRARY_PATH=$HOME/ffmpeg/lib



## Usage

`````go

import "github.com/giorgisio/goav/avformat"

func main() {

	filename := "/home/giorgis/media/sample2.mp4"

	// Register all formats and codecs
	avformat.Av_register_all()

	if avformat.Avformat_open_input(&pFormatCtx, filename, nil, nil) != 0 {
		return
	}

	//...

}
`````

## TODO

- [ ] Errors
- [ ] Tests
- [ ] Review library methods
- [ ] Possible restructuring packages
- [x] Tutorial01.c
- [ ] More Tutorial

## License
This library is under the [MIT License](http://opensource.org/licenses/MIT)