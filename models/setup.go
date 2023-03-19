package models

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	USER := "integratif"
	PASS := "G3rb4ng!"
	HOST := "10.199.14.47"
	PORT := "1433"
	DBNAME := "GATE_DEV"
	URL := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", USER, PASS, HOST, PORT, DBNAME)
	database, err := gorm.Open(sqlserver.Open(URL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DB = database
}
