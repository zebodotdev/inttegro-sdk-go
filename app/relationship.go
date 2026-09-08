package app

// AppRelationshipPolicy describes who manages a child app relationship.
type RelationshipPolicy struct {
	// ChildStanding is the initial standing assigned to the child app.
	ChildStanding string `json:"child_standing,omitempty"`

	// Management controls who can manage the child app and its resources.
	Management ManagementRole `json:"management,omitempty"`

	// Credentials controls who can create, rotate, or disable the child app's API keys.
	Credentials CredentialOwner `json:"credentials,omitempty"`
}

// AppRelationship is the placement relationship receipt for a child app.
type Relationship struct {
	ID                             string             `json:"id"`
	Kind                           RelationshipKind   `json:"kind"`
	PolicyVersion                  string             `json:"policy_version"`
	Status                         RelationshipStatus `json:"status"`
	ActorAppID                     string             `json:"actor_app_id"`
	CreatorAppID                   string             `json:"creator_app_id"`
	PlacementParentAppID           string             `json:"placement_parent_app_id"`
	SubjectAppID                   string             `json:"subject_app_id"`
	ChildAppID                     string             `json:"child_app_id"`
	ChildStanding                  string             `json:"child_standing"`
	RelationshipPolicy             RelationshipPolicy `json:"relationship_policy"`
	RetainedCreatorAuthorityExists bool               `json:"retained_creator_authority_exists"`
	CreatedAt                      string             `json:"created_at"`
}
