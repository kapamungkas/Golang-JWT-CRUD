package databases

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	db_user := os.Getenv("MYSQL_DB_USER")
	db_name := os.Getenv("MYSQL_DB_NAME")
	db_password := os.Getenv("MYSQL_DB_PASSWORD")

	db, err := sql.Open("mysql", db_user+":"+db_password+"@tcp("+host+":"+port+")/"+db_name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("CONECTED TO DATABASE")

	return db, nil
}
