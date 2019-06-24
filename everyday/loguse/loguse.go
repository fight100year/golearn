package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "hi: ", log.Lshortfile)
	)

	logger.Print("123")
	fmt.Print(&buf)
}
