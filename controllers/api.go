// Package controllers
// this package keeps the Methods and their implementation
package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/mpalop/test_golang_api/models"
	"github.com/mpalop/test_golang_api/persistence"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Initialize our warehouse
var warehouse *persistence.Warehouse = persistence.InitWarehouse()

// addOrder
// this is the endpoint that loads a POST request with an order and saves it
func addOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}
	theOrder, err := models.BuildOrder(body)
	if err == nil {
		if warehouse.SaveOrder(theOrder) == nil {
			fmt.Printf("Order %d stored successfully\n", theOrder.Id)
		} else {
			http.Error(w, fmt.Sprintf("Order %d already stored. Nothing to do", theOrder.Id ), http.StatusBadRequest)
		}
	} else {
		http.Error(w, fmt.Sprintf("Error loading order: %v", err), http.StatusBadRequest)
	}
}

// getOrder
// This is and endpoint to allows ask for an order and get the order back
// it expects a URL as /order/<order id>
// if the order id does not exists, fails with 404
// if the order id is something but an int, it fails with 400
func getOrder(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	theOrder, err := strconv.Atoi(ps.ByName("OrderId"))
	if err == nil {
		fullOrder, exists := warehouse.GetOrder(theOrder)
		if exists {
			w.Write([]byte(fmt.Sprintf("%s\n", fullOrder)))
		} else {
			http.Error(w, fmt.Sprintf("Order %d does not exists", theOrder), http.StatusNotFound)
		}
	} else {
		http.Error(w, fmt.Sprintf("Some problem arises with the Order Id: %d: %v", theOrder, err), http.StatusBadRequest)
	}
}

// listOrders

// The endpoint that returns the list of orders stored on the warehouse
// it is a simple implementation. On production, there will be a paging system to recover carefully the orders
func listOrders(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	theList := warehouse.GetOrderList()
	w.Write([]byte("List of Orders\n\n"))
	for order := range theList {
		w.Write([]byte(fmt.Sprintf("%v\n", theList[order])))
	}
}

// StartServer
// This is the main function.
//
// It just starts up the http server and server the URL defined:
//
// /order <POST> to add orders
//
// /order <GET> to return the list of orders
//
// /order/:orderId <GET> to get back the order <orderId> or 404 if not found, or 400 if <orderId> is not an int
func StartServer() {
	router := httprouter.New()
	router.POST("/order", addOrder)
	router.GET("/order/:OrderId", getOrder)
	router.GET("/order", listOrders)
	http.ListenAndServe(":8000", router)
}