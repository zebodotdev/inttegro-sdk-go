package uploadrequest

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/zebodotdev/inttegro-sdk-go/v5/file"
	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

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
