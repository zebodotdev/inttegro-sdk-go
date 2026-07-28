package commerce

import "context"

type MessageTemplateVariable struct {
	About    string                    `json:"about,omitempty"`
	Default  any                       `json:"default,omitempty"`
	Items    []MessageTemplateVariable `json:"items,omitempty"`
	Name     string                    `json:"name"`
	Required bool                      `json:"required,omitempty"`
	Type     string                    `json:"type"`
}

type MessageTemplateSMSContent struct {
	MessageTemplate string `json:"message_template"`
}

type MessageTemplateMailbox struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type MessageTemplateEmailContent struct {
	From    *MessageTemplateMailbox `json:"from,omitempty"`
	Headers map[string]string       `json:"headers,omitempty"`
	HTML    string                  `json:"html"`
	ReplyTo *MessageTemplateMailbox `json:"reply_to,omitempty"`
	Subject string                  `json:"subject"`
}

type MessageTemplate struct {
	ID                    string                       `json:"id"`
	About                 string                       `json:"about,omitempty"`
	ArchivedAt            string                       `json:"archived_at,omitempty"`
	Attachments           []string                     `json:"attachments,omitempty"`
	Channel               string                       `json:"channel"`
	CreatedAt             string                       `json:"created_at"`
	DraftVersion          int                          `json:"draft_version"`
	Email                 *MessageTemplateEmailContent `json:"email,omitempty"`
	HasUnpublishedChanges bool                         `json:"has_unpublished_changes"`
	Locale                string                       `json:"locale"`
	Name                  string                       `json:"name"`
	PublishedAt           string                       `json:"published_at,omitempty"`
	PublishedVersion      *int                         `json:"published_version,omitempty"`
	Purpose               string                       `json:"purpose"`
	SMS                   *MessageTemplateSMSContent   `json:"sms,omitempty"`
	Status                string                       `json:"status"`
	UpdatedAt             string                       `json:"updated_at"`
	Variables             []MessageTemplateVariable    `json:"variables,omitempty"`
	Version               int                          `json:"version"`
}

type MessageTemplateCreateParams struct {
	IdempotencyKey string                       `json:"-"`
	About          string                       `json:"about,omitempty"`
	Attachments    []string                     `json:"attachments,omitempty"`
	Channel        string                       `json:"channel"`
	Email          *MessageTemplateEmailContent `json:"email,omitempty"`
	Locale         string                       `json:"locale,omitempty"`
	Name           string                       `json:"name"`
	Purpose        string                       `json:"purpose"`
	SMS            *MessageTemplateSMSContent   `json:"sms,omitempty"`
	Variables      []MessageTemplateVariable    `json:"variables,omitempty"`
}

type MessageTemplateUpdateParams struct {
	IdempotencyKey string                       `json:"-"`
	ID             string                       `json:"id"`
	About          string                       `json:"about,omitempty"`
	Attachments    []string                     `json:"attachments,omitempty"`
	Channel        string                       `json:"channel,omitempty"`
	Email          *MessageTemplateEmailContent `json:"email,omitempty"`
	Locale         string                       `json:"locale,omitempty"`
	Name           string                       `json:"name,omitempty"`
	Purpose        string                       `json:"purpose,omitempty"`
	SMS            *MessageTemplateSMSContent   `json:"sms,omitempty"`
	Variables      []MessageTemplateVariable    `json:"variables,omitempty"`
}

type MessageTemplatePageParams struct {
	Channel string `json:"channel,omitempty"`
	Locale  string `json:"locale,omitempty"`
	Page    int    `json:"page,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Size    int    `json:"size,omitempty"`
	Status  string `json:"status,omitempty"`
}

type MessageTemplatePage struct {
	Number           int               `json:"number"`
	Size             int               `json:"size"`
	MessageTemplates []MessageTemplate `json:"message_templates"`
}

type MessageTemplateReference struct {
	TemplateID string         `json:"template_id"`
	Variables  map[string]any `json:"variables,omitempty"`
}

type MessageTemplateRenderPreviewParams struct {
	MessageTemplate MessageTemplateReference `json:"message_template"`
}

type MessageTemplateRenderedContent struct {
	Channel string         `json:"channel"`
	Email   map[string]any `json:"email,omitempty"`
	SMS     map[string]any `json:"sms,omitempty"`
}

type MessageTemplateRenderPreviewOutput struct {
	MessageTemplate *MessageTemplate                `json:"message_template,omitempty"`
	Rendered        *MessageTemplateRenderedContent `json:"rendered,omitempty"`
}

type MessageTemplatesService struct {
	client *Client
}

func (s *MessageTemplatesService) Create(ctx context.Context, params MessageTemplateCreateParams) (*MessageTemplate, error) {
	var resp struct {
		MessageTemplate MessageTemplate `json:"message_template"`
	}
	opts := requestOptions{IdempotencyKey: params.IdempotencyKey}
	if err := s.client.doJSON(ctx, "/message_templates/create", params, opts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *MessageTemplatesService) Update(ctx context.Context, params MessageTemplateUpdateParams) (*MessageTemplate, error) {
	var resp struct {
		MessageTemplate MessageTemplate `json:"message_template"`
	}
	opts := requestOptions{IdempotencyKey: params.IdempotencyKey}
	if err := s.client.doJSON(ctx, "/message_templates/update", params, opts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *MessageTemplatesService) Publish(ctx context.Context, id string, opts ...RequestOption) (*MessageTemplate, error) {
	var resp struct {
		MessageTemplate MessageTemplate `json:"message_template"`
	}
	requestOpts := requestOptionsFromOptions(opts)
	if err := s.client.doJSON(ctx, "/message_templates/publish", map[string]string{"id": id}, requestOpts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *MessageTemplatesService) Archive(ctx context.Context, id string, opts ...RequestOption) (*MessageTemplate, error) {
	var resp struct {
		MessageTemplate MessageTemplate `json:"message_template"`
	}
	requestOpts := requestOptionsFromOptions(opts)
	if err := s.client.doJSON(ctx, "/message_templates/archive", map[string]string{"id": id}, requestOpts, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *MessageTemplatesService) Lookup(ctx context.Context, id string) (*MessageTemplate, error) {
	var resp struct {
		MessageTemplate MessageTemplate `json:"message_template"`
	}
	if err := s.client.do(ctx, "POST", "/message_templates/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.MessageTemplate, nil
}

func (s *MessageTemplatesService) Page(ctx context.Context, params MessageTemplatePageParams) (*MessageTemplatePage, error) {
	var resp struct {
		Page MessageTemplatePage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/message_templates/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

func (s *MessageTemplatesService) RenderPreview(ctx context.Context, params MessageTemplateRenderPreviewParams) (*MessageTemplateRenderPreviewOutput, error) {
	var resp MessageTemplateRenderPreviewOutput
	if err := s.client.do(ctx, "POST", "/message_templates/render_preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func requestOptionsFromOptions(opts []RequestOption) requestOptions {
	var out requestOptions
	for _, opt := range opts {
		if opt != nil {
			opt(&out)
		}
	}
	return out
}
