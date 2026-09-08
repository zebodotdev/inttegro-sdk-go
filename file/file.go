// Package file provides file resources and operations.
package file

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
)

type Status string

const (
	StatusUploading Status = "uploading"
)

const (
	StatusProcessing Status = "processing"
)

const (
	StatusAvailable Status = "available"
)

const (
	StatusFailed Status = "failed"
)

const (
	StatusDeleted Status = "deleted"
)

type Disposition string

const (
	DispositionAttachment Disposition = "attachment"
)

const (
	DispositionInline Disposition = "inline"
)

type Delivery string

const (
	DeliveryStream Delivery = "stream"
)

const (
	DeliveryRedirect Delivery = "redirect"
)

type ScanStatus string

const (
	ScanStatusPending ScanStatus = "pending"
)

const (
	ScanStatusPassed ScanStatus = "passed"
)

const (
	ScanStatusFailed ScanStatus = "failed"
)

const (
	ScanStatusSkipped ScanStatus = "skipped"
)

type SourceType string

const (
	SourceTypeDirect SourceType = "direct"
)

const (
	SourceTypeUploadRequest SourceType = "upload_request"
)

const (
	SourceTypeService SourceType = "service"
)

type StorageEncoding string

const (
	StorageEncodingIdentity StorageEncoding = "identity"
)

const (
	StorageEncodingBrotli StorageEncoding = "br"
)

type CreateParams struct {
	File           string            `json:"-"`
	Purpose        string            `json:"purpose"`
	Title          string            `json:"title,omitempty"`
	CustomData     map[string]string `json:"custom_data,omitempty"`
	IdempotencyKey string            `json:"-"`
}

type PageParams struct {
	CreatedAfter  string `json:"created_after,omitempty"`
	CreatedBefore string `json:"created_before,omitempty"`
	PageNumber    int    `json:"page_number,omitempty"`
	PageSize      int    `json:"page_size,omitempty"`
	Purpose       string `json:"purpose,omitempty"`
	Status        Status `json:"status,omitempty"`
}

type ContentsParams struct {
	FileID      string      `json:"file_id"`
	Disposition Disposition `json:"disposition,omitempty"`
}

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

type Download struct {
	io.ReadCloser
}

func (d *Download) SaveTo(path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, d.ReadCloser)
	return err
}

type Service struct {
	client transport.Client
}

func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	file, err := os.Open(params.File)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	if err := writer.WriteField("purpose", params.Purpose); err != nil {
		return nil, err
	}
	if params.Title != "" {
		if err := writer.WriteField("title", params.Title); err != nil {
			return nil, err
		}
	}
	if params.CustomData != nil {
		raw, err := json.Marshal(params.CustomData)
		if err != nil {
			return nil, err
		}
		if err := writer.WriteField("custom_data", string(raw)); err != nil {
			return nil, err
		}
	}
	part, err := writer.CreateFormFile("file", filepath.Base(params.File))
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	var resp struct {
		File Resource `json:"file"`
	}
	if err := s.client.DoRaw(ctx, "POST", "/files/create", &body, writer.FormDataContentType(), params.IdempotencyKey, true, &resp, ""); err != nil {
		return nil, err
	}
	return &resp.File, nil
}

func (s *Service) Lookup(ctx context.Context, fileID string) (*Resource, error) {
	var resp struct {
		File Resource `json:"file"`
	}
	if err := s.client.Do(ctx, "POST", "/files/lookup", map[string]string{"file_id": fileID}, &resp); err != nil {
		return nil, err
	}
	return &resp.File, nil
}

func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/files/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

func (s *Service) Contents(ctx context.Context, params ContentsParams) (*Download, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.RawResponse(ctx, "POST", "/files/contents", bytes.NewReader(body), "application/json", "", true, "")
	if err != nil {
		return nil, err
	}
	return &Download{ReadCloser: resp.Body}, nil
}

func (s *Service) Delete(ctx context.Context, fileID string) (*Resource, error) {
	var resp struct {
		File Resource `json:"file"`
	}
	if err := s.client.Do(ctx, "POST", "/files/delete", map[string]string{"file_id": fileID}, &resp); err != nil {
		return nil, err
	}
	return &resp.File, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
