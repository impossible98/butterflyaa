package db

import (
	"fmt"
	"os"

	// import third-party packages
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	// import local packages
	"github.com/impossible98/butterfly/internal/model"
	"github.com/impossible98/butterfly/pkg/path"
)

func InitDb() *gorm.DB {
	exist := path.PathExists("./data")

	if !exist {
		err := os.Mkdir("./data", os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
	db, err := gorm.Open(sqlite.Open("data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 迁移 schema
	db.AutoMigrate(&model.Product{})
	// Create
	db.Create(&model.Product{Code: "D42", Price: 100})
	return db
}
