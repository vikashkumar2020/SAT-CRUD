package controller

import (
"github.com/gin-gonic/gin"
database "sat-result/internal/infra/postgres/database"
// types "sat-result/internal/model/types"
model "sat-result/internal/model/entity"
)

func GetAllResult(c *gin.Context) {
	
	// get the database connection
	db := database.GetDBInstance().GetDB();

	// get the result from the database
	var result []model.Result

	if err := db.Find(&result).Error; err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Internal Server Error",
				"status": "ERROR",
			},
		)
	} else {
		c.JSON(200, result)
	}

	






}

func AddResult(c *gin.Context) {
}

func GetRank(c *gin.Context) {
}

func UpdateResult(c *gin.Context) {
}
 
func DeleteResult(c *gin.Context) {
}