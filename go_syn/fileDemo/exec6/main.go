package main

import (
	"encoding/json"
	"fmt"
)

type Master struct {
	Name   string
	Age    int
	Work   string
	Nengli string
}

func unStruc() {
	str := "{\"Name\":\"牛魔王\",\"Age\":10000,\"Work\":\"妖怪\",\"Nengli\":\"泡妞\"}"

	var ms Master

	err := json.Unmarshal([]byte(str), &ms)

	if err != nil {
		fmt.Printf("失败err := %v\n", err)
	}
	fmt.Printf("反序列化 %v\n", ms)

}

func unMap() {

	str := " {\"age\":500,\"name\":\"孙悟空\",\"职业\":\"保镖\"}"

	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("失败err := %v\n", err)
	}
	fmt.Printf("序列化map %v \n", a)

}

func unSel() {
	str := "[{\"age\":111,\"name\":\"aaa\"},{\"age\":222,\"name\":\"bbbb\"}]"

	var se []map[string]interface{}

	err := json.Unmarshal([]byte(str), &se)
	if err != nil {
		fmt.Printf("cuowu err %v ", err)
	}
	fmt.Printf("切片反序列 %v", se)

}

func main() {
	unStruc()
	unMap()
	unSel()
}
