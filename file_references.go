package inttegro

import "context"

// FileReferencesService manages Inttegro resource file references.
type FileReferencesService struct {
	client *Client
}

type FileReferenceInput struct {
	FileID        string `json:"file_id"`
	Field         string `json:"field"`
	Reference     string `json:"reference,omitempty"`
	ReferenceKind string `json:"reference_kind,omitempty"`
	Purpose       string `json:"purpose,omitempty"`
}

type FileReferenceReconcileParams struct {
	ResourceType string               `json:"resource_type"`
	ResourceID   string               `json:"resource_id"`
	References   []FileReferenceInput `json:"references,omitempty"`
}

type FileReferenceReconcileResponse struct {
	Reconciled bool `json:"reconciled"`
}

// Reconcile replaces the live file references for a Inttegro resource.
func (s *FileReferencesService) Reconcile(ctx context.Context, params FileReferenceReconcileParams) (*FileReferenceReconcileResponse, error) {
	var resp FileReferenceReconcileResponse
	if err := s.client.do(ctx, "POST", "/file_references/reconcile", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
