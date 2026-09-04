package inttegro

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type RequestOption func(*requestOptions)

type requestOptions struct {
	IdempotencyKey string
}

func WithIdempotencyKey(key string) RequestOption {
	return func(o *requestOptions) {
		o.IdempotencyKey = key
	}
}

type FileCreateParams struct {
	File           string            `json:"-"`
	Purpose        string            `json:"purpose"`
	Title          string            `json:"title,omitempty"`
	CustomData     map[string]string `json:"custom_data,omitempty"`
	IdempotencyKey string            `json:"-"`
}

type FilePageParams struct {
	CreatedAfter  string     `json:"created_after,omitempty"`
	CreatedBefore string     `json:"created_before,omitempty"`
	PageNumber    int        `json:"page_number,omitempty"`
	PageSize      int        `json:"page_size,omitempty"`
	Purpose       string     `json:"purpose,omitempty"`
	Status        FileStatus `json:"status,omitempty"`
}

type FileContentsParams struct {
	FileID      string          `json:"file_id"`
	Disposition FileDisposition `json:"disposition,omitempty"`
}

type File struct {
	ID         string            `json:"id"`
	Purpose    string            `json:"purpose"`
	Status     FileStatus        `json:"status"`
	CustomData map[string]string `json:"custom_data,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

type FilesPage struct {
	Number int    `json:"number"`
	Size   int    `json:"size"`
	Files  []File `json:"files"`
}

type FileDownload struct {
	io.ReadCloser
}

func (d *FileDownload) SaveTo(path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, d.ReadCloser)
	return err
}

type FilesService struct {
	client *Client
}

func (s *FilesService) Create(ctx context.Context, params FileCreateParams) (*File, error) {
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
	if err := s.client.doRaw(ctx, "POST", "/files/create", &body, writer.FormDataContentType(), params.IdempotencyKey, true, &resp, ""); err != nil {
		return nil, err
	}
	return &resp.File, nil
}

func (s *FilesService) Lookup(ctx context.Context, fileID string) (*File, error) {
	var resp struct {
		File File `json:"file"`
	}
	if err := s.client.do(ctx, "POST", "/files/lookup", map[string]string{"file_id": fileID}, &resp); err != nil {
		return nil, err
	}
	return &resp.File, nil
}

func (s *FilesService) Page(ctx context.Context, params FilePageParams) (*FilesPage, error) {
	var resp struct {
		Page FilesPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/files/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

func (s *FilesService) Contents(ctx context.Context, params FileContentsParams) (*FileDownload, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.rawResponse(ctx, "POST", "/files/contents", bytes.NewReader(body), "application/json", "", true, "")
	if err != nil {
		return nil, err
	}
	return &FileDownload{ReadCloser: resp.Body}, nil
}

func (s *FilesService) Delete(ctx context.Context, fileID string) (*File, error) {
	var resp struct {
		File File `json:"file"`
	}
	if err := s.client.do(ctx, "POST", "/files/delete", map[string]string{"file_id": fileID}, &resp); err != nil {
		return nil, err
	}
	return &resp.File, nil
}

type FileActor struct {
	Type  string `json:"type,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type FileLinkDelivery struct {
	Mode        FileLinkDeliveryMode `json:"mode,omitempty"`
	Filename    string               `json:"filename,omitempty"`
	ContentType string               `json:"content_type,omitempty"`
	Disposition string               `json:"disposition,omitempty"`
}

type FileLinkAccess struct {
	MaxAccesses    int64    `json:"max_accesses,omitempty"`
	AllowDownload  bool     `json:"allow_download,omitempty"`
	AllowedOrigins []string `json:"allowed_origins,omitempty"`
}

type FileLinkCreateParams struct {
	Access     FileLinkAccess    `json:"access"`
	CreatedBy  FileActor         `json:"created_by"`
	Delivery   FileLinkDelivery  `json:"delivery"`
	ExpiresAt  string            `json:"expires_at,omitempty"`
	FileID     string            `json:"file_id"`
	CustomData map[string]string `json:"custom_data,omitempty"`
}

type FileLinkPageParams struct {
	FileID     string         `json:"file_id,omitempty"`
	PageNumber int            `json:"page_number,omitempty"`
	PageSize   int            `json:"page_size,omitempty"`
	Status     FileLinkStatus `json:"status,omitempty"`
}

type FileLinkRevokeParams struct {
	ID        string    `json:"id"`
	RevokedBy FileActor `json:"revoked_by"`
}

type FileLink struct {
	ID         string            `json:"id"`
	FileID     string            `json:"file_id"`
	Status     FileLinkStatus    `json:"status"`
	CustomData map[string]string `json:"custom_data,omitempty"`
	Metadata   map[string]string `json:"metadata,omitempty"`
}

type FileLinksPage struct {
	Number    int        `json:"number"`
	Size      int        `json:"size"`
	FileLinks []FileLink `json:"file_links"`
}

