package bridge

import "testing"

func Test(t *testing.T) {
	t.Run("large-milk: ", OrderLargeMilkCoffee)
	t.Run("large-sugar: ", OrderLargeSugarCoffee)
	t.Run("medium-milk: ", OrderMediumMilkCoffee)
	t.Run("smarll-sugar: ", OrderSmallSugarCoffee)
}

// 大杯加奶
func OrderLargeMilkCoffee(t *testing.T) {
	var coffeeAddtion ICoffeeAddtion
	coffeeAddtion = NewCoffeeAddtion(CoffeeAddtionTypeMilk)
	var coffee ICoffee
	coffee = NewCoffee(CoffeeCupTypeLarge, coffeeAddtion)
	// 最终下订单handler
	coffee.OrderCoffee()
}

// 大杯加糖咖啡
func OrderLargeSugarCoffee(t *testing.T) {
	var coffeeAddtion ICoffeeAddtion
	coffeeAddtion = NewCoffeeAddtion(CoffeeAddtionTypeSugar)
	var coffee ICoffee
	coffee = NewCoffee(CoffeeCupTypeLarge, coffeeAddtion)
	// 最终下订单handler
	coffee.OrderCoffee()
}

// 中杯牛奶咖啡
func OrderMediumMilkCoffee(t *testing.T) {
	var coffeeAddtion ICoffeeAddtion
	coffeeAddtion = NewCoffeeAddtion(CoffeeAddtionTypeMilk)
	var coffee ICoffee
	coffee = NewCoffee(CoffeeCupTypeMedium, coffeeAddtion)
	// 最终下订单handler
	coffee.OrderCoffee()
}

// 小杯加糖咖啡
func OrderSmallSugarCoffee(t *testing.T) {
	var coffeeAddtion ICoffeeAddtion
	coffeeAddtion = NewCoffeeAddtion(CoffeeAddtionTypeSugar)
	var coffee ICoffee
	coffee = NewCoffee(CoffeeCupTypeSmall, coffeeAddtion)
	// 最终下订单handler
	coffee.OrderCoffee()
}