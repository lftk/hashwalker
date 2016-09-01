package main

import (
	"testing"
)

func Test_file(t *testing.T) {
	f1 := &file{
		path: "file1",
		size: 1024,
		hash: []byte("0123456789"),
	}
	if f1.String() != "file1,30313233343536373839,1024" {
		t.Error(f1)
	}
}
