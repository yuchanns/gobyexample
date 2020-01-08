package server

import (
	"github.com/yuchanns/gobyexample/blog/dao"
	"github.com/yuchanns/gobyexample/blog/server/http"
	"github.com/yuchanns/gobyexample/blog/service"
)

func AppInit() (closeFunc func(), err error) {
	db, err := dao.NewDB()

	if err != nil {
		return
	}

	srv := service.New(db)

	closeFunc, err = http.New(srv)
	return
}
