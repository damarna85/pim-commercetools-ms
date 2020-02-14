package product

import (
	"encoding/json"
	"log"
	"net/http"

	"../../client"
	"github.com/labd/commercetools-go-sdk/commercetools"
)

// QueryProducts controller for the Home link
func QueryProducts(w http.ResponseWriter, r *http.Request) {
	products, err := client.GetClient().ProductQuery(&commercetools.QueryInput{})
	if err != nil {
		log.Fatal("Error getting the products")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(products.Results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
