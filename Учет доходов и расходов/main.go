package main

import "fmt"

func main() {
	// Slice для хранения всех транзакций (предаллоцируем под типичное кол-во)
	transactions := make([]float64, 0, 10)
	for {
		// Сканируем транзакцию от пользователя
		transaction := scanTransaction()
		if transaction == 0 {
			// Если введено 0, выходим из цикла
			break
		}
		// Добавляем транзакцию в список
		transactions = append(transactions, transaction)
	}
	// Рассчитываем общий баланс
	balance := calculateBalance(transactions)
	fmt.Printf("Ваш баланс: %.2f", balance)
}

// scanTransaction: Функция для ввода транзакции от пользователя
func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Введите транзакцию (0 для выхода): ")
	fmt.Scan(&transaction)
	return transaction
}

// calculateBalance: Функция для расчета общего баланса
func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	// Суммируем все транзакции
	for _, value := range transactions {
		balance += value
	}
	return balance
}
