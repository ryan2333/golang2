package transfer

import "io"

type TranseData interface {
	Read(r io.Reader) error
	Write(w io.Writer) error
}

type TranseFile struct {
	buf []byte
}

func (f TranseFile) Read(r io.Reader) error {

	return nil
}

func (f TranseFile) Write(w io.Writer) error {

	return nil
}
