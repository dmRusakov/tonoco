package main

import (
	"github.com/dmRusakov/tonoco/internal/entity/api"
	"reflect"
	"syscall/js"
)

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
	param := api.GridParam{
		Id:  json.Get("id").String(),
		Sku: json.Get("sku").String(),
	}

	// TODO get data from API
	item := api.GridItem{
		Id:               param.Id,
		Sku:              param.Sku,
		Brand:            "Futuro",
		Name:             "Test",
		ShortDescription: "hfdg",
		Url:              "wetwetr",
		MainImage: map[string]interface{}{
			"filename":  "test",
			"extension": "jpg",
			"is_webp":   false,
			"title":     "test",
			"alt":       "test",
		},
	}

	send(id, item)
	return id
}
