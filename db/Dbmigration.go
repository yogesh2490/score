package db

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  gorm "github.com/jinzhu/gorm"
)

func DbManager() *gorm.DB{
	fmt.Println("i am in db")
  db, _ := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
  return db
}