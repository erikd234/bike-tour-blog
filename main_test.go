package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGpxFilesNames(t *testing.T) {
	files, err := os.ReadDir("gpx")
	if err != nil {
		t.Error(err)
	}
	spew.Dump(files)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
