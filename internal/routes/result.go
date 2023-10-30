package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterResultRoutes(router *gin.RouterGroup) {
	router.GET("/result", GetAllResult)
	router.POST("/result", AddResult)
	router.GET("/rank", GetRank)
	router.PUT("/update", UpdateResult)
	router.DELETE("/delete", DeleteResult)
}