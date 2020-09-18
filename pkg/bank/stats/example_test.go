package stats


import (
	"github.com/ehsontjk/bank/pkg/bank/types"
)
func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 12,
			Amount: 40_000_00,
		},
		{
			ID: 13,
			Amount: 20_000_00,
		},
		{
			ID: 34,
			Amount: 48_000_00,	
		},
		{
ID: 5,
Amount: 48_000_00,
	},
	{
		ID: 45,
		Amount: 41_000_00,
			},

			
		}
	

	max := Max(payments)
	fmt.Println(max.ID)
	// Output: 34
}