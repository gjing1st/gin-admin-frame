// Path: internal/apiserver/store/database/initdata
// FileName: base_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/1$ 14:25$

package initdata

import (
	"encoding/json"
	"fmt"
	"testing"
)

type A struct {
	Server []B `json:"server"`
	Total  []C `json:"total"`
}

type B struct {
	Name  string
	Total int
}

type C struct {
	Name string
}

func TestA(t *testing.T) {
	var a A
	var b B
	b.Name = "aaa"
	b.Total = 11
	a.Server = append(a.Server, b)
	a.Server = append(a.Server, B{
		"qq",
		12,
	})
	js, _ := json.Marshal(a)
	fmt.Println(string(js))
	a1 := (-3) % (-2)
	fmt.Println(a1)
}
