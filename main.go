package main

import (
	"go-sample/calculator"
	"go-sample/enumeration"
	"go-sample/gomap"
	"go-sample/gostruct"
	"go-sample/slice"
)

func main() {
	calculator.RunCalculator()
	slice.RunSlice()
	gomap.RunMap()
	enumeration.RunEnum()
	gostruct.RunStruct()
}
