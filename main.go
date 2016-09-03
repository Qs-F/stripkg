package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		return
	} else if len(os.Args) > 1 {
		s := strings.TrimPrefix(os.Args[1], "https://")
		fmt.Fprintf(os.Stdout, s)
	}
	return
}
