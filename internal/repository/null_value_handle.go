package repository

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type NullString string

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		xs := fmt.Sprintf("%s", value)
		if xs == "" {
			return errors.New("column is not a string")
		}
		*s = NullString(xs)
		return nil
	}
	*s = NullString(strVal)
	return nil
}
func (s NullString) Value() (driver.Value, error) {
	if len(s) == 0 { // if NULL or EMPTY String
		return nil, nil
	}
	return string(s), nil
}
