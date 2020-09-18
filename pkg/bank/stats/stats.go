package stats

import (
	"github.com/ehsontjk/bank/pkg/bank"
)
func Avg(payments []types.Payment) types.Money  {
	max := payments[0]
	for _, payment := range payments {
		if max.Amount < payment.Amount { 
		

			max = payment
		}
	}


	return max
}