package stats

import (
	"fmt"
	"github.com/bahrom656/bank/pkg/types"
)

func ExampleAvg() {
	cards := []types.Payment{
		{
			ID:       1,
			Amount:   20_000_00,
			Category: "car",
		},
		{
			ID:       2,
			Amount:   10_000_00,
			Category: "phone",
		},
	}
	fmt.Println(Avg(cards))
	//Output:
	//1500000
}
