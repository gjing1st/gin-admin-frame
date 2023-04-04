// Path: internal/apiserver/store/database
// FileName: db.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/4/2$ 19:46$

package database

import (
	"database/sql"
	"fmt"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	DriverPostgresql = "postgresql"
	DriverMysql      = "mysql"
	DriverMongo      = "mongodb"
)

type Entity interface {
	TableName() string
}

// DBI 该接口主要模仿gorm实现
type DBI interface {
	CreateDB() (dsn string, err error)
	InitDB()
	GetDB() DBI
	Model(value interface{}) DBI
	Where(query string, args ...interface{}) DBI
	Update(column string, value interface{}) DBI
	First(interface{}) DBI
	Delete(value interface{}, conds ...interface{}) DBI
	Debug() DBI
	Find(dest interface{}, conds ...interface{}) DBI
	Limit(limit int) DBI
	Offset(offset int) DBI
	Select(query interface{}, args ...interface{}) DBI
	Joins(query string, args ...interface{}) DBI
	ExampleAddFunc(str string) DBI //新增方法，到base结构体实现一个空函数，避免其他结构体无法实现DBI接口

}

type Base struct {
	gorm.DB
}

// ExampleAddFunc 此处只为实现DBI接口ExampleAddFunc方法，使其子类可以集成该方法，避免无法实现DBI接口
// 子类可以重写该方法，丰富子类自己的功能
func (db *Base) ExampleAddFunc(str string) (i DBI) {
	return
}
func (db *Base) Debug() (i DBI) {
	return
}
func (db *Base) CreateDB() (dsn string, err error) {
	dsn = MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", config.Config.Mysql.DBName)
	// 创建数据库
	if err = createDatabase(dsn, "mysql", createSql); err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMysql + "数据库创建失败")
		return
	}

	//数据库驱动
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.Mysql.UserName,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port,
		config.Config.Mysql.DBName,
	)
	return
}

// MysqlEmptyDsn
// @description: mysql配置
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/26 18:15
// @success:
func MysqlEmptyDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.Config.Mysql.UserName,
		config.Config.Mysql.Password,
		config.Config.Mysql.Host,
		config.Config.Mysql.Port)
}

// createDatabase
// @description: 创建数据库
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/10/26 18:15
// @success:
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func GetOrm(orm string) DBI {
	switch orm {
	case "zorm":
		return GetZorm()
	case "gorm":
		return GetGorm()
	}
	panic("orm错误")
}
