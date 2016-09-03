package main

import (
	"fmt"
	"io"
	"path/filepath"
)

func write(w io.Writer, root string, files <-chan *file) error {
	for {
		select {
		case f, ok := <-files:
			if !ok {
				return nil
			}
			rel, err := filepath.Rel(root, f.path)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintf(w, "%s,%x,%d\n", rel, f.hash, f.size)
			if err != nil {
				return err
			}
		}
	}
}
