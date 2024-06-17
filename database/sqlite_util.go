package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"web_scraping/exceptions"
)

func GetSqlite3Conn(dbLocation string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{})
	exceptions.HandleError(err)

	return db
}

func GetDefaultSqlite3Conn() *gorm.DB {
	return GetSqlite3Conn(Sqlite3DBLocation)
}
