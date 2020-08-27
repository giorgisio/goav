package avcodec

import "C"
import (
	"github.com/giorgisio/goav/avutil"
	"unsafe"
)

func (f *Frame) Pts() int64 {
	return int64(f.pts)
}

func (f *Frame) GetPicType() avutil.AvPictureType {
	return (avutil.AvPictureType)(f.pict_type)
}

func (f *Frame) SetPicType(pictureType avutil.AvPictureType)  {
	f.pict_type =  uint32(pictureType)
}

func (f *Frame) GetPktSize() int {
	return int(f.pkt_size)
}

func (f *Frame) GetPts() int64 {
	return int64(f.pts)
}

func (f *Frame) GetKeyFrame() int {
	return int(f.key_frame)
}

func (f *Frame) GetCodedPictureNumber() int {
	return int(f.coded_picture_number)
}

func (f *Frame) GetWidth() int {
	return int(f.width)
}
func (f *Frame) GetHeight() int {
	return int(f.height)
}

func (f *Frame) GetData() [][]byte{
	var data [][]byte
	var slice []byte
	for i:=0;i<len(f.linesize);i++ {
		wrap := int(f.linesize[i])
		if f.data[i] != nil{
			slice = ((*[2<<31]byte)(unsafe.Pointer(f.data[i])))[:f.GetHeight()*wrap]
		}
		data = append(data,slice)
	}
	return data
}