package app

import (
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/ddd/app/dto"
	"github.com/yuchanns/gobyexample/ddd/domain"
)

type NewsService struct {
	dom *domain.NewsDomain
}

func NewNewsService(repo domain.INewsRepo) *NewsService {
	dom := domain.NewNewsDomain(repo)
	return &NewsService{
		dom: dom,
	}
}

func (srv *NewsService) Add(ctx *gin.Context, dto *dto.NewsDTO) error {
	// service call domain, the real business action
	return srv.dom.Add(ctx, dto)
}
