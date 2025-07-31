package main

import (
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
