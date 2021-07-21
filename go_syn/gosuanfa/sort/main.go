package main

import "fmt"

func Select(arr *[5]int) {
	for j := 0; j < len(arr); j++ {
		max := arr[j]
		maxIndex := j
		for i := 0 + j; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]

		}
		fmt.Printf("第%d循环 \n", j)
		fmt.Println(arr)
	}

}

func main() {
	arr := [5]int{10, 23, 12, 100, 80}
	Select(&arr)

}
