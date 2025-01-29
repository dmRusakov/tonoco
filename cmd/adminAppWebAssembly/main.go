//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"

	"github.com/dmRusakov/tonoco/pkg/utils/random"
)

// This function returns a UUID and stores the data in the session storage.
func returnData(data any, err error) int {
	if err != nil {
		js.Global().Get("sessionStorage").Call("setItem", 0, err.Error())
		return 0
	}

	uuid, err := random.Int(10000)

	if err != nil {
		js.Global().Get("sessionStorage").Call("setItem", 0, err.Error())
		return 0
	}

	js.Global().Get("sessionStorage").Call("setItem", uuid, data)
	return uuid
}

// This calls a JS function from Go.
func main() {}

//export status
func status() int {
	return returnData("WebAssembly Ready.", nil)
}
