package domain

type Ep struct {
	Balance   string
	UpdatedAt string
}

type EpConfig struct {
	Provider string
	Username string
	Password string
}

type EpRepository interface {
	GetBalance(username string, password string) (Ep, error)
}

type EpService interface {
	GetBalance() (Ep, error)
	Render() string
	RenderChannel(ch chan<- string)
}
