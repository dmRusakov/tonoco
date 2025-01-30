package main

import (
	"syscall/js"

	"github.com/dmRusakov/tonoco/internal/entity/api"
)

// This function returns a UUID and stores the data in the session storage.
func returnData(data any, err error) int {
	id := int(js.Global().Get("Math").Call("random").Float() * 256 * 4)

	if err != nil {
		js.Global().Get("sessionStorage").Call("setItem", id, err.Error())
		return 0
	}

	js.Global().Get("sessionStorage").Call("setItem", id, data)
	return id
}

// This calls a JS function from Go.
func main() {}

//export status
func status() int {
	return returnData("WebAssembly Ready!", nil)
}

//export getGrid
func getGrid(id string) int {
	item := api.GridItem{
		Sku:              "sku",
		Brand:            "brand",
		Name:             "name",
		ShortDescription: "short_description",
		Url:              "url",
		SalePrice:        "sale_price",
		Price:            "444.44",
		Currency:         "555.33",
		Quantity:         1,
	}

	return returnData(item, nil)
}
