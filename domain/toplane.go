package domain

type Toplane struct {
	Balance   string
	UpdatedAt string
}

type ToplaneConfig struct {
	Provider   string
	CustomerId string
}

type ToplaneRepository interface {
	GetBalance(customerId string) (Toplane, error)
}

type ToplaneService interface {
	GetBalance() (Toplane, error)
	Render() string
	RenderChannel(ch chan<- string)
}
