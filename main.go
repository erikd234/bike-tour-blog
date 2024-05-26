package main

import (
	"fmt"
	"os"

	sm "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
)

func main() {
	fmt.Println("the start")

	//providers := sm.GetTileProviders()
	//spew.Dump(providers)
	ctx := sm.NewContext()
	ctx.SetSize(3000, 3000)
	files, err := os.ReadDir("./gpx")
	if err != nil {
		panic(err)
	}
	paths := []*sm.Path{}
	for _, file := range files {
		s := fmt.Sprintf("color:blue|weight:4|gpx:./gpx/%s", file.Name())
		path, err := sm.ParsePathString(s)
		if err != nil || len(path) == 0 {
			panic(err)
		}
		paths = append(paths, path[0])
	}
	for _, p := range paths {
		ctx.AddObject(
			p,
		)
	}

	img, err := ctx.Render()
	if err != nil {
		panic(err)
	}

	if err := gg.SavePNG("my-map.png", img); err != nil {
		panic(err)
	}

}
