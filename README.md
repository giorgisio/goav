# goav
Golang binding for FFmpeg

A comprehensive binding to the ffmpeg video/audio manipulation library.

## Usage

`````go

import "github.com/giorgisio/goav/avformat"

func main() {

	filename := "/home/giorgis/media/sample.mp4"

	// Register all formats and codecs
	avformat.Av_register_all()

	if avformat.Avformat_open_input(&pFormatCtx, filename, nil, nil) != 0 {
		return
	}

	//...

}
`````

## Libraries

* `avcodec` corresponds to the ffmpeg library: libavcodec [provides implementation of a wider range of codecs]
* `avformat` corresponds to the ffmpeg library: libavformat [implements streaming protocols, container formats and basic I/O access]
* `avutil`` corresponds to the ffmpeg library: libavutil [includes hashers, decompressors and miscellaneous utility functions]
* `avfilter` corresponds to the ffmpeg library: libavfilter [provides a mean to alter decoded Audio and Video through chain of filters]
* `avdevice` corresponds to the ffmpeg library: libavdevice [provides an abstraction to access capture and playback devices]
* `swresample` corresponds to the ffmpeg library: libswresample [implements audio mixing and resampling routines]
* `swscale` corresponds to the ffmpeg library: libswscale [implements color conversion and scaling routines]


## Installation

``` sh
go get github.com/giorgisio/goav
``` 

[FFMPEG libraries] (https://github.com/FFmpeg/FFmpeg/blob/master/INSTALL.md)


``` sh
export FFMPEG_ROOT=$HOME/ffmpeg
export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
export CGO_CFLAGS="-I$FFMPEG_ROOT/include"
export LD_LIBRARY_PATH=$HOME/ffmpeg/lib
``` 

## TODO

- [ ] Returning Errors
- [ ] Review included/excluded functions from each library
- [ ] Go Tests
- [ ] Possible restructuring packages
- [x] Tutorial01.c
- [ ] More Tutorial

## License
This library is under the [MIT License](http://opensource.org/licenses/MIT)


## Notice
goav comes with absolutely no warranty.
