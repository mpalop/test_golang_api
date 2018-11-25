// Package persistence
// This package deals with all related with persist in memory the Orders
package persistence

import (
	"errors"
	"fmt"
	"github.com/mpalop/test_golang_api/models"
	"sort"
	"sync"
)

// Type Warehouse
// the base object to store all data
type Warehouse struct {
	Ids			map[int]*models.Order		// map <order_id> -> Order object
	mutex		sync.Mutex					// implements mutex objects to avoid sync problems
}

// InitWarehouse
// The constructor. Initializes the map and returns the new object Warehouse
func InitWarehouse() *Warehouse {
	newMap := make(map[int]*models.Order)
	return &Warehouse{
		Ids: newMap,
	}
}

// SaveOrder
// This method gets an Order and tries to save it.
//
// Returns a (possible) error if the Order is already Stored
func (w *Warehouse) SaveOrder(order *models.Order) error {
	if !w.Exists(order.Id) {
		w.mutex.Lock()
		defer w.mutex.Unlock()
		w.Ids[order.Id] = order
		return nil
	} else {
		return errors.New(fmt.Sprintf("Order %d already exists", order.Id))
	}
}

// Exists
// This method allows knowing if one Order id is already stored in the Warehouse
//
// Returns: True/False depending if the order id is already on the Warehouse
func (w *Warehouse) Exists(orderNumber int) bool {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	_, ok := w.Ids[orderNumber]
	return ok
}

// GetOrder
// This method returns an Order object with the Id passed by parameter
//
// Returns: The Order object that has this Id, and True/False if the Id exists in the Warehouse
func (w *Warehouse) GetOrder(orderNumber int) (*models.Order, bool) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	order, ok := w.Ids[orderNumber]
	return order, ok
}

// GetOrderList
// This method is a simplified version of a List of the Orders that are in the Warehouse
//
// Returns: It returns an ordered array of pointers to the orders that are in the Warehouse
func (w *Warehouse) GetOrderList() []*models.Order {
	var ret []*models.Order
	w.mutex.Lock()
	defer w.mutex.Unlock()

	// getting keys and sort them
	var keys []int
	for k := range w.Ids {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// build list in proper order
	for _, k := range keys {
		ret = append(ret, w.Ids[k])
	}
	return ret
}