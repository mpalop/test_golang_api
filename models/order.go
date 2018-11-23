package models

import (
	"encoding/json"
	"errors"
)

type MainOrder struct {
	Order		Order		`json:"order"`
}

type Order struct {
	Id 			int 		`json:"id"`
	StoreId		int			`json:"store_id"`
	Lines		[]Line 		`json:"lines"`
}

type Line struct {
	LineNumber	int			`json:"line_number"`
	Sku			string		`json:"sku"`
}

func BuildOrder(orderString []byte) (*MainOrder,error) {
	var NewOrder MainOrder
	err := json.Unmarshal(orderString, &NewOrder)

	if err == nil {
		err := NewOrder.validateOrder()
		return &NewOrder, err
	}
	return &NewOrder, err
}

func (order *MainOrder) validateOrder() error {
	// checking the line numbers

	ordered:=true
	counter:=1
	lines:=order.Order.Lines

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
