package auth

const (
	USER_ID_KEY = "user-id"
)

type Middleware struct {
	config Config
}

type Config struct {
	accessTokenSecret string
}

func New(accessSecret string) Middleware {
	return Middleware{
		config: Config{
			accessTokenSecret: accessSecret,
		},
	}
}
