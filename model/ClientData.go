package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ClientData struct {
	Scopes []string `json:"scope,omitempty"`
}

// Make the ClientData struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (c ClientData) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Make the ClientData struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (c *ClientData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &c)
}
