package db

import (
  "fmt"
  "github.com/go-sql-driver/mysql"
  "github.com/jinzhu/gorm"
)

func (db *gorm.DB) dbManager() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
  defer db.Close()
  return db
}