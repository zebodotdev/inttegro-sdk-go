package payment

type Error struct {
	Message string `json:"message"`
	DocsURL string `json:"docs_url"`
	Source  string `json:"source"`
	Type    string `json:"type"`
	Code    string `json:"code"`
}
