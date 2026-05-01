package router

import (
	"app/handler"

	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.GET("/animes", handler.GetAnimes)
	r.GET("/animes/:id", handler.GetAnimeByID)
	r.PUT("/animes/:id", handler.EditAnime)
	r.DELETE("/animes/:id", handler.DeleteAnime)
	r.POST("/animes", handler.CreateAnime)
	r.GET("/", handler.HelloWorld)
	return r
}
