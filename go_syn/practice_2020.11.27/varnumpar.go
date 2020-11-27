package main

import "fmt"

func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	//fmt.Println(min) 	获取第一个值
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func f1(s ...string) {

	fmt.Println(s)
	//f2()
	//f3()
	for i := 0; i < len(s); i++ {
		if i == 0 {
			f2(s[0])
		}
		if i == 4 {
			f3(s[4])
		}
	}

}

func f2(a string) {
	fmt.Printf("f2 num int %s \n", a)
}

func f3(b string) {
	fmt.Printf("f3 num int %s \n", b)
}

func ff(values ...interface{}) {
	//fmt.Println(values)  //[2 fff true 1]

	//2
    //fff
    //true
    //1
	for _, val := range values {
		switch v := val.(type) {
		case int:
			fmt.Println(v)
		case string:
			fmt.Println(v)
		case bool:
			fmt.Println(v)
		case float64:
			fmt.Println(v)
		default:
			fmt.Println("wu wu wu")
		}
	}


}

func main() {

	//x := min(1, 2, 3, 4, 5)
	//fmt.Printf("the minnum is:%d\n", x)

	//arr := []int{7, 9, 3, 0}
	//x = min(arr...)
	//fmt.Printf("min num in the array arr is :%d \n", x)

	//var a string = "1"
	//var b string = "dd"
	//var c string = "fff"
	//var d string = "4"
	//var e string = "5"
	//f1(a, b, c, d, e)

	//var ww int = 2
	//var cc string = "fff"
	//var vv bool = true
	//var aa int = 1
	//ff(ww, cc, vv, aa)
}
