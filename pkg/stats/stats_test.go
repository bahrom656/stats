package stats

import (
	"github.com/bahrom656/bank/v2/pkg/types"
	"testing"
	"reflect"
)

func TestCategoriesAvg(t *testing.T) {
	var excepted types.Money
	payments := []types.Payment {
		{ID: 1, Category: "auto", Amount: 10_000_00},
		{ID: 2, Category: "food", Amount: 20_000_00,},
		{ID: 3, Category: "auto", Amount: 30_000_00,},
		{ID: 4, Category: "auto", Amount: 40_000_00,},
		{ID: 5, Category: "fun", Amount: 50_000_00,},
	}
	excepted = 26_666_66
	result := CategoriesAvg(payments)
	
	if !reflect.DeepEqual(excepted, result["auto"]) {
		t.Error("invalid result")
	}
}
func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money {
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money {
		"auto": 5,
		"food": 3,
	}
	result := map[types.Category]types.Money {
		"auto": -5,
		"food": -17,
	}
	
	sum := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(result, sum) {
		t.Error("invalid result")
	}
}

