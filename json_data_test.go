package inttegro

import (
	"encoding/json"
	"testing"
)

func TestJSONDataValidatesAndDecodesValues(t *testing.T) {
	data := NewJSONData()
	if err := data.Set("attempts", 3); err != nil {
		t.Fatal(err)
	}
	if err := data.Set("unsafe", make(chan int)); err == nil {
		t.Fatal("expected non-JSON value to be rejected")
	}

	var attempts int
	ok, err := data.Decode("attempts", &attempts)
	if err != nil || !ok || attempts != 3 {
		t.Fatalf("decoded attempts=%d, present=%v, error=%v", attempts, ok, err)
	}
	encoded, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != `{"attempts":3}` {
		t.Fatalf("encoded = %s", encoded)
	}
}
