package inttegro

import (
	"encoding/json"
	"errors"
	"fmt"
)

// JSONValue is one validated arbitrary JSON value. It is reserved for API
// extension points whose value type is intentionally chosen by the caller.
type JSONValue struct {
	raw json.RawMessage
}

// NewJSONValue validates and stores value as JSON.
func NewJSONValue(value any) (JSONValue, error) {
	encoded, err := json.Marshal(value)
	if err != nil {
		return JSONValue{}, fmt.Errorf("inttegro: encode JSON value: %w", err)
	}
	return JSONValue{raw: append(json.RawMessage(nil), encoded...)}, nil
}

// Decode decodes the value into out.
func (v JSONValue) Decode(out any) error {
	if out == nil {
		return errors.New("inttegro: JSON value decode target is nil")
	}
	encoded := v.raw
	if len(encoded) == 0 {
		encoded = []byte("null")
	}
	return json.Unmarshal(encoded, out)
}

func (v JSONValue) MarshalJSON() ([]byte, error) {
	if len(v.raw) == 0 {
		return []byte("null"), nil
	}
	return append([]byte(nil), v.raw...), nil
}

func (v *JSONValue) UnmarshalJSON(encoded []byte) error {
	if v == nil {
		return errors.New("inttegro: cannot decode JSON value into nil receiver")
	}
	if !json.Valid(encoded) {
		return errors.New("inttegro: invalid JSON value")
	}
	v.raw = append(json.RawMessage(nil), encoded...)
	return nil
}

// JSONData is a deliberately extensible JSON object.
//
// It is used only where the API contract permits integration-defined keys.
// Unlike map[string]any, it rejects values that cannot be represented in JSON
// and does not expose mutable aliases to its backing data.
type JSONData struct {
	values map[string]json.RawMessage
}

// NewJSONData returns an empty extensible JSON object.
func NewJSONData() *JSONData {
	return &JSONData{values: make(map[string]json.RawMessage)}
}

// Set validates and stores a JSON value.
func (d *JSONData) Set(key string, value any) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("inttegro: encode JSON data %q: %w", key, err)
	}
	if d.values == nil {
		d.values = make(map[string]json.RawMessage)
	}
	d.values[key] = append(json.RawMessage(nil), encoded...)
	return nil
}

// Remove deletes a value. Removing a missing key is a no-op.
func (d *JSONData) Remove(key string) {
	if d != nil {
		delete(d.values, key)
	}
}

// Decode decodes one value into out and reports whether the key exists.
func (d *JSONData) Decode(key string, out any) (bool, error) {
	if d == nil {
		return false, nil
	}
	value, ok := d.values[key]
	if !ok {
		return false, nil
	}
	if out == nil {
		return true, errors.New("inttegro: JSON data decode target is nil")
	}
	if err := json.Unmarshal(value, out); err != nil {
		return true, fmt.Errorf("inttegro: decode JSON data %q: %w", key, err)
	}
	return true, nil
}

// Raw returns a defensive copy of a value's encoded JSON.
func (d *JSONData) Raw(key string) (json.RawMessage, bool) {
	if d == nil {
		return nil, false
	}
	value, ok := d.values[key]
	return append(json.RawMessage(nil), value...), ok
}

// Len returns the number of entries.
func (d *JSONData) Len() int {
	if d == nil {
		return 0
	}
	return len(d.values)
}

func (d JSONData) MarshalJSON() ([]byte, error) {
	if d.values == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(d.values)
}

func (d *JSONData) UnmarshalJSON(encoded []byte) error {
	if d == nil {
		return errors.New("inttegro: cannot decode JSON data into nil receiver")
	}
	if string(encoded) == "null" {
		d.values = nil
		return nil
	}
	var values map[string]json.RawMessage
	if err := json.Unmarshal(encoded, &values); err != nil {
		return fmt.Errorf("inttegro: decode JSON data: %w", err)
	}
	d.values = make(map[string]json.RawMessage, len(values))
	for key, value := range values {
		d.values[key] = append(json.RawMessage(nil), value...)
	}
	return nil
}
