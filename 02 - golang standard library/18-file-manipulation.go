package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(
		name,
		os.O_CREATE|os.O_WRONLY,
		0666,
	)

	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}

	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(message)
	return nil
}

func main() {
	/* ===== INTRO =====
	- Di package OS, terdapat File Management, namun sengaja ditunda pembahasannya, karena kita harus tahu
	  dulu tentang IO
	- Saat kita membuat atau membaca file menggunakan Package OS, struct File merupakan implementasi
	  dari io.Reader dan io.Writer
	- Oleh karena itu, kita bisa melakukan baca dan tulis terhadap File tersebut menggunakan Package io/bufio
	*/

	// ===== WRITE FILE =====
	//createNewFile("sample.log", "this is sample log")

	// ===== READ FILE =====
	//fmt.Println("===== READ FILE =====")
	//result, _ := readFile("./sample.log") // this is sample log
	//fmt.Println(result)

	// ===== READ AND WRITE FILE =====
	addToFile("sample.log", "\nthis is add message")
}
