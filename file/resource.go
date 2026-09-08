package file

type Resource struct {
	ID         string            `json:"id"`
	Purpose    string            `json:"purpose"`
	Status     Status            `json:"status"`
	CustomData map[string]string `json:"custom_data,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

type Page struct {
	Number int        `json:"number"`
	Size   int        `json:"size"`
	Files  []Resource `json:"files"`
}
