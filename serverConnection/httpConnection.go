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
	r.HandleFunc("/api/shopping-cart", entity.GetCarts).Methods(http.MethodGet)
	r.HandleFunc("/api/shopping-cart/{id}", entity.GetCart).Methods(http.MethodGet)
	r.HandleFunc("/api/shopping-cart", entity.CreateCart).Methods(http.MethodPost)
	r.HandleFunc("/api/shopping-cart/{id}", entity.UpdateCart).Methods(http.MethodPut)
	r.HandleFunc("/api/shopping-cart/{id}", entity.DeleteCart).Methods(http.MethodDelete)

	//Initilize the server
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
