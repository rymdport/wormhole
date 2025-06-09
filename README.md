# Wormhole

Wormhole is a Go implementation of [magic wormhole](https://magic-wormhole.readthedocs.io/en/latest/). It provides secure end-to-end encrypted file transfers between computers. The endpoints are connected using the same "wormhole code".
This client is compatible with the official Python and Rust clients for magic-wormhole.

This repository implements various improvements that were not accepted upstream and some that would have required breaking changes.
The goal here is to have a faster release cadance and be more actively maintained. Any patches that can be upstreamed will be upstreamed.

## Improvements over upstream
- Faster release cycle and more actively maintained.
- Contains various code improvements and moderinisations that would have required breaking changes.
  - Converted many global variables into constants.
  - Optimized the `wordlist` package for improved performance and memory usage.
  - Removed deprecated APIs.
- Added a fast path for getting the contents of text receives.
- Removed all usage of runtime reflection and replace it with faster type-checked code.
- Updated minimum Go version to 1.19 and updated code to use newer features.
- Updated dependencies to newer versions for performance and security reasons.
- Removed deprecated APIs.
- Various other improvements, cleanups and code optimizations.

## Future improvements
- Potentially decreasing the reliance on `io.ReadSeeker` for better mobile support.

## API Usage

Sending text:

```go
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/rymdport/wormhole/wormhole"
)

func sendText() {
	var c wormhole.Client

	msg := "Dillinger-entertainer"

	ctx := context.Background()

	code, status, err := c.SendText(ctx, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("On the other computer, please run: wormhole receive")
	fmt.Printf("Wormhole code is: %s\n", code)

	s := <-status

	if s.OK {
		fmt.Println("OK!")
	} else {
		log.Fatalf("Send error: %s", s.Error)
	}
}

func recvText(code string) {
	var c wormhole.Client

	ctx := context.Background()
	msg, err := c.Receive(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	if msg.Type != wormhole.TransferText {
		log.Fatalf("Expected a text message but got type %s", msg.Type)
	}

	msgBody, err := ioutil.ReadAll(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("got message:")
	fmt.Println(msgBody)
}
```

See the [examples](https://github.com/rymdport/wormhole/tree/master/examples) directory for working examples of how to use the API to send and receive text, files and directories.
