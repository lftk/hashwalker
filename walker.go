package main

import (
	"os"
	"path/filepath"
)

func walk(root string, ignores []string, files chan *file) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		b, err := filter(root, ignores, path)
		if err != nil {
			return err
		}
		if b == true {
			return nil
		}
		if info.IsDir() {
			if path != root {
				walk(path, ignores, files)
			}
		} else {
			files <- &file{
				path: path,
				size: info.Size(),
			}
		}
		return nil
	})
}

func filter(root string, patterns []string, path string) (b bool, err error) {
	for _, pattern := range patterns {
		pattern = filepath.Join(root, pattern)
		b, err = filepath.Match(pattern, path)
		if err != nil || b == true {
			break
		}
	}
	return
}
