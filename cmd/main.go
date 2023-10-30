package main

import (
	"sat-result/internal/common/register"
	"sat-result/internal/config"
	pgdatabase "sat-result/internal/infra/postgres/database"
	"sat-result/internal/utils"

	"github.com/gin-gonic/gin"
)

func main(){

	// Initialize the config
	config.LoadEnv()
	utils.LogInfo("env loaded")

	// Initialize the server
	serverConfig := config.NewServerConfig()
	utils.LogInfo("server config loaded")

	// Initialize the database
	dbConfig := config.NewDBConfig()
	utils.LogInfo("db config loaded")

	// initialize database
	database := pgdatabase.GetDBInstance();
	database.NewDBConnection(dbConfig);
	utils.LogInfo("db connection established")

	router := gin.Default()
	register.Routes(router, serverConfig)
	
	if err := router.Run(":" + serverConfig.Port); err != nil {
		utils.LogFatal(err)
	}
	utils.LogInfo("server started")
}