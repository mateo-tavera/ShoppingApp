package serverConnection

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateo-tavera/shoppingCart/service"
)

func GetServerConnection() {
	//Init Router
	r := mux.NewRouter()

	//Route handlers / Endpoints
	r.HandleFunc("/api/shopping-cart", service.GetCarts).Methods(http.MethodGet)
	r.HandleFunc("/api/shopping-cart/{id}", service.GetCart).Methods(http.MethodGet)
	r.HandleFunc("/api/shopping-cart", service.CreateCart).Methods(http.MethodPost)
	r.HandleFunc("/api/shopping-cart/{id}", service.UpdateCart).Methods(http.MethodPut)
	r.HandleFunc("/api/shopping-cart/{id}", service.DeleteCart).Methods(http.MethodDelete)

	//Initilize the server
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
