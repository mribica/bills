package http

import (
	"net/url"

	"github.com/mribica/bills/domain"
	"github.com/mribica/bills/internal/encoding/html/table"
	"github.com/mribica/bills/internal/http/client"
)

type httpToplaneRepository struct {
	c *client.Client
}

func NewHttpToplaneRepository() domain.ToplaneRepository {
	return &httpToplaneRepository{
		c: client.NewClient("https://www.toplanesarajevo.ba"),
	}
}

func (r *httpToplaneRepository) GetBalance(customerId string) (domain.Toplane, error) {
	data := url.Values{"sifra_korisnika": {customerId}}
	res, err := r.c.FormPost("korisnici/stanje-racuna", data)
	if err != nil {
		return domain.Toplane{}, err
	}

	defer res.Close()

	var items []string
	table.NewDecoder(res).Decode(&items)
	return domain.Toplane{Balance: items[7], UpdatedAt: items[9]}, nil
}
