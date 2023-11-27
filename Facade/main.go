package main

import (
	"fmt"
	facade "go_design_patterns/Facade/pkg"
)

var (
	bank = facade.Bank{
		Name:  "Банк",
		Cards: []facade.Card{},
	}

	card1 = facade.Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = facade.Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}

	user1 = facade.User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = facade.User{
		Name: "Покупатель-2",
		Card: &card2,
	}

	prod = facade.Product{
		Name:  "Сыр",
		Price: 150,
	}

	shop = facade.Shop{
		Name:     "Shop",
		Products: []facade.Product{prod},
	}
)

func main() {
	fmt.Println("[Банк] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]\n", user1.Name)
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
