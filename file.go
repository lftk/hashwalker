package main

import (
	"fmt"
)

type file struct {
	path string
	size int64
	hash []byte
}

func (f *file) String() string {
	return fmt.Sprintf("%s,%x,%d", f.path, f.hash, f.size)
}
