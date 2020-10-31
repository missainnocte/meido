package store

import (
	"encoding/json"
	"github.com/ivanh/meido/persistence"
	"net/http"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	//err := createItemTable(persistence.GetInstance().GetDb())
	//if err != nil {
	//	w.WriteHeader(http.StatusForbidden)
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	items, err := getItems(persistence.GetInstance().GetDb())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jItems, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(jItems)
}
