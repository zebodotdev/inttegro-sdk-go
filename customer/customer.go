// Package customer provides customer resources and operations.
package customer

import inttegro "github.com/zebodotdev/inttegro-sdk-go/v4"

type (
	Service      = inttegro.CustomersService
	Data         = inttegro.CustomerData
	CreateParams = inttegro.CreateCustomerParams
	UpdateParams = inttegro.UpdateCustomerParams
	LookupParams = inttegro.LookupCustomerParams
	PageParams   = inttegro.PageCustomersParams
	Resource     = inttegro.Customer
	Page         = inttegro.CustomersPage
)
