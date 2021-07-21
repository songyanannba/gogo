package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	file, err := os.Open("/Users/syn/test.txt")
	if err != nil {
		fmt.Printf("打开文件错误 err=%v", err)
		return
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		for _, v := range str {
			fmt.Println(string(v))
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v >= ' ':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			}
		}
	}
	fmt.Printf("字符个数 %v, 数字个数 %v, 空格个数 %v", count.ChCount, count.SpaceCount, count.NumCount)

}
