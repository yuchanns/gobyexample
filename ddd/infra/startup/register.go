package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/ddd/domain"
	"github.com/yuchanns/gobyexample/ddd/server"
)

func RegisterHttpRoute(engine *gin.Engine) error {
	news := engine.Group("/news")
	{
		newsRepo := &domain.NewsRepo{}
		newsSrv := server.NewNewsServer(newsRepo)
		//news.GET("/:id")
		news.POST("/", newsSrv.Add)
		//news.PUT("/:id")
		//news.DELETE("/:id")
	}
	return nil
}
