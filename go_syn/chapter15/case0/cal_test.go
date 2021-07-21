package cal

import (
	"testing"
)

//测试
func TestAddUpper(t *testing.T) {
	res := AddUpper(10)
	if res != 45 {
		//fmt.Println("Addupper(10) 错误 期望值 %v ,实际值 %v \n", 55, res)
		t.Fatalf("Addupper(10) 错误 期望值 %v ,实际值 %v \n", 55, res)
	}

	t.Logf("执行正确...")
}
