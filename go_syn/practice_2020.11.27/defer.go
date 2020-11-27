package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func function1() int {
	fmt.Printf("111111111 \n")
	defer function2()
	fmt.Printf("222222222 \n")
	return 11
	//defer function2() //放在这里报错 missing return at end of function
}

func function2() {
	fmt.Printf("333333333 \n")
}

func function3() {
	i := 5
	defer fmt.Println("bbb", i)
	i++
	fmt.Println("aaa ", i)
	defer fmt.Println("ccc", i)
	return
}

func function4() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d \n", i)
	}
}

func function5() {
	inputFile, inputError := os.Open("/Users/syn/go/src/practice_2020.11.27/1.go")

	//fmt.Println(inputFile)
	if inputError != nil {
		fmt.Printf("error xxxxxxx %s \n", inputError)
		return
	}

	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("input :%s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}

func main() {
	//a := function1()
	//fmt.Print(a)

	//function3() //defer 的语句同样可以接受参数

	//function4()//当有多个 defer 行为被注册时，它们会以逆序执行(类似栈，即后进先出)

	//function5() //读文件
}
