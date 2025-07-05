package _3_channel_parameter

import (
	"fmt"
	"testing"
	"time"
)

// By default, ketika menggunakan channel sbg parameter maka akan menggunakan pass by reference

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Nathan"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	/*
		=== RUN   TestChannelAsParameter
		Nathan
		--- PASS: TestChannelAsParameter (2.00s)
		PASS
	*/
}
