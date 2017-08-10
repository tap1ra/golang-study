package main

import (
	"fmt"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Event struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Flag *bool  `json:"flag"`
}

func gormConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root@tcp(localhost:3306)/sample")

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	db := gormConnect()
	ptrue := &[]bool{true}[0]
	pfalse := &[]bool{false}[0]

	// 初期データ作成
	eventEx := Event{Id: 1, Name: "taro", Flag: ptrue}
	result := db.Create(eventEx)
	fmt.Println(result)

	// 初期データのID=1を更新する
	eventExBefore := Event{Id: 1}
	db.First(&eventExBefore)
	eventExAfter := Event{Name: "jiro", Flag: pfalse, Id: 10}
	db.Model(&eventExBefore).UpdateColumns(&eventExAfter)
}
