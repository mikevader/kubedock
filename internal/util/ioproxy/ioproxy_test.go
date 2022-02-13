package ioproxy

import (
	"bytes"
	"testing"
	"time"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		write string
		read  []byte
		flush []byte
	}{
		{
			write: "hello\n\nto the bat-mobile\nlet's go",
			read: []byte{
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0xa,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xa,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x12, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x61, 0x74, 0x2d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0xa,
			},
			flush: []byte{
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0xa,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0xa,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x12, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x62, 0x61, 0x74, 0x2d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0xa,
				0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x6c, 0x65, 0x74, 0x27, 0x73, 0x20, 0x67, 0x6f,
			},
		},
	}
	for i, tst := range tests {
		// with manual flush
		buf := &bytes.Buffer{}
		iop := New(buf, Stdout)
		iop.Write([]byte(tst.write))
		if !bytes.Equal(buf.Bytes(), tst.read) {
			t.Errorf("failed read %d - expected %v, but got %v", i, tst.read, buf.Bytes())
		}
		iop.Flush()
		if !bytes.Equal(buf.Bytes(), tst.flush) {
			t.Errorf("failed flush %d - expected %v, but got %v", i, tst.flush, buf.Bytes())
		}
		// without manual flushing
		buf = &bytes.Buffer{}
		iop = New(buf, Stdout)
		iop.Write([]byte(tst.write))
		time.Sleep(110 * time.Millisecond)
		if !bytes.Equal(buf.Bytes(), tst.flush) {
			t.Errorf("failed read %d - expected %v, but got %v", i, tst.read, buf.Bytes())
		}
	}
}