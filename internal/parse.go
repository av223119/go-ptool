package internal

import (
	"os"

	"github.com/evanoberholster/imagemeta"
	"github.com/evanoberholster/imagemeta/meta/exif"
)

func ParseFile(p string) (*exif.Exif, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ex, err := imagemeta.Decode(f)
	if err != nil {
		return nil, err
	}

	return &ex, nil

}
