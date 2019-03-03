// +build sample

package main

import (
	"fmt"

	"github.com/kyoh86/ask"
)

func main() {
	fmt.Printf("A version of the Package %s is %s\n", "ask", ask.Version())
}
