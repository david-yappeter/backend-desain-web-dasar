package migration

import (
	"myapp/config"
	"myapp/graph/model"
)

func MigrateTable() {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.Migrator().CreateTable(&model.User{})
}
