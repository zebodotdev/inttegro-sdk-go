package product

import (
	"bytes"
	"encoding/json"
)

// ProductCategory describes a product category.
type Category struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// UnmarshalJSON accepts both the canonical category string and the legacy
// expanded category object so existing callers keep their field accessors.
func (c *Category) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if bytes.Equal(trimmed, []byte("null")) {
		*c = Category{}
		return nil
	}
	if len(trimmed) > 0 && trimmed[0] == '"' {
		return json.Unmarshal(trimmed, &c.Name)
	}
	type categoryAlias Category
	return json.Unmarshal(trimmed, (*categoryAlias)(c))
}
