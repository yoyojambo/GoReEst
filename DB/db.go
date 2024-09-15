package db

import (
	"database/sql"
    "os"
    "fmt"
    "sync"
	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once
var dbConnect *sql.DB 

func DB() *sql.DB {
    once.Do(initDB)
    return dbConnect
}

var(
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
	hostname = os.Getenv("HOSTNAME")
	dbname   = os.Getenv("DBNAME")
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func initDB(){
    db, err := sql.Open("mysql", dsn(dbname))

    if err != nil {
	    fmt.Printf("Error %s when opening DB\n", err)
	    return

    }
    dbConnect = db
}