type FileLinksService struct {
	client *Client
}

func (s *FileLinksService) Create(ctx context.Context, params FileLinkCreateParams, opts ...RequestOption) (*FileLink, string, error) {
	var resp struct {
		FileLink FileLink `json:"file_link"`
		URL      string   `json:"url"`
	}
	if err := s.client.doJSON(ctx, "/file_links/create", params, applyRequestOptions(opts), &resp); err != nil {
		return nil, "", err
	}
	return &resp.FileLink, resp.URL, nil
}

func (s *FileLinksService) Lookup(ctx context.Context, id string) (*FileLink, error) {
	var resp struct {
		FileLink FileLink `json:"file_link"`
	}
	if err := s.client.do(ctx, "POST", "/file_links/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.FileLink, nil
}

func (s *FileLinksService) Page(ctx context.Context, params FileLinkPageParams) (*FileLinksPage, error) {
	var resp struct {
		Page FileLinksPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/file_links/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

func (s *FileLinksService) Revoke(ctx context.Context, params FileLinkRevokeParams, opts ...RequestOption) (*FileLink, error) {
	var resp struct {
		FileLink FileLink `json:"file_link"`
	}
	if err := s.client.doJSON(ctx, "/file_links/revoke", params, applyRequestOptions(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.FileLink, nil
}

func (s *FileLinksService) Open(ctx context.Context, url string) (*FileDownload, error) {
	resp, err := s.client.rawResponse(ctx, "GET", url, nil, "", "", false, "file_links.download")
	if err != nil {
		return nil, err
	}
	return &FileDownload{ReadCloser: resp.Body}, nil
}

type UploadRequestConstraints struct {
	ContentTypes []string `json:"content_types,omitempty"`
	ExactSize    int64    `json:"exact_size,omitempty"`
	Extensions   []string `json:"extensions,omitempty"`
	Filename     string   `json:"filename,omitempty"`
	MaxSize      int64    `json:"max_size,omitempty"`
	MinSize      int64    `json:"min_size,omitempty"`
}

type UploadRequestDisplay struct {
	Description string `json:"description,omitempty"`
	HelpText    string `json:"help_text,omitempty"`
	Title       string `json:"title,omitempty"`
}

type FileParty struct {
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

type UploadRequestCreateParams struct {
	Constraints UploadRequestConstraints `json:"constraints"`
	Display     UploadRequestDisplay     `json:"display"`
	ExpiresAt   string                   `json:"expires_at,omitempty"`
	CustomData  map[string]string        `json:"custom_data,omitempty"`
	Purpose     string                   `json:"purpose"`
	Recipient   FileParty                `json:"recipient"`
	Requester   FileActor                `json:"requester"`
	Resource    FileResource             `json:"resource"`
	Subject     FileParty                `json:"subject"`
}

type UploadRequestPageParams struct {
	PageNumber int                 `json:"page_number,omitempty"`
	PageSize   int                 `json:"page_size,omitempty"`
	Purpose    string              `json:"purpose,omitempty"`
	Resource   FileResource        `json:"resource"`
	Status     UploadRequestStatus `json:"status,omitempty"`
}

type UploadRequestCancelParams struct {
	CanceledBy FileActor `json:"canceled_by"`
	ID         string    `json:"id"`
}

type UploadRequestReviewReason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
}

type UploadRequestReviewParams struct {
	AttemptID      string                      `json:"attempt_id,omitempty"`
	AttemptOrdinal int64                       `json:"attempt_ordinal,omitempty"`
	Decision       string                      `json:"decision"`
	ID             string                      `json:"id"`
	PublicMessage  string                      `json:"public_message,omitempty"`
	Reasons        []UploadRequestReviewReason `json:"reasons,omitempty"`
}

type UploadRequestFulfillParams struct {
	File      string
	UploadURL string
}

type UploadRequest struct {
	ID         string              `json:"id"`
	Purpose    string              `json:"purpose"`
	Status     UploadRequestStatus `json:"status"`
	UploadURL  string              `json:"upload_url,omitempty"`
	CustomData map[string]string   `json:"custom_data,omitempty"`
	Metadata   map[string]string   `json:"metadata,omitempty"`
}

type UploadRequestsPage struct {
	Number         int             `json:"number"`
	Size           int             `json:"size"`
	UploadRequests []UploadRequest `json:"upload_requests"`
}

type UploadRequestsService struct {
	client *Client
}

func (s *UploadRequestsService) Create(ctx context.Context, params UploadRequestCreateParams, opts ...RequestOption) (*UploadRequest, error) {
	var resp struct {
		UploadRequest UploadRequest `json:"upload_request"`
	}
	if err := s.client.doJSON(ctx, "/upload_requests/create", params, applyRequestOptions(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.UploadRequest, nil
}

func (s *UploadRequestsService) Lookup(ctx context.Context, id string) (*UploadRequest, error) {
	var resp struct {
		UploadRequest UploadRequest `json:"upload_request"`
	}
	if err := s.client.do(ctx, "POST", "/upload_requests/lookup", map[string]string{"id": id}, &resp); err != nil {
		return nil, err
	}
	return &resp.UploadRequest, nil
}

func (s *UploadRequestsService) Page(ctx context.Context, params UploadRequestPageParams) (*UploadRequestsPage, error) {
	var resp struct {
		Page UploadRequestsPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/upload_requests/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

func (s *UploadRequestsService) Cancel(ctx context.Context, params UploadRequestCancelParams, opts ...RequestOption) (*UploadRequest, error) {
	var resp struct {
		UploadRequest UploadRequest `json:"upload_request"`
	}
	if err := s.client.doJSON(ctx, "/upload_requests/cancel", params, applyRequestOptions(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.UploadRequest, nil
}

// Review records a manual approval or rejection for an upload attempt.
func (s *UploadRequestsService) Review(ctx context.Context, params UploadRequestReviewParams, opts ...RequestOption) (*UploadRequest, error) {
	var resp struct {
		UploadRequest UploadRequest `json:"upload_request"`
	}
	if err := s.client.doJSON(ctx, "/upload_requests/review", params, applyRequestOptions(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.UploadRequest, nil
}

func (s *UploadRequestsService) Fulfill(ctx context.Context, params UploadRequestFulfillParams) (*UploadRequest, *File, error) {
	file, err := os.Open(params.File)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("file", filepath.Base(params.File))
	if err != nil {
		return nil, nil, err
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	var resp struct {
		UploadRequest UploadRequest `json:"upload_request"`
		File          File          `json:"file"`
	}
	if err := s.client.doRaw(ctx, "POST", params.UploadURL, &body, writer.FormDataContentType(), "", false, &resp, "upload_requests.upload"); err != nil {
		return nil, nil, err
	}
	return &resp.UploadRequest, &resp.File, nil
}

func applyRequestOptions(opts []RequestOption) requestOptions {
	var out requestOptions
	for _, opt := range opts {
		opt(&out)
	}
	return out
}

func (c *Client) doJSON(ctx context.Context, path string, body any, opts requestOptions, out any) error {
	raw, err := c.jsonRequestBody("POST", path, body, opts.IdempotencyKey)
	if err != nil {
		return err
	}
	return c.doRaw(ctx, "POST", path, bytes.NewReader(raw), "application/json", opts.IdempotencyKey, true, out, "")
}

func (c *Client) doRaw(ctx context.Context, method, pathOrURL string, body io.Reader, contentType, idempotencyKey string, authenticated bool, out any, operation string) error {
	resp, err := c.rawResponse(ctx, method, pathOrURL, body, contentType, idempotencyKey, authenticated, operation)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if out == nil {
		return nil
	}
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respBytes) == 0 {
		return nil
	}
	return json.Unmarshal(respBytes, out)
}

func (c *Client) rawResponse(ctx context.Context, method, pathOrURL string, body io.Reader, contentType, idempotencyKey string, authenticated bool, operation string) (*http.Response, error) {
	ctx, telemetry := c.startRequestTelemetry(ctx, method, pathOrURL, operation)
	defer telemetry.end()

	url := pathOrURL
	if len(pathOrURL) == 0 || pathOrURL[0] == '/' {
		url = c.BaseURL + pathOrURL
	}
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		telemetry.fail("request_error")
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	if authenticated {
		req.Header.Set("Authorization", "Bearer "+c.APIKey)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if idempotencyKey != "" {
		req.Header.Set("Idempotency-Key", idempotencyKey)
	} else if authenticated && strings.EqualFold(method, "POST") && isIdempotentMutationPath(pathOrURL) && !strings.HasPrefix(contentType, "application/json") {
		req.Header.Set("Idempotency-Key", generateIdempotencyKey())
	}
	telemetry.inject(ctx, req.Header)
	telemetry.attempt()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		telemetry.fail(classifyTelemetryError(err, "transport_error"))
		return nil, err
	}
	telemetry.response(resp)
	if resp.StatusCode < 400 {
		return resp, nil
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		telemetry.fail("read_error")
		return nil, err
	}
	apiErr := &APIError{StatusCode: resp.StatusCode, Body: respBytes}
	var parsed APIError
	if err := json.Unmarshal(respBytes, &parsed); err == nil && hasAPIErrorDetails(&parsed) {
		copyAPIError(apiErr, &parsed)
	} else if len(respBytes) > 0 {
		apiErr.Message = string(respBytes)
	}
	telemetry.fail(fmt.Sprintf("http_%d", resp.StatusCode))
	return nil, fmt.Errorf("inttegro api error: %w", apiErr)
}
