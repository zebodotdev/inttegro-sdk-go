package filelink

type Resource struct {
	ID         string            `json:"id"`
	FileID     string            `json:"file_id"`
	Status     Status            `json:"status"`
	CustomData map[string]string `json:"custom_data,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

type Page struct {
	Number    int        `json:"number"`
	Size      int        `json:"size"`
	FileLinks []Resource `json:"file_links"`
}
