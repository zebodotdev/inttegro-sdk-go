package messagetemplate

import (
	"context"

	"github.com/zebodotdev/inttegro-sdk-go/v5/internal/transport"
	"github.com/zebodotdev/inttegro-sdk-go/v5/request"
)

type Service struct {
	client transport.Client
}

func (s *Service) Create(ctx context.Context, params CreateParams) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	opts := request.Options{IdempotencyKey: params.IdempotencyKey}
	if err := s.client.DoJSON(ctx, "/message_templates/create", params, opts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Update(ctx context.Context, params UpdateParams) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	opts := request.Options{IdempotencyKey: params.IdempotencyKey}
	if err := s.client.DoJSON(ctx, "/message_templates/update", params, opts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Publish(ctx context.Context, id string, opts ...request.Option) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	requestOpts := request.Apply(opts)
	if err := s.client.DoJSON(ctx, "/message_templates/publish", map[string]string{"id": id}, requestOpts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Archive(ctx context.Context, id string, opts ...request.Option) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	requestOpts := request.Apply(opts)
	if err := s.client.DoJSON(ctx, "/message_templates/archive", map[string]string{"id": id}, requestOpts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Lookup(ctx context.Context, id string) (*Resource, error) {
	var resp struct {
		MessageTemplate Resource `json:"message_template"`
	}
	if err := s.client.Do(ctx, "POST", "/message_templates/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *Service) Page(ctx context.Context, params PageParams) (*Page, error) {
	var resp struct {
		Page Page `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/message_templates/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

func (s *Service) RenderPreview(ctx context.Context, params RenderPreviewParams) (*RenderPreviewOutput, error) {
	var resp RenderPreviewOutput
	if err := s.client.Do(ctx, "POST", "/message_templates/render_preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// NewService constructs the resource service used by inttegro.Client.
func NewService(client transport.Client) *Service { return &Service{client: client} }
