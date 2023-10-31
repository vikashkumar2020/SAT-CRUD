package controller

import (
"github.com/gin-gonic/gin"
database "sat-result/internal/infra/postgres/database"
types "sat-result/internal/model/types"
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

	// get the request body
	var requestBody types.CreateSatResultRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Invalid Request Body",
				"status": "ERROR",
			},
		)
		return
	}

	// get the database connection
	db := database.GetDBInstance().GetDB();

	var Passed string

	if requestBody.SatScore >= 30 {
		Passed = "PASS"
	} else {
		Passed = "FAIL"
	}

	// create the result
	result := model.Result{
		Name: requestBody.Name,
		Address: requestBody.Address,
		City: requestBody.City,
		Country: requestBody.Country,
		Pincode: requestBody.Pincode,
		SatScore: requestBody.SatScore,
		Passed: Passed,
	}

	// save the result
	if err := db.Create(&result).Error; err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Internal Server Error",
				"status": "ERROR",
			},
		)
		return
	} 

		c.JSON(
			200,
			gin.H{
				"message": "Result Added Successfully",
				"status": "SUCCESS",
			},
		)
}

func GetRank(c *gin.Context) {
}

func UpdateResult(c *gin.Context) {
}
 
func DeleteResult(c *gin.Context) {
}