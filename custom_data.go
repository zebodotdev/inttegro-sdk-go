package inttegro

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// MaxCustomDataKeyBytes is the maximum UTF-8 byte length accepted for a
	// custom-data key by the Commerce API.
	MaxCustomDataKeyBytes = 256
	// MaxCustomDataBytes is the maximum encoded JSON size accepted for a
	// custom-data object by the Commerce API.
	MaxCustomDataBytes = 25 * 1024
)

var (
	ErrCustomDataKeyTooLong = errors.New("inttegro: custom data key exceeds 256 bytes")
	ErrCustomDataTooLarge   = errors.New("inttegro: custom data exceeds 25 KiB")
)

// CustomData is merchant-defined string data attached to an API resource.
//
// Its backing map is private so callers cannot bypass validation or mutate a
// response through an alias. Use Set and Remove to change it, and Values when
// a detached standard map is needed.
type CustomData struct {
	values map[string]string
}

// NewCustomData validates and copies values into a new CustomData value.
func NewCustomData(values ...map[string]string) (*CustomData, error) {
	if len(values) > 1 {
		return nil, fmt.Errorf("inttegro: NewCustomData accepts at most one map")
	}
	data := &CustomData{values: make(map[string]string)}
	if len(values) == 0 || values[0] == nil {
		return data, nil
	}
	for key, value := range values[0] {
		if err := data.Set(key, value); err != nil {
			return nil, err
		}
	}
	return data, nil
}

// Set validates and stores a key/value pair.
func (d *CustomData) Set(key, value string) error {
	if len([]byte(key)) > MaxCustomDataKeyBytes {
		return fmt.Errorf("%w: %q", ErrCustomDataKeyTooLong, key)
	}
	if d.values == nil {
		d.values = make(map[string]string)
	}
	previous, existed := d.values[key]
	d.values[key] = value
	if err := validateCustomDataSize(d.values); err != nil {
		if existed {
			d.values[key] = previous
		} else {
			delete(d.values, key)
		}
		return err
	}
	return nil
}

// Remove deletes a key. Removing a missing key is a no-op.
func (d *CustomData) Remove(key string) {
	if d != nil {
		delete(d.values, key)
	}
}

// Get returns a value without exposing the backing map.
func (d *CustomData) Get(key string) (string, bool) {
	if d == nil {
		return "", false
	}
	value, ok := d.values[key]
	return value, ok
}

// Len returns the number of custom-data entries.
func (d *CustomData) Len() int {
	if d == nil {
		return 0
	}
	return len(d.values)
}

// Values returns a defensive copy of the stored values.
func (d *CustomData) Values() map[string]string {
	if d == nil {
		return nil
	}
	values := make(map[string]string, len(d.values))
	for key, value := range d.values {
		values[key] = value
	}
	return values
}

// Clone returns an independent copy.
func (d *CustomData) Clone() *CustomData {
	clone, _ := NewCustomData(d.Values())
	return clone
}

func (d CustomData) MarshalJSON() ([]byte, error) {
	if d.values == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(d.values)
}

func (d *CustomData) UnmarshalJSON(encoded []byte) error {
	if d == nil {
		return errors.New("inttegro: cannot decode custom data into nil receiver")
	}
	if string(encoded) == "null" {
		d.values = nil
		return nil
	}
	var values map[string]string
	if err := json.Unmarshal(encoded, &values); err != nil {
		return fmt.Errorf("inttegro: decode custom data: %w", err)
	}
	validated, err := NewCustomData(values)
	if err != nil {
		return err
	}
	d.values = validated.values
	return nil
}

func validateCustomDataSize(values map[string]string) error {
	encoded, err := json.Marshal(values)
	if err != nil {
		return fmt.Errorf("inttegro: encode custom data: %w", err)
	}
	if len(encoded) > MaxCustomDataBytes {
		return ErrCustomDataTooLarge
	}
	return nil
}

// CustomDataPatch records merge operations for endpoints that distinguish
// setting a value from explicitly removing it.
type CustomDataPatch struct {
	changes map[string]*string
}

// NewCustomDataPatch returns an empty custom-data patch.
func NewCustomDataPatch() *CustomDataPatch {
	return &CustomDataPatch{changes: make(map[string]*string)}
}

// Set records a value assignment.
func (p *CustomDataPatch) Set(key, value string) error {
	if len([]byte(key)) > MaxCustomDataKeyBytes {
		return fmt.Errorf("%w: %q", ErrCustomDataKeyTooLong, key)
	}
	if p.changes == nil {
		p.changes = make(map[string]*string)
	}
	copyOfValue := value
	previous, existed := p.changes[key]
	p.changes[key] = &copyOfValue
	if err := validateCustomDataPatchSize(p.changes); err != nil {
		if existed {
			p.changes[key] = previous
		} else {
			delete(p.changes, key)
		}
		return err
	}
	return nil
}

// Unset records an explicit JSON null, which removes the key server-side.
func (p *CustomDataPatch) Unset(key string) error {
	if len([]byte(key)) > MaxCustomDataKeyBytes {
		return fmt.Errorf("%w: %q", ErrCustomDataKeyTooLong, key)
	}
	if p.changes == nil {
		p.changes = make(map[string]*string)
	}
	previous, existed := p.changes[key]
	p.changes[key] = nil
	if err := validateCustomDataPatchSize(p.changes); err != nil {
		if existed {
			p.changes[key] = previous
		} else {
			delete(p.changes, key)
		}
		return err
	}
	return nil
}

// Remove discards a pending change without affecting server-side data.
func (p *CustomDataPatch) Remove(key string) {
	if p != nil {
		delete(p.changes, key)
	}
}

// Len returns the number of pending changes.
func (p *CustomDataPatch) Len() int {
	if p == nil {
		return 0
	}
	return len(p.changes)
}

func (p CustomDataPatch) MarshalJSON() ([]byte, error) {
	if p.changes == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(p.changes)
}

func validateCustomDataPatchSize(changes map[string]*string) error {
	encoded, err := json.Marshal(changes)
	if err != nil {
		return fmt.Errorf("inttegro: encode custom data patch: %w", err)
	}
	if len(encoded) > MaxCustomDataBytes {
		return ErrCustomDataTooLarge
	}
	return nil
}
