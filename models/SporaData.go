package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SporaData struct {
	Scopes []string `json:"scope,omitempty"`
}

// Make the SporaData struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (s SporaData) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Make the SporaData struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (s *SporaData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &s)
}
