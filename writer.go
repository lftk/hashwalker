package main

import (
	"fmt"
	"io"
)

func write(w io.Writer, files <-chan *file) error {
	for {
		select {
		case f, ok := <-files:
			if !ok {
				return nil
			}
			_, err := fmt.Fprintln(w, f)
			if err != nil {
				return err
			}
		}
	}
}
