package commerce

// RequestMeta carries per-request controls that do not change the operation payload.
type RequestMeta struct {
	// IdempotencyKey safely identifies retries for mutation requests.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

// Money represents a monetary amount in a specific currency.
//
// Amounts are always expressed in minor units (cents, pesewas, centimes, etc).
// For example, $10.50 USD is represented as Currency: "usd", Value: 1050.
//
// Example:
//
//	// Ten dollars and fifty cents
//	amount := commerce.Money{Currency: "usd", Value: 1050}
//
//	// Five thousand Ghanaian cedis (GHS 50.00)
//	amount := commerce.Money{Currency: "ghs", Value: 5000}
//
// Currency codes follow ISO 4217 (lowercase). Value must be a positive integer
// for most operations (charges, payouts, line items).
type Money struct {
	// Currency is the three-letter ISO 4217 currency code (lowercase).
	// Supported currencies: usd, ghs, kes, ugx, tzs, xof, xaf, and more.
	// Query /spec/countries to see supported currencies by country.
	Currency string `json:"currency"`

	// Value is the amount in minor units (cents, pesewas, etc).
	// For zero-decimal currencies (e.g., JPY), this represents the actual amount.
	// Must be positive for charges and payouts.
	Value int64 `json:"value"`
}

// PaymentMethodType enumerates supported payment rails.
//
// Each type represents a distinct payment method category with different
// capabilities, verification requirements, and regional availability.
type PaymentMethodType string

const (
	// PaymentMethodTypeMobileMoney represents mobile money wallets.
	// Includes MTN Mobile Money, Vodafone Cash, Airtel Money, Tigo Cash, and others.
	// Requires network and account number.
	// Supports OTP verification via confirms_use setting.
	PaymentMethodTypeMobileMoney PaymentMethodType = "mobile_money"

	// PaymentMethodTypeBankAccount represents bank accounts.
	// Currently limited availability. Check country specifications.
	PaymentMethodTypeBankAccount PaymentMethodType = "bank_account"

	// PaymentMethodTypeCard represents credit and debit cards.
	// Currently limited availability. Check country specifications.
	PaymentMethodTypeCard PaymentMethodType = "card"

	// PaymentMethodTypeMotito represents Zebo's branded payment method.
	// Internal use only.
	PaymentMethodTypeMotito PaymentMethodType = "motito"
)

// MobileMoneyParams describes a mobile money wallet.
//
// Used when creating orders or tokenizing payment methods with inline
// mobile money data instead of referencing a saved payment method.
type MobileMoneyParams struct {
	// Network is the mobile money network code.
	// Examples: "mtn", "vodafone", "airteltigo", "airtel", "telecel".
	Network string `json:"network"`

	// AccountNumber is the mobile money account phone number.
	// Must include country code. Example: "+233244123456"
	// Used as the payment source and for sending OTP verification codes.
	AccountNumber string `json:"account_number"`
}

// PaymentMethodData represents inline payment method data for one-time use.
//
// Use this to charge a payment method without saving it for future use.
// For repeat customers, tokenize the payment method first using
// PaymentMethods.Tokenize, then reference it by ID.
//
// Example (mobile money):
//
//	paymentData := &commerce.PaymentMethodData{
//	    Type: commerce.PaymentMethodTypeMobileMoney,
//	    MobileMoney: &commerce.MobileMoneyParams{
//	        Network: "mtn",
//	        AccountNumber: "+233244123456",
//	    },
//	}
type PaymentMethodData struct {
	// Type specifies the payment method category.
	Type PaymentMethodType `json:"type"`

	// MobileMoney provides mobile money wallet details.
	// Required when Type is PaymentMethodTypeMobileMoney.
	MobileMoney *MobileMoneyParams `json:"mobile_money,omitempty"`

	// Additional fields for other payment method types (bank, card)
	// will be added as those types become available.
}

// CustomerData captures inline customer information for order creation.
//
// Use this when creating orders for new customers who don't have a customer ID yet.
// The API will create a customer record automatically and return it in the response.
//
// For existing customers, pass CustomerID instead to link the order to their record.
//
// Example:
//
//	customerData := &commerce.CustomerData{
//	    Name:        "Jane Doe",
//	    Email:       "jane@example.com",
//	    PhoneNumber: "+233244123456",
//	    Reference:   "customer_123_in_my_system",
//	}
type CustomerData struct {
	// Name is the customer's full name (required).
	// Used for billing records and invoice generation.
	Name string `json:"name"`

	// Email is the customer's email address (optional but recommended).
	// Used for sending invoice links and payment notifications.
	// Validated as RFC 5322 email format.
	Email string `json:"email_address,omitempty"`

	// PhoneNumber is the customer's phone number with country code (required).
	// Used for SMS notifications and payment confirmations.
	// Example: "+233244123456"
	PhoneNumber string `json:"phone_number"`

	// Reference is your internal customer identifier (optional).
	// Use this to link Commerce customer records to your system's users.
	// Maximum 255 characters. Must be unique across your customers.
	Reference string `json:"reference,omitempty"`

	// CustomData holds arbitrary key-value pairs about the customer (optional).
	// Both keys and values must be strings. Maximum 25KB when serialized.
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// ProductCategory describes a product category.
type ProductCategory struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// ProductPrice represents a catalog product price.
type ProductPrice struct {
	Amount   int64  `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
}

// ProductPriceAmount represents a product price amount in smallest currency units.
type ProductPriceAmount struct {
	Currency string `json:"currency"`
	Value    int64  `json:"value"`
}

// ProductDefaultUnitPrice represents a product's loaded default unit price.
type ProductDefaultUnitPrice struct {
	ID         string              `json:"id,omitempty"`
	ProductID  string              `json:"product_id,omitempty"`
	Label      string              `json:"label,omitempty"`
	About      string              `json:"about,omitempty"`
	Nominal    *ProductPriceAmount `json:"nominal,omitempty"`
	CreatedAt  string              `json:"created_at,omitempty"`
	UpdatedAt  string              `json:"updated_at,omitempty"`
	ArchivedAt string              `json:"archived_at,omitempty"`
}

// ProductPriceSummary represents a product price listed with a product response.
type ProductPriceSummary struct {
	ID      string              `json:"id,omitempty"`
	Label   string              `json:"label,omitempty"`
	Nominal *ProductPriceAmount `json:"nominal,omitempty"`
}

// ProductShipmentDimensions describes physical dimensions.
type ProductShipmentDimensions struct {
	Length float64 `json:"length,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`
	Weight float64 `json:"weight,omitempty"`
}

// ProductShipment describes fulfillment details.
type ProductShipment struct {
	Type       string                     `json:"type,omitempty"`
	Carrier    string                     `json:"carrier,omitempty"`
	Dimensions *ProductShipmentDimensions `json:"dimensions,omitempty"`
}

// ProductMediaItem represents product media.
type ProductMediaItem struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}

// CreateProductParams creates a catalog product.
type CreateProductParams struct {
	Type        string             `json:"type"`
	Reference   string             `json:"reference,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description,omitempty"`
	About       string             `json:"about,omitempty"`
	TaxCode     string             `json:"tax_code,omitempty"`
	Category    *ProductCategory   `json:"category,omitempty"`
	Price       *ProductPrice      `json:"price,omitempty"`
	Shipment    *ProductShipment   `json:"shipment,omitempty"`
	Media       []ProductMediaItem `json:"media,omitempty"`
	Attributes  map[string]string  `json:"attributes,omitempty"`
	CustomData  map[string]string  `json:"custom_data,omitempty"`
}

// LookupProductParams looks up a product by ID.
type LookupProductParams struct {
	ProductID string `json:"product_id"`
}

// UpdateProductParams updates a product.
type UpdateProductParams struct {
	ProductID   string             `json:"product_id"`
	Reference   string             `json:"reference,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	About       string             `json:"about,omitempty"`
	TaxCode     string             `json:"tax_code,omitempty"`
	Category    *ProductCategory   `json:"category,omitempty"`
	Price       *ProductPrice      `json:"price,omitempty"`
	Shipment    *ProductShipment   `json:"shipment,omitempty"`
	Media       []ProductMediaItem `json:"media,omitempty"`
	Attributes  map[string]string  `json:"attributes,omitempty"`
	CustomData  map[string]string  `json:"custom_data,omitempty"`
}

// ProductActionParams performs an action on a product.
type ProductActionParams struct {
	ProductID string `json:"product_id"`
}

// AddProductPriceParams creates a new price for an existing product.
type AddProductPriceParams struct {
	ProductID    string             `json:"product_id"`
	Label        string             `json:"label,omitempty"`
	About        string             `json:"about,omitempty"`
	Amount       ProductPriceAmount `json:"amount"`
	SetAsDefault bool               `json:"set_as_default,omitempty"`
}

// SetDefaultUnitPriceParams changes the price used as a product's default.
type SetDefaultUnitPriceParams struct {
	ProductID string `json:"product_id"`
	PriceID   string `json:"price_id"`
}

// PageProductsParams pages through products.
type PageProductsParams struct {
	PageNumber int `json:"page_number,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}

// Product represents a catalog product.
type Product struct {
	ID               string                   `json:"id,omitempty"`
	ApplicationID    string                   `json:"application_id,omitempty"`
	Type             string                   `json:"type,omitempty"`
	Reference        string                   `json:"reference,omitempty"`
	Name             string                   `json:"name,omitempty"`
	Description      string                   `json:"description,omitempty"`
	About            string                   `json:"about,omitempty"`
	TaxCode          string                   `json:"tax_code,omitempty"`
	Category         *ProductCategory         `json:"category,omitempty"`
	Price            *ProductPrice            `json:"price,omitempty"`
	DefaultUnitPrice *ProductDefaultUnitPrice `json:"default_unit_price,omitempty"`
	Prices           []ProductPriceSummary    `json:"prices,omitempty"`
	Shipment         *ProductShipment         `json:"shipment,omitempty"`
	Media            []ProductMediaItem       `json:"media,omitempty"`
	Attributes       map[string]string        `json:"attributes,omitempty"`
	CustomData       map[string]string        `json:"custom_data,omitempty"`
	Active           bool                     `json:"active,omitempty"`
	Archived         bool                     `json:"archived,omitempty"`
	CreatedAt        string                   `json:"created_at,omitempty"`
	UpdatedAt        string                   `json:"updated_at,omitempty"`
	ArchivedAt       string                   `json:"archived_at,omitempty"`
}

// ProductsPage holds a page of products.
type ProductsPage struct {
	Number   int       `json:"number,omitempty"`
	Size     int       `json:"size,omitempty"`
	Products []Product `json:"products,omitempty"`
}

// CreatePriceParams creates a catalog price.
type CreatePriceParams struct {
	ProductID string `json:"product_id,omitempty"`
	Label     string `json:"label,omitempty"`
	About     string `json:"about,omitempty"`
	Currency  string `json:"currency"`
	Amount    int64  `json:"amount"`
}

// LookupPriceParams looks up a price by ID.
type LookupPriceParams struct {
	PriceID string `json:"price_id"`
}

// UpdatePriceParams updates price metadata.
type UpdatePriceParams struct {
	PriceID   string `json:"price_id"`
	ProductID string `json:"product_id,omitempty"`
	Label     string `json:"label,omitempty"`
	About     string `json:"about,omitempty"`
}

// PriceNominal represents a price amount.
type PriceNominal struct {
	Currency string `json:"currency"`
	Value    int64  `json:"value"`
	Sign     int    `json:"sign"`
}

// Price represents a catalog price.
type Price struct {
	ID         string        `json:"id,omitempty"`
	ProductID  string        `json:"product_id,omitempty"`
	Label      string        `json:"label,omitempty"`
	About      string        `json:"about,omitempty"`
	Nominal    *PriceNominal `json:"nominal,omitempty"`
	CreatedAt  string        `json:"created_at,omitempty"`
	UpdatedAt  string        `json:"updated_at,omitempty"`
	ArchivedAt string        `json:"archived_at,omitempty"`
}

// CreateCustomerParams creates a customer record.
type CreateCustomerParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`

	// Name is the customer's full name (required).
	Name string `json:"name"`

	// Title is an honorific or title (optional).
	Title string `json:"title,omitempty"`

	// Suffix is a name suffix like Jr. (optional).
	Suffix string `json:"suffix,omitempty"`

	// Reference is your internal customer identifier (optional).
	Reference string `json:"reference,omitempty"`

	// Email is the customer's email address (optional).
	Email string `json:"email_address,omitempty"`

	// PhoneNumber is the customer's phone number (optional).
	PhoneNumber string `json:"phone_number,omitempty"`

	// CustomData holds arbitrary key-value pairs (optional).
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// LookupCustomerParams looks up a customer by ID.
type LookupCustomerParams struct {
	CustomerID string `json:"customer_id"`
}

// PageCustomersParams pages through customers.
type PageCustomersParams struct {
	PageNumber int `json:"page_number,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}

// Customer represents a customer record.
type Customer struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name,omitempty"`
	Title       string            `json:"title,omitempty"`
	Suffix      string            `json:"suffix,omitempty"`
	Reference   string            `json:"reference,omitempty"`
	Email       string            `json:"email_address,omitempty"`
	PhoneNumber string            `json:"phone_number,omitempty"`
	CustomData  map[string]string `json:"custom_data,omitempty"`
	CreatedAt   string            `json:"created_at,omitempty"`
}

// CustomersPage holds a page of customers.
type CustomersPage struct {
	Number    int        `json:"number,omitempty"`
	Size      int        `json:"size,omitempty"`
	Customers []Customer `json:"customers,omitempty"`
}

// Address represents a postal address.
//
// Used for billing and shipping addresses on orders. All orders require
// a billing address. Shipping address is optional and only needed for
// physical goods delivery.
type Address struct {
	// Name is the recipient's name at this address (required).
	// For billing: typically matches customer name.
	// For shipping: may differ if shipping to someone else.
	Name string `json:"name"`

	// PhoneNumber is the contact number at this address (required).
	// Include country code. Example: "+233244123456"
	PhoneNumber string `json:"phone_number"`

	// Line1 is the first address line (required).
	// Street address, building number, apartment/unit.
	// Example: "123 Main Street, Apt 4B"
	Line1 string `json:"line1"`

	// Line2 is the second address line (optional).
	// Additional address details if needed.
	Line2 string `json:"line2,omitempty"`

	// Town is the city or town name (required).
	// Example: "Accra", "Nairobi", "Kampala"
	Town string `json:"town"`

	// Region is the state or province (optional).
	// Example: "Greater Accra", "Nairobi County"
	Region string `json:"region,omitempty"`

	// District is the district or locality (optional).
	// Finer geographic subdivision than region.
	District string `json:"district,omitempty"`

	// Country is the two-letter ISO 3166-1 alpha-2 country code (required).
	// Uppercase. Examples: "GH" (Ghana), "KE" (Kenya), "UG" (Uganda)
	Country string `json:"country"`

	// PostCode is the postal/ZIP code (optional).
	// Format varies by country. May be required in some jurisdictions.
	PostCode string `json:"post_code,omitempty"`
}

// BillingDetails captures billing contact information for an order.
//
// Required for all orders. This information appears on invoices and
// is used for payment authorization and dispute resolution.
//
// The billing phone number and email are used for payment notifications
// and confirmation codes.
type BillingDetails struct {
	// Name is the billing contact's full name (required).
	// Typically matches the customer name.
	Name string `json:"name"`

	// Email is the billing email address (required).
	// Used for invoice delivery and payment receipts.
	// Validated as RFC 5322 email format.
	Email string `json:"email_address"`

	// PhoneNumber is the billing contact phone with country code (required).
	// Used for payment confirmations and delivery of OTP codes.
	// Example: "+233244123456"
	PhoneNumber string `json:"phone_number"`

	// Address is the billing postal address (required).
	// Must be a complete, valid address.
	Address Address `json:"address"`
}

// Shipping holds shipping information for physical goods.
//
// Only required when selling physical products that need delivery.
// Digital goods and services don't require shipping information.
type Shipping struct {
	// Address is the delivery postal address (required for physical goods).
	// Must be a complete, deliverable address.
	Address Address `json:"address"`
}

// ProductLineItem represents a product in the order cart.
//
// Products are goods or services sold to customers. Each product has a unit
// price, quantity, and optional descriptive information.
//
// Example:
//
//	item := &commerce.ProductLineItem{
//	    Type:     "physical",
//	    Name:     "Wireless Headphones",
//	    About:    "Bluetooth 5.0, 30-hour battery",
//	    Quantity: 2,
//	    Price:    commerce.Money{Currency: "usd", Value: 7999}, // $79.99 each
//	    Reference: "SKU-12345",
//	}
type ProductLineItem struct {
	// ID is assigned by the API when the order is created (read-only).
	// Don't set this when creating orders.
	ID string `json:"id,omitempty"`

	// Type indicates whether the product is physical or digital (required).
	// Values: "physical" or "digital"
	// Physical products require shipping address.
	// Digital products can be delivered electronically.
	Type string `json:"type"`

	// Name is the product name (required).
	// Displayed on invoices and checkout pages.
	// Maximum 255 characters.
	Name string `json:"name"`

	// About is a short product description (optional).
	// Additional details about the product.
	// Maximum 1000 characters.
	About string `json:"about,omitempty"`

	// Quantity is the number of units (required).
	// Must be positive. Total line item amount = Price * Quantity.
	Quantity int64 `json:"quantity"`

	// Price is the unit price per item (required).
	// In minor units. The line item total is Price.Value * Quantity.
	Price Money `json:"price"`

	// Reference is your internal product identifier (optional).
	// Link to your inventory system's SKU or product ID.
	// Maximum 255 characters.
	Reference string `json:"reference,omitempty"`

	// TaxCode specifies the tax treatment (optional).
	// Used for tax calculation when tax integration is enabled.
	// Format depends on your tax provider.
	TaxCode string `json:"tax_code,omitempty"`

	// CustomData holds arbitrary key-value custom data (optional).
	// Both keys and values must be strings.
	// Maximum 25KB when serialized.
	// Learn more: https://studio.zebo.dev/custom-data
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// FeeLineItem represents an additional charge like service or processing fees.
//
// Fees are one-time charges added to the order subtotal. Unlike products,
// fees don't have quantities—they're always a fixed amount.
//
// Example:
//
//	fee := &commerce.FeeLineItem{
//	    Label:       "Service Fee",
//	    Description: "Platform usage fee",
//	    Amount:      commerce.Money{Currency: "usd", Value: 299}, // $2.99
//	}
type FeeLineItem struct {
	// ID is assigned by the API when the order is created (read-only).
	// Don't set this when creating orders.
	ID string `json:"id,omitempty"`

	// Label is the fee name (optional but recommended).
	// Displayed on invoices. Example: "Service Fee", "Processing Fee"
	// Maximum 255 characters.
	Label string `json:"label,omitempty"`

	// Description explains the fee (optional).
	// Additional context about why this fee is charged.
	// Maximum 1000 characters.
	Description string `json:"description,omitempty"`

	// TaxCode specifies the tax treatment (optional).
	// Used for tax calculation when tax integration is enabled.
	TaxCode string `json:"tax_code,omitempty"`

	// CustomData holds arbitrary key-value custom data (optional).
	// Both keys and values must be strings.
	// Maximum 25KB when serialized.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// Amount is the total fee charge (required).
	// In minor units. Not multiplied by any quantity.
	Amount Money `json:"amount"`
}

// ShippingLineItem represents the delivery charge for physical goods.
//
// Only needed when selling physical products. Automatically omitted for
// orders containing only digital products.
//
// Example:
//
//	shipping := &commerce.ShippingLineItem{
//	    Fee: commerce.Money{Currency: "usd", Value: 500}, // $5.00
//	}
type ShippingLineItem struct {
	// ID is assigned by the API when the order is created (read-only).
	// Don't set this when creating orders.
	ID string `json:"id,omitempty"`

	// Fee is the total shipping charge (required).
	// In minor units. Not multiplied by any quantity.
	Fee Money `json:"fee"`

	// TaxCode specifies the tax treatment (optional).
	// Used for tax calculation when tax integration is enabled.
	TaxCode string `json:"tax_code,omitempty"`

	// CustomData holds arbitrary key-value custom data (optional).
	// Both keys and values must be strings.
	// Maximum 25KB when serialized.
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// LineItemType enumerates the three types of line items.
type LineItemType string

const (
	// LineItemTypeProduct represents a product (good or service).
	LineItemTypeProduct LineItemType = "product"

	// LineItemTypeFee represents an additional charge.
	LineItemTypeFee LineItemType = "fee"

	// LineItemTypeShipping represents a delivery charge.
	LineItemTypeShipping LineItemType = "shipping"
)

// OrderLineItem is a discriminated union of line item types.
//
// Each order line item is one of three types: product, fee, or shipping.
// Set Type and the corresponding field (Product, Fee, or Shipping).
// Leave the other fields nil.
//
// Example (product):
//
//	lineItem := commerce.OrderLineItem{
//	    Type: commerce.LineItemTypeProduct,
//	    Product: &commerce.ProductLineItem{
//	        Type:     "digital",
//	        Name:     "Premium Subscription",
//	        Quantity: 1,
//	        Price:    commerce.Money{Currency: "usd", Value: 999},
//	    },
//	}
//
// Example (fee):
//
//	lineItem := commerce.OrderLineItem{
//	    Type: commerce.LineItemTypeFee,
//	    Fee: &commerce.FeeLineItem{
//	        Label:  "Platform Fee",
//	        Amount: commerce.Money{Currency: "usd", Value: 299},
//	    },
//	}
type OrderLineItem struct {
	// Type specifies which variant is active (required).
	Type LineItemType `json:"type"`

	// Product is populated when Type is LineItemTypeProduct.
	// Nil for other types.
	Product *ProductLineItem `json:"product,omitempty"`

	// Fee is populated when Type is LineItemTypeFee.
	// Nil for other types.
	Fee *FeeLineItem `json:"fee,omitempty"`

	// Shipping is populated when Type is LineItemTypeShipping.
	// Nil for other types.
	Shipping *ShippingLineItem `json:"shipping,omitempty"`
}

// CheckoutSettings configures checkout page behavior and redirect URLs.
//
// When you finalize an order (finalize: true), Commerce
// generates a hosted checkout page. These settings control where customers
// are redirected after completing or canceling payment.
type CheckoutSettings struct {
	// RedirectURL is where customers go after successful payment (optional).
	// Must be HTTPS in production. Can be HTTP for testing.
	// If omitted, customers see a generic success page.
	// Example: "https://example.com/orders/thank-you"
	RedirectURL string `json:"redirect_url,omitempty"`

	// CancelURL is where customers go if they cancel payment (optional).
	// Must be HTTPS in production. Can be HTTP for testing.
	// If omitted, customers see a generic cancellation page.
	// Example: "https://example.com/cart"
	CancelURL string `json:"cancel_url,omitempty"`
}

// OrderPayoutSettings overrides payout configuration for a single order.
type OrderPayoutSettings struct {
	// Destination specifies where payout funds should be sent (optional).
	Destination *OrderPayoutDestination `json:"destination,omitempty"`

	// EnableFX controls foreign exchange conversion for the payout (optional).
	EnableFX *bool `json:"enable_fx,omitempty"`
}

// OrderPayoutDestination sets the payout destination for the order.
type OrderPayoutDestination struct {
	// FinancialAccountID references an existing financial account (optional).
	FinancialAccountID string `json:"financial_account_id,omitempty"`

	// FinancialAccountData provides inline financial account details (optional).
	FinancialAccountData *OrderPayoutFinancialAccount `json:"financial_account_data,omitempty"`
}

// OrderPayoutFinancialAccount defines inline payout destination details.
type OrderPayoutFinancialAccount struct {
	// Type specifies the account type (wallet, bank_account, dosh_account).
	Type FinancialAccountType `json:"type"`

	// Wallet contains mobile money details when Type is "wallet".
	Wallet *WalletConfig `json:"wallet,omitempty"`

	// BankAccount contains bank account details when Type is "bank_account".
	BankAccount *BankAccountConfig `json:"bank_account,omitempty"`

	// DoshAccount contains Dosh wallet details when Type is "dosh_account".
	DoshAccount map[string]any `json:"dosh_account,omitempty"`
}

// OrderCreateParams contains all parameters for creating an order.
//
// Orders represent a purchase transaction with line items, customer info,
// and payment details. The order creation flow is flexible:
//
// 1. Create draft order (no payment method) → finalize → customer pays from hosted page
// 2. Create order with payment method → execute payment immediately
// 3. Create order with checkout settings → finalize → redirect customer to hosted page
//
// Customer specification (exactly one required):
//   - CustomerData: for new customers (API creates customer record)
//   - CustomerID: for existing customers
//
// Payment method specification (optional):
//   - PaymentMethodID: reference saved payment method
//   - PaymentMethodData: inline payment method for one-time use
//
// Example (new customer, immediate payment):
//
//	params := commerce.OrderCreateParams{
//	    CustomerData: &commerce.CustomerData{
//	        Name:        "Jane Doe",
//	        Email:       "jane@example.com",
//	        PhoneNumber: "+233244123456",
//	    },
//	    PaymentMethodData: &commerce.PaymentMethodData{
//	        Type: commerce.PaymentMethodTypeMobileMoney,
//	        MobileMoney: &commerce.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    LineItems: []commerce.OrderLineItem{
//	        {
//	            Type: commerce.LineItemTypeProduct,
//	            Product: &commerce.ProductLineItem{
//	                Type:     "digital",
//	                Name:     "Premium Plan",
//	                Quantity: 1,
//	                Price:    commerce.Money{Currency: "ghs", Value: 10000},
//	            },
//	        },
//	    },
//	    BillingDetails: commerce.BillingDetails{
//	        Name:        "Jane Doe",
//	        Email:       "jane@example.com",
//	        PhoneNumber: "+233244123456",
//	        Address: commerce.Address{
//	            Name:        "Jane Doe",
//	            PhoneNumber: "+233244123456",
//	            Line1:       "123 Main St",
//	            Town:        "Accra",
//	            Country:     "GH",
//	        },
//	    },
//	    ExecutePayment: commerce.Bool(true),
//	    RequestMeta: &commerce.RequestMeta{IdempotencyKey: "order_20231215_jane_001"},
//	}
type OrderCreateParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	// Prefer RequestMeta.IdempotencyKey over the legacy top-level IdempotencyKey.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`

	// CustomerData provides inline customer information for new customers (optional).
	// Use this when the customer doesn't have a Commerce customer ID yet.
	// The API creates a customer record and returns it in the response.
	// Mutually exclusive with CustomerID (provide exactly one).
	CustomerData *CustomerData `json:"customer_data,omitempty"`

	// CustomerID references an existing customer by ID (optional).
	// Use this for repeat customers who already have a Commerce customer record.
	// Mutually exclusive with CustomerData (provide exactly one).
	CustomerID string `json:"customer_id,omitempty"`

	// PaymentMethodID references a saved payment method by ID (optional).
	// Used for charging repeat customers with tokenized payment methods.
	// Mutually exclusive with PaymentMethodData.
	// If provided, you can set ExecutePayment: true to charge immediately.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// PaymentMethodData provides inline payment details for one-time use (optional).
	// Use this when you don't want to save the payment method.
	// Mutually exclusive with PaymentMethodID.
	// If provided, you can set ExecutePayment: true to charge immediately.
	PaymentMethodData *PaymentMethodData `json:"payment_method_data,omitempty"`

	// StatementDescriptor appears on the customer's payment statement (optional).
	// Maximum 22 characters. Only letters, numbers, and spaces allowed.
	// Example: "ACME INC SUBSCRIPTION"
	// If omitted, uses your business name from settings.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// StatementDescriptorPrefix builds a descriptor from a 2-10 character prefix and generated order ID.
	// The API formats it as prefix*order_id and truncates the order ID to fit.
	// Mutually exclusive with StatementDescriptor.
	StatementDescriptorPrefix string `json:"statement_descriptor_prefix,omitempty"`

	// ExecutePayment triggers immediate payment attempt (optional, default: false).
	// Only valid when PaymentMethodID or PaymentMethodData is provided.
	// If false (default), order is created but payment must be initiated separately.
	// If true, attempts to charge the payment method immediately.
	ExecutePayment *bool `json:"execute_payment,omitempty"`

	// Finalize seals the order and generates checkout page (optional, default: false).
	// If true, order is finalized and checkout URL is returned in response.
	// Finalized orders cannot be modified—line items and amounts are locked.
	// Required for hosted checkout flow.
	Finalize *bool `json:"finalize,omitempty"`

	// IdempotencyKey prevents duplicate order creation (optional but recommended).
	// Deprecated: use RequestMeta.IdempotencyKey instead.
	// If the same key is used twice, the second request returns the original order
	// instead of creating a duplicate. Keys can be reused if the original request failed.
	// Maximum 255 characters. Use a unique value per order attempt.
	// Example: "order_20231215_customer_123_attempt_1"
	// Learn more: https://studio.zebo.dev/idempotency
	IdempotencyKey string `json:"idempotency_key,omitempty"`

	// CheckoutSettings configures hosted checkout page redirects (optional).
	// Only relevant when Finalize is true.
	// Specifies where customers are redirected after payment or cancellation.
	CheckoutSettings *CheckoutSettings `json:"checkout_settings,omitempty"`

	// PayoutSettings overrides payout configuration for this order (optional).
	PayoutSettings *OrderPayoutSettings `json:"payout_settings,omitempty"`

	// Number is a custom order number for your records (optional).
	// If omitted, Commerce generates a unique order number automatically.
	// Maximum 255 characters. Must be unique across your orders.
	// Example: "ORD-2023-00123"
	Number string `json:"number,omitempty"`

	// LineItems is the list of products, fees, and shipping charges (required).
	// Must have at least one line item. Orders without line items are invalid.
	// The order total is calculated from all line items.
	LineItems []OrderLineItem `json:"line_items"`

	// CustomData holds arbitrary key-value metadata for the order (optional).
	// Both keys and values must be strings.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// BillingDetails captures billing contact and address (required).
	// Used for invoicing and payment authorization.
	// The email and phone number receive payment notifications.
	BillingDetails BillingDetails `json:"billing_details"`

	// Shipping provides delivery address for physical goods (optional).
	// Required only when line items include physical products.
	// Omit for digital-only orders.
	Shipping *Shipping `json:"shipping,omitempty"`
}

// OrderLookupParams specifies which order to retrieve.
type OrderLookupParams struct {
	// OrderID is the unique order identifier (required).
	// Starts with "or_". Example: "or_abc123def456"
	OrderID string `json:"order_id"`
}

// OrderPayParams initiates payment for an existing order.
//
// Use this to charge an order after creation, or to retry payment
// on an order with a failed payment attempt.
//
// Provide either PaymentMethodID (for saved payment methods) or
// PaymentMethodData (for one-time payment methods).
//
// Example (with saved payment method):
//
//	params := commerce.OrderPayParams{
//	    OrderID:         order.ID,
//	    PaymentMethodID: "pm_abc123",
//	}
//	response, err := client.Orders.Pay(ctx, params)
type OrderPayParams struct {
	// OrderID is the order to charge (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`

	// PaymentMethodID references a saved payment method (optional).
	// Mutually exclusive with PaymentMethodData.
	// Use this for repeat customers with tokenized payment methods.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// PaymentMethodData provides inline payment details (optional).
	// Mutually exclusive with PaymentMethodID.
	// Use this for one-time payments without saving the method.
	PaymentMethodData *PaymentMethodData `json:"payment_method_data,omitempty"`

	// PaidOutOfBand marks payment as completed outside Commerce (optional).
	// Set to true if customer paid via cash, bank transfer, or other method.
	// The order is marked paid without actually charging the payment method.
	// Use carefully—this bypasses payment processing entirely.
	PaidOutOfBand *bool `json:"paid_out_of_band,omitempty"`
}

// OrderConfirmParams confirms a payment using a verification token.
//
// After initiating payment on an order with confirms_use enabled,
// the customer receives an OTP. Use this to submit their OTP and
// complete the payment.
//
// Example:
//
//	params := commerce.OrderConfirmParams{
//	    OrderID: order.ID,
//	    Token:   "123456", // OTP from customer
//	}
//	order, err := client.Orders.ConfirmPayment(ctx, params)
type OrderConfirmParams struct {
	// OrderID is the order awaiting confirmation (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`

	// Token is the OTP or verification code (required).
	// Typically 4-6 digits sent to the customer's phone via SMS.
	Token string `json:"token"`
}

// OrderRequestConfirmationParams requests a new OTP for payment confirmation.
//
// Use this if the customer didn't receive the original OTP or if it expired.
// Triggers a new OTP to be sent to the customer's phone number.
type OrderRequestConfirmationParams struct {
	// OrderID is the order needing a new confirmation token (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`
}

// OrderFinalizeParams seals an order and generates checkout assets.
//
// Finalizing an order:
// - Locks the order (no more modifications)
// - Calculates final totals
// - Generates hosted checkout page
// - Creates invoice documents
//
// After finalization, the order cannot be edited. Use this when you're
// ready for the customer to pay.
type OrderFinalizeParams struct {
	// OrderID is the order to finalize (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`
}

// OrderSendInvoiceParams sends an invoice link for an order.
type OrderSendInvoiceParams struct {
	// OrderID is the order whose invoice should be sent (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`
}

// OrderSendReceiptParams sends a receipt link for a paid order.
type OrderSendReceiptParams struct {
	// OrderID is the paid order whose receipt should be sent (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`
}

// OrderCompleteParams marks an order as completed.
//
// Completing an order:
// - Marks order as fulfilled
// - Triggers payout eligibility (after aging period)
// - Creates balance transaction
//
// Only valid for paid orders. Use PaidOutOfBand if payment happened
// outside the Commerce platform.
type OrderCompleteParams struct {
	// OrderID is the order to complete (required).
	OrderID string `json:"order_id"`

	// PaidOutOfBand indicates payment happened externally (optional).
	// If true, marks order as paid without charging the payment method.
	// Use for cash, bank transfer, or other offline payment methods.
	PaidOutOfBand *bool `json:"paid_out_of_band,omitempty"`
}

// OrderCancelParams cancels an order.
//
// Canceling an order:
// - Prevents any future payment attempts
// - Refunds any captured payment (if applicable)
// - Marks order as permanently closed
//
// Cannot cancel orders that are already completed or refunded.
type OrderCancelParams struct {
	// OrderID is the order to cancel (required).
	OrderID string `json:"order_id"`

	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`
}

// OrderRefundParams issues a refund for a paid order.
//
// Refunding an order:
// - Returns the full order amount to the customer
// - Reverses the balance transaction
// - Marks order as refunded
//
// Only valid for paid, completed orders. Refunds are final and cannot
// be reversed. Partial refunds are not currently supported—refunds are
// always for the full order amount.
type OrderRefundParams struct {
	// OrderID is the order to refund (required).
	OrderID string `json:"order_id"`
}

// OrderPageParams specifies pagination for listing orders.
//
// Use this to fetch recent orders with pagination support.
//
// Example:
//
//	params := commerce.OrderPageParams{
//	    PageNumber: 1,
//	    PageSize:   50,
//	}
//	orders, err := client.Orders.Page(ctx, params)
type OrderPageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed. First page is 1, not 0.
	PageNumber int `json:"page_number,omitempty"`

	// PageSize is the number of orders per page (optional, default: 20).
	// Maximum 100. Minimum 1.
	PageSize int `json:"page_size,omitempty"`
}

// OrderDocumentDeliveryResponse contains an order and document delivery result.
type OrderDocumentDeliveryResponse struct {
	Order    Order                 `json:"order"`
	Delivery OrderDocumentDelivery `json:"delivery"`
}

// OrderDocumentDelivery describes a delivered invoice or receipt link.
type OrderDocumentDelivery struct {
	DocumentKind   string                         `json:"document_kind"`
	DocumentURL    string                         `json:"document_url"`
	SentChannels   []string                       `json:"sent_channels,omitempty"`
	FailedChannels []string                       `json:"failed_channels,omitempty"`
	Deliveries     []OrderDocumentDeliveryAttempt `json:"deliveries,omitempty"`
	Failures       []OrderDocumentDeliveryAttempt `json:"failures,omitempty"`
}

// OrderDocumentDeliveryAttempt describes one delivery channel result.
type OrderDocumentDeliveryAttempt struct {
	Channel string `json:"channel"`
	ChimeID string `json:"chime_id,omitempty"`
	Error   string `json:"error,omitempty"`
}

// PaymentMethodObject represents a tokenized payment method.
//
// Payment methods are saved customer payment instruments that can be
// charged repeatedly without re-entering payment details.
//
// Payment methods must be verified before use (except when confirms_use
// is false). Verification sends an OTP to confirm the customer owns the
// payment instrument.
type PaymentMethodObject struct {
	// ID is the unique payment method identifier.
	// Starts with "pm_". Example: "pm_abc123def456"
	ID string `json:"id"`

	// CustomerID is the owning customer's ID.
	CustomerID string `json:"customer_id"`

	// Type is the payment method category.
	// Values: "mobile_money", "bank_account", "card", "motito"
	Type PaymentMethodType `json:"type"`

	// MobileMoney contains wallet details when Type is "mobile_money".
	MobileMoney *struct {
		AccountNumber string `json:"account_number,omitempty"`
		Network       string `json:"network,omitempty"`
	} `json:"mobile_money,omitempty"`

	// BankAccount contains bank details when Type is "bank_account".
	BankAccount *struct {
		GhanaBankAccount *struct {
			Branch        string `json:"branch,omitempty"`
			Name          string `json:"name,omitempty"`
			AccountNumber string `json:"account_number,omitempty"`
			SortCode      string `json:"sort_code,omitempty"`
			SwiftCode     string `json:"swift_code,omitempty"`
		} `json:"ghana_bank_account,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"bank_account,omitempty"`

	// Card contains card details when Type is "card".
	Card *struct {
		Brand     string  `json:"brand,omitempty"`
		ExpiresOn *string `json:"expires_on,omitempty"`
		Issuer    *struct {
			EmailAddress string `json:"email_address,omitempty"`
			Name         string `json:"name,omitempty"`
			PhoneNumber  string `json:"phone_number,omitempty"`
			Type         string `json:"type,omitempty"`
		} `json:"issuer,omitempty"`
		Owner *struct {
			EmailAddress string `json:"email_address,omitempty"`
			Name         string `json:"name,omitempty"`
			PhoneNumber  string `json:"phone_number,omitempty"`
		} `json:"owner,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"card,omitempty"`

	// Verification contains verification metadata when available.
	Verification *struct {
		CompletedAt *string `json:"completed_at,omitempty"`
		InitiatedAt *string `json:"initiated_at,omitempty"`
		Mechanism   string  `json:"mechanism,omitempty"`
		RequestID   string  `json:"request_id,omitempty"`
		Type        string  `json:"type,omitempty"`
	} `json:"verification,omitempty"`

	// CustomData holds arbitrary metadata attached to the payment method.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// ExpiresOn is when this payment method expires (ISO 8601), if applicable.
	ExpiresOn *string `json:"expires_on,omitempty"`

	// CreatedAt is the tokenization timestamp (ISO 8601).
	CreatedAt string `json:"created_at"`

	// Verified indicates whether the payment method has been verified.
	// Unverified payment methods cannot be charged (unless confirms_use is false).
	Verified bool `json:"verified"`

	// VerifiedAt is the verification completion timestamp (ISO 8601).
	// Nil if not yet verified.
	VerifiedAt *string `json:"verified_at,omitempty"`
}

// PaymentAttempt captures details of a single payment attempt.
//
// Each payment may have multiple attempts if initial attempts fail.
// This tracks the most recent attempt's status and timing.
type PaymentAttempt struct {
	// PaymentMethodType is the type of payment method used.
	PaymentMethodType string `json:"payment_method_type,omitempty"`

	// PaymentMethodID is the ID of the payment method charged.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// Reference is the external transaction reference from the payment provider.
	// Use this when contacting support or investigating payment issues.
	Reference string `json:"reference,omitempty"`

	// Status is the attempt's current state.
	// Values: "initiated", "succeeded", "failed"
	Status string `json:"status,omitempty"`

	// InitiatedAt is when the attempt started (ISO 8601).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// SucceededAt is when the attempt succeeded (ISO 8601).
	// Nil if not yet succeeded.
	SucceededAt *string `json:"succeeded_at,omitempty"`

	// FailedAt is when the attempt failed (ISO 8601).
	// Nil if not yet failed.
	FailedAt *string `json:"failed_at,omitempty"`
}

// PaymentNextAction describes the next step required to complete payment.
//
// When a payment requires customer action (OTP confirmation, redirect to
// bank page, etc), this field describes what needs to happen.
//
// Check Type to determine the required action:
//   - "confirm_payment": Customer must provide OTP
//   - "redirect": Customer must visit a URL
//   - "execute": Internal processing (no action needed)
type PaymentNextAction struct {
	// Type specifies the action category.
	// Values: "confirm_payment", "redirect", "execute"
	Type string `json:"type"`

	// ConfirmPayment contains OTP confirmation details.
	// Only present when Type is "confirm_payment".
	ConfirmPayment *struct {
		// ExpiresAt is when the OTP expires (ISO 8601).
		// Customer must confirm before this time.
		ExpiresAt string `json:"expires_at"`

		// Scheme describes the confirmation method.
		// Example: "otp"
		Scheme string `json:"scheme,omitempty"`

		// Request contains OTP delivery details.
		Request *struct {
			// ID is the OTP request identifier.
			ID string `json:"id"`

			// Recipient is where the OTP was sent.
			// For SMS: the phone number.
			Recipient string `json:"recipient"`

			// SentVia is the delivery channel.
			// Values: "sms", "email"
			SentVia string `json:"sent_via"`

			// TokenSize is the number of OTP digits.
			// Typically 4 or 6.
			TokenSize int `json:"token_size"`

			// SenderID is the SMS sender name.
			SenderID string `json:"sender_id"`
		} `json:"request,omitempty"`
	} `json:"confirm_payment,omitempty"`

	// Execute is present when Type is "execute" (internal processing).
	Execute any `json:"execute,omitempty"`

	// Redirect contains redirect details when Type is "redirect".
	Redirect *struct {
		// URL is where the customer should be redirected.
		// Open this URL in a browser for the customer to complete payment.
		URL string `json:"url"`
	} `json:"redirect,omitempty"`
}

// Payment represents payment details and status for an order.
//
// Every paid order has an associated payment object tracking the charge
// lifecycle, attempts, and any required customer actions.
type Payment struct {
	// ID is the unique payment identifier.
	// Starts with "py_". Example: "py_abc123def456"
	ID string `json:"id,omitempty"`

	// Status is the payment's current state.
	// Values: "initiated", "requires_action", "processing", "paid", "failed"
	Status string `json:"status,omitempty"`

	// StatementDescriptor is what appears on the customer's statement.
	// Maximum 22 characters.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// Amount is the charged amount.
	Amount *Money `json:"amount,omitempty"`

	// PaymentMethod is the charged payment method details.
	PaymentMethod *PaymentMethodObject `json:"payment_method,omitempty"`

	// LatestAttempt is the most recent payment attempt.
	// Nil if no attempts yet.
	LatestAttempt *PaymentAttempt `json:"latest_attempt,omitempty"`

	// NextAction describes any required customer action.
	// Nil if no action required (payment is processing or complete).
	NextAction *PaymentNextAction `json:"next_action,omitempty"`

	// BalanceTransaction is the resulting balance entry when payment succeeds.
	// Used for tracking payouts and available balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction,omitempty"`

	// PayoutConfiguration is the payout setup used for this payment (if applicable).
	PayoutConfiguration *PayoutConfiguration `json:"payout_configuration,omitempty"`

	// InitiatedAt is when payment was first attempted (ISO 8601).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// ExecutedAt is when payment was submitted to the network (ISO 8601).
	// Nil if not yet executed.
	ExecutedAt *string `json:"executed_at,omitempty"`

	// PaidAt is when payment was confirmed successful (ISO 8601).
	// Nil if not yet paid.
	PaidAt *string `json:"paid_at,omitempty"`

	// FailedAt is when payment was marked failed (ISO 8601).
	// Nil if not failed.
	FailedAt *string `json:"failed_at,omitempty"`
}

// InvoiceFormat contains URLs for invoice documents in different formats.
type InvoiceFormat struct {
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
	Format *InvoiceFormat `json:"format,omitempty"`

	// Deliveries tracks invoice email/SMS delivery attempts.
	Deliveries any `json:"deliveries,omitempty"`
}

// LineItemGroup contains line items grouped by type with totals.
//
// The API returns this grouped structure in order responses to make
// it easier to display cart breakdowns.
type LineItemGroup struct {
	// LineItems is the list of cart items.
	LineItems []OrderLineItem `json:"line_items"`

	// Total is the sum of all line items.
	// This is the order amount before any fees or discounts.
	Total Money `json:"total"`
}

// Order represents a complete order object.
//
// Orders are the central resource in Commerce, representing a purchase
// transaction from cart to fulfillment. Orders go through several states:
//
// 1. draft: Created but not finalized
// 2. sealed: Finalized and ready for payment
// 3. paid: Payment succeeded
// 4. completed: Fulfilled and settled
// 5. cancelled: Permanently canceled
// 6. refunded: Payment returned to customer
type Order struct {
	// ID is the unique order identifier (read-only).
	// Starts with "or_". Example: "or_abc123def456"
	ID string `json:"id"`

	// Status is the order's current state (read-only).
	// Values: "draft", "sealed", "paid", "completed", "cancelled", "refunded"
	Status string `json:"status"`

	// Number is the human-readable order number.
	// Auto-generated if not provided during creation.
	// Example: "ORD-2023-00123"
	Number string `json:"number,omitempty"`

	// CustomerID is the owning customer's ID.
	CustomerID string `json:"customer_id,omitempty"`

	// Customer contains full customer details.
	// Only populated in responses when customer exists.
	Customer *CustomerData `json:"customer,omitempty"`

	// BillingDetails holds billing contact and address.
	BillingDetails *BillingDetails `json:"billing_details,omitempty"`

	// Shipping holds delivery address for physical goods.
	// Nil for digital-only orders.
	Shipping *Shipping `json:"shipping,omitempty"`

	// LineItems is the list of products, fees, and shipping charges.
	LineItems []OrderLineItem `json:"line_items,omitempty"`

	// LineItemGroup contains line items grouped with totals.
	// Useful for displaying cart breakdowns.
	LineItemGroup *LineItemGroup `json:"line_item_group,omitempty"`

	// Payment contains payment details and status.
	// Nil if order hasn't been charged yet.
	Payment *Payment `json:"payment,omitempty"`

	// PaymentStatus is a summary of payment state (read-only).
	// Values: "unpaid", "requires_action", "processing", "paid", "failed"
	PaymentStatus string `json:"payment_status,omitempty"`

	// PaymentMethodID is the attached payment method's ID.
	// May be set but not yet charged.
	PaymentMethodID string `json:"payment_method_id,omitempty"`

	// StatementDescriptor appears on payment statements.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`

	// CheckoutSettings contains hosted checkout configuration.
	CheckoutSettings *CheckoutSettings `json:"checkout_settings,omitempty"`

	// InitiatedAt is when the order was created (ISO 8601, read-only).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// SealedAt is when the order was finalized (ISO 8601, read-only).
	// Nil if not yet finalized.
	SealedAt *string `json:"sealed_at,omitempty"`

	// CompletedAt is when the order was marked complete (ISO 8601, read-only).
	// Nil if not yet completed.
	CompletedAt *string `json:"completed_at,omitempty"`

	// ExpiresAt is when the order expires if unpaid (ISO 8601, read-only).
	// Nil for completed orders.
	ExpiresAt *string `json:"expires_at,omitempty"`

	// CreatedAt is the creation timestamp (ISO 8601, read-only).
	CreatedAt *string `json:"created_at,omitempty"`

	// UpdatedAt is the last modification timestamp (ISO 8601, read-only).
	UpdatedAt *string `json:"updated_at,omitempty"`

	// PaidAt is when payment was confirmed (ISO 8601, read-only).
	// Nil if not yet paid.
	PaidAt *string `json:"paid_at,omitempty"`

	// CancelledAt is when the order was canceled (ISO 8601, read-only).
	// Nil if not canceled.
	CancelledAt *string `json:"cancelled_at,omitempty"`

	// Invoice contains invoice document links and delivery status.
	// Nil if order not finalized.
	Invoice *Invoice `json:"invoice,omitempty"`
}

// PaymentResponse is returned from the Orders.Pay method.
//
// Provides a summary of the payment initiation result, including whether
// customer action (like OTP confirmation) is required.
type PaymentResponse struct {
	// PaymentID is the created payment's ID.
	PaymentID string `json:"payment_id,omitempty"`

	// OrderID is the charged order's ID.
	OrderID string `json:"order_id,omitempty"`

	// Status is the payment's current state.
	// Values: "initiated", "requires_action", "processing", "paid", "failed"
	Status string `json:"status,omitempty"`

	// RequiresConfirmation indicates whether customer must provide OTP.
	RequiresConfirmation bool `json:"requires_confirmation,omitempty"`

	// ConfirmationSent indicates whether OTP was sent to customer.
	// True when confirms_use is enabled and OTP was delivered.
	ConfirmationSent bool `json:"confirmation_sent,omitempty"`
}

// ChimeRecipientType specifies how to address a chime recipient.
type ChimeRecipientType string

const (
	// ChimeRecipientTypePhone indicates the recipient is identified by phone number.
	// Used for SMS delivery.
	ChimeRecipientTypePhone ChimeRecipientType = "phone"

	// ChimeRecipientTypeEmail indicates the recipient is identified by email address.
	// Used for email delivery.
	ChimeRecipientTypeEmail ChimeRecipientType = "email"
)

// ChimeTransport specifies the delivery channel for a chime.
type ChimeTransport string

const (
	// ChimeTransportSMS sends the chime via SMS text message.
	ChimeTransportSMS ChimeTransport = "sms"

	// ChimeTransportEmail sends the chime via email.
	ChimeTransportEmail ChimeTransport = "email"
)

// ChimeRecipient specifies who should receive a notification chime.
//
// Set Type and the corresponding field (Phone or Email).
//
// Example (SMS):
//
//	recipient := commerce.ChimeRecipient{
//	    Type: commerce.ChimeRecipientTypePhone,
//	    Name: "Jane Doe",
//	    Phone: &struct{Number string `json:"number"`}{Number: "+233244123456"},
//	}
type ChimeRecipient struct {
	// Type specifies how the recipient is identified (required).
	// Values: "phone" or "email"
	Type ChimeRecipientType `json:"type"`

	// Name is the recipient's display name (optional).
	// Used in email subject lines and SMS sender names.
	Name string `json:"name,omitempty"`

	// Phone contains the phone number (required when Type is "phone").
	// Must include country code. Example: "+233244123456"
	Phone *struct {
		Number string `json:"number"`
	} `json:"phone,omitempty"`

	// Email contains the email address (required when Type is "email").
	// Validated as RFC 5322 email format.
	Email *struct {
		Address string `json:"address"`
	} `json:"email,omitempty"`
}

// SendChimeParams sends a notification immediately.
//
// Chimes are transactional notifications sent to customers via SMS or email.
// Unlike invoice emails, chimes give you full control over message content.
//
// Example (SMS notification):
//
//	params := commerce.SendChimeParams{
//	    Recipient: commerce.ChimeRecipient{
//	        Type: commerce.ChimeRecipientTypePhone,
//	        Phone: &struct{Number string}{Number: "+233244123456"},
//	    },
//	    FullMessage: "Your order #12345 has shipped!",
//	    Transport:   commerce.ChimeTransportSMS,
//	    Purpose:     "order_shipped",
//	    IdempotencyKey: "chime_order_12345_shipped",
//	}
type SendChimeParams struct {
	// Recipient specifies who receives the chime (required).
	Recipient ChimeRecipient `json:"recipient"`

	// FullMessage is the complete message content (required).
	// For SMS: maximum 160 characters for single message, 1530 for concatenated.
	// For email: becomes the email body (plain text).
	FullMessage string `json:"full_message"`

	// Transport specifies delivery channel (optional, default: inferred from recipient type).
	// Values: "sms" or "email"
	// If omitted, uses SMS for phone recipients, email for email recipients.
	Transport ChimeTransport `json:"transport,omitempty"`

	// Sender is the sender ID shown to recipient (optional).
	// For SMS: sender name (max 11 characters). Example: "AcmeStore"
	// For email: sender name in "From" header. Example: "Acme Support"
	// If omitted, uses your business name from settings.
	Sender string `json:"sender,omitempty"`

	// Purpose categorizes the chime for analytics (optional).
	// Examples: "order_confirmation", "payment_reminder", "otp"
	// Use consistent values for reporting and filtering.
	Purpose string `json:"purpose,omitempty"`

	// CustomData holds arbitrary key-value custom data (optional).
	// Both keys and values must be strings. Maximum 25KB when serialized.
	// Useful for linking chimes to your internal records.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// IdempotencyKey prevents duplicate sends (optional but strongly recommended).
	// If the same key is used twice, the second request returns the original chime
	// instead of sending a duplicate. Keys can be reused if the original failed.
	// Example: "chime_order_12345_shipped"
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

// ScheduleChimeParams schedules a notification for future delivery.
//
// Use this to queue notifications to one or more recipient addresses.
// Useful for reminders, follow-ups, or coordinated messaging campaigns.
//
// Example:
//
//	params := commerce.ScheduleChimeParams{
//	    Recipients:  []string{"+233244123456", "user@example.com"},
//	    FullMessage: "Your subscription renews tomorrow.",
//	    SendAfter:   "2024-01-15T09:00:00Z",
//	    SenderID:    "YourBrand",
//	}
type ScheduleChimeParams struct {
	// Recipients specifies all recipient addresses (required).
	Recipients []string `json:"recipients,omitempty"`

	// FullMessage is the complete message content (required).
	// For SMS: maximum 160 characters for single message.
	// For email: becomes the email body.
	FullMessage string `json:"full_message"`

	// SendAfter is when to send the chime (required).
	// ISO 8601 timestamp. Must be in the future.
	// Example: "2024-01-15T09:00:00Z"
	SendAfter string `json:"send_after,omitempty"`

	// SenderID is the sender identifier displayed to recipients (optional).
	SenderID string `json:"sender_id,omitempty"`

	// Purpose categorizes the chime for analytics (optional).
	Purpose string `json:"purpose,omitempty"`

	// IdempotencyKey prevents duplicate scheduling (optional but recommended).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

// ScheduleDetail describes a scheduled chime and its execution state.
type ScheduleDetail struct {
	ID         string          `json:"id,omitempty"`
	Recipients []string        `json:"recipients,omitempty"`
	Content    string          `json:"content,omitempty"`
	SenderID   string          `json:"sender_id,omitempty"`
	Purpose    *string         `json:"purpose,omitempty"`
	SendAfter  string          `json:"send_after,omitempty"`
	CreatedAt  string          `json:"created_at,omitempty"`
	ExecutedAt *string         `json:"executed_at,omitempty"`
	CanceledAt *string         `json:"canceled_at,omitempty"`
	Errors     []ScheduleError `json:"errors,omitempty"`
	ChimeIDs   []string        `json:"chime_ids,omitempty"`
}

// ScheduleError reports a per-recipient scheduling error.
type ScheduleError struct {
	Recipient string `json:"recipient,omitempty"`
	FixCode   string `json:"fix_code,omitempty"`
	Type      string `json:"type,omitempty"`
}

// ScheduledChime is returned after scheduling a chime.
type ScheduledChime struct {
	ID          string   `json:"id,omitempty"`
	Recipients  []string `json:"recipients,omitempty"`
	FullMessage string   `json:"full_message,omitempty"`
	SenderID    string   `json:"sender_id,omitempty"`
	Purpose     *string  `json:"purpose,omitempty"`
	SendAfter   string   `json:"send_after,omitempty"`
	CreatedAt   string   `json:"created_at,omitempty"`
	ExecutedAt  *string  `json:"executed_at,omitempty"`
}

// ScheduleResponse wraps a scheduled chime response.
type ScheduleResponse struct {
	ScheduledChime ScheduledChime `json:"scheduled_chime"`
}

// BroadcastChimeParams sends a broadcast chime to multiple recipients.
type BroadcastChimeParams struct {
	Recipients       []string `json:"recipients,omitempty"`
	MessageTemplate  string   `json:"message_template,omitempty"`
	ServiceName      string   `json:"service_name,omitempty"`
	Sender           string   `json:"sender,omitempty"`
	Purpose          string   `json:"purpose,omitempty"`
	PreferredGateway string   `json:"preferred_gateway,omitempty"`
	IdempotencyKey   string   `json:"idempotency_key,omitempty"`
}

// BroadcastResponse summarizes a broadcast enqueue request.
type BroadcastResponse struct {
	BroadcastID     string `json:"broadcast_id,omitempty"`
	Status          string `json:"status,omitempty"`
	RecipientsCount int    `json:"recipients_count,omitempty"`
	QueuedAt        string `json:"queued_at,omitempty"`
}

// BroadcastDetail describes a broadcast chime and its execution state.
type BroadcastDetail struct {
	ID         string           `json:"id,omitempty"`
	Recipients []string         `json:"recipients,omitempty"`
	Content    string           `json:"content,omitempty"`
	SenderID   string           `json:"sender_id,omitempty"`
	Purpose    *string          `json:"purpose,omitempty"`
	SendAfter  string           `json:"send_after,omitempty"`
	CreatedAt  string           `json:"created_at,omitempty"`
	ExecutedAt *string          `json:"executed_at,omitempty"`
	CanceledAt *string          `json:"canceled_at,omitempty"`
	Errors     []BroadcastError `json:"errors,omitempty"`
	ChimeIDs   []string         `json:"chime_ids,omitempty"`
}

// BroadcastError reports a per-recipient broadcast error.
type BroadcastError struct {
	Recipient string `json:"recipient,omitempty"`
	FixCode   string `json:"fix_code,omitempty"`
	Type      string `json:"type,omitempty"`
}

// LookupScheduleParams specifies which scheduled chime to retrieve.
type LookupScheduleParams struct {
	ScheduleID string `json:"schedule_id"`
}

// CancelScheduleParams specifies which scheduled chime to cancel.
type CancelScheduleParams struct {
	ScheduleID string `json:"schedule_id"`
}

// ScheduleLookupResponse wraps a scheduled chime lookup response.
type ScheduleLookupResponse struct {
	ScheduledChime ScheduleDetail `json:"scheduled_chime"`
}

// ScheduleCancelResponse wraps a scheduled chime cancel response.
type ScheduleCancelResponse struct {
	ScheduledChime ScheduleDetail `json:"scheduled_chime"`
}

// LookupBroadcastParams specifies which broadcast to retrieve.
type LookupBroadcastParams struct {
	BroadcastID string `json:"broadcast_id"`
}

// CancelBroadcastParams specifies which broadcast to cancel.
type CancelBroadcastParams struct {
	BroadcastID string `json:"broadcast_id"`
}

// LookupBroadcastResponse wraps a broadcast lookup response.
type LookupBroadcastResponse struct {
	Broadcast BroadcastDetail `json:"broadcast"`
}

// BroadcastCancelResponse wraps a broadcast cancel response.
type BroadcastCancelResponse struct {
	Broadcast BroadcastDetail `json:"broadcast"`
}

// LookupChimeParams specifies which chime to retrieve.
type LookupChimeParams struct {
	// ChimeID is the unique chime identifier (required).
	// Starts with "chm_". Example: "chm_abc123def456"
	ChimeID string `json:"chime_id"`
}

// Chime represents a notification message.
//
// Chimes track delivery status, transmission details, and any errors
// that occurred during sending.
type Chime struct {
	// ID is the unique chime identifier (read-only).
	// Starts with "chm_". Example: "chm_abc123def456"
	ID string `json:"id,omitempty"`

	// CreatedAt is when the chime was created (ISO 8601, read-only).
	CreatedAt string `json:"created_at,omitempty"`

	// FullMessage is the message content that was sent.
	FullMessage string `json:"full_message,omitempty"`

	// Recipient contains the recipient details.
	Recipient *ChimeRecipient `json:"recipient,omitempty"`

	// SenderID is the sender name shown to the recipient.
	SenderID string `json:"sender_id,omitempty"`

	// Purpose is the chime category.
	Purpose string `json:"purpose,omitempty"`

	// CustomData contains attached custom data.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// Delivery contains delivery status and timestamps.
	// Tracks whether the message was successfully delivered.
	Delivery any `json:"delivery,omitempty"`

	// Transmission contains low-level transmission details.
	// Includes carrier responses, error codes, etc.
	Transmission any `json:"transmission,omitempty"`
}

// FinancialAccountType enumerates types of payout destination accounts.
type FinancialAccountType string

const (
	// FinancialAccountTypeWallet represents mobile money wallet accounts.
	// Supports MTN Mobile Money, Vodafone Cash, Airtel Money, and others.
	FinancialAccountTypeWallet FinancialAccountType = "wallet"

	// FinancialAccountTypeBank represents traditional bank accounts.
	// Availability varies by country. Check country specifications.
	FinancialAccountTypeBank FinancialAccountType = "bank_account"

	// FinancialAccountTypeDosh represents Zebo's Dosh wallet accounts.
	// Internal payout destination for Zebo ecosystem users.
	FinancialAccountTypeDosh FinancialAccountType = "dosh_account"
)

// BankAccountOwnerAddress captures the account holder's address.
type BankAccountOwnerAddress struct {
	// ID is the unique address identifier (read-only).
	ID string `json:"id,omitempty"`

	// ApplicationID is the owning application ID (read-only).
	ApplicationID string `json:"application_id,omitempty"`

	// Name is a label for the address (required).
	Name string `json:"name"`

	// Phone is a contact phone number (optional).
	Phone string `json:"phone,omitempty"`

	// Line1 is the first line of the address (required).
	Line1 string `json:"line_1"`

	// Line2 is the second line of the address (optional).
	Line2 string `json:"line_2,omitempty"`

	// City is the city or town (required).
	City string `json:"city"`

	// Region is the region or state (required).
	Region string `json:"region"`

	// PostCode is the postal or ZIP code (optional).
	PostCode string `json:"post_code,omitempty"`

	// Country is the country name or code (required).
	Country string `json:"country"`
}

// BankAccountOwner captures account holder information.
type BankAccountOwner struct {
	// Name is the account holder's full name (required).
	Name string `json:"name"`

	// Address is the account holder's address (required).
	Address BankAccountOwnerAddress `json:"address"`
}

// GhanaBankAccount contains Ghana bank account details.
type GhanaBankAccount struct {
	// BankName is the name of the banking institution (optional).
	BankName string `json:"bank_name,omitempty"`

	// Branch is the bank branch identifier or name (optional).
	Branch string `json:"branch,omitempty"`

	// Number is the bank account number (required).
	Number string `json:"number"`

	// SortCode is required if SwiftCode is not provided.
	SortCode string `json:"sort_code,omitempty"`

	// SwiftCode is required if SortCode is not provided.
	SwiftCode string `json:"swift_code,omitempty"`

	// Holder is the account holder information (required).
	Holder BankAccountOwner `json:"holder"`
}

// BankAccountConfig describes bank account configuration.
type BankAccountConfig struct {
	// ID is the unique bank account identifier (read-only).
	ID string `json:"id,omitempty"`

	// Type specifies the bank account type (e.g., "ghana_bank_account").
	Type string `json:"type"`

	// GhanaBankAccount is required when Type is "ghana_bank_account".
	GhanaBankAccount *GhanaBankAccount `json:"ghana_bank_account,omitempty"`
}

// WalletConfig describes mobile money wallet configuration.
//
// Used when creating wallet-type financial accounts for receiving payouts.
type WalletConfig struct {
	// Type specifies the wallet category.
	// Currently only "mobile_money" is supported.
	Type string `json:"type"`

	// MobileMoney contains mobile money wallet details.
	MobileMoney *struct {
		// ID is assigned by the API (read-only).
		ID string `json:"id,omitempty"`

		// AccountNumber is the mobile money wallet number with country code.
		// Example: "+233244123456"
		AccountNumber string `json:"account_number"`

		// Network is the mobile money network code.
		// Examples: "mtn", "vodafone", "airteltigo", "airtel", "telecel"
		Network string `json:"network"`
	} `json:"mobile_money,omitempty"`
}

// PullPushConfig configures whether an account can send or receive funds.
//
// Pull configuration controls whether Commerce can debit the account.
// Push configuration controls whether Commerce can credit the account.
// Most payout destinations only need push enabled.
type PullPushConfig struct {
	// Enabled indicates whether the operation is allowed.
	// For payout destinations, PushConfiguration.Enabled should be true.
	// For payment sources, PullConfiguration.Enabled should be true.
	Enabled *bool `json:"enabled,omitempty"`

	// EnabledAt indicates when this configuration was enabled (read-only).
	EnabledAt string `json:"enabled_at,omitempty"`

	// Mandate contains mandate details for pull authorization (optional).
	Mandate map[string]any `json:"mandate,omitempty"`
}

// FinancialAccountCreateParams creates a payout destination account.
//
// Financial accounts are where your payouts are sent. Connect mobile money
// wallets, bank accounts, or Dosh wallets to receive settlement funds.
//
// Example (mobile money):
//
//	params := commerce.FinancialAccountCreateParams{
//	    Label:       "Primary Payout Account",
//	    Type:        commerce.FinancialAccountTypeWallet,
//	    Reference:   "main_wallet",
//	    Currency:    "ghs",
//	    Description: "Main MTN wallet for receiving payouts",
//	    PushConfiguration: &commerce.PullPushConfig{
//	        Enabled: commerce.Bool(true),
//	    },
//	    Wallet: &commerce.WalletConfig{
//	        Type: "mobile_money",
//	        MobileMoney: &struct{
//	            AccountNumber string `json:"account_number"`
//	            Network string `json:"network"`
//	        }{
//	            AccountNumber: "+233244123456",
//	            Network: "mtn",
//	        },
//	    },
//	}
type FinancialAccountCreateParams struct {
	// Label is a descriptive name for this account (required).
	// Displayed in dashboard and payout reports.
	// Example: "Primary Payout Account", "GHS Mobile Money"
	Label string `json:"label"`

	// Type specifies the account category (required).
	// Values: "wallet", "bank_account", "dosh_account"
	Type FinancialAccountType `json:"type"`

	// Reference is your internal identifier for this account (required).
	// Must be unique across your financial accounts.
	// Example: "primary_wallet", "backup_bank_01"
	Reference string `json:"reference"`

	// Currency is the account's currency (required).
	// Three-letter ISO 4217 code (lowercase).
	// Must match the currency you want to receive payouts in.
	// Example: "ghs", "usd", "kes"
	Currency string `json:"currency"`

	// Description provides additional context (optional).
	// Not displayed to customers, for internal use only.
	Description string `json:"description,omitempty"`

	// PullConfiguration controls whether Commerce can debit this account (optional).
	// Typically false for payout destinations.
	PullConfiguration *PullPushConfig `json:"pull_configuration,omitempty"`

	// PushConfiguration controls whether Commerce can credit this account (required).
	// Must be enabled for payout destinations.
	PushConfiguration *PullPushConfig `json:"push_configuration,omitempty"`

	// Wallet contains mobile money wallet details (required when Type is "wallet").
	Wallet *WalletConfig `json:"wallet,omitempty"`

	// BankAccount contains bank account details (required when Type is "bank_account").
	BankAccount *BankAccountConfig `json:"bank_account,omitempty"`

	// DoshAccount contains Dosh wallet details (required when Type is "dosh_account").
	DoshAccount map[string]any `json:"dosh_account,omitempty"`

	// CustomData contains optional key-value metadata for tracking.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// Owner contains financial account owner information (required for payouts).
	Owner *BankAccountOwner `json:"owner,omitempty"`
}

// FinancialAccountDisablePushParams disables push configuration for a financial account.
type FinancialAccountDisablePushParams struct {
	// AccountID is the financial account identifier (required).
	AccountID string `json:"account_id"`

	// UnsetAsPayoutDestination removes the account from payout destinations first.
	// Defaults to false when omitted.
	UnsetAsPayoutDestination *bool `json:"unset_as_payout_destination,omitempty"`
}

// FinancialAccountDisconnectParams disconnects a financial account.
type FinancialAccountDisconnectParams struct {
	// AccountID is the financial account identifier (required).
	AccountID string `json:"account_id"`

	// UnsetAsPayoutDestination removes the account from payout destinations first.
	// Defaults to false when omitted.
	UnsetAsPayoutDestination *bool `json:"unset_as_payout_destination,omitempty"`
}

// FinancialAccount represents a connected payout destination account.
//
// Financial accounts must be verified before use. Some account types
// require document verification or test deposits.
type FinancialAccount struct {
	// ID is the unique financial account identifier (read-only).
	// Starts with "fa_". Example: "fa_abc123def456"
	ID string `json:"id,omitempty"`

	// Label is the account's descriptive name.
	Label string `json:"label,omitempty"`

	// Type is the account category.
	// Values: "wallet", "bank_account", "dosh_account"
	Type FinancialAccountType `json:"type,omitempty"`

	// Reference is your internal identifier.
	Reference string `json:"reference,omitempty"`

	// Currency is the account's currency (ISO 4217, lowercase).
	Currency string `json:"currency,omitempty"`

	// Description provides additional context.
	Description string `json:"description,omitempty"`

	// PullConfiguration indicates whether Commerce can debit this account.
	PullConfiguration *PullPushConfig `json:"pull_configuration,omitempty"`

	// PushConfiguration indicates whether Commerce can credit this account.
	PushConfiguration *PullPushConfig `json:"push_configuration,omitempty"`

	// Wallet contains mobile money wallet details (when Type is "wallet").
	Wallet *WalletConfig `json:"wallet,omitempty"`

	// BankAccount contains bank account details (when Type is "bank_account").
	BankAccount *BankAccountConfig `json:"bank_account,omitempty"`

	// DoshAccount contains Dosh wallet details (when Type is "dosh_account").
	DoshAccount map[string]any `json:"dosh_account,omitempty"`

	// CustomData contains optional key-value metadata for tracking.
	CustomData map[string]string `json:"custom_data,omitempty"`

	// Owner contains financial account owner information.
	Owner *BankAccountOwner `json:"owner,omitempty"`

	// Verification contains verification status and requirements.
	// Nil if verification not started or not required.
	Verification any `json:"verification,omitempty"`

	// ArchivedAt is when the account was archived (ISO 8601, read-only).
	// Archived accounts cannot receive new payouts.
	// Nil if not archived.
	ArchivedAt *string `json:"archived_at,omitempty"`

	// DisconnectedAt is when the account was disconnected (ISO 8601, read-only).
	// Nil if the account is still active.
	DisconnectedAt *string `json:"disconnected_at,omitempty"`

	// CreatedAt is when the account was created (ISO 8601, read-only).
	CreatedAt *string `json:"created_at,omitempty"`
}

// PageFinancialAccountsParams paginates financial accounts.
type PageFinancialAccountsParams struct {
	PageNumber int `json:"page_number,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}

// FinancialAccountsPage holds paginated account data.
type FinancialAccountsPage struct {
	Number   int                `json:"number,omitempty"`
	Size     int                `json:"size,omitempty"`
	Accounts []FinancialAccount `json:"accounts,omitempty"`
}

// PageFinancialAccountsResponse wraps a page of accounts.
type PageFinancialAccountsResponse struct {
	Page FinancialAccountsPage `json:"page"`
}

// PaymentMethodSettings contains acceptance configuration for each payment method type.
//
// Controls which payment methods customers can use and whether OTP confirmation
// is required. Retrieve with PaymentMethods.Settings().
type PaymentMethodSettings struct {
	// MobileMoney configures mobile money wallet acceptance.
	MobileMoney *PaymentMethodTypeSetting `json:"mobile_money,omitempty"`

	// BankAccount configures bank account acceptance.
	BankAccount *PaymentMethodTypeSetting `json:"bank_account,omitempty"`

	// Card configures card payment acceptance.
	Card *PaymentMethodTypeSetting `json:"card,omitempty"`

	// Motito configures Zebo's branded payment method.
	Motito *PaymentMethodTypeSetting `json:"motito,omitempty"`
}

// PaymentMethodTypeSetting controls availability and confirmation for a payment method type.
type PaymentMethodTypeSetting struct {
	// Type is the payment method category (read-only).
	Type string `json:"type,omitempty"`

	// Name is the display name (read-only).
	Name string `json:"name,omitempty"`

	// Description explains the payment method (read-only).
	Description string `json:"description,omitempty"`

	// Enabled indicates whether this payment method is accepted (read-only).
	// Customers cannot use disabled payment methods.
	Enabled bool `json:"enabled,omitempty"`

	// ConfirmsUse indicates whether OTP confirmation is required (read-only).
	// When true, customers must provide an OTP before charges complete.
	// Adds security but creates friction. Most merchants keep this enabled.
	ConfirmsUse bool `json:"confirms_use,omitempty"`
}

// PayoutScheduleSpec describes the balance transaction aging period.
//
// Balance transactions must age before becoming eligible for payout.
// This protects against late-arriving disputes and chargebacks.
type PayoutScheduleSpec struct {
	// ID is the spec identifier.
	ID string `json:"id,omitempty"`

	// TPlus indicates the aging period in days.
	// Example: "t+7" means 7 days after transaction.
	TPlus string `json:"t_plus,omitempty"`

	// Label is a human-readable description.
	Label string `json:"label,omitempty"`

	// Abide is the formal aging rule specification.
	Abide string `json:"abide,omitempty"`
}

// PayoutSchedule describes your payout timing configuration.
//
// Payouts can be automatic (weekly, daily) or manual (on-demand).
// Automatic schedules trigger payouts at regular intervals for eligible
// balance transactions. Manual mode requires explicit payout initiation.
type PayoutSchedule struct {
	// ID is the schedule identifier (read-only).
	ID string `json:"id,omitempty"`

	// Name is the schedule's display name (read-only).
	// Example: "Weekly Automatic", "Manual"
	Name string `json:"name,omitempty"`

	// Type indicates automatic or manual mode (read-only).
	// Values: "automatic", "manual"
	Type string `json:"type,omitempty"`

	// Interval is the payout frequency for automatic schedules (read-only).
	// Values: "weekly", "daily", nil (for manual)
	Interval string `json:"interval,omitempty"`

	// ScheduleOn specifies when automatic payouts run (read-only).
	// For weekly: day of week (e.g., "monday")
	// For daily: time of day
	ScheduleOn string `json:"schedule_on,omitempty"`

	// Description explains the schedule behavior (read-only).
	Description string `json:"description,omitempty"`

	// Spec contains the balance transaction aging rules (read-only).
	Spec *PayoutScheduleSpec `json:"spec,omitempty"`
}

// PayoutSettings contains your complete payout configuration.
//
// Controls when payouts happen, where funds go, and whether currency
// conversion is enabled.
type PayoutSettings struct {
	// ID is the settings identifier (read-only).
	ID string `json:"id,omitempty"`

	// FxEnabled indicates whether currency conversion is enabled (read-only).
	// When true, can receive payouts in different currency than source funds.
	// Requires FX-enabled destination accounts.
	FxEnabled bool `json:"fx_enabled,omitempty"`

	// Destinations maps currencies to financial account IDs.
	// Key: currency code (e.g., "ghs", "usd")
	// Value: financial account ID (e.g., "fa_abc123")
	// Example: {"ghs": "fa_abc123", "usd": "fa_def456"}
	Destinations map[string]string `json:"destinations,omitempty"`

	// Schedule describes payout timing and frequency.
	Schedule *PayoutSchedule `json:"schedule,omitempty"`
}

// Payout represents a settlement transfer to your bank or mobile money account.
//
// Payouts move funds from your Commerce balance to your financial accounts.
// Each payout contains one or more balance transactions that have aged
// past the dispute window.
type Payout struct {
	// ID is the unique payout identifier (read-only).
	// Starts with "po_". Example: "po_abc123def456"
	ID string `json:"id,omitempty"`

	// ApplicationID is your application's ID (read-only).
	ApplicationID string `json:"application_id,omitempty"`

	// DestinationID is the receiving financial account's ID (read-only).
	// Corresponds to a financial account you've connected.
	DestinationID string `json:"destination_id,omitempty"`

	// Amount is the payout total (read-only).
	Amount *Money `json:"amount,omitempty"`

	// Status is the payout's current state (read-only).
	// Values include "scheduled", "initiated", "processing", "succeeded", "failed", "canceled"
	Status string `json:"status,omitempty"`

	// InitiatedBy indicates who triggered the payout (read-only).
	// Values: "schedule" (automatic), "manual" (you initiated)
	InitiatedBy string `json:"initiated_by,omitempty"`

	// LatestAttemptID is the most recent execution attempt's ID (read-only).
	LatestAttemptID string `json:"latest_attempt_id,omitempty"`

	// LatestError contains error details if payout failed (read-only).
	// Nil if payout succeeded or is still processing.
	LatestError any `json:"latest_error,omitempty"`

	// InitiatedAt is when the payout was created (ISO 8601, read-only).
	InitiatedAt string `json:"initiated_at,omitempty"`

	// ExecuteAfter is the scheduled execution timestamp for queued payouts (ISO 8601, read-only).
	// Nil for immediate/manual payouts that are not scheduled.
	ExecuteAfter *string `json:"execute_after,omitempty"`

	// ScheduledAt is when the payout was queued for execution (ISO 8601, read-only).
	// Nil when not scheduled.
	ScheduledAt *string `json:"scheduled_at,omitempty"`

	// CanceledAt is when a scheduled payout was canceled (ISO 8601, read-only).
	// Nil unless the payout has status "canceled".
	CanceledAt *string `json:"canceled_at,omitempty"`

	// MaxAmount is the maximum amount authorized for scheduled payouts (read-only).
	// This may differ from Amount when payout execution has not started.
	MaxAmount *Money `json:"max_amount,omitempty"`

	// ExecutedAt is when the payout was submitted to the network (ISO 8601, read-only).
	// Nil if not yet executed.
	ExecutedAt *string `json:"executed_at,omitempty"`

	// ExpectedAt is when the payout should arrive (ISO 8601, read-only).
	// Estimate based on network speed. Actual arrival may vary.
	ExpectedAt *string `json:"expected_at,omitempty"`

	// SucceededAt is when the payout was confirmed (ISO 8601, read-only).
	// Nil if not yet succeeded.
	SucceededAt *string `json:"succeeded_at,omitempty"`

	// BalanceTransactionIDs lists the included balance transactions (read-only).
	// These are the source funds being paid out.
	BalanceTransactionIDs []string `json:"balance_transaction_ids,omitempty"`
}

// BalanceTransaction represents funds from a completed payment.
//
// When an order payment succeeds, a balance transaction is created representing
// the available funds. Balance transactions age for 7 days (or your configured
// aging period) before becoming eligible for payout. This protects against
// late-arriving disputes.
type BalanceTransaction struct {
	// ID is the unique balance transaction identifier (read-only).
	// Starts with "bt_". Example: "bt_abc123def456"
	ID string `json:"id,omitempty"`

	// PaymentID is the source payment's ID (read-only).
	PaymentID string `json:"payment_id,omitempty"`

	// OrderID is the source order's ID (read-only).
	OrderID string `json:"order_id,omitempty"`

	// AmountExpected is the gross amount before fees (read-only).
	AmountExpected *Money `json:"amount_expected,omitempty"`

	// AmountAvailable is the net amount after fees (read-only).
	// This is the amount that will be paid out to you.
	AmountAvailable *Money `json:"amount_available,omitempty"`

	// AvailableAt is when funds become eligible for payout (ISO 8601, read-only).
	// Typically 7 days after CreatedAt. Nil if already paid out.
	AvailableAt *string `json:"available_at,omitempty"`

	// CreatedAt is when the balance transaction was created (ISO 8601, read-only).
	CreatedAt *string `json:"created_at,omitempty"`

	// PayoutConfiguration describes how this balance transaction will be paid out.
	PayoutConfiguration *PayoutConfiguration `json:"payout_configuration,omitempty"`
}

// PayoutConfiguration describes payout routing and FX settings for a payment or balance transaction.
type PayoutConfiguration struct {
	// EnableFX indicates whether FX conversion is enabled for this payout.
	EnableFX *bool `json:"enable_fx,omitempty"`

	// Destination specifies the financial account receiving the payout.
	Destination *PayoutDestination `json:"destination,omitempty"`
}

// PayoutDestination identifies the payout financial account.
type PayoutDestination struct {
	// FinancialAccountID is the ID of the destination financial account.
	FinancialAccountID string `json:"financial_account_id,omitempty"`
}

// TokenizePaymentMethodParams saves a payment method for future use.
//
// Tokenization stores customer payment details securely for repeat charges.
// The customer owns the payment method—only they can delete it.
//
// Example:
//
//	params := commerce.TokenizePaymentMethodParams{
//	    CustomerID: "cu_abc123",
//	    PaymentMethodData: commerce.PaymentMethodData{
//	        Type: commerce.PaymentMethodTypeMobileMoney,
//	        MobileMoney: &commerce.MobileMoneyParams{
//	            Network: "mtn",
//	            AccountNumber: "+233244123456",
//	        },
//	    },
//	    VerifyImmediately: commerce.Bool(true),
//	}
type TokenizePaymentMethodParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`

	// CustomerID is who owns this payment method (required).
	CustomerID string `json:"customer_id"`

	// PaymentMethodData contains the payment details to save (required).
	PaymentMethodData PaymentMethodData `json:"payment_method_data"`

	// VerifyImmediately triggers verification right after tokenization (optional).
	// If true, sends OTP immediately for customer to confirm ownership.
	// If false (default), must call Verify() separately before first use.
	VerifyImmediately *bool `json:"verify_immediately,omitempty"`
}

// VerifyPaymentMethodParams starts verification for a payment method.
//
// Sends an OTP to the payment method (phone number for mobile money)
// to confirm the customer owns it. Required before first use.
type VerifyPaymentMethodParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`

	// PaymentMethodID is the payment method to verify (required).
	PaymentMethodID string `json:"payment_method_id"`
}

// VerificationStatusResponse contains verification state and delivery details.
type VerificationStatusResponse struct {
	// Verification contains verification progress and OTP delivery info.
	Verification *struct {
		// PaymentMethodID is the payment method being verified.
		PaymentMethodID string `json:"payment_method_id,omitempty"`

		// Status is the verification state.
		// Values: "pending", "verified", "failed"
		Status string `json:"status,omitempty"`

		// TokenSentAt is when OTP was sent (ISO 8601).
		TokenSentAt *string `json:"token_sent_at,omitempty"`

		// ExpiresAt is when OTP expires (ISO 8601).
		// Customer must confirm before this time.
		ExpiresAt *string `json:"expires_at,omitempty"`

		// Delivery contains OTP delivery details.
		Delivery *struct {
			// Recipient is where OTP was sent.
			Recipient string `json:"recipient,omitempty"`

			// Channel is the delivery method.
			// Values: "sms", "ussd"
			Channel string `json:"channel,omitempty"`

			// SenderID is the SMS sender name.
			SenderID string `json:"sender_id,omitempty"`
		} `json:"delivery,omitempty"`
	} `json:"verification,omitempty"`
}

// ConfirmPaymentMethodVerificationParams submits the OTP to complete verification.
type ConfirmPaymentMethodVerificationParams struct {
	// PaymentMethodID is the payment method being verified (required).
	PaymentMethodID string `json:"payment_method_id"`

	// Token is the OTP from the customer (required).
	// Typically 4-6 digits sent via SMS.
	Token string `json:"token"`
}

// LookupPaymentMethodParams specifies which payment method to retrieve.
type LookupPaymentMethodParams struct {
	// PaymentMethodID is the payment method identifier (required).
	// Starts with "pm_". Example: "pm_abc123def456"
	PaymentMethodID string `json:"payment_method_id"`
}

// DeletePaymentMethodParams permanently removes a payment method.
//
// Deleted payment methods cannot be restored. Customers can delete
// their own payment methods; merchants cannot force re-enablement.
type DeletePaymentMethodParams struct {
	// RequestMeta carries per-request controls such as idempotency.
	RequestMeta *RequestMeta `json:"request_meta,omitempty"`

	// PaymentMethodID is the payment method to delete (required).
	PaymentMethodID string `json:"payment_method_id"`
}

// DeletePaymentMethodResponse confirms deletion.
type DeletePaymentMethodResponse struct {
	// Deleted indicates whether deletion succeeded.
	Deleted bool `json:"deleted"`

	// PaymentMethodID is the deleted payment method's ID.
	PaymentMethodID string `json:"payment_method_id,omitempty"`
}

// PayoutPageParams specifies pagination for listing payouts.
type PayoutPageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed.
	PageNumber int `json:"page_number,omitempty"`

	// PageSize is the number of payouts per page (optional, default: 20).
	// Maximum 100.
	PageSize int `json:"page_size,omitempty"`
}

// BalanceTransactionPageParams specifies pagination for listing balance transactions.
type BalanceTransactionPageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed.
	PageNumber int `json:"page_number,omitempty"`

	// PageSize is the number of transactions per page (optional, default: 20).
	// Maximum 100.
	PageSize int `json:"page_size,omitempty"`
}

// CountryBankBranch describes a bank branch available in a country bank directory.
type CountryBankBranch struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	SortCode string `json:"sort_code,omitempty"`
}

// CountryBank describes a banking institution available in a country bank directory.
type CountryBank struct {
	ID             string              `json:"id,omitempty"`
	Name           string              `json:"name,omitempty"`
	SwiftCode      string              `json:"swift_code,omitempty"`
	SortCodePrefix string              `json:"sort_code_prefix,omitempty"`
	Branches       []CountryBankBranch `json:"branches,omitempty"`
}

// CountryBankDirectory describes bank reference data for country-specific bank accounts.
type CountryBankDirectory struct {
	BankAccountType string        `json:"bank_account_type,omitempty"`
	CodeScheme      string        `json:"code_scheme,omitempty"`
	Items           []CountryBank `json:"items,omitempty"`
}

// CountrySpecification describes supported Commerce features for a country.
//
// Use this to discover supported currencies, payment methods, payout schedules,
// and other country-specific capabilities before integrating.
//
// Query with Spec.Countries() to get all country specifications.
type CountrySpecification struct {
	// CountryCode is the two-letter ISO 3166-1 alpha-2 code (read-only).
	// Example: "GH", "KE", "UG", "US"
	CountryCode string `json:"country_code,omitempty"`

	// CountryName is the full country name (read-only).
	// Example: "Ghana", "Kenya", "Uganda"
	CountryName string `json:"country_name,omitempty"`

	// Currencies lists supported currency codes (read-only).
	// Example: ["ghs", "usd"] for Ghana
	Currencies []string `json:"currencies,omitempty"`

	// PaymentMethods lists supported payment method types (read-only).
	// Example: ["mobile_money", "bank_account"]
	PaymentMethods []string `json:"payment_methods,omitempty"`

	// PayoutSchedules lists available payout schedule types (read-only).
	// Example: ["weekly", "manual"]
	PayoutSchedules []string `json:"payout_schedules,omitempty"`

	// BTAgingSpecs lists balance transaction aging options (read-only).
	// Example: ["t+7", "t+14"] for 7-day or 14-day aging
	BTAgingSpecs []string `json:"bt_aging_specs,omitempty"`

	// LegalEntityTypes lists supported business types (read-only).
	// Structure varies by country requirements.
	LegalEntityTypes []map[string]any `json:"legal_entity_types,omitempty"`

	// FinancialAccountTypes lists supported payout destination types (read-only).
	// Details wallet, bank_account, and dosh_account configurations.
	FinancialAccountTypes []map[string]any `json:"financial_account_types,omitempty"`

	// IDDocumentTypes lists accepted identification documents (read-only).
	// Used for KYC/verification requirements.
	IDDocumentTypes []map[string]any `json:"id_document_types,omitempty"`

	// Banks lists country-specific bank reference data, when available.
	// Ghana uses bank_account_type "ghana_bank_account" and sort code branches.
	Banks *CountryBankDirectory `json:"banks,omitempty"`
}
