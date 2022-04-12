package service

import (
	"fmt"

	"github.com/mribica/bills/domain"
)

type toService struct {
	toConfig domain.ToplaneConfig
	toRepo   domain.ToplaneRepository
}

func NewToplaneService(tr domain.ToplaneRepository, tc domain.ToplaneConfig) domain.ToplaneService {
	return toService{toConfig: tc, toRepo: tr}
}

func (ts toService) GetBalance() (domain.Toplane, error) {
	return ts.toRepo.GetBalance(ts.toConfig.CustomerId)
}

func (ts toService) Render() string {
	to, err := ts.GetBalance()
	if err != nil {
		return fmt.Sprintf(domain.ToplaneRenderTemplate, err, "")
	}

	return fmt.Sprintf(domain.ToplaneRenderTemplate, to.Balance, to.UpdatedAt)
}

func (ts toService) RenderChannel(ch chan<- string) {
	ch <- ts.Render()
}
