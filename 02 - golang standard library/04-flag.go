package main

import (
	"flag"
	"fmt"
)

func main() {
	/* ===== INTRO =====
	- Package FLAG berisikan fungsionalitas utk memparsing command line argument
	*/
	var username *string = flag.String("username", "localhost", "Database username")
	password := flag.String("password", "password", "Database password")
	host := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 3306, "Database port")

	flag.Parse()

	// go run 04-flag.go -host=123.231.23.3 -password="rahasia banget" -username=admin -port=3307
	fmt.Printf(
		"Username: %s\nPassword: %s\nHost: %s\nPort: %d\n",
		*username,
		*password,
		*host,
		*port,
	)

}
