package dto

type HistoryTrx struct {
	TransactionDate 	string	`json:"transaction_date"`
	TotalTransaction	int64	`json:"total_transaction"`
	Status 				string	`json:"status"`
	MerchantName		string 	`json:"merchant_name"`
}
