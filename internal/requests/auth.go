package requests

type RegisterRequest struct {
	Name  string `json:"name" validate:"required"`
	NIK   string `json:"nik" validate:"required"`
	Phone string `json:"phone" validate:"required"`
}

type LoginRequest struct {
	AccountNumber string `json:"no_rekening" validate:"required"`
	NIK           string `json:"nik" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
}
