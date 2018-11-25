package tests

import (
	"github.com/mpalop/test_golang_api/models"
	"github.com/mpalop/test_golang_api/persistence"
	"reflect"
	"testing"
)

func TestPersistence(t *testing.T) {
	//Test1: load file.
	var warehouse *persistence.Warehouse = persistence.InitWarehouse()

	test1 := loadFile("./fixtures/sample1.json")
	t.Log("Testing store normal file")
	theOrder, err := models.BuildOrder(test1)
	if err != nil {
		t.Errorf("%v, %v\n", theOrder, err)
	}

	err2 := warehouse.SaveOrder(theOrder)
	if err2 == nil {
		savedOrder, exists := warehouse.GetOrder(theOrder.Id)
		if exists {
			if reflect.DeepEqual(theOrder, savedOrder) {
				t.Logf("Order %d was saved correctly. OK", theOrder.Id)
			} else {
				t.Errorf("Order %d was saved wronly. KO", theOrder.Id)
			}
		} else {
			t.Errorf("Order %d was not saved. KO", theOrder.Id)
		}
	} else {
		t.Errorf("Order %d was not saved. The Error is %v", theOrder.Id, err2)
	}

	t.Log("Testing store normal file twice")
	err3 := warehouse.SaveOrder(theOrder)
	if err3 != nil {
		t.Logf("Order %d returns error when tried to save twice: %v. OK", theOrder.Id, err3)
	} else {
		t.Errorf("Order %d was saved twice. KO", theOrder.Id)
	}
}







