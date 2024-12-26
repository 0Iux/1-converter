package main

import "fmt"

const usd_eur = 0.96
const usd_rub = 99.96
const eur_rub = usd_rub / usd_eur

func main() {
}

func get_user_input() (string, float64, string) {
	var start_currency, need_currency string
	var ammount float64
	fmt.Println("Какая валюта у вас сейчас? ")
	fmt.Scan((&start_currency))
	fmt.Println("Сколько у вас этой валюты?: ")
	fmt.Scan((&ammount))
	fmt.Println("Какую валюту вы хотите получить? ")
	fmt.Scan((&need_currency))
	return start_currency, ammount, need_currency
}

func calculate_exchange(start_currency, need_currency string, ammount float64) float64 {

}
