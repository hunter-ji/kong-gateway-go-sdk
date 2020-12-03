package kong_gateway_go_sdk

type Config struct {
	Url string `json:"url"`
}

func Connect(kongUrl string) *Config {
	return &Config{
		Url: kongUrl,
	}
}
