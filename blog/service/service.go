package service

import "github.com/yuchanns/gobyexample/blog/dao"

type Service struct {
	Dao *dao.Dao
}

func New(d *dao.Dao) *Service {
	return &Service{Dao:d}
}
