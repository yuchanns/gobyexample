package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yuchanns/gobyexample/utils/stdresp"
	"net/http"
	"strings"
)

type Service struct {
	DB *gorm.DB
	c  *gin.Context
}

func (s *Service) SetContext(c *gin.Context) {
	s.c = c
}

func (s *Service) Json(resp *stdresp.DefaultResp) {
	s.c.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"message": resp.Msg,
		"data":    resp.Data,
	})

	s.c.Abort()
}

func (s *Service) Ping() {
	name := s.c.DefaultQuery("name", "Stranger")
	resp := s.NewResp("pong", s.WithCode(200), s.WithMsg(strings.Join([]string{"hello", name}, " ")))
	s.Json(resp)
}

func (Service) NewResp(data interface{}, opts ...stdresp.IOption) *stdresp.DefaultResp {
	return stdresp.NewStdResp(data, opts...)
}

func (Service) WithCode(c int) stdresp.IOption {
	return stdresp.WithCode(c)
}

func (Service) WithMsg(m string) stdresp.IOption {
	return stdresp.WithMsg(m)
}

func (Service) WithData(d interface{}) stdresp.IOption {
	return stdresp.WithData(d)
}

func New(db *gorm.DB) (srv *Service) {
	srv = &Service{
		DB: db,
	}

	return
}
