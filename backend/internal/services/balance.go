// internal/services/balance.go
package services

// var (
// 	ErrInvalidTransactionType = errors.New("invalid transaction type")
// 	ErrInsufficientBalance    = errors.New("insufficient balance")
// )

func ApplyTransactionToBalance(
	currentBalance float64,
	transactionType string,
	amount float64,
) (float64, error) {
	switch transactionType {
	case "income":
		return currentBalance + amount, nil

	case "expense":
		if currentBalance < amount {
			return 0, ErrInsufficientBalance
		}

		return currentBalance - amount, nil

	default:
		return 0, ErrInvalidTransactionType
	}
}

func ReverseTransactionFromBalance(
	currentBalance float64,
	transactionType string,
	amount float64,
) (float64, error) {
	switch transactionType {
	case "income":
		if currentBalance < amount {
			return 0, ErrInsufficientBalance
		}

		return currentBalance - amount, nil

	case "expense":
		return currentBalance + amount, nil

	default:
		return 0, ErrInvalidTransactionType
	}
}