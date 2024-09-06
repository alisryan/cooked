package types

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"net/url"

	go_json "github.com/goccy/go-json"
)

// A URL is a wrapper around url.URL that provides JSON marshalling and unmarshalling.
type URL url.URL

func (u URL) MarshalJSON() ([]byte, error) {
	url := u.Into()
	return go_json.Marshal(url.String())
}

var ErrInvalidURL = errors.New("invalid url")

func (u *URL) UnmarshalJSON(data []byte) error {
	var s string
	if err := go_json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsedURL, err := url.ParseRequestURI(s)
	if err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	*u = URL(*parsedURL)
	return nil
}

func (u URL) Value() (driver.Value, error) {
	url := u.Into()
	return driver.Value(url.String()), nil
}

func (u *URL) Scan(raw interface{}) error {
	switch v := raw.(type) {
	case []byte:
		return u.UnmarshalJSON(v)
	case string:
		parsedURL, err := url.Parse(v)
		if err != nil {
			return fmt.Errorf("invalid URL: %v", err)
		}
		*u = URL(*parsedURL)
		return nil
	case nil:
		*u = URL{}
		return nil
	default:
		return fmt.Errorf("cannot sql.Scan() URL, unexpected type: %T", v)
	}
}

func (u URL) Into() url.URL {
	return url.URL(u)
}
