package routes

import (
	"sat-result/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterResultRoutes(router *gin.RouterGroup) {
	router.GET("/result", controller.GetAllResult)
	router.POST("/result", controller.AddResult)
	router.GET("/rank/:name", controller.GetRank)
	router.PUT("/update", controller.UpdateResult)
	router.DELETE("/delete/:name", controller.DeleteResult)
}