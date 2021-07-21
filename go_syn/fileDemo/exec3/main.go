package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("参数有：", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("ages[%v]=%v \n", i, v)
	}
}
