package main

import (
	_ "embed"
	"fmt"
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
