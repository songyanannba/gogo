package main

import "fmt"

func read(ch1 chan int, ch2 chan bool) {
	for {
		v, ok := <-ch1
		if ok {
			fmt.Printf("read a int is %d\n", v)
		} else {
			ch2 <- true
		}
	}
}

func write(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}
	//从执行结果可以看到，当我们将10个数写完之后，如果不close()ch1,read就会阻塞，
	//程序中所有的协程都被阻塞，ch2无法写入，也无法读取，系统这时候检测到这种错误就会报错
	close(ch) //不管不会报错
}

func example1() {
	//ch1 := make(chan int)
	//ch2 := make(chan bool)

	//go write(ch1)
	//go read(ch1, ch2)
	//<-ch2
	//fmt.Println("sdjhkfhskj")
}

func example2() {
	//对于同一个通道，发送操作在接收者准备好之前是阻塞的。
	//如果通道中的数据无人接收，就无法再给通道传入其他数据。
	//新的输入无法在通道非空的情况下传入，所以发送操作会等待channel再次变为可用状态，即通道值被接收后
	//ch3 := make(chan bool, 1)
	//ch3 <- true
	//fmt.Println("1111")
	//<-ch3
	//fmt.Println("sdjhkfhskj")
}

func example3() {

	//recover()捕获的是祖父级调用时的异常，直接调用是无效的
	//recover()
	//panic(1)

	//直接调用defer也是无效的
	//defer recover()
	//panic(1)

	//必须在defer函数中直接调用才有效
	//defer func() {
	//	recover()
	//}()
	//panic(1)

}

func example4() {
	//切片的长度是切片中元素的数量。
	//切片的容量是从创建切片的索引开始的底层数组中元素的数量。
	//切片可以通过len()方法获取长度，可以通过cap()方法获取容量
	//数组计算cap()结果与len()相同

	//数组
	arr0 := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Println("cap(arr0)=", cap(arr0), arr0)

	//切片
	s01 := arr0[2:8]
	fmt.Printf("%T \n", s01)
	fmt.Println("cap(01)=", cap(s01), s01)

	s02 := arr0[4:7]
	fmt.Println("cap(s02)=", cap(s02), s02)

	s03 := s01[3:9]
	fmt.Println("截取s01[3:9]后形成s03:", s03)

	s04 := s02[4:7]
	fmt.Println("截取s02[4:7]", cap(s04), s04)

	//切片是饮用类型
	arr0[0] = "x"
	fmt.Println(arr0, s01, s02, s03, s04)

}

func main() {

	//example1() //内置函数clone
	//example2() //通道阻塞问题
	//example3() //recover和panic

	example4()

}
