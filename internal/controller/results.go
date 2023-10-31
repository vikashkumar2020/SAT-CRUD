package controller

import (
	"github.com/gin-gonic/gin"
	database "sat-result/internal/infra/postgres/database"
	model "sat-result/internal/model/entity"
	types "sat-result/internal/model/types"
)

func GetAllResult(c *gin.Context) {

	// get the database connection
	db := database.GetDBInstance().GetDB()

	// get the result from the database
	var result []model.Result

	if err := db.Find(&result).Error; err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Internal Server Error",
				"status":  "ERROR",
			},
		)
	} else {
		c.JSON(200, gin.H{
			"message": "Results Fetched Successfully",
			"status":  "SUCCESS",
			"results": result,
		})
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
				"status":  "ERROR",
			},
		)
		return
	}

	// get the database connection
	db := database.GetDBInstance().GetDB()

	var Passed string

	if requestBody.SatScore >= 30 {
		Passed = "PASS"
	} else {
		Passed = "FAIL"
	}

	// create the result
	result := model.Result{
		Name:     requestBody.Name,
		Address:  requestBody.Address,
		City:     requestBody.City,
		Country:  requestBody.Country,
		Pincode:  requestBody.Pincode,
		SatScore: requestBody.SatScore,
		Passed:   Passed,
	}

	// save the result
	if err := db.Create(&result).Error; err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Result Creation Failed",
				"status":  "ERROR",
			},
		)
		return
	}

	c.JSON(
		200,
		gin.H{
			"message": "Result Added Successfully",
			"status":  "SUCCESS",
		},
	)
}

func GetRank(c *gin.Context) {
	
}

func UpdateResult(c *gin.Context) {

	// get the request body
	var requestBody types.UpdateSatResultRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Invalid Request Body",
				"status":  "ERROR",
			},
		)
		return
	}

	// get the database connection
	db := database.GetDBInstance().GetDB()

	// get the result from the database
	var result model.Result

	if err := db.Where("name = ?", requestBody.Name).First(&result).Error; err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Result Not Found",
				"status":  "ERROR",
			},
		)
		return
	}

	// update the result
	result.SatScore = requestBody.SatScore
	
	if result.SatScore >= 30 {
		result.Passed = "PASS"
	} else {
		result.Passed = "FAIL"
	}

	if err := db.Save(&result).Error; err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Result Updation Failed",
				"status":  "ERROR",
			},
		)
		return
	}

	c.JSON(
		200,
		gin.H{
			"message": "Result Updated Successfully",
			"status":  "SUCCESS",
		},
	)
}

func DeleteResult(c *gin.Context) {

	// get the name from the request url 

	name := c.Param("name")

	// get the database connection
	db := database.GetDBInstance().GetDB()

	// get the result from the database
	var result model.Result

	if err := db.Where("name = ?", name).First(&result).Error; err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Result Not Found",
				"status":  "ERROR",
			},
		)
		return
	}

	// delete the result
	if err := db.Delete(&result).Error; err != nil {
		c.JSON(
			500,
			gin.H{
				"message": "Result Deletion Failed",
				"status":  "ERROR",
			},
		)
		return
	}

	c.JSON(
		200,
		gin.H{
			"message": "Result Deleted Successfully",
			"status":  "SUCCESS",
		},
	)
}
