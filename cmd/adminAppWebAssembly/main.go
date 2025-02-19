package main

import (
	"reflect"
	"syscall/js"
)

type GridParam struct {
	Id  string `json:"id"`
	Sku string `json:"sku"`
}

type GridItem struct {
	Id               string `json:"id"`
	No               int32  `json:"no"`
	Sku              string `json:"sku"`
	Brand            string `json:"brand"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	Url              string `json:"url"`
	SalePrice        string `json:"sale_price"`
	Price            string `json:"price"`
	Currency         string `json:"currency"`
	Quantity         int64  `json:"quantity"`
}

type Status struct {
	Status string `json:"status"`
}

func main() {}

// send saves data to sessionStorage
func send(id int, d any) {
	data := make(map[string]interface{})
	itemValue := reflect.ValueOf(d)
	for i := 0; i < itemValue.NumField(); i++ {
		jsonKey := itemValue.Type().Field(i).Tag.Get("json")
		data[jsonKey] = itemValue.Field(i).Interface()
	}
	js.Global().Get("sessionStorage").Call("setItem", id,
		js.Global().Get("JSON").Call("stringify", data))
}

// get retrieves data from sessionStorage
func get(id int) js.Value {
	paramString := js.Global().Get("sessionStorage").Call("getItem", id).String()
	if paramString == "null" {
		return js.Null()
	}
	return js.Global().Get("JSON").Call("parse", paramString)
}

//export status
func status(id int) int {
	send(id, Status{
		Status: "WebAssembly Ready!",
	})
	return id
}

//export grid
func grid(id int) int {
	json := get(id)
	param := GridParam{
		Id:  json.Get("id").String(),
		Sku: json.Get("sku").String(),
	}

	// TODO get data from API
	item := GridItem{
		Id:               param.Id,
		No:               3,
		Sku:              param.Sku,
		Brand:            "",
		Name:             "",
		ShortDescription: "",
		Url:              "",
		SalePrice:        "",
		Price:            "",
		Currency:         "",
		Quantity:         4444,
	}

	send(id, item)
	return id
}
