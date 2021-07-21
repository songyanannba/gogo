package mast

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := monster{
		Name:  "孙悟空",
		Age:   111,
		Skill: "72变",
	}

	res := monster.Store()

	if !res {
		t.Fatalf("monster store err 期望 %v 其实 %v", true, res)
	}
	t.Logf("888 monster store 成功")
}

func TestReStore(t *testing.T) {
	var monster monster

	res := monster.ReStore()

	if !res {
		t.Fatalf("err %v", res)
	}

	if monster.Name != "孙悟空" {
		t.Fatalf("sdkfgkhjdf")
	}
	t.Logf("trues")

}
