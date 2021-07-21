package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//打开的文件路径
	file := "/Users/syn/sssyyymmm.php"

	//O_TRUNC 清空写
	//O_APPEND 追加写
	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("打开文件错误")
	}

	defer f.Close()

	//读取文件内容
	reader := bufio.NewReader(f)

	for {
		strr, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(strr)
	}

	str := "asdh一款\n"

	w := bufio.NewWriter(f)

	for i := 0; i < 10; i++ {
		w.WriteString(str)
	}

	w.Flush()

}
