package utils

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

type Reddit struct {
	gorm.Model
	Code  string
	Price uint
}

type HistoryEntry struct {
	gorm.Model
	url     string
	title   string
	visited *time.Time
	browser string
}

type Movie struct {
	gorm.Model
	title       string
	releaseYear *time.Time
	director    string
}

type Book struct {
	gorm.Model
	title       string
	releaseYear *time.Time
	author      string
	source      string
	links       string
}

type Bookmark struct {
	gorm.Model
	location     string
	orgHierarchy string
	source       string
}

func init() {

	db, err := gorm.Open("sqlite3", "funnl.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Reddit{})
	db.AutoMigrate(&HistoryEntry{})
	db.AutoMigrate(&Movie{})
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Bookmark{})

}
