// Path: pkg/utils/errcode
// FileName: errors.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/11/18$ 14:24$

package errcode

import (
	"github.com/bluele/gcache"
	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
	ErrKeyNotFound    = gcache.KeyNotFoundError
)
