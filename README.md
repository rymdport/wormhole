# wormhole-william

Wormhole is a Go implementation of [magic wormhole](https://magic-wormhole.readthedocs.io/en/latest/). It provides secure end-to-end encrypted file transfers between computers. The endpoints are connected using the same "wormhole code".
This client is compatible with the official Python and Rust clients for magic-wormhole.

This repository implements some improvements that were not accepted upstream. As well as some improvements that make the library more suited for non cli usages.
We will try to upstream as much as possible and recommend users to use upstream unless specifically needing to use this project.

## Improvements over upstream
- Removes the cli library for smaller module downloads (a minimalistic cli will be implemented in a different repository).
- Optimized the wordlist for improved performance.
- Add a fast path for getting the contents of text recieves.
- Updated module dependencies.
- Various other minor cleanups and code improvements.
- Removed deprecated transfer fields.
- Removed all uses of runtime reflection from the API.
- Update to newer minimum Go version and switch to new APIs.

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
