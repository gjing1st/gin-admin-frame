// $
// internal/apiserver/store
// Created by dkedTeam.
// Author: GJing
// Date: 2022/10/28$ 15:11$

package store

import (
	"gorm.io/gorm"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/database"
)

var DB *gorm.DB
var GC = database.GetCache()
