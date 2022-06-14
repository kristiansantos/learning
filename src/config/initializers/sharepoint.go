package initializers

type Sharepoint struct {
	URL          string `env:"sp_url"`
	ClientID     string `env:"sp_client_id"`
	ClientSecret string `env:"sp_client_secret"`
}
