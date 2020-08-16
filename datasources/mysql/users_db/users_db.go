package users_db

// users_db.go

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "admin"
	mysql_users_password = "19Omega25"
	mysql_users_host     = "127.0.0.1:3306"
	mysql_users_schema   = "users"
)

var (
	Client *sql.DB

	username = os.Getenv("mysql_users_username")
	password = os.Getenv("mysql_users_password")
	hostname = os.Getenv("mysql_users_host")
	schema   = os.Getenv("mysql_users_schema")
)

func init() {

	fmt.Println(username, password, hostname, schema)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		mysql_users_username,
		mysql_users_password,
		mysql_users_host,
		mysql_users_schema,
	)

	var err error

	Client, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Connected to DB users_db!")
}
