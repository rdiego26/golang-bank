package accounts

type Message string

const (
	OperationSuccessfully Message = "Operation made successfully, thanks for using our services!"
	InsufficientFunds             = "Insufficient funds, we couldn't achieve your request!"
)
