package sample

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	UserId       int
	Title        string
	Content      string
	UpdatedTimes int
}

func Sample() {
	db, err := gorm.Open("mysql", "root:@/goose_practice")

	if err != nil {
		panic("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(&Article{})

	// Create
	db.Create(&Article{UserId: 1, Title: "gorm", Content: "gotrm", UpdatedTimes: 2})

	// Read
	var article Article
	db.First(&article, 1)
	db.First(&article, "title = ?", "gorm")

	db.Model(&article).Update("Title", "gorm2")

	db.Delete(&article)
}
