package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	dbc "github.com/mateo-tavera/shoppingCart/dbConnection"
)

type ArticleList struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
	Qty   int    `json:"qty"`
}

func GetArticles() func(qty, item int) []ArticleList {

	var listOfArticles []ArticleList

	//Using the GET method, we obtain the article list form API
	url := "http://challenge.getsandbox.com/articles"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Cannot Get the url", err)
	}
	defer response.Body.Close()

	//Store the data in a byte slice
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Cannot read the body", err)
	}

	if err := json.Unmarshal(content, &listOfArticles); err != nil {
		panic(err)
	}

	//Convert into type ArticleList
	err = json.Unmarshal(content, &listOfArticles)
	if err != nil {
		log.Fatal("Cannot parse the JSON", err)
	}

	SentToDatabase(listOfArticles)
	var itemsForCart []ArticleList

	//Now that we have the list, we select which and how many articles we want
	return func(qty, item int) []ArticleList {
		listOfArticles[item].Qty = qty
		itemsForCart = append(itemsForCart, listOfArticles[item])
		return itemsForCart
	}

}

func SentToDatabase(listOfArticles []ArticleList) {
	Db, err := dbc.GetDBConnection()
	if err != nil {
		log.Fatal("Cannot get DB connection", err)
	}

	for _, item := range listOfArticles {
		_, err = Db.Exec("INSERT INTO articles (id_article, title, price) VALUES (?, ?, ?)",
			item.Id, item.Title, item.Price)
		if err != nil {
			log.Fatal("Cannot execute query:", err)
		}
	}
	log.Println("Data successfully sent to database")

}
