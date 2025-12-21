package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/SQU1DMAN6/ftrchat/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

var db *bun.DB

func ConnectDatabase() {
	// Open database
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:database.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}
	if err = sqldb.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connected, Quan can cook.")
	db = bun.NewDB(sqldb, sqlitedialect.New())

	model.ModelUser(db)
}

func GetDB() *bun.DB {
	return db
}
