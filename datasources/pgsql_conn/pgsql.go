package pgsql_conn

import (
	"github.com/dollarkillerx/beegoorm"
	_ "github.com/lib/pq"
	"ninemanga-reptile/config"
	"ninemanga-reptile/datamodels"
)

func init() {
	err := beegoorm.RegisterDataBase("default", "postgres", config.MyConfig.Pgsql.Dsn, config.MyConfig.Pgsql.MaxIdle, config.MyConfig.Pgsql.MaxOpen)
	if err != nil {
		panic(err)
	}

	mapping()
}

func PgSql() beegoorm.Ormer {
	newOrm := beegoorm.NewOrm()
	return newOrm
}

// 数据库映射
func mapping() {
	// register model
	beegoorm.RegisterModel(
		&datamodels.PreComic{},
		&datamodels.PreComicItem{},
	)
}
