package main

import (
	"fmt"
	"io/ioutil"
)

//把一个文件的内容读取到另一个文件
func main() {

	fiel1Path := "/Users/syn/sssyyymmm.php"
	fiel2Path := "/Users/syn/shangtian"

	data, err := ioutil.ReadFile(fiel1Path)

	if err != nil {
		fmt.Printf("读取目标文件错误 err=%v\n", err)
		return
	}
	err = ioutil.WriteFile(fiel2Path, data, 0666)
	if err != nil {
		fmt.Printf("写目标文件错误 err=%v\n", err)
		return
	}

}
