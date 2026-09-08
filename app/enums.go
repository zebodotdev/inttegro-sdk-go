package app

type ManagementRole string

const (
	ManagementRoleParent ManagementRole = "parent"
	ManagementRoleChild  ManagementRole = "child"
)

type CredentialOwner string

const (
	CredentialOwnerChild  CredentialOwner = "child"
	CredentialOwnerParent CredentialOwner = "parent"
)

type RelationshipKind string

const RelationshipKindPlacement RelationshipKind = "placement"

type RelationshipStatus string

const (
	RelationshipStatusActive    RelationshipStatus = "active"
	RelationshipStatusInactive  RelationshipStatus = "inactive"
	RelationshipStatusSuspended RelationshipStatus = "suspended"
	RelationshipStatusRevoked   RelationshipStatus = "revoked"
)
