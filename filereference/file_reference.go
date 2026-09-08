// Package filereference provides file-reference reconciliation resources and operations.
package filereference

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service         = inttegro.FileReferencesService
	Input           = inttegro.FileReferenceInput
	ReconcileParams = inttegro.FileReferenceReconcileParams
	Reconciliation  = inttegro.FileReferenceReconciliation
	Resource        = inttegro.FileReferenceReconciliation
)
