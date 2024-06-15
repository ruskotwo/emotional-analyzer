package config

type OAuth2 struct {
	Client OAuth2Client
}

type OAuth2Client struct {
	ID     string
	Secret string
	Domain string
}
