package parser

import (
	"fmt"
	"strings"
)

var traceLevel int

func trace(name string) string {
	fmt.Printf("%s-> %s\n", strings.Repeat(" ", traceLevel*2), name)
	traceLevel++
	return name
}

func untrace(name string) {
	traceLevel--
	fmt.Printf("%s<- %s\n", strings.Repeat(" ", traceLevel*2), name)
}
