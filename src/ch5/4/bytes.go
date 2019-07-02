package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("Hello "))

	fmt.Fprintf(&b, "World !\n")

	io.Copy(os.Stdout, &b)
}
