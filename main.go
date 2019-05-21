package main

import (
	"fmt"
	"go-crawler-modules/drivers"
	"go-crawler-modules/models"
)

const (
	host     = "tcp(127.0.0.1:3306)"
	user     = "root"
	password = "root"
	dbname   = "web_crawler_go"
)

func main() {
	// Connect MySQL DB
	db, err := drivers.Connect(host, user, password, dbname)
	if err != nil {
		panic(err)
	}

	// Auto Migrate DB
	db.SQL.LogMode(false)
	db.SQL.AutoMigrate(&models.Url{}, &models.Article{})

	// TEST CONNECT DB MYSQL
	fmt.Println("====================> Successfully.")
}
