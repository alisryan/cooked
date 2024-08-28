package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"

	go_json "github.com/goccy/go-json"
)

// A Duration is a wrapper around time.Duration that provides JSON marshalling and unmarshalling.
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return go_json.Marshal(d.Into().String())
}

var ErrInvalidDuration = errors.New("invalid duration")

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := go_json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return ErrInvalidDuration
	}
}

func (d Duration) Value() (driver.Value, error) {
	return driver.Value(int64(d)), nil
}

func (d *Duration) Scan(raw interface{}) error {
	switch v := raw.(type) {
	case int64:
		*d = Duration(v)
	case nil:
		*d = Duration(0)
	default:
		return fmt.Errorf("cannot sql.Scan() struct, unexpected type for Duration: %T", v)
	}
	return nil
}

func (d Duration) Into() time.Duration {
	return time.Duration(d)
}
