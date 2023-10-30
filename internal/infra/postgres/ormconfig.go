package postgres


import (
	"fmt"
	config "sat-result/internal/config"
)

// GetOrmConfig Get DB string
func GetOrmConfig(dbConfig *config.DBConfig) string {
	configString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host, dbConfig.Username,
		dbConfig.Password, dbConfig.Dbname,
		dbConfig.Port,
	)
	fmt.Println(configString)
	return configString
}