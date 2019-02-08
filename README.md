# SQL Helpers

This package provides some utilities for **database/sql**.

## JSON

The functions `JSONValue` and `JSONScan` provide helpers that store a type as JSON in a string column.

```go
type MyType map[string]string

// Value encodes a sql value
func (m MyType) Value() (driver.Value, error) {
	return sqlhelpers.JSONValue(m)
}

// Scan decodes a sql value
func (m *MyType) Scan(value interface{}) error {
	return sqlhelpers.JSONScan(m, value)
}
```


