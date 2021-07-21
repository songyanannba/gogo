package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("开始")
	file, err := os.Open("/Users/syn/tt.php")
	if err != nil {
		fmt.Println("open file err", err)
	}

	defer file.Close()

	reder := bufio.NewReader(file)

	for {
		str, err := reder.ReadString('\n')

		if err == io.EOF {
			break
		}

		fmt.Print(str)
	}
	fmt.Println("文件读取结束。。。")

}
