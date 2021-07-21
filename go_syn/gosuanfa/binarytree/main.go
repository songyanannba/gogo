package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

//前序遍历 中左右
func PreOreder(node *Hero) {
	if node != nil {
		fmt.Printf("no = %v ,name = %v \n", node.No, node.Name)
		PreOreder(node.Left)
		PreOreder(node.Right)
	}
}

//中序 左中右
func Oreder(node *Hero) {
	if node != nil {
		Oreder(node.Left)
		fmt.Printf("no = %v ,name = %v \n", node.No, node.Name)
		Oreder(node.Right)
	}
}

func ROreder(node *Hero) {
	if node != nil {
		ROreder(node.Left)
		ROreder(node.Right)
		fmt.Printf("no = %v ,name = %v \n", node.No, node.Name)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "aa",
	}
	l1 := &Hero{
		No:   2,
		Name: "wy",
	}
	r1 := &Hero{
		No:   3,
		Name: "ljy",
	}

	r2 := &Hero{
		No:   4,
		Name: "lc",
	}
	r3 := &Hero{
		No:   6,
		Name: "gg",
	}

	l2 := &Hero{
		No:   5,
		Name: "电话",
	}

	root.Left = l1
	root.Right = r1
	r1.Right = r2
	l1.Left = l2
	l1.Right = r3

	//PreOreder(root)
	//Oreder(root)
	ROreder(root)

}
