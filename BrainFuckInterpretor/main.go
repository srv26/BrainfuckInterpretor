package main

import (
	"strings"

	"main.go/interpretor"
)

func main() {
	r := strings.NewReader("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
	interpretor.Initialize(r)
	interpretor.Execute()
}
