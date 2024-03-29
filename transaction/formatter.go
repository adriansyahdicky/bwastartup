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

type UserTransactionFormatter struct {
	ID        int               `json:"id"`
	Amount    int               `json:"amount"`
	Status    string            `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transactions) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transactions) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}

type TransactionFormatter struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transactions Transactions) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transactions.ID
	formatter.CampaignID = transactions.CampaignID
	formatter.UserID = transactions.UserID
	formatter.Amount = transactions.Amount
	formatter.Status = transactions.Status
	formatter.Code = transactions.Code
	formatter.PaymentURL = transactions.PaymentURL
	return formatter
}
