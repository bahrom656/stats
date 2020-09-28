package stats

import "github.com/bahrom656/bank/pkg/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	if len(payments) == 0 {
		return payments[0].Amount
	}
	var sum types.Money
	var count int
	for i := 0; i < len(payments); i++ {
		sum = sum + payments[i].Amount
		count++
	}
	return sum / types.Money(count)
}
//TotalInCategory находят сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for i := 0; i < len(payments); i++ {
		if payments[i].Category == category{
			sum = sum + payments[i].Amount			
		}
	}
	return sum
}