package main

import (
	"job-api/utils"
)

var Utils = new(utils.Utils)

func main() {
	type node struct {
		name string
		age  byte
	}
	a := node{name: "xwj"}
	println(a)
}
