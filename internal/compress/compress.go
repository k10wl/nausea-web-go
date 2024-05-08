package compress

import (
	"compress/gzip"
	"io"
)

type Compressor interface {
	Compress(w io.Writer, r io.Reader) error
	Ext() string
}

type Gzip struct{}

func NewGzip() *Gzip {
	return &Gzip{}
}

func (c *Gzip) Compress(w io.Writer, r io.Reader) error {
	zw, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	_, err = zw.Write(data)
	if err != nil {
		return (err)
	}
	err = zw.Flush()
	if err != nil {
		return err
	}
	err = zw.Close()
	return err
}

func (c *Gzip) Ext() string {
	return ".gzip"
}
