package routes

import (
	"fmt"
	"integrasia/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Routing() {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Fibo).Methods("POST")

	fmt.Println("PORT: 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
