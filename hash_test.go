package main

import (
	"os"
	"testing"
)

func Test_calc(t *testing.T) {
	_, err := calc(os.Args[0])
	if err != nil {
		t.Error(err)
	}
}
