package main

import (
	"fmt"
	"testing"
)

func TestCombine(t *testing.T) {
	folderPath := "./gpx" // Change this to your folder path if different
	outputFile := "combined_output.gpx"

	inputFiles, err := getGPXFilesFromFolder(folderPath)
	if err != nil {
		fmt.Println("Error getting GPX files:", err)
		return
	}

	if err := combineGPXFiles(inputFiles, outputFile); err != nil {
		fmt.Println("Error combining GPX files:", err)
		return
	}

	fmt.Printf("Combined GPX file saved as %s\n", outputFile)
}

func TestReduceSizeOfFile(t *testing.T) {
	gpxFile := "modified_modified_modified_combined_output.gpx"
	reduceSizeOfFile(gpxFile)
}
