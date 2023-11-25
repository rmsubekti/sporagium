package service

import (
	"github.com/rmsubekti/sporagium/database"
	"github.com/rmsubekti/sporagium/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	conn, sql := database.Connect()
	db = conn

	// init extensions, schemas, and functions
	sql.ExecFile("before_create_table", db)
	db.Exec(`
		CREATE SCHEMA IF NOT EXISTS "master";
		CREATE SCHEMA IF NOT EXISTS "user";
		CREATE SCHEMA IF NOT EXISTS "spora";
	`)

	db.AutoMigrate(
		&models.User{},
		&models.Spora{},
		&models.Secret{},
	)

	// dbname := helper.GetEnv("POSTGRES_DB", "sporagium")
	// db.Exec("ALTER DATABASE " + dbname + " SET search_path TO spora;")

	sql.ExecFile("after_create_table", db)
}
