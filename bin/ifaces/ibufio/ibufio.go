package ibufio

import (
	"bufio"
	"io"
)

const defaultBufSize = 4096

type Writer struct {
	base *bufio.Writer
}

var NewWriter = func(w io.Writer) *Writer {
	b := bufio.NewWriterSize(w, defaultBufSize)
	return &Writer{base: b}
}

var Flush = func(w *Writer) error {
	return w.base.Flush()
}

func (b *Writer) Flush() error { return Flush(b) }

var WriteString = func(s string, w *Writer) (int, error) {
	return w.base.WriteString(s)
}

func (b *Writer) WriteString(s string) (int, error) {
	return WriteString(s, b)
}
