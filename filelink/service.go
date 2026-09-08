package filelink

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/file"
	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

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
