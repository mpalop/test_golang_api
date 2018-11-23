package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mpalop/test_golang_api/models"
	"io/ioutil"
	"net/http"
)

func addOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}
	theOrder, err := models.BuildOrder(body)
	fmt.Printf("%v, %v\n", theOrder, err)
}


func StartServer() {
	router := httprouter.New()
	router.POST("/order", addOrder)
	http.ListenAndServe(":8000", router)
}