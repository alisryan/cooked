package types_test

import (
	"net/url"
	"testing"

	"github.com/GenerateNU/cooked/backend/internal/types"
)

func TestMarshalURL(t *testing.T) {
	t.Parallel()

	url, err := url.ParseRequestURI("https://example.com/image.jpg")
	if err != nil {
		t.Fatalf("failed to parse url: %v", err)
	}

	u := types.URL(*url)

	marshalledURL, err := u.MarshalJSON()
	if err != nil {
		t.Fatalf("failed to marshal url: %v", err)
	}

	if string(marshalledURL) != `"https://example.com/image.jpg"` {
		t.Fatalf("unexpected marshalled url: %s", string(marshalledURL))
	}
}

func TestUnmarshalURL(t *testing.T) {
	t.Parallel()

	var u types.URL
	if err := u.UnmarshalJSON([]byte(`"https://example.com/image.jpg"`)); err != nil {
		t.Fatalf("failed to unmarshal url: %v", err)
	}

	url := u.Into()
	if url.String() != "https://example.com/image.jpg" {
		t.Fatalf("unexpected url: %s", url.String())
	}
}
