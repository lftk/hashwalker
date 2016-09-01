package main

import (
	"os"
	"path/filepath"
	"strings"
)

func walk(root string, ignores []string, files chan *file) error {
	var ignored []string
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// get relative path
		path, err = filepath.Rel(root, path)
		if err != nil {
			return err
		}
		// filter file
		b, err := filter(ignores, path)
		if err != nil || b == true {
			if err == nil && info.IsDir() {
				ignored = append(ignored, path)
			}
			return err
		}
		for _, dir := range ignored {
			if strings.HasPrefix(path, dir) {
				return nil
			}
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
