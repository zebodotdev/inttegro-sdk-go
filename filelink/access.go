package filelink

type Actor struct {
	Type  string `json:"type,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Delivery struct {
	Mode        DeliveryMode `json:"mode,omitempty"`
	Filename    string       `json:"filename,omitempty"`
	ContentType string       `json:"content_type,omitempty"`
	Disposition string       `json:"disposition,omitempty"`
}

type Access struct {
	MaxAccesses    int64    `json:"max_accesses,omitempty"`
	AllowDownload  bool     `json:"allow_download,omitempty"`
	AllowedOrigins []string `json:"allowed_origins,omitempty"`
}
