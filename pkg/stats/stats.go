package stats

import (
	"github.com/bahodurnazarov/goTask/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	for _, payment := range payments {
		sum += payment.Amount
	}
	return sum / types.Money(len(payments))
}