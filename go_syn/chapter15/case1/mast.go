package mast

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *monster) Store() bool {
	//序列化
	data, err := json.Marshal(this)
	if err != err {
		fmt.Printf("marshal err := %v", err)
		return false
	}

	filePath := "/Users/syn/go/src/chapter15/case1/monstr.txt"
	errf := ioutil.WriteFile(filePath, data, 0666)

	if err != nil {
		fmt.Printf("write err := %v", errf)
		return false
	}
	//保存在文件中
	return true
}

func (this *monster) ReStore() bool {
	fileDir := "/Users/syn/go/src/chapter15/case1/monstr.txt"

	data, err := ioutil.ReadFile(fileDir)

	if err != nil {
		fmt.Printf("restore err:= %v ", err)
		return false
	}

	error := json.Unmarshal(data, this)

	if error != nil {
		fmt.Printf("反序列化 err :+=%v", err)
		return false
	}

	return true

}
