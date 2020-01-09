package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yuchanns/gobyexample/blog/model"
)

func (s *Service) AddMarkdown(c *gin.Context) {
	var md model.Markdown

	if err := c.ShouldBind(&md); err != nil {
		s.Json(c, s.NewRespErr(s.WithMsg(err.Error())))
		return
	}

	validate := validator.New()
	if err := validate.Struct(&md); err != nil {
		s.Json(c, s.NewRespErr(s.WithMsg(err.Error())))
		return
	}

	if err := s.DB.Create(&md).Error; err != nil {
		s.Json(c, s.NewRespErr(s.WithMsg(err.Error())))
		return
	}

	s.Json(c, s.NewResp(&md))
}

func (s *Service) GetMarkdown(c *gin.Context) {
	var md model.Markdown

	id := c.DefaultQuery("id", "0")

	if err := s.DB.Where("id = ?", id).Take(&md).Error; err != nil {
		s.Json(c, s.NewRespErr(s.WithMsg(err.Error())))
		return
	}

	s.Json(c, s.NewResp(&md))
}
