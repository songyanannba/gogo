package main

import (
	"fmt"
	"reflect"
)

func reftest(b interface{}) {

	//通过反射获取type kind value
	//现获取 反射类型
	rt := reflect.TypeOf(b)
	fmt.Printf("rt = %T\n", rt)

	rval := reflect.ValueOf(b)
	fmt.Printf("jjj = %v tt = %T \n", rval, rval)

}

func ref(b interface{}) {

	rt := reflect.TypeOf(b)
	fmt.Printf("rt = %T\n", rt)

	rval := reflect.ValueOf(b)
	fmt.Printf("vv = %v tt = %T \n", rval, rval)

	fmt.Printf("kind %v %v \n", rt.Kind(), rval.Kind())

	iiv := rval.Interface()

	fmt.Printf("iivv = %v  %T\n", iiv, iiv)

	ss, ok := iiv.(Stu)

	if ok {
		fmt.Printf("stu.Name %v\n", ss.Name)
	}

}

type Stu struct {
	Name string
	Age  int
}

func main() {

	//var num int = 100
	//reftest(num)

	//stu := Stu{
	//	Name: "回家",
	//	Age:  11,
	//}

	const (
		a    = iota
		b    = iota
		c, d = iota, iota
		f    = iota
		g
		h
	)
	fmt.Println(a, b, c, d, f, g, h)
}
