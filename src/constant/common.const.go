package constant

type CtxKey string

const (
	DB                        CtxKey = "db"
	TRANSACTION_TYPE_TOPUP           = "TOPUP"
	TRANSACTION_TYPE_PAYMENT         = "PAYMENT"
	TRANSACTION_TYPE_TRANSFER        = "TRANSFER"
)
