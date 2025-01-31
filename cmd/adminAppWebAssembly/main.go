package main

import (
	"syscall/js"
)

func id() int {
	return int(js.Global().Get("Math").Call("random").Float() * 256 * 4)
}

// This calls a JS function from Go.
func main() {}

//export status
func status() int {
	id := id()
	status := js.ValueOf(map[string]interface{}{
		"status": "WebAssembly Ready!",
	})
	js.Global().Get("sessionStorage").Call("setItem", id, js.Global().Get("JSON").Call("stringify", status))
	return id
}

//export grid
func grid() int {
	id := id()
	item := js.ValueOf(map[string]interface{}{
		"Sku":              "sku",
		"Brand":            "brand",
		"Name":             "name",
		"ShortDescription": "short_description",
		"Url":              "url",
		"SalePrice":        "sale_price",
		"Price":            "444.44",
		"Currency":         "555.33",
		"Quantity":         644,
	})

	js.Global().Get("sessionStorage").Call("setItem", id, js.Global().Get("JSON").Call("stringify", item))
	return id
}
