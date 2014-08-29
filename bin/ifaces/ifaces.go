package ifaces

import (
	"bufio"
	"io"
	"os"
)

const defaultBufSize = 4096

var GetCurDir = func() (string, error) {
	return os.Getwd()
}

var MkDir = func(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

var Create = func(name string) (*os.File, error) {
	return os.Create(name)
}

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
