package handler

import (
	"app/db"

	"github.com/gin-gonic/gin"
)

func LivenessProbe(r *gin.Engine) {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
}

// Testar se o banco de dados está acessível
func ReadinessProbe(r *gin.Engine) {
	r.GET("/ready", func(ctx *gin.Context) {
		sqlDB, err := db.GetDB().DB()
		if err != nil || sqlDB.Ping() != nil {
			ctx.JSON(503, gin.H{"status": "not ready", "database": "offline"})
			return
		}
		ctx.JSON(200, gin.H{"status": "ready", "database": "online"})
	})
}
