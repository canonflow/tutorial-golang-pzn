package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	/* ===== INTRO =====
	- Golang menyediakan package encoding utk melakukan encode dan decode
	- Golang menyediakan berbagai macam algoritma utk encoding, contohnya yg populer adalah base64, csv,dan json
	*/

	// BASE64
	fmt.Println("===== BASE64 =====")
	value := "nathan garzya santoso"
	var encoded string = base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded) // bmF0aGFuIGdhcnp5YSBzYW50b3Nv

	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(string(decoded)) // nathan garzya santoso
	}

	// CSV READER
	fmt.Println("\n===== CSV READER =====")
	csvString := "nathan,garzya,santoso\n" +
		"budi,pratama,nugraha\n" +
		"joko,morro,diah"

	// Buat reader
	reader := csv.NewReader(strings.NewReader(csvString))

	// Baca
	for {
		record, err := reader.Read()

		// Kalo udh EOF (End of File)
		if err == io.EOF {
			break
		}

		fmt.Println(record)
		/*
			[nathan garzya santoso]
			[budi pratama nugraha]
			[joko morro diah]
		*/
	}

	// CSV WRITER
	fmt.Println("\n===== CSV WRITER =====")
	// Write ke terminal dlu
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"nathan", "garzya", "santoso"})
	_ = writer.Write([]string{"joko", "pratama", "nugraha"})
	_ = writer.Write([]string{"joko", "morro", "diah"})

	// Flush
	writer.Flush()
	/*
		nathan,garzya,santoso
		joko,pratama,nugraha
		joko,morro,diah
	*/
}
