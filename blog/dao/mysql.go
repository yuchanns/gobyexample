package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/yuchanns/gobyexample/utils/uviper"
)

type dbConfig struct {
	Dsn    string
	Active int
	Idle   int
}

func NewDB() (db *gorm.DB, err error) {
	var config dbConfig

	if err = uviper.Get("db").Unmarshal(&config); err != nil {
		return
	}

	db, err = gorm.Open("mysql", config.Dsn)

	if err == nil {
		db.DB().SetMaxOpenConns(config.Active)
		db.DB().SetMaxIdleConns(config.Idle)
	}

	return
}
