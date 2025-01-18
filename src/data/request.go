package data

type RegisterReq struct {
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required,e164" `
	Address     string `json:"address" binding:"required"`
	Pin         string `json:"pin" binding:"required,max=6,min=6"`
}

type LoginReq struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Pin         string `json:"pin" binding:"required,max=6,min=6"`
}

type TopUpReq struct {
	Amount int64 `json:"amount"`
}

type PaymentReq struct {
	Amount  int64  `json:"amount"`
	Remarks string `json:"remarks"`
}

type TransferReq struct {
	Amount  int64  `json:"amount"`
	Remarks string `json:"remarks"`
}
