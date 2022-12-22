package p_decorator

import "fmt"

/*
コンポーネントのインターフェース
*/
type IPizza interface {
	getPrice() int
}

/*
具象コンポーネント
*/
type VeggeMania struct{}

func (p *VeggeMania) getPrice() int {
	return 15
}

/*
具象デコレーター
*/
type CheeseTopping struct {
	pizze IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizze.getPrice()
	return pizzaPrice + 10
}

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

/*
クライアントコード
*/
func main() {
	pizza := &VeggeMania{}

	// add cheese topping
	pizzaWithCheese := &CheeseTopping{pizze: pizza}

	// add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithCheese}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
