# HTTP TESTER
This is a package with basic functions to help test http response and can be used and recreated in various forms.

### Simeple usage
```go
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
		DelayTime(1 * time.Second).				// How long to stop after each request per client
		Build()                                 // Prepared test object and channel for returning information

	testor.Do() // Start the test

	for {
		for response := range rchan {
			//type Resp struct {Err error, Status int, Latency int64, Size int}
			fmt.Println(response)
		}
	}
}

```
Note: Upon initial creation, each attribute value is assigned a default value.

### Diverse tester support
- NewFastHttpTest(uri string) : Http client tester capable of handling common Get Request requests.
- NewPipelineHttpTest(uri string) : Http client capable of handling TCP requests.

### Custom tester support
- NewHttpTest(client Client)

You only need to implement a client with a function that can put a response-type(*resp) channel as a parameter.
For more type information, see ./client/types.go