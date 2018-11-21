package main

import (
	"fmt"
	"test_golang_api/orders"
	"testing"
)

func TestParseOrder(t *testing.T) {

	var order_sample = []byte(`
{
  "order": {
    "id": 1,
    "store_id": 20,
    "lines": [
      {
        "line_number": 1,
        "sku": "blue_sock"
      },
      {
        "line_number": 2,
        "sku": "red_sock"
      }
    ]
  }
}
`)
	var order_sample_bad_numbers = []byte(`
{
  "order": {
    "id": 1,
    "store_id": 20,
    "lines": [
      {
        "line_number": 3,
        "sku": "blue_sock"
      },
      {
        "line_number": 1,
        "sku": "red_sock"
      }
    ]
  }
}
`)


	var order_bad = []byte(`
{
  "order": {
    "id": 1,
    "store_id": 20,
    "lines": [
      {
        "line_number": 1
        "sku": "blue_sock"
      },
      {
        "line_number": 2,
        "sku": "red_sock"
      }
    ]
  }
}
`)

	theOrder, err := orders.Build(order_sample)
	if err != nil {
		t.Errorf("%v, %v\n", theOrder, err)
	}
	theOrder2, err := orders.Build(order_sample_bad_numbers)
	if err != nil {
		t.Errorf("%v %v\n", theOrder2, err)
	}

	fail, err := orders.Build(order_bad)
	if err != nil {
		fmt.Printf("Detecting bad JSON OK\n")
	} else {
		t.Errorf("No fail detected\n %v", fail)
	}

}