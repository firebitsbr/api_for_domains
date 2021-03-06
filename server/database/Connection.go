package database

import (
    "database/sql"
    "net/http"
    "time"

    _ "github.com/lib/pq"
    errors "github.com/kliver98/api_for_domains/server/error"
)

const SQL string = "postgresql://root@localhost:26257/reto_prueba?sslmode=disable"
const PING_DATABASE = "http://localhost:5001"


func GetConnection() (*sql.DB, error) {

	// Connect to the "reto_prueba" database.
    db, err := sql.Open("postgres", SQL)

    if IsDown() {
    	err = &errors.NoPingError{Message: "Unreachable to database "+PING_DATABASE}
    }

    return db, err

}

func IsDown() bool {
	timeout := time.Duration(400 * time.Millisecond)
	client := http.Client{
	    Timeout: timeout,
	}
	_, err := client.Get(PING_DATABASE)
	if err != nil {
	    return true
	} else {
	    return false
	}
}