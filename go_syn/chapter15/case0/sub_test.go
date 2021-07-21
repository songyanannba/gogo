package cal

import (
	"testing"
)

//测试
func TestGetSub(t *testing.T) {
	res := GetSub(10, 7)
	if res != 3 {
		//fmt.Println("Addupper(10) 错误 期望值 %v ,实际值 %v \n", 55, res)
		t.Fatalf("Addupper(10) 错误 期望值 %v ,实际值 %v \n", 2, res)
	}

	t.Logf("执行正确...")
}
