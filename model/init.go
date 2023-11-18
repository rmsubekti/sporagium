package model

import (
	"github.com/rmsubekti/sporagium/data"
	"github.com/rmsubekti/sporagium/database"

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
		&Gender{},
		&Account{},
		&User{},
		&Email{},
		&Phone{},
		&Spora{},
		&Client{},
	)

	// dbname := helper.GetEnv("POSTGRES_DB", "sporagium")
	// db.Exec("ALTER DATABASE " + dbname + " SET search_path TO spora;")

	sql.ExecFile("after_create_table", db)

	//init data csv
	csv := data.LoadCsvData()
	gender, _ := csv.Read("gender")
	initGender(gender, db)

}

func initGender(data [][]string, db *gorm.DB) (err error) {
	if (db.First(&Gender{}).RowsAffected > 0) {
		return
	}
	for id, v := range data {
		if err := db.FirstOrCreate(&Gender{ID: uint(id + 1), Name: v[0]}).Error; err != nil {
			break
		}
	}
	return
}
