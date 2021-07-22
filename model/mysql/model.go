package mysql

import (
	"blog/model/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(config.SqlConfig.Dialect, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.SqlConfig.Name, config.SqlConfig.Pwd, config.SqlConfig.Host, config.SqlConfig.Port, config.SqlConfig.Db))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
}

func DB() *gorm.DB {
	return db
}
