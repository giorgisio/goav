# goav
Golang 语言与Ffmpeg绑定


与 ffmpeg video/audio 综合绑定的操作库。

[![GoDoc](https://godoc.org/github.com/Turing-Chu/goav?status.svg)](https://godoc.org/github.com/Turing-Chu/goav)

## 用法

### 1. go 语言代码

```go

import "github.com/giorgisio/goav/avformat"

func main() {

	filename := "sample.mp4"

	// Register all formats and codecs
	avformat.AvRegisterAll()

	ctx := avformat.AvformatAllocContext()

	// Open video file
	if avformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

	//...

}
```

### 运行

```shell
# 设置PKG_CONFIG_PATH以在编译时能找到相应头文件
export PKG_CONFIG_PATH=${PKG_CONFIG_PATH}:YOUR_FFMPEG_PKG_CONFIG_PATH
go run main.go
```

## 库

* `avcodec` 对应 ffmpeg 库: libavcodec [提供各种编码实现]
* `avformat` 对应 ffmpeg 库: libavformat 实现流协议，容器格式以及基本IO访问]
* `avutil` 对应 ffmpeg 库: libavutil [包括哈希化、解压以及各种各样的应用函数]
* `avfilter` 对应 ffmpeg 库: libavfilter [提供了一种通过滤波器链来更改解码的音频和视频的方法]
* `avdevice` 对应 ffmpeg 库: libavdevice [提供访问捕获和回放设备的抽象]
* `swresample` 对应 ffmpeg 库: libswresample [实现音频混合和重采样例程]
* `swscale` 对应 ffmpeg 库: libswscale [实现颜色转换和缩放例程]


## 安装

[FFMPEG INSTALL INSTRUCTIONS](https://github.com/FFmpeg/FFmpeg/blob/master/INSTALL.md)

``` sh
sudo apt-get -y install autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev

sudo apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev

sudo apt-get install yasm

export FFMPEG_ROOT=$HOME/ffmpeg
export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
export CGO_CFLAGS="-I$FFMPEG_ROOT/include"
export LD_LIBRARY_PATH=$HOME/ffmpeg/lib
``` 

``` 
go get github.com/giorgisio/goav

``` 

## 更多示例

更多可用代码示例都在 `examples/`目录下。

## 注意
- 该库中 Go 中的函数名为常量以便于搜索
- [cgo: 用 C 扩展 Go](http://blog.giorgis.io/cgo-examples)

## 贡献
- 克隆该仓库并创建你自己的特性分支。
- 遵从 Go 标准约定。
- 测试代码。
- 创建 PR 请求。

## TODO

- [ ] 错误返回
- [ ] 垃圾收集
- [X] 每个库的导入/导出函数的代码审查
- [ ] Go 测试
- [ ] 可能重构包
- [x] Tutorial01.c
- [ ] 更多教程

## License
该库在 [MIT License](http://opensource.org/licenses/MIT) 许可下。
