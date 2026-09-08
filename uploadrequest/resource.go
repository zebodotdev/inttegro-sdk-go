package uploadrequest

type Constraints struct {
	ContentTypes []string `json:"content_types,omitempty"`
	ExactSize    int64    `json:"exact_size,omitempty"`
	Extensions   []string `json:"extensions,omitempty"`
	Filename     string   `json:"filename,omitempty"`
	MaxSize      int64    `json:"max_size,omitempty"`
	MinSize      int64    `json:"min_size,omitempty"`
}

type Display struct {
	Description string `json:"description,omitempty"`
	HelpText    string `json:"help_text,omitempty"`
	Title       string `json:"title,omitempty"`
}

type Party struct {
	Type  string `json:"type,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type FileResource struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ReviewReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
}

// UploadRequest represents a request for a user to upload a file.
type UploadRequest struct {
	ID         string            `json:"id"`
	Purpose    string            `json:"purpose"`
	Status     Status            `json:"status"`
	UploadURL  string            `json:"upload_url,omitempty"`
	CustomData map[string]string `json:"custom_data,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

type Page struct {
	Number         int             `json:"number"`
	Size           int             `json:"size"`
	UploadRequests []UploadRequest `json:"upload_requests"`
}
