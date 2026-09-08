package file

import (
	"io"
	"os"
)

type Download struct {
	io.ReadCloser
}

func (d *Download) SaveTo(path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, d.ReadCloser)
	return err
}
