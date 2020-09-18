package stats

import (
	"github.com/ehsontjk/bank/pkg/bank/types"
)
func Avg(payments []types.Payment) types.Money  {
	sum := types.Money(0)
	max := types.Money(0)
	for _, payment := range payments{
	 
		sum += payment.Amount
		max = sum / types.Money(len(payments))
	}

return max

}