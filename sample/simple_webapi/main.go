package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Example Post WebAPI
func Example(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer r.Body.Close()

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type ExampleParameter struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	var param ExampleParameter

	err = json.Unmarshal(bodyBytes, &param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%+v\n", param))
}

func main() {
	router := httprouter.New()

	router.POST("/Example", Example)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
