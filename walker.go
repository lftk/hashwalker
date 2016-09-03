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
		// get relative path
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		// filter path
		b, err := filter(ignores, rel)
		if err != nil || b == true {
			if err == nil && info.IsDir() {
				err = filepath.SkipDir
			}
			return err
		}
		// save to files(chan *file)
		if info.IsDir() == false {
			files <- &file{
				path: path,
				size: info.Size(),
			}
		}
		return nil
	})
}

func filter(patterns []string, path string) (b bool, err error) {
	for _, pattern := range patterns {
		b, err = filepath.Match(pattern, path)
		if err != nil || b == true {
			break
		}
	}
	return
}
