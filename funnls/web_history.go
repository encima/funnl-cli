package funnls

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func WbGet(browser string) {

	switch browser {
	case "safari":
		historyDBPath := os.ExpandEnv(viper.GetString("funnls.web_history.browsers.safari.path"))
		getHist(historyDBPath, viper.GetString("funnls.web_history.browsers.safari.statement"))
	}
}

func getHist(path string, query string) {
	historyDB, _ := sql.Open("sqlite3", path)
	defer historyDB.Close()
	row, err := historyDB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var url string
		var title string
		var datetime string
		row.Scan(&url, &title, &datetime)
		log.Println("Entry: ", url, " ", title, " ", datetime)
	}

}
