package serverConnection

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateo-tavera/shoppingCart/entity"
)

func GetServerConnection() {
	//Init Router
	r := mux.NewRouter()

	//Route handlers / Endpoints
	r.HandleFunc("/api/shopping-cart", entity.GetCarts).Methods("GET")
	r.HandleFunc("/api/shopping-cart/{id}", entity.GetCart).Methods("GET")
	r.HandleFunc("/api/shopping-cart", entity.CreateCart).Methods("POST")
	r.HandleFunc("/api/shopping-cart/{id}", entity.UpdateCart).Methods("PUT")
	r.HandleFunc("/api/shopping-cart/{id}", entity.DeleteCart).Methods("DELETE")

	//Initilize the server
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
