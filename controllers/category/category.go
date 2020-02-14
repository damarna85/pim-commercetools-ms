package category

import (
	"encoding/json"
	"log"
	"net/http"

	"../../client"
	"github.com/labd/commercetools-go-sdk/commercetools"
)

// QueryCategories controller for the Home link
func QueryCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := client.GetClient().CategoryQuery(&commercetools.QueryInput{})
	if err != nil {
		log.Fatal("Error getting the categories")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(categories.Results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
