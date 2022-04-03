package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "David"
	password = "David123"
	dbname   = "neobar"
)

// Opening a database and save the reference to `Database` struct.
func Init() *sql.DB {
	//connStr := fmt.Sprintf("%s://%s:%s@%s:%d/lcp?sslmode=verify-full", dbname, user, password, host, port)

	herokuEnv := os.Getenv("DATABASE_URL")
	var connStr string
	if herokuEnv == "" {
		connStr = "user=David password=David123 dbname=neobar port=5432 sslmode=disable"
	} else {
		connStr = herokuEnv
	}

	//connStr := fmt.Sprintf("host=%s port=%d user=%s "+
	//"password=%s dbname=%s sslmode=disable",
	//host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	db.SetMaxIdleConns(10)
	return db
}
