package main

import (
	"fmt"
	"github.com/shabbyrobe/xmlwriter"
	"godocx/element"
	"os"
)

func main() {
	line := new(element.Line)
	xw := xmlwriter.Open(os.Stdout)
	_ = xw.StartDoc(xmlwriter.Doc{})
	err := line.Write(xw)
	if err != nil {
		fmt.Println(err)
	}
	xw.EndAllFlush()

}
