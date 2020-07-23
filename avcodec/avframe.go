package avcodec

import "github.com/giorgisio/goav/avutil"

func (f *Frame) Pts() int64 {
	return int64(f.pts)
}

func (f *Frame) GetPicType() avutil.AvPictureType {
	return (avutil.AvPictureType)(f.pict_type)
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
