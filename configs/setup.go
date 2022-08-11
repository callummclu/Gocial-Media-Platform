package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDB() (*sql.DB, error) {
	host := getEnvByName("DB_HOST")
	port := getEnvByName("DB_PORT")
	user := getEnvByName("DB_USER")
	password := getEnvByName("DB_PASS")
	dbname := getEnvByName("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db, err
}
