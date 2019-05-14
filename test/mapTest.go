package test

import "fmt"

func TestMap() {
	var mapTest map[string]string
	mapTest = make(map[string]string)
	mapTest["data"] = "ok"
	fmt.Println(mapTest["data"])
}