// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 15:49$

package database

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/database/initdata"
	"time"
)

var (
	db *gorm.DB
)

const (
	DriverPostgresql = "postgresql"
	DriverMysql      = "mysql"
	DriverMongo      = "mongodb"
)

// InitDB
// @description: 初始化数据库
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 22:37
// @success:
func InitDB() {
	var err error
	var dsn = MysqlEmptyDsn()
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
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMysql + "数据库连接失败")
		return
	}
	//db.Use(dbresolver.Register(dbresolver.Config{
	//	// `db2` 作为 sources，`db3`、`db4` 作为 replicas
	//	Sources:  []gorm.Dialector{mysql.Open("dsn")},
	//	Replicas: []gorm.Dialector{mysql.Open("db3_dsn"), mysql.Open("db4_dsn")},
	//	// sources/replicas 负载均衡策略
	//	Policy: dbresolver.RandomPolicy{},
	//}))

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMysql + "数据库连接失败")
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(DriverMysql + "数据库连接失败")
		return
	}

	// SetMaxIdleConns 设置MySQL的最大空闲连接数。
	sqlDB.SetMaxIdleConns(config.Config.Mysql.MinConns)
	// SetMaxOpenConns 设置MySQL的最大连接数。
	sqlDB.SetMaxOpenConns(config.Config.Mysql.MaxConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	log.Info("init db success")
}

// GetDB
// @description: 获取数据库连接
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 22:38
// @success:
func GetDB() *gorm.DB {
	if db == nil {
		InitDB()
	}
	// 初始化表和表数据
	initdata.InitData(db)

	return db
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
