package http

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mribica/bills/domain"
	"github.com/mribica/bills/internal/http/client"
)

type httpTelemachRepository struct {
	c *client.Client
}

type authenticationResponse struct {
	Status                string
	StatusMessage         string
	UserTokenAuthenticate domain.TelemachToken
}

type balanceResponse struct {
	Status        string
	StatusMessage string
	Saldo         struct {
		CustomerName string
		Message      string
		SaldoNumber  string
	}
}

func NewHttpTelemachRepository() domain.TelemachRepository {
	return &httpTelemachRepository{
		c: client.NewClient("https://mojtelemach.ba/gateway"),
	}
}

func (r *httpTelemachRepository) Authenticate(username string, encryptedPassword string, applicationToken string) (domain.TelemachToken, error) {
	data := map[string]string{
		"username":      username,
		"password":      encryptedPassword,
		"grantType":     "password",
		"domainId":      "TBA",
		"applicationId": "69cj",
	}

	headers := map[string]string{"Authorization": applicationToken}
	res, err := r.c.Post("SCAuthAPI/1.0/scauth/auth/authentication", data, headers)
	if err != nil {
		return domain.TelemachToken{}, err
	}

	defer res.Close()

	var authResponse authenticationResponse
	if err = json.NewDecoder(res).Decode(&authResponse); err != nil {
		return domain.TelemachToken{}, err
	}

	return authResponse.UserTokenAuthenticate, nil
}

func (r *httpTelemachRepository) GetBalance(token domain.TelemachToken, applicationToken string) (domain.Telemach, error) {
	path := fmt.Sprintf("SelfCareAPI/1.0/selfcareapi/%s/subscriber/%s/saldo", token.SiteId, token.Identity)

	headers := map[string]string{
		"accessToken":   token.AccessToken,
		"Authorization": applicationToken,
	}
	res, err := r.c.Get(path, headers)
	if err != nil {
		return domain.Telemach{}, err
	}

	defer res.Close()

	var balanceResponse balanceResponse
	if err = json.NewDecoder(res).Decode(&balanceResponse); err != nil {
		return domain.Telemach{}, err
	}

	now := time.Now().Local()
	endOfPreviousMonth := now.AddDate(0, 0, -now.Day())

	return domain.Telemach{
		Balance:   balanceResponse.Saldo.SaldoNumber,
		UpdatedAt: endOfPreviousMonth.Format("02.01.2006."),
	}, nil
}
