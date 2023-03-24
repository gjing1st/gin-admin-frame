// Path: internal/pkg/dict
// FileName: func_test.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/12$ 19:25$

package dict

import (
	"fmt"
	"testing"
)

func TestSearchResources(t *testing.T) {
	r := SearchResources("节点")
	fmt.Println("============r", r)
}
