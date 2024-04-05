package image

import (
	"github.com/kolesa-team/goexiv"
)

func Exif(p string) (*goexiv.ExifData, error) {
	// open
	img, err := goexiv.Open(p)
	if err != nil {
		return nil, err
	}

	// read data
	err = img.ReadMetadata()
	if err != nil {
		return nil, err
	}

	return img.GetExifData(), nil
}

func init() {
	goexiv.SetLogMsgLevel(goexiv.LogMsgMute)
}
