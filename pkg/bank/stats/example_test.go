package stats


import (
	"github.com/ehsontjk/bank/pkg/bank/types"
	"fmt"
)
func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 12,
			Amount: 40_000_00,
			Category: "bank",
		},
		{
			ID: 13,
			Amount: 20_000_00,
			Category: "bank",
		},
		{
			ID: 34,
			Amount: 48_000_00,
			Category: "bank",	
		},
		{
ID: 5,
Amount: 48_000_00,
	},
	{
		ID: 45,
		Amount: 41_000_00,
		Category: "bank",
			},

			
		}
	

	max := Avg(payments)
	fmt.Println(max)
	// Output: 3940000
}