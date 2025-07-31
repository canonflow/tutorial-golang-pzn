package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
	/*
		1.0.0-SNAPSHOT
		--- PASS: TestEmbed (0.00s)
		PASS
	*/
}

//go:embed logo.jpg
var logo []byte

func TestByte(t *testing.T) {
	// Deprecated
	//err := ioutil.WriteFile("logo.png", logo, fs.ModePerm)

	err := os.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	/*
		--- PASS: TestByte (0.00s)
		PASS
	*/
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a)) // AAA

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b)) // BBB

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c)) // CCC
}
