package main

import (
	"log"
	"net/http"

	"./controllers/category"
	"./controllers/product"
	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", product.QueryProducts)
	myRouter.HandleFunc("/categories", category.QueryCategories)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
