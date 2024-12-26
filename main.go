package main

import "fmt"

const usd_eur = 0.96
const eur_usd = 1 / usd_eur
const usd_rub = 99.96
const rub_usd = 1 / usd_rub
const eur_rub = usd_rub / usd_eur
const rub_eur = 1 / eur_rub

func main() {
	fmt.Println("\n---------------Калькулятор обменника---------------")
	for {
		start_currency, ammount, need_currency := get_user_input()
		result := calculate_exchange(start_currency, need_currency, ammount)
		output_result(result, need_currency)
		is_repeat := check_repeat()
		if !is_repeat {
			break
		}
	}
}

func check_repeat() bool {
	fmt.Println("\n------------------------Меню------------------------")
	var choise string
	fmt.Println("Хотите ли вы еще раз рассчитать? (y/n): ")
	fmt.Scan(&choise)
	if choise == "n" || choise == "N" {
		return false
	}
	return true
}

func output_result(res float64, cur string) {
	fmt.Printf("Вы получили %.2f %s", res, cur)

}

func get_user_input() (string, float64, string) {
	var start_currency, need_currency string
	var ammount float64

	for {
		fmt.Println("Какая валюта у вас сейчас? (usd, rub, eur): ")
		fmt.Scan((&start_currency))
		if start_currency != "usd" && start_currency != "rub" && start_currency != "eur" {
			fmt.Println("Нет такого варианта!")
			continue
		}
		break
	}

	for {
		fmt.Println("Сколько у вас этой валюты?: ")
		fmt.Scan((&ammount))
		if ammount <= 0 {
			fmt.Println("Введенно некоректное число")
			continue
		}
		break
	}

	for {
		fmt.Println("Какую валюту вы хотите получить? (usd, rub, eur):")
		fmt.Scan((&need_currency))
		if need_currency == start_currency {
			fmt.Println("Введена одинаковая валюта!")
			continue
		}
		if need_currency != "usd" && need_currency != "rub" && need_currency != "eur" {
			fmt.Println("Нет такого варианта!")
			continue
		}
		break
	}
	return start_currency, ammount, need_currency
}

func calculate_exchange(start_currency, need_currency string, ammount float64) float64 {
	switch {
	case start_currency == "rub":
		if need_currency == "eur" {
			return ammount * rub_eur
		}
		return ammount * rub_usd

	case start_currency == "eur":
		if need_currency == "rub" {
			return ammount * eur_rub
		}
		return ammount * eur_usd

	case start_currency == "usd":
		if need_currency == "eur" {
			return ammount * usd_eur
		}
	}
	return ammount * usd_rub
}
