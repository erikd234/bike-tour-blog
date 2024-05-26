package main

import (
	"encoding/xml"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
)

// GPX represents the structure of a GPX file
type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Version string   `xml:"version,attr"`
	Creator string   `xml:"creator,attr"`
	Tracks  []Track  `xml:"trk"`
}

// Track represents a track in the GPX file
type Track struct {
	Name    string   `xml:"name"`
	TrkSegs []TrkSeg `xml:"trkseg"`
}

// TrkSeg represents a track segment
type TrkSeg struct {
	TrkPts []TrkPt `xml:"trkpt"`
}

// TrkPt represents a track point
type TrkPt struct {
	Lat  string `xml:"lat,attr"`
	Lon  string `xml:"lon,attr"`
	Elev string `xml:"ele,attr"`
	Time string `xml:"time"`
}

func getGPXFilesFromFolder(folderPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".gpx" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func addXMLHeader(data []byte) []byte {
	xmlHeader := []byte(`<?xml version="1.0" encoding="UTF-8"?>`)
	return append(xmlHeader, data...)
}

func combineGPXFiles(inputFiles []string, outputFile string) error {
	var combinedGPX GPX
	combinedGPX.Version = "1.1"
	combinedGPX.Creator = "GPX Combiner"

	for _, file := range inputFiles {
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		var gpx GPX
		if err := xml.Unmarshal(data, &gpx); err != nil {
			return err
		}

		combinedGPX.Tracks = append(combinedGPX.Tracks, gpx.Tracks...)
	}

	outputDataNoHeader, err := xml.MarshalIndent(combinedGPX, "", "  ")
	if err != nil {
		return err
	}
	outputData := addXMLHeader(outputDataNoHeader)
	err = os.WriteFile(outputFile, outputData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func formatFloat(val float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Round(val*shift) / shift
}

func reduceSizeOfFile(gpxFile string) error {
	data, err := os.ReadFile(gpxFile)
	if err != nil {
		return fmt.Errorf("Error Reading GPX file: %v", err)
	}

	var gpx GPX
	if err := xml.Unmarshal(data, &gpx); err != nil {
		return fmt.Errorf("Error unmarshalling GPX file: %v", err)
	}

	for i, trk := range gpx.Tracks {
		for j, trkseg := range trk.TrkSegs {
			var reducedTrkPts []TrkPt
			for k, trkpt := range trkseg.TrkPts {
				if k%2 == 0 {
					lat, err := strconv.ParseFloat(trkpt.Lat, 64)
					if err != nil {
						return fmt.Errorf("Error parsing latitude: %v", err)
					}
					lon, err := strconv.ParseFloat(trkpt.Lon, 64)
					if err != nil {
						return fmt.Errorf("Error parsing longitude: %v", err)
					}
					lat = formatFloat(lat, 5)
					lon = formatFloat(lon, 5)
					trkpt.Lat = strconv.FormatFloat(lat, 'f', 5, 64)
					trkpt.Lon = strconv.FormatFloat(lon, 'f', 5, 64)
					reducedTrkPts = append(reducedTrkPts, trkpt)
				}
			}
			gpx.Tracks[i].TrkSegs[j].TrkPts = reducedTrkPts
		}
	}

	outputDataNoHeader, err := xml.MarshalIndent(gpx, "", "  ")
	if err != nil {
		return err
	}
	outputData := addXMLHeader(outputDataNoHeader)
	outputFile := "modified_" + gpxFile
	err = os.WriteFile(outputFile, outputData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing modified GPX file: %v", err)
	}

	fmt.Printf("Reduced size GPX file saved as %s\n", outputFile)
	return nil
}
