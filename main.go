package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type City struct {
	Code         string
	Name         string
	ProvinceCode string `gorm:"column:provinceCode"`
}

func main() {
	fmt.Println("test")
	location := "/Users/zhujiajun/Develop/data.sqlite"
	db, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var city City
	db.Raw("select code, name, provinceCode from city where code = ?", 1301).Scan(&city)

	fmt.Println(city)
}
