package main

import (
	"crypto/sha1"
	"io"
	"os"
)

func hash(out chan *file, in <-chan *file) error {
	for {
		select {
		case f, ok := <-in:
			if !ok {
				return nil
			}
			hash, err := calc(f.path)
			if err != nil {
				return err
			}
			f.hash = hash
			out <- f
		}
	}
}

func calc(path string) (b []byte, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return
	}
	b = h.Sum(nil)
	return
}
