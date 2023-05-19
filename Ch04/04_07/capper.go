package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	writer io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	upper := strings.ToUpper(string(p))
	return c.writer.Write([]byte(upper))
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
