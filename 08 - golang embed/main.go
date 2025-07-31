package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version string

//go:embed logo.jpg
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	// Version
	fmt.Println(version)

	// Byte-slice
	err := os.WriteFile("logo_new_main.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	// Path Matcher
	dir, _ := path.ReadDir("files")

	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(content))
		}
	}
}
