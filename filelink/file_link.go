// Package filelink provides filelink resources and operations.
package filelink

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/file"
	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

type Status string

const (
	StatusActive Status = "active"
)

const (
	StatusRevoked Status = "revoked"
)

const (
	StatusExpired Status = "expired"
)

const (
	StatusDisabled Status = "disabled"
)

type Kind string

const KindPublic Kind = "public"

type DeliveryMode string

const (
	DeliveryModeRedirect DeliveryMode = "redirect"
)

const (
	DeliveryModeDownload DeliveryMode = "download"
)

const (
	DeliveryModeInline DeliveryMode = "inline"
)

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

type CreateParams struct {
	Access     Access            `json:"access"`
	CreatedBy  Actor             `json:"created_by"`
	Delivery   Delivery          `json:"delivery"`
	ExpiresAt  string            `json:"expires_at,omitempty"`
	FileID     string            `json:"file_id"`
	CustomData map[string]string `json:"custom_data,omitempty"`
}

type PageParams struct {
	FileID     string `json:"file_id,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	Status     Status `json:"status,omitempty"`
}

type RevokeParams struct {
	ID        string `json:"id"`
	RevokedBy Actor  `json:"revoked_by"`
}

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

type Service struct {
	client transport.Client
}

func (s *Service) Create(ctx context.Context, params CreateParams, opts ...request.Option) (*Resource, string, error) {
	var resp struct {
		FileLink Resource `json:"file_link"`
		URL      string   `json:"url"`
	}
	if err := s.client.DoJSON(ctx, "/file_links/create", params, request.Apply(opts), &resp); err != nil {
		return nil, "", err
	}
	return &resp.FileLink, resp.URL, nil
}

func (s *Service) Lookup(ctx context.Context, id string) (*Resource, error) {
	var resp struct {
		FileLink Resource `json:"file_link"`
	}
	if err := s.client.Do(ctx, "POST", "/file_links/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.FileLink, nil
}

func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/file_links/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

func (s *Service) Revoke(ctx context.Context, params RevokeParams, opts ...request.Option) (*Resource, error) {
	var resp struct {
		FileLink Resource `json:"file_link"`
	}
	if err := s.client.DoJSON(ctx, "/file_links/revoke", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.FileLink, nil
}

func (s *Service) Open(ctx context.Context, url string) (*file.Download, error) {
	resp, err := s.client.RawResponse(ctx, "GET", url, nil, "", "", false, "file_links.download")
	if err != nil {
		return nil, err
	}
	return &file.Download{ReadCloser: resp.Body}, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
