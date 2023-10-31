package migrations

import (
	"context"
	"database/sql"
	"sat-result/internal/config"

	database "sat-result/internal/infra/postgres/database"
	model "sat-result/internal/model/entity"
	"github.com/pressly/goose/v3"
	"gorm.io/gorm"
)

var db *gorm.DB


func init() {
	config.LoadEnv()
		config := config.NewDBConfig()
		database := database.Database{}
		database.NewDBConnection(config)
		db = database.DB
	goose.AddMigrationContext(upCreateResultTable, downCreateResultTable)
}

func upCreateResultTable(ctx context.Context, tx *sql.Tx) error {
	return db.Migrator().CreateTable(&model.Result{})
}

func downCreateResultTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return db.Migrator().DropTable(&model.Result{})
}
