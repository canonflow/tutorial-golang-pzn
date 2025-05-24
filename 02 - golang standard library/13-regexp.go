package main

import (
	"fmt"
	"regexp"
)

func main() {
	/* ===== INTRO =====
	- Package regexp adalah utilitas di golang utk melakukan regular expression
	- Regular expression di golang menggunakan library C yg dibuat Google bernama RE2
	- Beberapa function:
		- regexp.MustCompile(string): membuat regexp
		- MatchString(string) bool: mengecek apakah regexp match dengan string
		- FindAllString(string, max): mencari string yg match dengan maksimum jumlah hasil
	*/

	var regex *regexp.Regexp = regexp.MustCompile(`e[a-z]o`)

	fmt.Println(regex.MatchString("eko")) // true
	fmt.Println(regex.MatchString("edo")) // true
	fmt.Println(regex.MatchString("eKo")) // false

	fmt.Println(regex.FindAllString("eko edo edi ego e1o, eto eKo", 10)) // [eko edo ego eto]
}
