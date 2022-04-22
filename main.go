package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/create", createHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}

type User struct {
	Name string
}

func createHandler(res http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	var user User

	err = json.Unmarshal(reqBody, &user)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", user.Name)
}
