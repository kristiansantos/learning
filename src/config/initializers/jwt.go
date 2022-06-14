package initializers

type Jwt struct {
	Issuer          string `env:"jwt_issuer"`
	SecretKey       string `env:"jwt_secretKey"`
	ExpirationHours int64  `env:"jwt_expirationHours"`
}
