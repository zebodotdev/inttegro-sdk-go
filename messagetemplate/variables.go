package messagetemplate

type Variable struct {
	About    string         `json:"about,omitempty"`
	Default  any            `json:"default,omitempty"`
	Items    []VariableItem `json:"items,omitempty"`
	Name     string         `json:"name"`
	Required bool           `json:"required,omitempty"`
	Type     VariableType   `json:"type"`
}

type VariableItem struct {
	About    string           `json:"about,omitempty"`
	Default  any              `json:"default,omitempty"`
	Name     string           `json:"name"`
	Required bool             `json:"required,omitempty"`
	Type     VariableItemType `json:"type"`
}
