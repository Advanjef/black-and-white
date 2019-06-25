package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	openingdao "github.com/hunhunyosshy/black-and-white/pkg/dao/opening"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/opening/", showOpeningIndex)
	r.HandleFunc("/opening/{id:[0-9]+}/", showOpeningByKey)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func jsonEncode(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	return string(bytes)
}

func showOpeningIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(jsonEncode(openingdao.FetchIndex())))
}

func showOpeningByKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Write([]byte(jsonEncode(openingdao.FetchByKey(id))))
}
