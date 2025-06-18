package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/* ===== INTRO =====
	- Bufio atau Buffered IO adalah sebuah package untuk membuat IO seperti Reader dan Writer
	*/

	// ===== READER =====
	inputs := strings.NewReader("This is long string\nWith new line\n")

	reader := bufio.NewReader(inputs)

	fmt.Println("===== READER =====")
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
		/*
			This is long string
			With new line
		*/
	}

	// ===== WRITER =====
	fmt.Println("\n===== WRITER =====")
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("Selamat belajar")
	writer.Flush()
	/*
		hello world
		Selamat belajar
	*/
}
