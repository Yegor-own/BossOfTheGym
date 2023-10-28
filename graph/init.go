package graph

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func init() {
	DBConn, _ = gorm.Open(sqlite.Open("database.sqlite3"), &gorm.Config{})
}
