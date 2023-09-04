package main

import (
	"fmt"
	"godocx/element"
	"os"
)

func main() {
	line := new(element.Line)
	err := line.Write(os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
