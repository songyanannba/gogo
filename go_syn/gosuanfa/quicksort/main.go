package main

import "fmt"

func Qsort(left int, right int, array *[10]int) {
	l := left
	r := right
	pivot := array[(left+right)/2]
	temp := 0

	for l < r {
		for array[l] < pivot {
			l++
		}
		for array[r] > pivot {
			r--
		}
		if l >= r {
			break
		}

		temp = array[l]
		array[l] = array[r]
		array[r] = temp

		if l == r {
			l++
			r--
		}

		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}

		if left < r {
			Qsort(left, r, array)
		}

		if right > l {
			Qsort(l, right, array)
		}

	}
}

func main() {
	arr := [10]int{9, 12, 3, 4, 77, 8, 100, 23, 445, 1}

	fmt.Println(arr)

	Qsort(0, len(arr)-1, &arr)

	fmt.Println(arr)
}
