package file

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/zebodotdev/inttegro-sdk-go/v6/internal/transport"
)

type Service struct {
	client transport.Client
}

func (s *Service) Create(ctx context.Context, params CreateParams) (*File, error) {
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
		File File `json:"file"`
	}
	if err := s.client.DoRaw(ctx, "POST", "/files/create", &body, writer.FormDataContentType(), params.IdempotencyKey, true, &resp, ""); err != nil {
		return nil, err
	}
	return &resp.File, nil
}

func (s *Service) Lookup(ctx context.Context, fileID string) (*File, error) {
	var resp struct {
		File File `json:"file"`
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

func (s *Service) Delete(ctx context.Context, fileID string) (*File, error) {
	var resp struct {
		File File `json:"file"`
	}
	if err := s.client.Do(ctx, "POST", "/files/delete", map[string]string{"file_id": fileID}, &resp); err != nil {
		return nil, err
	}
	return &resp.File, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
