package stats

import "github.com/bahrom656/bank/v2/pkg/types"

// Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	if len(payments) == 0 {
		return payments[0].Amount
	}
	var sum types.Money
	var count int
	for i := 0; i < len(payments); i++ {
		if payments[i].Status != types.StatusFail {
			sum = sum + payments[i].Amount
			count++
		}
	}
	return sum / types.Money(count)
}

//TotalInCategory находят сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for i := 0; i < len(payments); i++ {

		if payments[i].Category == category && payments[i].Status != types.StatusFail{
			sum = sum + payments[i].Amount			
		}
	}
	return sum
}

//CategoriesAvg считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	counts := map[types.Category] types.Money {}

	for _, payment := range payments {
		counts[payment.Category]++ 
		categories[payment.Category] += payment.Amount 
	}
	
	for value := range categories { 
		categories[value] /= counts[value] 
	}
	return categories
}

//PeriodsDynamic сравнивает расходы по котегориям за два периода
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	result := map[types.Category]types.Money{}
	if len(first) > len(second) {
		for key := range first {
			result[key] = second[key] - first[key]
		}
	}else {
		for key := range second {
			result[key] = second[key] - first[key]
		}
	}
	return result
}