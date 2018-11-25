package tests

import (
	"github.com/mpalop/test_golang_api/models"
	"testing"
)

func TestParseOrder(t *testing.T) {
	//Test1: normal file.
	test1 := loadFile("./fixtures/sample1.json")
	t.Log("Testing normal file")
	theOrder, err := models.BuildOrder(test1)
	if err != nil {
		t.Errorf("%v, %v\n", theOrder, err)
	}
	t.Logf("Order processed OK:\n %s", theOrder)

	//Test2: bad syntax file.
	test2 := loadFile("./fixtures/sample_bad_build.json")
	t.Log("Testing bad syntax file")
	theOrder2, err := models.BuildOrder(test2)
	if err == nil {
		t.Errorf("Not detected bad syntax order: %s", theOrder2)
	} else {
		t.Logf("Detected bad syntax OK. Error %v", err)
	}

	//Test3: bad order lines file.
	test3 := loadFile("./fixtures/sample_bad_order.json")
	t.Log("Testing bad order lines file")
	theOrder3, err := models.BuildOrder(test3)
	if err == nil {
		t.Errorf("Not detected bad lines order order: %s", theOrder3)
	} else {
		t.Logf("Detected bad order lines OK. Error %v", err)
	}

	//Test4: bad structure file.
	test4 := loadFile("./fixtures/sample_bad_struct.json")
	t.Log("Testing bad structure file")
	theOrder4, err := models.BuildOrder(test4)
	if err == nil {
		t.Errorf("Not detected bad structure: %s", theOrder4)
	} else {
		t.Logf("Detected bad structure OK. Error %v", err)
	}
}
