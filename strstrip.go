// Strips whitespace from provided string
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	fmt.Println(strings.Replace(input, " ", "", -1))
}
