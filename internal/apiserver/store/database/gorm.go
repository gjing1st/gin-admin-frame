// $
// Created by dkedTeam.
// Author: GJing
// Date: 2022/9/9$ 15:49$

package database

import (
	"fmt"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/store/database/initdata"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	db *gorm.DB
)

type Gorm struct {
	Base
	DB *gorm.DB
}

// InitDB
// @description: 初始化数据库
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2022/4/6 22:37
// @success:
func InitDB() {
	var err error
	dsn := MysqlEmptyDsn()
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
	//dsn, err := dsnAndCreate()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库创建失败")
		return
	}
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库连接失败")
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
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库连接失败")
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库连接失败")
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
func GetGorm() *Gorm {
	return &Gorm{}
}
func (g *Gorm) InitDB() {
	var err error
	dsn := MysqlEmptyDsn()
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
	//dsn, err := dsnAndCreate()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库创建失败")
		return
	}
	g.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库连接失败")
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
	sqlDB, err := g.DB.DB()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库连接失败")
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库连接失败")
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

func (g *Gorm) GetDB() DBI {
	if g.DB == nil {
		g.DB = GetDB()
		//g.InitDB()
	}
	// 初始化表和表数据
	//initdata.InitData(g.DB)
	return g

}
func (g *Gorm) Model(entity interface{}) DBI {
	g.DB = g.DB.Model(entity)
	return g
}

func (g *Gorm) Where(query string, args ...interface{}) DBI {
	g.DB = g.DB.Where(query, args)
	return g
}
func (g *Gorm) Update(column string, value interface{}) DBI {
	g.DB = g.DB.Update(column, value)
	return g
}
func (g *Gorm) First(entity interface{}) DBI {
	g.DB = g.DB.First(entity)
	return g
}
func (g *Gorm) Delete(value interface{}, conds ...interface{}) DBI {
	g.DB = g.DB.Delete(value, conds)
	return g
}

func (g *Gorm) Debug() DBI {
	g.DB = g.DB.Debug()
	return g
}

func (g *Gorm) Find(dest interface{}, conds ...interface{}) DBI {
	g.DB = g.DB.Find(dest, conds)
	return g
}

func (g *Gorm) Limit(limit int) DBI {
	g.DB = g.DB.Limit(limit)
	return g
}

func (g *Gorm) Offset(offset int) DBI {
	g.DB = g.DB.Offset(offset)
	return g
}

func (g *Gorm) Select(query interface{}, args ...interface{}) DBI {
	g.DB = g.DB.Select(query, args)
	return g

}
func (g *Gorm) Joins(query string, args ...interface{}) DBI {
	g.DB = g.DB.Joins(query, args)
	return g
}
