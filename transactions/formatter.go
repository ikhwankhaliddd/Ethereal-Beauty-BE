package transactions

import (
	"time"
)

type ProductTransactionsFormatJSON struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionsFormatJSON struct {
	ID        int              `json:"id"`
	Amount    int              `json:"amount"`
	Status    string           `json:"paid"`
	CreatedAt time.Time        `json:"created_at"`
	Product   ProductFormatter `json:"product"`
}

type ProductFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
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

func FormatUserTransaction(transactions Transactions) UserTransactionsFormatJSON {

	formatter := UserTransactionsFormatJSON{
		ID:        transactions.ID,
		Amount:    transactions.Amount,
		Status:    transactions.Status,
		CreatedAt: transactions.CreatedAt,
	}
	productFormatter := ProductFormatter{
		Name:     transactions.Product.Name,
		ImageUrl: "",
	}
	if len(transactions.Product.ProductImages) > 0 {
		productFormatter.ImageUrl = transactions.Product.ProductImages[0].FileName
	}
	formatter.Product = productFormatter
	return formatter

}

func FormatUserTransactions(transactions []Transactions) []UserTransactionsFormatJSON {
	if len(transactions) == 0 {
		return []UserTransactionsFormatJSON{}
	}

	var transactionsFormatter []UserTransactionsFormatJSON

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}
	return transactionsFormatter
}
