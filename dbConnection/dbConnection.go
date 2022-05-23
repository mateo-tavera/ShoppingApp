package dbConnection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func GetDBParameters() {
	envErr := godotenv.Load("./dbConnection/.env")
	if envErr != nil {
		log.Fatal("Could not load .env file:", envErr)
	}

}

func GetDBConnection() (*sql.DB, error) {

	GetDBParameters()

	var (
		host     = os.Getenv("GLBHOST")
		port     = os.Getenv("GLBPORT")
		user     = os.Getenv("GLBUSER")
		password = os.Getenv("GLBPASSWORD")
		dbname   = os.Getenv("GLBDBNAME")
	)

	mysqlConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)

	return sql.Open("mysql", mysqlConnection)
}
