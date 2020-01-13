package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/blog/dao"
	"net/http"
)

type Service struct {
	Dao *dao.Dao
}

func (Service) Json(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

func (Service) JsonErr(c *gin.Context, msg string, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": msg,
	})
}

func New(d *dao.Dao) *Service {
	return &Service{Dao: d}
}
