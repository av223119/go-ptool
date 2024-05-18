package internal

import (
	"os"

	"github.com/tajtiattila/metadata/exif"
)

func ParseFile(p string) (*exif.Exif, error) {
	f, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return exif.Decode(f)
}
