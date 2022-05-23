package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	dbc "github.globant.com/mateo-tavera/shoppingCart/dbConnection"
	class "github.globant.com/mateo-tavera/shoppingCart/entity"
	server "github.globant.com/mateo-tavera/shoppingCart/serverConnection"
)

var Db *sql.DB
var err error

func main() {

	//Set db connection
	Db, err = dbc.GetDBConnection()
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	fmt.Println("connected to database")
	defer Db.Close()

	//Delete previus data
	_, err = Db.Exec("TRUNCATE articles")
	if err != nil {
		log.Fatal("Cannot execute query:", err)
	}
	_, err = Db.Exec("TRUNCATE cart")
	if err != nil {
		log.Fatal("Cannot execute query:", err)
	}

	//Get the list of articles provided from API
	addItems := class.GetArticles()
	//Add articles manually
	var items []class.ArticleList
	items = addItems(3, 3)
	items = addItems(4, 4)
	items = addItems(2, 0)
	items = addItems(1, 4)

	//Data to create a cart manually
	class.CartList = append(class.CartList, class.Cart{
		IdCart: "1",
		Items:  items})

	class.CartList = append(class.CartList, class.Cart{
		IdCart: "2",
		Items:  items})

	//Start the server
	server.GetServerConnection()

}
