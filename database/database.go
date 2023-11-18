package database

import (
	"embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rmsubekti/sporagium/helper"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//go:embed sql
var sqls embed.FS

type sqlEmbed struct {
	files embed.FS
}

func (e *sqlEmbed) ExecFile(fileName string, db *gorm.DB) {
	if sql, err := e.files.ReadFile("sql/" + fileName + ".sql"); err == nil {
		db.Exec(string(sql))
	} else {
		log.Fatal(err)
	}
}

// DATABASE_URL=postgres://{user}:{password}@{hostname}:{port}/{database-name}

func GetDBConnectionString() (dsn string, uri string) {
	host := helper.GetEnv("POSTGRES_HOSTNAME", "localhost")
	user := helper.GetEnv("POSTGRES_USER", "postgres")
	password := helper.GetEnv("POSTGRES_PASSWORD", "5up3rSP0r4")
	db := helper.GetEnv("POSTGRES_DB", "sporagium")
	port := helper.GetEnv("POSTGRES_PORT", "5432")
	ssl := helper.GetEnv("POSTGRES_SSLMODE", "disable")

	//database dsn
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s ", host, user, password, db, port, ssl)
	uri = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, db)
	return
}

func Connect() (db *gorm.DB, embed *sqlEmbed) {
	var err error
	if e := godotenv.Load(); e != nil {
		log.Println(e)
	}

	logLevel := logger.Silent
	if os.Getenv("SPORAGIUM_DEBUG") == "debug" {
		logLevel = logger.Info
	}

	//database dsn
	dsn, _ := GetDBConnectionString()

	//connect to database
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					TablePrefix:   "oauth2.",
					SingularTable: false,
				},
				Logger: logger.Default.LogMode(logLevel),
			},
		)
		if err == nil {
			log.Println("Connected to database host.")
			break
		}
		log.Println("\nReconnecting to your database host....")
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("No database found : %s", err)
	}

	return db, &sqlEmbed{files: sqls}
}
