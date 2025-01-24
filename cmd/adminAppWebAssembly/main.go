//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/dmRusakov/tonoco/pkg/utils/random"
	"syscall/js"
)

// This calls a JS function from Go.
func main() {
	println("WebAssembly module loaded.")
}

//export add
func add(x, y int) int {
	val := x + y
	uuid, err := random.Int(10000)

	if err != nil {
		js.Global().Get("sessionStorage").Call("setItem", "error", err.Error())
		return 0
	}

	js.Global().Get("sessionStorage").Call("setItem", uuid, val)
	return uuid
}

//export hi
func hi() string {
	return "Hello from Go!"
}

//export getService
func getService() *map[string]string {
	person := map[string]string{"Service": "Alice!", "Aria": "30", "Serviced": "Go"}

	return &person
}
