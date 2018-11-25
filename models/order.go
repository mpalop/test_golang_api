// Package models
// this package deals with the definition and main functions to handle the Orders
//
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// MainOrder represents the Order as it comes to the system.
// is just a container to the real Order
//
type MainOrder struct {
	Order		Order		`json:"order"`			// Order struct
}

// Order is the real order
type Order struct {
	Id 			int 		`json:"id"`				// Order Id <int>
	StoreId		int			`json:"store_id"`		// Stored Id <int>
	Lines		[]Line 		`json:"lines"`			// Array of order lines
}

// Line represent each line of the order
type Line struct {
	LineNumber	int			`json:"line_number"`	// LineNumber <int>
	Sku			string		`json:"sku"`			// Sku is the product identifier <string>
}

// BuildOrder
// The constructor. This function gets the bunch of bytes with the encoded json with the Order,
// unmarshals it, checks the order of the lines and returns and Order object
//
// Returns: A pointer to the Order with the info and a (possible) error or nil
//
// Possible errors:
//
// * object is not properly formatted, it will fail the marshalling
//
// * order's line numbers are not well set
func BuildOrder(orderString []byte) (*Order,error) {
	var NewOrder MainOrder
	err := json.Unmarshal(orderString, &NewOrder)

	if err == nil {
		err := NewOrder.Order.validateOrder()
		return &NewOrder.Order, err
	}
	return &NewOrder.Order, err
}

// String
// this allows this Order object to implement the interface Stringer.
//
// Returns a string representing the object
func (order *Order) String() string {
	var sb strings.Builder
	for _,line := range order.Lines {
		sb.WriteString(fmt.Sprintf("-line:%3d SKU: %-20s\n", line.LineNumber, line.Sku))
	}
	return fmt.Sprintf("Order: %5d, StoreId: %5d\n%s", order.Id, order.StoreId, sb.String())
}

// validateOrder
// This internal method checks the order of the lines to validate if the order is correct or no
//
// Returns a (possible) Error if lines are disordered or unconsecutive
func (order *Order) validateOrder() error {
	ordered:=true
	counter:=1
	lines:=order.Lines

	for pos:=0; pos< len(lines); pos++ {
		if lines[pos].LineNumber == counter {
			counter++
		} else {
			ordered = false
			break
		}
	}
	if !ordered {
		return errors.New("line numbers are not ordered")
	} else {
		return nil
	}
}
