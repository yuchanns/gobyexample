package dao

import (
	"database/sql"
	"github.com/yuchanns/gobyexample/utils/uviper"
)

type dbConfig struct {
	DSN    string
	Active int
	Idle   int
}

func NewDB() (db *sql.DB, err error) {
	var config dbConfig
	err = uviper.Get("mysql").Unmarshal(&config)

	if err != nil {
		return
	}

	db, err = sql.Open("mysql", config.DSN)

	if err == nil {
		db.SetMaxOpenConns(config.Active)
		db.SetMaxIdleConns(config.Idle)
	}

	return
}
