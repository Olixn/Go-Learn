/**
 * @Author: Ne-21
 * @Description:
 * @File:  main_test
 * @Version: 1.0.0
 * @Date: 2022/1/5 18:15
 */

package main

import "testing"

func TestMonster_Store(t *testing.T) {
	monster := &Monster{
		Name:  "小白",
		Age:   56,
		Skill: "学习",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("Monster_Store 错误")
	}
	t.Logf("Monster_Store 成功")
}

func TestMonster_ReStore(t *testing.T) {
	monster := &Monster{}

	res := monster.ReStore()
	if !res {
		t.Fatalf("Monster_ReStore 错误")
	}

	if monster.Name != "小白" {
		t.Fatalf("Monster_ReStore 错误")
	}
	t.Logf("Monster_ReStore 成功")
}
