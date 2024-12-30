package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ConfigConnectionDB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func ConnectDB(configConnection ConfigConnectionDB) (*sql.DB, error) {
	strConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		configConnection.Host, configConnection.Port, configConnection.User, configConnection.Password, configConnection.DBName)

	db, err := sql.Open("postgres", strConn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado com o banco: " + configConnection.DBName)

	return db, nil
}
