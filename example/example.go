package main

import (
	"fmt"
	"time"

	"github.com/kdh6429/httptester"
)

func main() {
	testor, rchan := httptester.
		NewFastHttpTest("https://google.com/"). // Server Uri to request
		NClient(5).                             // How many clients will you request
		NFactor(2).                             // How many requests will each client make
		Timeout(3 * time.Second).               // When test should be ended
		DelayTime(1 * time.Second).             // How long to stop after each request per client
		Build()                                 // Prepared test object and channel for returning information

	testor.Do() // Start the test

	for {
		for response := range rchan {
			//type Resp struct {Err error, Status int, Latency int64, Size int}
			fmt.Println(response)
		}
	}
}
