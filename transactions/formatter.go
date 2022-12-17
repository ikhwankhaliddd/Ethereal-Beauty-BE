package transactions

import "time"

type ProductTransactionsFormatJSON struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatProductTransaction(transaction Transactions) ProductTransactionsFormatJSON {
	formatter := ProductTransactionsFormatJSON{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}

	return formatter
}

func FormatProductTransactions(transactions []Transactions) []ProductTransactionsFormatJSON {
	if len(transactions) == 0 {
		return []ProductTransactionsFormatJSON{}
	}

	var transactionsFormatter []ProductTransactionsFormatJSON

	for _, transaction := range transactions {
		formatter := FormatProductTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}