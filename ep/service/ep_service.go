package service

import (
	"fmt"

	"github.com/mribica/bills/domain"
)

type epService struct {
	epConfig domain.EpConfig
	epRepo   domain.EpRepository
}

func NewEpService(er domain.EpRepository, ec domain.EpConfig) domain.EpService {
	return epService{epConfig: ec, epRepo: er}
}

func (es epService) GetBalance() (domain.Ep, error) {
	return es.epRepo.GetBalance(es.epConfig.Username, es.epConfig.Password)
}

func (es epService) Render() string {
	ep, err := es.GetBalance()
	if err != nil {
		return fmt.Sprintf(domain.EpRenderTemplate, err, "")
	}

	return fmt.Sprintf(domain.EpRenderTemplate, ep.Balance, ep.UpdatedAt)
}

func (es epService) RenderChannel(ch chan<- string) {
	ch <- es.Render()
}
