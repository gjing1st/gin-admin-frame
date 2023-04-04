// $
// internal/apiserver/store
// Created by dkedTeam.
// Author: GJing
// Date: 2022/10/28$ 15:11$

package store

import (
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/database"
	"gorm.io/gorm"
)

var DB *gorm.DB
var GC = database.GetCache()
var DBI database.DBI
