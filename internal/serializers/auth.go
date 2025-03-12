package serializers

type RegisterResponse struct {
	AccountNumber string `json:"no_rekening"`
}

type LoginResponse struct {
	AuthToken string `json:"auth_token"`
}
