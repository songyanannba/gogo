package main

import (
	"bufio"
	"fmt"
	"os"
)

func main0() {
	//打开的文件路径
	file := "/Users/syn/sssyyymmm.php"

	//用写的方法 打开一个文件
	f, err := os.OpenFile(file, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print("文件打开错误")
	}

	//程序结束关闭
	defer f.Close()
	str := "gogogogoo"

	//网结构体对象写入数据 缓存的方法 对象有文件具柄
	w := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		w.WriteString(str)
	}

	w.Flush()

}
