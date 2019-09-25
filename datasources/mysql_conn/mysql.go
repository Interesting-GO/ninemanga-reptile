/**
* Created by GoLand
* User: dollarkiller
* Date: 19-7-13
* Time: 上午8:39
* */
package mysql_conn

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"ninemanga-reptile/config"
	"ninemanga-reptile/datamodels"
	"time"
)

var (
	MysqlEngine *xorm.Engine
	e           error
)

func init() {
	MysqlEngine, e = xorm.NewEngine("mysql", config.MyConfig.Mysql.Dsn)
	if e != nil {
		panic(e.Error())
	}

	if config.MyConfig.App.Debug {
		MysqlEngine.ShowSQL(true)
		log.Println(config.MyConfig.Mysql.Dsn)
	}

	ping := MysqlEngine.Ping()
	if ping != nil {
		panic(ping.Error())
	}

	if config.MyConfig.Mysql.Cache {
		cacher2 := xorm.NewLRUCacher2(xorm.NewMemoryStore(), time.Hour*4, 1024)
		MysqlEngine.SetDefaultCacher(cacher2)
	}
	MysqlEngine.SetMaxIdleConns(50)
	MysqlEngine.SetMaxOpenConns(100)

	// 数据库表映射
	mapping()
}

func mapping() {
	sync2 := MysqlEngine.Sync2(
		new(datamodels.PreCartoon),
		new(datamodels.PreCartoonItem),
	)

	if sync2 != nil {
		panic(sync2.Error())
	}
}
