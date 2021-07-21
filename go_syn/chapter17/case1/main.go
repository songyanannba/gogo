package main

import (
	"fmt"
	"reflect"
)

//定义monster结构体
type Monster struct {
	Name  string `json:"name" `
	Age   int
	Score float32
	Sex   string
}

//输出 显示 值
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, Score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = Score
	s.Sex = sex
}

func testStrct(b interface{}) {

	typ := reflect.TypeOf(b)
	val := reflect.ValueOf(b)
	kd := val.Kind()

	//如果不是结构体 不玩了
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField()
	for i := 0; i < num; i++ {
		fmt.Printf("field %d 值为=%v\n", i, val.Field(i))

		tagVal := typ.Field(i).Tag.Get("json")

		if tagVal != "" {
			fmt.Printf("Filed_val %d  tag ==%v\n", i, tagVal)
		}
	}

	funcN := val.NumMethod()

	fmt.Printf("funcN %v\n", funcN)

	//for i := 0; i < funcN; i++ {
	//	fmt.Printf("sss %v\n", val.Method().Field())
	//}

	//val.Method(0).Call(1, 2)
	val.Method(1).Call(nil)

	var p []reflect.Value

	p = append(p, reflect.ValueOf(10))
	p = append(p, reflect.ValueOf(20))

	res := val.Method(0).Call(p)
	fmt.Println("res ==", res[0].Int())

}

func main() {

	var a Monster = Monster{
		Name:  "黄鼠狼",
		Age:   111,
		Score: 200.0,
		Sex:   "男",
	}

	testStrct(a)
}
