package main

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

var testHash = []byte("0123456789")

type testWriter struct {
	idx int
}

func (w *testWriter) Write(b []byte) (int, error) {
	s := fmt.Sprintf("path%d,%x,%d\n", w.idx, testHash, w.idx)
	if i := bytes.Compare([]byte(s), b); i != 0 {
		return 0, errors.New(s)
	}
	w.idx++
	return len(b), nil
}

func Test_write(t *testing.T) {
	files := make(chan *file, 16)
	for i := 0; i < 16; i++ {
		files <- &file{
			path: fmt.Sprintf("path%d", i),
			size: int64(i),
			hash: testHash,
		}
	}
	close(files)
	err := write(new(testWriter), files)
	if err != nil {
		t.Error(err)
	}
}
