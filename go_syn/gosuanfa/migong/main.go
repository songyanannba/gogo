package main

import "fmt"

func Setway(myMap *[8][7]int, i int, j int) bool {
	//什么情况下找到出路 myMap[6][5]
	if myMap[6][5] == 2 {
		return true
	} else {
		if myMap[i][j] == 0 {
			//探测 上下左右
			myMap[i][j] = 2
			if Setway(myMap, i+1, j) {
				return true
			} else if Setway(myMap, i, j+1) {
				return true
			} else if Setway(myMap, i-1, j) {
				return true
			} else if Setway(myMap, i, j-1) {
				return true
			} else {
				myMap[i][j] = 3
				return false
			}
		} else {
			//是墙 不可以探测
			return false
		}
	}
}

func main() {

	//迷宫 栈的应用
	var myMap [8][7]int
	//元素为1 表示墙
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}

	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	myMap[3][1] = 1
	myMap[3][2] = 1

	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	Setway(&myMap, 1, 1)

	fmt.Println("探测完毕...")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

}
