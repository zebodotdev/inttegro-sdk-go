// Package app provides application resources and operations.
package app

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service            = inttegro.AppsService
	ManagementRole     = inttegro.AppManagementRole
	CredentialOwner    = inttegro.AppCredentialOwner
	RelationshipKind   = inttegro.AppRelationshipKind
	RelationshipStatus = inttegro.AppRelationshipStatus
	RelationshipPolicy = inttegro.AppRelationshipPolicy
	CreateParams       = inttegro.CreateAppParams
	UpdateParams       = inttegro.UpdateAppParams
	SecretKey          = inttegro.AppSecretKey
	Relationship       = inttegro.AppRelationship
	Resource           = inttegro.App
)

const (
	ManagementRoleParent = inttegro.AppManagementRoleParent
	ManagementRoleChild  = inttegro.AppManagementRoleChild

	CredentialOwnerChild  = inttegro.AppCredentialOwnerChild
	CredentialOwnerParent = inttegro.AppCredentialOwnerParent

	RelationshipKindPlacement = inttegro.AppRelationshipKindPlacement

	RelationshipStatusActive    = inttegro.AppRelationshipStatusActive
	RelationshipStatusInactive  = inttegro.AppRelationshipStatusInactive
	RelationshipStatusSuspended = inttegro.AppRelationshipStatusSuspended
	RelationshipStatusRevoked   = inttegro.AppRelationshipStatusRevoked
)
