package main

import "fmt"

type valNode struct {
	row int
	col int
	val int
}

func main() {

	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//转成悉数数组
	var spacseArr []valNode

	//标准的稀疏数组 记录原始的行和列
	va := valNode{
		row: 11,
		col: 11,
		val: 0,
	}
	spacseArr = append(spacseArr, va)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				val := valNode{
					row: i,
					col: j,
					val: v2,
				}
				spacseArr = append(spacseArr, val)
			}
		}
	}

	//打印稀疏数组
	fmt.Println("打印稀疏数组...")
	for i, v := range spacseArr {
		fmt.Printf(" i == %d : row== %d \t col == %v \t val == %d \n", i, v.row, v.col, v.val)
	}

	var chessMap2 [11][11]int

	//遍历稀疏数组
	for i, v := range spacseArr {
		if i != 0 {
			chessMap2[v.row][v.col] = v.val
		}
	}

	fmt.Println("恢复后的数据..")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
