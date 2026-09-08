// Package filereference provides filereference resources and operations.
package filereference

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
)

// FileReferencesService manages Inttegro resource file references.
type Service struct {
	client transport.Client
}

type Input struct {
	FileID        string `json:"file_id"`
	Field         string `json:"field"`
	Reference     string `json:"reference,omitempty"`
	ReferenceKind string `json:"reference_kind,omitempty"`
	Purpose       string `json:"purpose,omitempty"`
}

type ReconcileParams struct {
	ResourceType string  `json:"resource_type"`
	ResourceID   string  `json:"resource_id"`
	References   []Input `json:"references,omitempty"`
}

// FileReference reports whether the reference set was reconciled.
type FileReference struct {
	Reconciled bool `json:"reconciled"`
}

// Reconcile replaces the live file references for a Inttegro resource.
func (s *Service) Reconcile(ctx context.Context, params ReconcileParams) (*FileReference, error) {
	var resp FileReference
	if err := s.client.Do(ctx, "POST", "/file_references/reconcile", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
