package main

import "fmt"

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func func1() {
	result := 0
	for i := 0; i < 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is : %d \n", i, result)
	}
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {

	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

func func2() {
	fmt.Printf("%d is even: is %t\n", 16, even(16))
	fmt.Printf("%d is even: is %t\n", 17, even(17))

	fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	fmt.Printf("%d is odd: is %t\n", 19, odd(19))
}

func main2() {
	//func1()

	//
	func2()

}
