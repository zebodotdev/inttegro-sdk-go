package inttegro

import "context"

// AppsService manages applications associated with Inttegro API keys.
type AppsService struct {
	client *Client
}

// AppRelationshipPolicy describes who manages a child app relationship.
type AppRelationshipPolicy struct {
	// ChildStanding is the initial standing assigned to the child app.
	ChildStanding string `json:"child_standing,omitempty"`

	// Management controls who can manage the child app and its resources.
	Management AppManagementRole `json:"management,omitempty"`

	// Credentials controls who can create, rotate, or disable the child app's API keys.
	Credentials AppCredentialOwner `json:"credentials,omitempty"`
}

// CreateAppParams creates an Inttegro child app.
type CreateAppParams struct {
	// Name is the app's display name. It must not be empty after trimming.
	Name string `json:"name"`

	// Alias is an optional short label for dashboards, logs, or tooling.
	Alias string `json:"alias,omitempty"`

	// Description describes what the app is used for.
	Description string `json:"description,omitempty"`

	// LegalEntityType is a caller-defined legal-entity label.
	LegalEntityType string `json:"legal_entity_type,omitempty"`

	// PlacementParentApplicationID places the child under another app when separately authorized.
	PlacementParentApplicationID string `json:"placement_parent_application_id,omitempty"`

	// RelationshipPolicy customizes the direct parent-child relationship.
	RelationshipPolicy *AppRelationshipPolicy `json:"relationship_policy,omitempty"`
}

// UpdateAppParams changes mutable metadata on the authenticated app.
type UpdateAppParams struct {
	// Name changes the app's display name. It must not be empty after trimming when supplied.
	Name *string `json:"name,omitempty"`

	// Alias changes or clears the app's short label.
	Alias *string `json:"alias,omitempty"`

	// Description changes or clears the app description.
	Description *string `json:"description,omitempty"`

	// LegalEntityType changes or clears the caller-defined legal-entity label.
	LegalEntityType *string `json:"legal_entity_type,omitempty"`
}

// AppSecretKey is the initial secret key returned when an app is created.
// The token is returned only once and is not included in later app lookups.
type AppSecretKey struct {
	ID        string             `json:"id,omitempty"`
	TokenType SecretKeyTokenType `json:"token_type,omitempty"`
	IssuedAt  string             `json:"issued_at,omitempty"`
	Token     string             `json:"token,omitempty"`
}

// AppRelationship is the placement relationship receipt for a child app.
type AppRelationship struct {
	ID                             string                `json:"id"`
	Kind                           AppRelationshipKind   `json:"kind"`
	PolicyVersion                  string                `json:"policy_version"`
	Status                         AppRelationshipStatus `json:"status"`
	ActorAppID                     string                `json:"actor_app_id"`
	CreatorAppID                   string                `json:"creator_app_id"`
	PlacementParentAppID           string                `json:"placement_parent_app_id"`
	SubjectAppID                   string                `json:"subject_app_id"`
	ChildAppID                     string                `json:"child_app_id"`
	ChildStanding                  string                `json:"child_standing"`
	RelationshipPolicy             AppRelationshipPolicy `json:"relationship_policy"`
	RetainedCreatorAuthorityExists bool                  `json:"retained_creator_authority_exists"`
	CreatedAt                      string                `json:"created_at"`
}

// App represents an Inttegro account.
type App struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	Alias        string           `json:"alias,omitempty"`
	Description  string           `json:"description,omitempty"`
	CreatedAt    string           `json:"created_at"`
	UpdatedAt    string           `json:"updated_at,omitempty"`
	ArchivedAt   string           `json:"archived_at,omitempty"`
	SecretKey    *AppSecretKey    `json:"secret_key,omitempty"`
	Relationship *AppRelationship `json:"relationship,omitempty"`
}

// Create creates an Inttegro child app.
func (s *AppsService) Create(ctx context.Context, params CreateAppParams) (*App, error) {
	var resp struct {
		App App `json:"app"`
	}
	if err := s.client.do(ctx, "POST", "/apps/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.App, nil
}

// Lookup retrieves the application associated with the API key used for the request.
func (s *AppsService) Lookup(ctx context.Context) (*App, error) {
	var resp struct {
		App App `json:"app"`
	}
	if err := s.client.do(ctx, "POST", "/apps/lookup", map[string]any{}, &resp); err != nil {
		return nil, err
	}
	return &resp.App, nil
}

// Update changes attributes of the application associated with the API key used for the request.
func (s *AppsService) Update(ctx context.Context, params UpdateAppParams, opts ...RequestOption) (*App, error) {
	var resp struct {
		App App `json:"app"`
	}
	if err := s.client.doJSON(ctx, "/apps/update", params, applyRequestOptions(opts), &resp); err != nil {
		return nil, err
	}
	return &resp.App, nil
}
