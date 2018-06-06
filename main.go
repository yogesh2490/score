package main

import (
  "fmt"
  "score/db"
)

func main() {
	manager := db.DbManager()
	defer manager.Close()
	fmt.Println("db connection ", manager)
}
