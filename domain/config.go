package domain

import "encoding/json"

type Config struct {
	Type string
	Data json.RawMessage
}
