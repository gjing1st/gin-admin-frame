// Path: internal/apiserver/store/database
// FileName: db1.go
// Created by dkedTeam
// Author: GJing
// Date: 2023/4/2$ 19:40$

package database

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/chunanyong/zorm"
	"github.com/gjing1st/gin-admin-frame/internal/apiserver/config"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type Zorm struct {
	Base
	DB      *zorm.DBDao
	Selects string
	Finder  *zorm.Finder
	Entity  interface{}
	Page    *zorm.Page
	Err     error
}

func GetZorm() *Zorm {
	var z Zorm
	z.InitDB()
	return &Zorm{}
}

// InitDB
// @Description 初始化数据库
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/4/4 10:08
func (z *Zorm) InitDB() {
	dsn, err := z.CreateDB()
	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库创建失败")
		return
	}
	dbConfig := &zorm.DataSourceConfig{
		//DSN 数据库的连接字符串,parseTime=true会自动转换为time格式,默认查询出来的是[]byte数组.&loc=Local用于设置时区
		DSN: dsn,
		//sql.Open(DriverName,DSN) DriverName就是驱动的sql.Open第一个字符串参数,根据驱动实际情况获取
		DriverName:            config.Config.DBDriver, //数据库驱动名称
		Dialect:               config.Config.DBDriver, //数据库类型
		SlowSQLMillis:         0,                      //慢sql的时间阈值,单位毫秒.小于0是禁用SQL语句输出;等于0是只输出SQL语句,不计算执行时间;大于0是计算SQL执行时间,并且>=SlowSQLMillis值
		MaxOpenConns:          200,                    //数据库最大连接数,默认50
		MaxIdleConns:          200,                    //数据库最大空闲连接数,默认50
		ConnMaxLifetimeSecond: 0,                      //连接存活秒时间. 默认600(10分钟)后连接被销毁重建.
		//避免数据库主动断开连接,造成死连接.MySQL默认wait_timeout 28800秒(8小时)
		DefaultTxOptions: nil, //事务隔离级别的默认配置,默认为nil
	}
	z.DB, err = zorm.NewDBDao(dbConfig)

	if err != nil {
		log.WithFields(log.Fields{"err": err.Error()}).Panic(config.Config.DBDriver + "数据库连接异常")
		return
	}
}

func (z *Zorm) GetDB() DBI {
	if z.DB == nil {
		z.InitDB()
	}

	return &Zorm{}
}

func (z *Zorm) GetTableName() (tableName string) {
	if z.Entity == nil {
		z.Err = errors.New("entity空")
		return
	}
	entity := z.Entity.(Entity)
	tableName = entity.TableName()
	if tableName == "" {
		t := reflect.TypeOf(entity)
		tableName = t.Name()
	}
	return
}
func (z *Zorm) Where(query string, args ...interface{}) DBI {
	if z.Finder == nil {
		z.Finder = zorm.NewFinder()
	}

	sql, _ := z.Finder.GetSQL()
	if strings.Index(sql, "where") == -1 {
		z.Finder = z.Finder.Append("where "+query, args)
	} else {
		z.Finder = z.Finder.Append("and "+query, args)
	}
	return z
	//sqlArr := strings.Split(sql, " ")
	//if len(sqlArr) < 1 {
	//	z.Err = errors.New("sql语句错误：" + sql)
	//	return z
	//}
	//switch sqlArr[0] {
	//case "SELECT":
	//	if strings.Index(sql, "where") == -1 {
	//		z.Finder = z.Finder.Append(" where "+query, args)
	//	} else {
	//		z.Finder = z.Finder.Append(" and "+query, args)
	//	}
	//case "UPDATE":
	//	if strings.Index(sql, "set") == -1 {
	//		z.Finder = z.Finder.Append("set").Append(query, args)
	//	} else if strings.Index(sql, "where") == -1 {
	//		z.Finder = z.Finder.Append(" where "+query, args)
	//	} else {
	//		z.Finder = z.Finder.Append(" and "+query, args)
	//	}
	//case "DELETE":
	//}
	//
	//return z
}
func (z *Zorm) Model(entity interface{}) DBI {
	z.Entity = entity
	if z.Finder != nil {
		z.Finder = zorm.NewFinder()
	}
	return z
}

func (z *Zorm) Update(column string, value interface{}) DBI {
	if z.Err != nil {
		return z
	}
	_, z.Err = zorm.Transaction(context.Background(), func(ctx context.Context) (interface{}, error) {
		finder := zorm.NewUpdateFinder(z.GetTableName()).Append(column+"=?", value)
		//拼接更新
		z.Finder, _ = finder.AppendFinder(z.Finder)
		_, err := zorm.UpdateFinder(ctx, z.Finder)
		// 如果返回的err不是nil,事务就会回滚
		return nil, err
	})

	return z
}

func (z *Zorm) First(entity interface{}) DBI {
	z.Entity = entity
	if z.Err != nil {
		return z
	}
	finder := zorm.NewSelectFinder(z.GetTableName())
	z.Finder, _ = finder.AppendFinder(z.Finder)
	_, err := zorm.QueryRow(context.Background(), finder, entity)
	z.Err = err
	return z
}

func (z *Zorm) Delete(entity interface{}, conds ...interface{}) DBI {
	z.Entity = entity
	if z.Err != nil {
		return z
	}
	_, z.Err = zorm.Transaction(context.Background(), func(ctx context.Context) (interface{}, error) {
		finder := zorm.NewDeleteFinder(z.GetTableName())
		//拼接更新
		z.Finder, _ = finder.AppendFinder(z.Finder)
		_, err := zorm.UpdateFinder(ctx, z.Finder)
		// 如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	return z
}
func (z *Zorm) Debug() DBI {
	sql, _ := z.Finder.GetSQL()
	fmt.Println("sql==", sql)
	return z
}
func (z *Zorm) Find(dest interface{}, conds ...interface{}) DBI {
	if z.Err != nil {
		return z
	}
	if z.Selects == "" {
		finder := zorm.NewSelectFinder(z.GetTableName())
		z.Finder, _ = finder.AppendFinder(z.Finder)
	} else {

	}

	err := zorm.Query(context.Background(), z.Finder, dest, z.Page)
	z.Err = err
	return z
}

func (z *Zorm) Limit(limit int) DBI {
	z.Page = zorm.NewPage()
	z.Page.PageSize = limit
	return z
}
func (z *Zorm) Offset(offset int) DBI {
	z.Page.PageNo = offset/z.Page.PageSize + 1
	return z
}
func (z *Zorm) Select(query interface{}, args ...interface{}) DBI {
	switch v := query.(type) {
	case []string:
	case string:
		z.Selects = "SELECT " + v
		if len(args) > 0 {
			for _, arg := range args {
				switch arg := arg.(type) {
				case string:
					z.Selects = z.Selects + "," + arg

				}

			}
		}
	}
	return z

}

func (z *Zorm) Joins(query string, args ...interface{}) DBI {
	if z.Err != nil {
		return z
	}
	tn := z.GetTableName()
	sql := z.Selects + " FROM " + tn + " " + query
	z.Finder = zorm.NewFinder().Append(sql)

	return z
}
