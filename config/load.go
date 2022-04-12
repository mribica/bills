package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mribica/bills/domain"
)

func Load() ([]domain.Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return []domain.Config{}, err
	}

	fileBytes, err := ioutil.ReadFile(filepath.Join(homeDir, ".bills", "providers.json"))
	if err != nil {
		return []domain.Config{}, errors.New("can't find .bills/providers.json in your home directory")
	}

	var config []domain.Config
	err = json.Unmarshal(fileBytes, &config)
	if err != nil {
		return []domain.Config{}, fmt.Errorf("can't parse providers.json\n %s", err.Error())
	}

	return config, nil
}
