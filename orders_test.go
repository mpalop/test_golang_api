package main

import (
	"fmt"
	"github.com/mpalop/test_golang_api/models"
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

	theOrder, err := models.BuildOrder(order_sample)
	if err != nil {
		t.Errorf("%v, %v\n", theOrder, err)
	}
	fmt.Printf("%s", theOrder)

	theOrder2, err := models.BuildOrder(order_sample_bad_numbers)
	if err == nil {
		t.Errorf("%v %v\n", "Not detected lack of line number order", theOrder2)
	} else {
		fmt.Printf("%v %v\n", theOrder2, "Detected lack of line number order")
	}

	fail, err := models.BuildOrder(order_bad)
	if err != nil {
		fmt.Printf("Detecting bad JSON OK\n")
	} else {
		t.Errorf("No fail detected\n %v", fail)
	}

}
