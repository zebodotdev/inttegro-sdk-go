// Package invoice provides invoice resources and operations.
package invoice

type DocumentKind string

const (
	DocumentKindInvoice DocumentKind = "invoice"
	DocumentKindReceipt DocumentKind = "receipt"
)

type DeliveryChannel string

const (
	DeliveryChannelEmail DeliveryChannel = "email"
	DeliveryChannelSMS   DeliveryChannel = "sms"
)

// OrderDocumentDelivery describes a delivered invoice or receipt link.
type Delivery struct {
	DocumentKind   DocumentKind      `json:"document_kind"`
	DocumentURL    string            `json:"document_url"`
	SentChannels   []DeliveryChannel `json:"sent_channels,omitempty"`
	FailedChannels []DeliveryChannel `json:"failed_channels,omitempty"`
	Deliveries     []DeliveryAttempt `json:"deliveries,omitempty"`
	Failures       []DeliveryAttempt `json:"failures,omitempty"`
}

// OrderDocumentDeliveryAttempt describes one delivery channel result.
type DeliveryAttempt struct {
	Channel DeliveryChannel `json:"channel"`
	ChimeID string          `json:"chime_id,omitempty"`
	Error   string          `json:"error,omitempty"`
}

// InvoiceFormat contains URLs for invoice documents in different formats.
type Format struct {
	// Web contains the hosted invoice page URL.
	Web *struct {
		// URL is the HTTPS link to view the invoice in a browser.
		// Includes line items, totals, payment button.
		URL string `json:"url"`
	} `json:"web,omitempty"`

	// PDF contains the downloadable PDF invoice URL.
	PDF *struct {
		// URL is the HTTPS link to download PDF invoice.
		URL string `json:"url"`
	} `json:"pdf,omitempty"`
}

// Invoice represents an order's invoice document and delivery status.
//
// Invoices are generated when orders are finalized. They provide
// customer-facing links for viewing and paying orders.
type Invoice struct {
	// ID is the unique invoice identifier.
	// Starts with "inv_". Example: "inv_abc123def456"
	ID string `json:"id,omitempty"`

	// Number is the human-readable invoice number.
	// Example: "INV-2023-00123"
	Number string `json:"number,omitempty"`

	// Format contains links to invoice documents.
	Format *Format `json:"format,omitempty"`

	// Deliveries tracks invoice email/SMS delivery attempts.
	Deliveries any `json:"deliveries,omitempty"`
}
