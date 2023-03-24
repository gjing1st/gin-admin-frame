// Path: internal/apiserver/store/database
// FileName: cache.go
// Created by dkedTeam
// Author: GJing
// Date: 2022/12/27$ 11:10$

package database

import (
	"github.com/bluele/gcache"
	"time"
)

var gc gcache.Cache

func init() {
	gc = gcache.New(200).
		ARC().
		Expiration(time.Hour * 8).
		Build()
}

func GetCache() gcache.Cache {
	return gc
}
