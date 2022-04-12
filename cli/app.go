package cli

import (
	"encoding/json"
	"fmt"

	"github.com/mribica/bills/config"
	"github.com/mribica/bills/domain"
	epRepo "github.com/mribica/bills/ep/repository/http"
	epService "github.com/mribica/bills/ep/service"
	tmRepo "github.com/mribica/bills/telemach/repository/http"
	tmService "github.com/mribica/bills/telemach/service"
	toRepo "github.com/mribica/bills/toplane/repository/http"
	toService "github.com/mribica/bills/toplane/service"
)

type App struct {
	Config    domain.Config
	Providers []domain.Presenter
}

func (app *App) LoadProviders() error {
	configData, err := config.Load()
	if err != nil {
		return err
	}

	for _, config := range configData {
		switch config.Type {
		case "Ep":
			var ec domain.EpConfig
			err := json.Unmarshal(config.Data, &ec)
			if err != nil {
				return fmt.Errorf("can't parse Ep provider confif.\n %s", err.Error())
			}

			er := epRepo.NewHttpEpRepository()
			es := epService.NewEpService(er, ec)
			app.Providers = append(app.Providers, es)
		case "Telemach":
			var tc domain.TelemachConfig
			err := json.Unmarshal(config.Data, &tc)
			if err != nil {
				return fmt.Errorf("can't parse Telemach provider confif.\n %s", err.Error())
			}

			tr := tmRepo.NewHttpTelemachRepository()
			ts := tmService.NewEpService(tr, tc)
			app.Providers = append(app.Providers, ts)
		case "Toplane":
			var tc domain.ToplaneConfig
			err := json.Unmarshal(config.Data, &tc)
			if err != nil {
				return fmt.Errorf("can't parse Toplane provider confif.\n %s", err.Error())
			}

			tr := toRepo.NewHttpToplaneRepository()
			ts := toService.NewToplaneService(tr, tc)
			app.Providers = append(app.Providers, ts)
		}
	}
	return nil
}

func (app *App) Run(ch chan<- string) {
	for _, service := range app.Providers {
		go service.RenderChannel(ch)
	}
}
