package main

import (
	"fmt"
	"studygo/base"
	"studygo/lib"
)

func main() {
	lib.Yaml()
	base.Test()
	base.TestPanicRecover()
	fmt.Println("some")
}
