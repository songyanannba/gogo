package main

import (
	"fmt"
)

func InserSort(arr *[5]int) {
	//给第二个元素找到合适的位置
	for i := 1; i < len(arr); i++ {
		inserval := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < inserval {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入
		if insertIndex+1 != 1 {
			arr[insertIndex+1] = inserval
		}
	}

}

func main() {
	arr := [5]int{23, 89, 4, 123, 0}
	fmt.Println(arr)

	InserSort(&arr)

	fmt.Println(arr)
}
