package entity

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	dbc "github.com/mateo-tavera/shoppingCart/dbConnection"

	"github.com/gorilla/mux"
)

type Cart struct {
	IdCart string        `json:"id_cart"`
	Items  []ArticleList `json:"items"`
}

//Init cart variable
var CartList []Cart

//Get all carts
func GetCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CartList)
}

//Get single cart
func GetCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get parameters
	//Loop through carts and find with Id
	for _, item := range CartList {
		if item.IdCart == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Cart{})
}

//Create a new cart
func CreateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cart Cart
	_ = json.NewDecoder(r.Body).Decode(&cart)
	neoID, _ := strconv.Atoi(CartList[len(CartList)-1].IdCart)
	cart.IdCart = strconv.Itoa(neoID + 1)

	CartList = append(CartList, cart)
	json.NewEncoder(w).Encode(cart)
	fmt.Printf("cart %v added\n", neoID+1)

	//Get db
	Db, err := dbc.GetDBConnection()
	if err != nil {
		log.Fatal("Cannot get DB connection", err)
	}

	//Send each id article
	for _, value := range cart.Items {
		_, err = Db.Exec("INSERT INTO cart (id_cart, id_article, qty_item) VALUES (?, ?, ?)",
			cart.IdCart, value.Id, value.Qty)
		if err != nil {
			log.Fatal("Cannot execute query:", err)
		}
	}
}

//Update a cart
func UpdateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//Loop through carts and find with Id
	var cart Cart
	for index, item := range CartList {
		if item.IdCart == params["id"] {
			CartList = append(CartList[:index], CartList[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&cart)
			cart.IdCart = params["id"]
			CartList = append(CartList, cart)
			json.NewEncoder(w).Encode(cart)

			fmt.Println("Cesperando a seguir..")
			fmt.Scanln()

			//Get db
			Db, err := dbc.GetDBConnection()
			if err != nil {
				log.Fatal("Cannot get DB connection", err)
			}
			//Update article parameters
			fmt.Printf("cart %v updated\n", params["id"])
			for _, value := range cart.Items {
				_, err = Db.Exec("UPDATE cart SET id_article = ?, qty_item = ? WHERE id_cart = ?",
					value.Id, value.Qty, params["id"])
				if err != nil {
					log.Fatal("Canot execute query:", err)
				}

			}

			return
		}
	}
	json.NewEncoder(w).Encode(CartList)

}

//Delete a cart
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//Loop through cart and find with Id
	for index, item := range CartList {
		if item.IdCart == params["id"] {
			CartList = append(CartList[:index], CartList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(CartList)

	//Get db
	Db, err := dbc.GetDBConnection()
	if err != nil {
		log.Fatal("Cannot get DB connection", err)
	}

	//Delete each id article
	fmt.Printf("cart %v deleted\n", params["id"])
	_, err = Db.Exec("DELETE FROM cart WHERE id_cart = ?", params["id"])
	if err != nil {
		log.Fatal("Canot execute query:", err)
	}
}
