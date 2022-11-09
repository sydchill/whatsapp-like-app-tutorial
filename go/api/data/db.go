package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var ConnectionDB = make(map[string]*sql.DB)

type Connect struct {
	DB *sql.DB
	TX *sql.Tx
}

func Connection() error {
	// load .env
	var err = godotenv.Load(".env")
	if err != nil {
		return err
	}

	dns := fmt.Sprintf("user=%s host=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DBUSER"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBNAME"))

	connect, err := sql.Open("postgres", dns)

	if err != nil {
		return err
	}
	connect.SetMaxIdleConns(10)
	connect.SetMaxOpenConns(30)

	err = connect.Ping()
	if err != nil {
		return err
	}
	ConnectionDB["postgres"] = connect
	return nil
}

func Open() (*Connect, error) {

	connect := &Connect{DB: ConnectionDB["postgres"]}

	tx, err := connect.DB.Begin()

	if err != nil {

		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	connect.TX = tx

	return connect, nil
}

type Mapping[T any] func(*sql.Rows) (T, error)

func Query[T any, S []T](da *Connect, queryString string, mapping Mapping[T], args ...any) (S, error) {

	if mapping == nil {

		_, err := da.TX.Exec(queryString, args...)

		return nil, err
	}

	rows, err := da.TX.Query(queryString, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make(S, 0)
	for rows.Next() {
		result, err := mapping(rows)

		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil

}
