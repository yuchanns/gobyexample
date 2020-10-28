package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/ddd/app"
	dto2 "github.com/yuchanns/gobyexample/ddd/app/dto"
	"github.com/yuchanns/gobyexample/ddd/domain"
)

type NewsServer struct {
	srv *app.NewsService
}

func NewNewsServer(repo domain.INewsRepo) *NewsServer {
	srv := app.NewNewsService(repo)
	return &NewsServer{
		srv: srv,
	}
}

func (s *NewsServer) Add(ctx *gin.Context) {
	dto := &dto2.NewsDTO{}
	// bind data to dto
	if err := ctx.ShouldBind(dto); err != nil {
		ctx.JSON(200, gin.H{
			"msg":  fmt.Sprintf("%+v", err),
			"code": 0,
		})
		return
	}
	// do dto validate
	if err := dto.Validate(); err != nil {
		ctx.JSON(200, gin.H{
			"msg":  fmt.Sprintf("%s", err),
			"code": 0,
		})
		return
	}
	// call service
	if err := s.srv.Add(ctx, dto); err != nil {
		ctx.JSON(200, gin.H{
			"msg": fmt.Sprintf("%+v", err),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"msg":  "success",
		"code": 200,
	})
	return
}
