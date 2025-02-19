package main

import (
	"syscall/js"
)

type GridParam struct {
	Id  string `json:"id"`
	Sku string `json:"sku"`
}

type GridItem struct {
	Id               string `json:"id" db:"id"`
	No               int32  `json:"no" db:"no"`
	Sku              string `json:"sku" db:"sku"`
	Brand            string `json:"brand" db:"brand"`
	Name             string `json:"name" db:"name"`
	ShortDescription string `json:"short_description" db:"short_description"`
	Url              string `json:"url" db:"url"`
	SalePrice        string `json:"sale_price" db:"price"`
	Price            string `json:"price" db:"price"`
	Currency         string `json:"currency" db:"currency"`
	Quantity         int64  `json:"quantity" db:"quantity"`
}

func main() {}

// send saves data to sessionStorage
func send(id int, data js.Value) {
	js.Global().Get("sessionStorage").Call("setItem", id, js.Global().Get("JSON").Call("stringify", data))
}

// get retrieves data from sessionStorage
func get(id int) js.Value {
	paramString := js.Global().Get("sessionStorage").Call("getItem", id).String()
	return js.Global().Get("JSON").Call("parse", paramString)
}

//export status
func status(id int) int {
	status := js.ValueOf(map[string]interface{}{
		"status": "WebAssembly Ready!",
	})
	send(id, status)
	return id
}

//export grid
func grid(id int) int {
	parsed := get(id)
	param := GridParam{
		Id:  parsed.Get("Id").String(),
		Sku: parsed.Get("sku").String(),
	}

	item := js.ValueOf(map[string]interface{}{
		"Id":               param.Id,
		"Sku":              param.Sku,
		"Brand":            "brand",
		"Name":             "name",
		"ShortDescription": "short_description",
		"Url":              "url",
		"SalePrice":        "sale_price",
		"Price":            "444.44",
		"Currency":         "555.33",
		"Quantity":         888,
	})
	send(id, item)
	return id
}
