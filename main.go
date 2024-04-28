package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Magic(readStream io.Reader, writeStream io.Writer) {
	in := bufio.NewReader(readStream)
	out := bufio.NewWriter(writeStream)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	fmt.Fprint(out, n)
}

func main() {
	Magic(os.Stdin, os.Stdout)
}
