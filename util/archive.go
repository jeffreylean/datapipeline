package util

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

func Archive(data []byte, fileName string, destination string) error {
	f, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer f.Close()

	w := zip.NewWriter(f)
	defer w.Close()

	wr, err := w.Create(fileName)
	if err != nil {
		return err
	}
	r := bytes.NewReader(data)

	if _, err := io.Copy(wr, r); err != nil {
		return err
	}
	return nil

}
