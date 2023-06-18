package transaction

import (
	"time"
)

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transactions Transactions) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{}
	formatter.ID = transactions.ID
	formatter.Name = transactions.User.Name
	formatter.Amount = transactions.Amount
	formatter.CreatedAt = transactions.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []Transactions) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}
