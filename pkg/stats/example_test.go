package stats

import (
	"fmt"
	"github.com/bahodurnazarov/goTask/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 4047,
			Amount: 200,
		},
		{
			ID: 4043,
			Amount: 100,
		},
		{
			ID: 4041,
			Amount: 300,
		},
		}
	avg := Avg(payments)
	fmt.Println(avg)

	// Output: 200
}
