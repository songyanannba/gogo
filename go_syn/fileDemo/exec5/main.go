package main

import (
	"encoding/json"
	"fmt"
)

type Master struct {
	Name   string `json:"dd"`
	Age    int    `json:"a"`
	Work   string `json:"work"`
	Nengli string `json:"cc"`
}

func jiegouti() {
	master := Master{
		Name:   "牛魔王",
		Age:    10000,
		Work:   "妖怪",
		Nengli: "泡妞",
	}

	data, error := json.Marshal(&master)

	if error != nil {
		fmt.Printf("序列化错误 err:=%v\n", error)
	}

	fmt.Printf("master序列化 = %v\n", string(data))
}

func testMap() {
	var a map[string]interface{}

	a = make(map[string]interface{})

	a["name"] = "孙悟空"
	a["age"] = 500
	a["zhiye"] = "保镖"

	data, err := json.Marshal(&a)

	if err != nil {
		fmt.Printf("map序列化错误 err=%v\n", err)
	}
	fmt.Printf("map序列化 %v\n", string(data))
}

func qiepian() {

	var slice []map[string]interface{}
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "aaa"
	m1["age"] = 111
	slice = append(slice, m1)

	var m2 map[string]interface{}

	m2 = make(map[string]interface{})
	m2["name"] = "bbbb"
	m2["age"] = 222
	slice = append(slice, m2)

	data, err := json.Marshal(&slice)

	if err != nil {
		fmt.Printf("map切片序列化错误 err=%v\n", err)
	}
	fmt.Printf("map切片序列化 %v\n", string(data))
}

func main() {
	jiegouti() //结构体
	testMap()
	qiepian()
}
