package domain

type Telemach struct {
	Balance   string
	UpdatedAt string
}

type TelemachToken struct {
	Identity     string
	AccessToken  string
	RefreshToken string
	TokenType    string
	SiteId       string
}

type TelemachConfig struct {
	Provider         string
	Username         string
	Password         string
	AESPassphrase    string
	ApplicationToken string
	ApplicationId    string
}

type TelemachRepository interface {
	Authenticate(username string, encryptedPassword string, applicationToken string) (TelemachToken, error)
	GetBalance(token TelemachToken, applicationToken string) (Telemach, error)
}

type TelemachService interface {
	GetBalance() (Telemach, error)
	Render() string
	RenderChannel(ch chan<- string)
}
