package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	files := "/Users/syn/tt.php"
	content, err := ioutil.ReadFile(files)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	fmt.Printf("%v", string(content))
}
