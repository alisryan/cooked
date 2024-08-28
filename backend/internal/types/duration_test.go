package types_test

import (
	"testing"
	"time"

	"github.com/GenerateNU/cooked/backend/internal/types"
)

func TestMarshalDuration(t *testing.T) {
	t.Parallel()

	d := types.Duration(time.Minute * 30)

	b, err := d.MarshalJSON()
	if err != nil {
		t.Fatalf("failed to marshal duration: %v", err)
	}

	if string(b) != `"30m0s"` {
		t.Errorf("unexpected duration: %s", string(b))
	}
}

func TestUnmarshalDuration(t *testing.T) {
	t.Parallel()

	var d types.Duration
	if err := d.UnmarshalJSON([]byte(`"30m0s"`)); err != nil {
		t.Fatalf("failed to unmarshal duration: %v", err)
	}

	if d.Into() != time.Minute*30 {
		t.Errorf("unexpected duration: %v", d.Into())
	}
}
