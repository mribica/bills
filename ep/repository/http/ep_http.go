package http

import (
	"net/url"

	"github.com/mribica/bills/domain"
	"github.com/mribica/bills/internal/encoding/html/table"
	"github.com/mribica/bills/internal/http/client"
)

type httpEpRepository struct {
	c *client.Client
}

func NewHttpEpRepository() domain.EpRepository {
	return &httpEpRepository{
		c: client.NewClient("https://www.epbih.ba"),
	}
}

func (r *httpEpRepository) GetBalance(username string, password string) (domain.Ep, error) {
	data := url.Values{}
	data.Add("username", username)
	data.Add("password", password)

	res, err := r.c.FormPost("profil/login", data)
	if err != nil {
		return domain.Ep{}, err
	}

	defer res.Close()

	var items []string
	table.NewDecoder(res).Decode(&items)
	return domain.Ep{Balance: items[2], UpdatedAt: items[1]}, nil
}
