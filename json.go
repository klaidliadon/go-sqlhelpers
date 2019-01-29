package sqlhelpers

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strings"
)

// JSONValue encodes a SQL value as JSON
func JSONValue(i interface{}) (driver.Value, error) {
	b := strings.Builder{}
	if err := json.NewEncoder(&b).Encode(i); err != nil {
		return nil, err
	}
	return b.String(), nil
}

// JSONScan decodes a JSON as an SQL value
func JSONScan(i, value interface{}) error {
	return json.NewDecoder(bytes.NewReader(value.([]byte))).Decode(i)
}
