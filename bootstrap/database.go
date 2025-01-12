package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabaseConnection(env *Env) *gorm.DB {
	var db *gorm.DB
	var err error

	log.Default().Printf("Connecting to %s database", env.DBType)

	if env.DBType == "postgres" {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		// check if the connection is working
		err = db.Exec("SELECT 1").Error
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		log.Default().Printf("Connected to %s database", env.DBType)
		// Enable foreign key constraints for PostgreSQL and self referential constraints
		db.Exec("SET CONSTRAINTS ALL DEFERRED;")
	} else if env.DBType == "sqlite" {
		dbFilePath := fmt.Sprintf("%s.db", env.DBName)
		db, err = gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{})
		// Enable foreign key constraints for SQLite
		db.Exec("PRAGMA foreign_keys = ON;")
	} else {
		log.Fatal("Unsupported DB type")
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}
