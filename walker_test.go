package main

import (
	"testing"
)

func Test_filter(t *testing.T) {
	testFilter(t, "/hello/world1", []string{"/hello/world1"}, true)
	testFilter(t, "/hello/world2", []string{"/hello/*"}, true)
	testFilter(t, "/hello/world3", []string{"/*/world3"}, true)
	testFilter(t, "/hello/world4", []string{"/*/world123"}, false)
}

func testFilter(t *testing.T, path string, patterns []string, b bool) {
	b2, err := filter(patterns, path)
	if err != nil {
		t.Error(err)
	} else if b2 != b {
		t.Error(path, b)
	}
}
