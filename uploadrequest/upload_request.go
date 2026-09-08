// Package uploadrequest provides uploadrequest resources and operations.
package uploadrequest

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/zebodotdev/inttegro-sdk-go/v5/file"
	"github.com/zebodotdev/inttegro-sdk-go/v5/filelink"
	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

type Status string

const (
	StatusPending Status = "pending"
)

const (
	StatusUploading Status = "uploading"
)

const (
	StatusFulfilled Status = "fulfilled"
)

const (
	StatusExpired Status = "expired"
)

const (
	StatusCanceled Status = "canceled"
)

const (
	StatusFailed Status = "failed"
)

type ReviewDecision string

const (
	ReviewDecisionApproved ReviewDecision = "approved"
)

const (
	ReviewDecisionRejected ReviewDecision = "rejected"
)

type ReviewType string

const (
	ReviewTypeAutomatic ReviewType = "automatic"
)

const (
	ReviewTypeManual ReviewType = "manual"
)

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

type CreateParams struct {
	Constraints Constraints       `json:"constraints"`
	Display     Display           `json:"display"`
	ExpiresAt   string            `json:"expires_at,omitempty"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
	Purpose     string            `json:"purpose"`
	Recipient   Party             `json:"recipient"`
	Requester   filelink.Actor    `json:"requester"`
	Resource    FileResource      `json:"resource"`
	Subject     Party             `json:"subject"`
}

type PageParams struct {
	PageNumber int          `json:"page_number,omitempty"`
	PageSize   int          `json:"page_size,omitempty"`
	Purpose    string       `json:"purpose,omitempty"`
	Resource   FileResource `json:"resource"`
	Status     Status       `json:"status,omitempty"`
}

type CancelParams struct {
	CanceledBy filelink.Actor `json:"canceled_by"`
	ID         string         `json:"id"`
}

type ReviewReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
}

type ReviewParams struct {
	AttemptID      string         `json:"attempt_id,omitempty"`
	AttemptOrdinal int64          `json:"attempt_ordinal,omitempty"`
	Decision       string         `json:"decision"`
	ID             string         `json:"id"`
	PublicMessage  string         `json:"public_message,omitempty"`
	Reasons        []ReviewReason `json:"reasons,omitempty"`
}

type FulfillParams struct {
	File      string
	UploadURL string
}

type Resource struct {
	ID         string            `json:"id"`
	Purpose    string            `json:"purpose"`
	Status     Status            `json:"status"`
	UploadURL  string            `json:"upload_url,omitempty"`
	CustomData map[string]string `json:"custom_data,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

type Page struct {
	Number         int        `json:"number"`
	Size           int        `json:"size"`
	UploadRequests []Resource `json:"upload_requests"`
}

type Service struct {
	client transport.Client
}

func (s *Service) Create(ctx context.Context, params CreateParams, opts ...request.Option) (*Resource, error) {
	var resp struct {
		UploadRequest Resource `json:"upload_request"`
	}
	if err := s.client.DoJSON(ctx, "/upload_requests/create", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.UploadRequest, nil
}

func (s *Service) Lookup(ctx context.Context, id string) (*Resource, error) {
	var resp struct {
		UploadRequest Resource `json:"upload_request"`
	}
	if err := s.client.Do(ctx, "POST", "/upload_requests/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.UploadRequest, nil
}

func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/upload_requests/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

func (s *Service) Cancel(ctx context.Context, params CancelParams, opts ...request.Option) (*Resource, error) {
	var resp struct {
		UploadRequest Resource `json:"upload_request"`
	}
	if err := s.client.DoJSON(ctx, "/upload_requests/cancel", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.UploadRequest, nil
}

// Review records a manual approval or rejection for an upload attempt.
func (s *Service) Review(ctx context.Context, params ReviewParams, opts ...request.Option) (*Resource, error) {
	var resp struct {
		UploadRequest Resource `json:"upload_request"`
	}
	if err := s.client.DoJSON(ctx, "/upload_requests/review", params, request.Apply(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.UploadRequest, nil
}

func (s *Service) Fulfill(ctx context.Context, params FulfillParams) (*Resource, *file.Resource, error) {
	uploadFile, err := os.Open(params.File)
	if err != nil {
		return nil, nil, err
	}
	defer uploadFile.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", filepath.Base(params.File))
	if err != nil {
		return nil, nil, err
	}
	if _, err := io.Copy(part, uploadFile); err != nil {
		return nil, nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	var resp struct {
		UploadRequest Resource      `json:"upload_request"`
		File          file.Resource `json:"file"`
	}
	if err := s.client.DoRaw(ctx, "POST", params.UploadURL, &body, writer.FormDataContentType(), "", false, &resp, "upload_requests.upload"); err != nil {
		return nil, nil, err
	}
	return &resp.UploadRequest, &resp.File, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
