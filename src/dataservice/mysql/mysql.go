package mysql

import (
	"fmt"
	"golang-backend-microservice/container/log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	Host string
	User string
	Pass string
}

func OpenMySqlConnection() *sqlx.DB {
	return Connection{
		Host: os.Getenv("MYSQL_HOST"),
		User: os.Getenv("MYSQL_USER"),
		Pass: os.Getenv("MYSQL_PASS"),
	}.Open()
}

func (c Connection) Open() *sqlx.DB {
	// Open connection
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@/%s", c.User, c.Pass, c.Host))
	if err != nil {
		log.Error("Error connecting to MySQL - %s", err.Error())
		return nil
	}
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(3000)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	// Test connection by pinging
	if err := db.Ping(); err != nil {
		log.Error("Error connecting to MySQL - %s", err.Error())
		return nil
	}

	return db
}
