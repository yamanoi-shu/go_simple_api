package controller

import (
	"encoding/json"
	"fmt"
	"go_simple_api/model"
	"io/ioutil"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (*UserController) Create(w http.ResponseWriter, r *http.Request) {

	reqBody := make(map[string]string)
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bodyBytes, &reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errRes := map[string]string{
			"massage": "requset params is invalid",
		}
		resBody, _ := json.Marshal(errRes)
		fmt.Fprintln(w, string(resBody))
		return
	}
	firstName := reqBody["first_name"]
	lastName := reqBody["last_name"]
	user, err := model.CreateUser(firstName, lastName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errRes := map[string]string{
			"massage": "user could not be created",
		}
		resBody, _ := json.Marshal(errRes)
		fmt.Fprintln(w, string(resBody))
		return
	}
	fmt.Fprintln(w, user)
}
