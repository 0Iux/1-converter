package main

import "fmt"

// У меня осталась привычка с питона, что многие типы автоматически являются глобальными.
// основываясь на этом я сразу и создал мапу как глобальную переменную
// есть ли какие-то проблемы связанные с этим? или это просто не принято?
// для дз с указателями все переделал под них

// var currency_map = map[string]float64{
// 	"usd_eur": 0.96,
// 	"usd_rub": 99.96,
// }

func main() {
	var currency_map = map[string]float64{
		"usd_eur": 0.96,
		"usd_rub": 99.96,
	}
	make_map(&currency_map)
	fmt.Println("\n---------------Калькулятор обменника---------------")
	for {
		start_currency, ammount, need_currency := get_user_input()
		result := calculate_exchange(start_currency, need_currency, ammount, &currency_map)
		output_result(result, need_currency)
		is_repeat := check_repeat()
		if !is_repeat {
			break
		}
	}
}

func make_map(currency_map *map[string]float64) {
	// Мапа вроде и так является ссылочным типом, так что не было смысла, но для упражнения вот
	(*currency_map)["eur_usd"] = 1 / (*currency_map)["usd_eur"]
	(*currency_map)["rub_usd"] = 1 / (*currency_map)["usd_rub"]
	(*currency_map)["eur_rub"] = (*currency_map)["usd_rub"] / (*currency_map)["usd_eur"]
	(*currency_map)["rub_eur"] = 1 / (*currency_map)["eur_rub"]
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
		fmt.Scan(&start_currency)
		if start_currency != "usd" && start_currency != "rub" && start_currency != "eur" {
			fmt.Println("Нет такого варианта!")
			continue
		}
		break
	}

	for {
		fmt.Println("Сколько у вас этой валюты?: ")
		fmt.Scan(&ammount)
		if ammount <= 0 {
			fmt.Println("Введенно некоректное число")
			continue
		}
		break
	}

	for {
		fmt.Println("Какую валюту вы хотите получить? (usd, rub, eur):")
		fmt.Scan(&need_currency)
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

func calculate_exchange(start_currency, need_currency string, ammount float64, currency_map *map[string]float64) float64 {
	need_value := start_currency + "_" + need_currency
	return ammount * (*currency_map)[need_value]
}
