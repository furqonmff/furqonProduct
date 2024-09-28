package Config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "root:@/furqon_product?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database")
	DB = db
}
