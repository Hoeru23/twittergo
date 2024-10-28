package models

type Secret struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	JWTSign  string `json:"jwtsign"`
	Password string `json:"password"`
	Database string `json:"database"`
}
