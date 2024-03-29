package serverConnection

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	svc "github.com/mateo-tavera/shoppingCart/service"
)

func GetServerConnection() {
	//Init Router
	r := mux.NewRouter()

	//Route handlers / Endpoints
	r.HandleFunc("/api/shopping-cart", svc.GetCarts).Methods(http.MethodGet)
	r.HandleFunc("/api/shopping-cart/{id}", svc.GetCart).Methods(http.MethodGet)
	r.HandleFunc("/api/shopping-cart", svc.CreateCart).Methods(http.MethodPost)
	r.HandleFunc("/api/shopping-cart/{id}", svc.UpdateCart).Methods(http.MethodPut)
	r.HandleFunc("/api/shopping-cart/{id}", svc.DeleteCart).Methods(http.MethodDelete)

	//Initilize the server
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
