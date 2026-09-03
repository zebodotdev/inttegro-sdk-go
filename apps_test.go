package inttegro

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestAppsServiceUsesTypedContracts(t *testing.T) {
	type capturedRequest struct {
		body           map[string]any
		idempotencyKey string
		path           string
	}

	var captured []capturedRequest
	client, close := newTestClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil && err != io.EOF {
			t.Fatalf("decode request: %v", err)
		}
		captured = append(captured, capturedRequest{
			body:           body,
			idempotencyKey: r.Header.Get("Idempotency-Key"),
			path:           r.URL.Path,
		})

		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/apps/create":
			_, _ = io.WriteString(w, createAppResponseJSON)
		case "/apps/lookup":
			_, _ = io.WriteString(w, lookupAppResponseJSON)
		case "/apps/update":
			_, _ = io.WriteString(w, updateAppResponseJSON)
		default:
			t.Fatalf("unexpected path %q", r.URL.Path)
		}
	}))
	if client == nil {
		return
	}
	defer close()
	if client.Apps == nil {
		t.Fatal("NewClient() did not initialize Apps")
	}

	ctx := context.Background()
	created, err := client.Apps.Create(ctx, CreateAppParams{
		Name:                         "Acme Production API",
		Alias:                        "acme-prod-api",
		Description:                  "Production Inttegro API for Acme Marketplace",
		LegalEntityType:              "business",
		PlacementParentApplicationID: "app_parent",
		RelationshipPolicy: &AppRelationshipPolicy{
			ChildStanding: "controlled",
			Credentials:   AppCredentialOwnerChild,
			Management:    AppManagementRoleParent,
		},
	})
	if err != nil {
		t.Fatalf("Apps.Create() error = %v", err)
	}
	lookedUp, err := client.Apps.Lookup(ctx)
	if err != nil {
		t.Fatalf("Apps.Lookup() error = %v", err)
	}
	updated, err := client.Apps.Update(ctx, UpdateAppParams{Alias: String("acme-api")}, WithIdempotencyKey("apps-update-123"))
	if err != nil {
		t.Fatalf("Apps.Update() error = %v", err)
	}

	wantPaths := []string{"/apps/create", "/apps/lookup", "/apps/update"}
	if len(captured) != len(wantPaths) {
		t.Fatalf("captured requests = %#v", captured)
	}
	for i, want := range wantPaths {
		if captured[i].path != want {
			t.Fatalf("request %d path = %q, want %q", i, captured[i].path, want)
		}
	}

	requestMeta, ok := captured[0].body["request_meta"].(map[string]any)
	if !ok {
		t.Fatalf("Apps.Create() body missing request_meta: %#v", captured[0].body)
	}
	generatedID, ok := requestMeta["idempotency_key"].(string)
	if !ok || !uuidV7Pattern.MatchString(generatedID) {
		t.Fatalf("generated idempotency key = %#v", requestMeta["idempotency_key"])
	}
	delete(captured[0].body, "request_meta")
	assertJSONMapEqual(t, captured[0].body, map[string]any{
		"alias":                           "acme-prod-api",
		"description":                     "Production Inttegro API for Acme Marketplace",
		"legal_entity_type":               "business",
		"name":                            "Acme Production API",
		"placement_parent_application_id": "app_parent",
		"relationship_policy": map[string]any{
			"child_standing": "controlled",
			"credentials":    "child",
			"management":     "parent",
		},
	})
	assertJSONMapEqual(t, captured[1].body, map[string]any{})
	assertJSONMapEqual(t, captured[2].body, map[string]any{"alias": "acme-api"})
	if captured[2].idempotencyKey != "apps-update-123" {
		t.Fatalf("Apps.Update() Idempotency-Key = %q", captured[2].idempotencyKey)
	}

	if created.ID != "app_child" || created.Name != "Acme Production API" {
		t.Fatalf("decoded created app = %#v", created)
	}
	if created.SecretKey == nil || created.SecretKey.TokenType != SecretKeyTokenTypeBearer || created.SecretKey.Token != "sk_test_child" {
		t.Fatalf("decoded create secret key = %#v", created.SecretKey)
	}
	if created.Relationship == nil ||
		created.Relationship.Kind != AppRelationshipKindPlacement ||
		created.Relationship.Status != AppRelationshipStatusActive ||
		created.Relationship.RelationshipPolicy.Management != AppManagementRoleParent ||
		created.Relationship.RelationshipPolicy.Credentials != AppCredentialOwnerChild {
		t.Fatalf("decoded relationship = %#v", created.Relationship)
	}
	if lookedUp.UpdatedAt == "" || lookedUp.ArchivedAt != "" {
		t.Fatalf("decoded lookup app = %#v", lookedUp)
	}
	if updated.Alias != "acme-api" {
		t.Fatalf("decoded updated app = %#v", updated)
	}
}

func TestUpdateAppParamsCanClearOptionalFields(t *testing.T) {
	raw, err := json.Marshal(UpdateAppParams{
		Alias:           String(""),
		Description:     String(""),
		LegalEntityType: String(""),
	})
	if err != nil {
		t.Fatalf("marshal update params: %v", err)
	}
	if string(raw) != `{"alias":"","description":"","legal_entity_type":""}` {
		t.Fatalf("UpdateAppParams JSON = %s", raw)
	}
}

const createAppResponseJSON = `{
  "app": {
    "id": "app_child",
    "name": "Acme Production API",
    "alias": "acme-prod-api",
    "description": "Production Inttegro API for Acme Marketplace",
    "created_at": "2026-09-02T10:00:00Z",
    "secret_key": {
      "id": "sk_child",
      "token_type": "bearer",
      "issued_at": "2026-09-02T10:00:00Z",
      "token": "sk_test_child"
    },
    "relationship": {
      "id": "rel_123",
      "kind": "placement",
      "policy_version": "app_relationship_authority.v1",
      "status": "active",
      "actor_app_id": "app_parent",
      "creator_app_id": "app_parent",
      "placement_parent_app_id": "app_parent",
      "subject_app_id": "app_child",
      "child_app_id": "app_child",
      "child_standing": "controlled",
      "relationship_policy": {
        "child_standing": "controlled",
        "credentials": "child",
        "management": "parent"
      },
      "retained_creator_authority_exists": true,
      "created_at": "2026-09-02T10:00:00Z"
    }
  }
}`

const lookupAppResponseJSON = `{
  "app": {
    "id": "app_child",
    "name": "Acme Production API",
    "alias": "acme-prod-api",
    "description": "Production Inttegro API for Acme Marketplace",
    "created_at": "2026-09-02T10:00:00Z",
    "updated_at": "2026-09-02T10:05:00Z"
  }
}`

const updateAppResponseJSON = `{
  "app": {
    "id": "app_child",
    "name": "Acme Production API",
    "alias": "acme-api",
    "description": "Production Inttegro API for Acme Marketplace",
    "created_at": "2026-09-02T10:00:00Z",
    "updated_at": "2026-09-02T10:10:00Z"
  }
}`
