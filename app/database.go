package app

import (
	"database/sql"
	"time"

	"github.com/rikurunico/go-restfull-api/helper"
)

func NewDB() *sql.DB {
	// db, err := sql.Open("mysql", "root@tcp(localhost:3307)/go_restapi")
	db, err := sql.Open("mysql", "root:root@tcp(172.19.16.1:3307)/go_restapi")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
