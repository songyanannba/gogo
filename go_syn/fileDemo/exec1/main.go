package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dsFlile string, srcFile string) (written int64, err error) {

	srcObj, err := os.Open(srcFile)

	if err != nil {
		fmt.Printf("打开文件错误 err= %v\n", err)
	}
	reader := bufio.NewReader(srcObj)

	dsObj, err := os.OpenFile(dsFlile, os.O_WRONLY|os.O_CREATE, 0666)

	write := bufio.NewWriter(dsObj)
	if err != nil {
		fmt.Printf("打开目标文件错误 err= %v\n", err)
	}

	defer srcObj.Close()
	defer dsObj.Close()
	return io.Copy(write, reader)
}

func main() {

	//拷贝文件
	srcFile := "/Users/syn/sssyyymmm.php"
	dsFlile := "/Users/syn/akl.wocao"
	_, err := CopyFile(dsFlile, srcFile)
	if err != nil {
		fmt.Printf("拷贝错误 err =%v\n", err)
	} else {
		fmt.Printf("拷贝成功")
	}

}
