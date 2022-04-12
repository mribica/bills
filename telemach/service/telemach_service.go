package service

import (
	"fmt"

	"github.com/mribica/bills/domain"
	"github.com/mribica/bills/internal/encryption"
)

type telemachService struct {
	tmConfig domain.TelemachConfig
	tmRepo   domain.TelemachRepository
}

func NewEpService(tr domain.TelemachRepository, tc domain.TelemachConfig) domain.TelemachService {
	return telemachService{tmConfig: tc, tmRepo: tr}
}

func (ts telemachService) GetBalance() (domain.Telemach, error) {
	encPassword, _ := encryption.Encrypt(ts.tmConfig.Password, ts.tmConfig.AESPassphrase)
	token, _ := ts.tmRepo.Authenticate(ts.tmConfig.Username, string(encPassword), ts.tmConfig.ApplicationToken)
	return ts.tmRepo.GetBalance(token, ts.tmConfig.ApplicationToken)
}

func (ts telemachService) Render() string {
	tm, err := ts.GetBalance()
	if err != nil {
		return fmt.Sprintf(domain.TelemachRenderTemplate, err, "")
	}

	return fmt.Sprintf(domain.TelemachRenderTemplate, tm.Balance, tm.UpdatedAt)
}

func (ts telemachService) RenderChannel(ch chan<- string) {
	ch <- ts.Render()
}
