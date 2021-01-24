package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var (
	Client *sql.DB
)

const (
	host     = "database"
	user     = "root"
	password = "toor"
	dbname   = "users_db"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, password, host, dbname)
	var err error

	for i := 0; i < 5; i++ {
		Client, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Println("Cant Open Database == ", err)
			time.Sleep(2 * time.Second)
			if i == 5 {
				panic(err)
			}
		} else {
			break
		}
	}

	for j := 0; j < 5; j++ {
		if err = Client.Ping(); err != nil {
			log.Println(fmt.Sprintf("Cant Ping Database retry %d:%d", j, 5))
			time.Sleep(2500 * time.Millisecond)
			if j == 5 {
				panic(err)
			}
		} else {
			break
		}
	}

	log.Println("Database successfully configured")
}
