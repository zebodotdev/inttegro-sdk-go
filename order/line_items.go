package order

import (
	"github.com/zebodotdev/inttegro-sdk-go/v7/money"
	"github.com/zebodotdev/inttegro-sdk-go/v7/price"
	"github.com/zebodotdev/inttegro-sdk-go/v7/product"
)

// ProductLineItemParams represents a product supplied in an order request.
//
// Products are goods or services sold to customers. Each product has a unit
// price, quantity, and optional descriptive information.
//
// Example:
//
//	item := &order.ProductLineItemParams{
//	    Type:     "physical",
//	    Name:     "Wireless Headphones",
//	    About:    "Bluetooth 5.0, 30-hour battery",
//	    Quantity: 2,
//	    Price:    price.InlineParams{AmountParams: money.AmountParams{Currency: money.USD, Value: 7999}}, // $79.99 each
//	    Reference: "SKU-12345",
//	}
type ProductLineItemParams struct {
	// Type indicates whether the product is physical or digital (required).
	// Values: "physical" or "digital"
	// Physical products require shipping address.
	// Digital products can be delivered electronically.
	Type product.Type `json:"type"`

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
	Price price.InlineParams `json:"price"`

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
	// Learn more: https://studio.inttegro.com/custom-data
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// ProductLineItem is a product returned in an order.
type ProductLineItem struct {
	ID         string            `json:"id"`
	ProductID  string            `json:"product_id,omitempty"`
	PriceID    string            `json:"price_id,omitempty"`
	Reference  string            `json:"reference,omitempty"`
	About      string            `json:"about,omitempty"`
	CustomData map[string]string `json:"custom_data,omitempty"`
	TaxCode    string            `json:"tax_code,omitempty"`
	Name       string            `json:"name"`
	Category   string            `json:"category,omitempty"`
	Type       product.Type      `json:"type,omitempty"`
	Price      price.Inline      `json:"price"`
	Quantity   int64             `json:"quantity"`
}

// FeeLineItemParams represents an additional charge supplied in a request.
//
// Fees are one-time charges added to the order subtotal. Unlike products,
// fees don't have quantities—they're always a fixed amount.
//
// Example:
//
//	fee := &order.FeeLineItemParams{
//	    Label:       "Service Fee",
//	    Description: "Platform usage fee",
//	    Amount:      money.AmountParams{Currency: money.USD, Value: 299}, // $2.99
//	}
type FeeLineItemParams struct {
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
	Amount money.AmountParams `json:"amount"`
}

// FeeLineItem is an additional charge returned in an order.
type FeeLineItem struct {
	ID          string       `json:"id"`
	Description string       `json:"description,omitempty"`
	TaxCode     string       `json:"tax_code,omitempty"`
	Amount      money.Amount `json:"amount"`
	Label       string       `json:"label"`
}

// ShippingLineItemParams represents a delivery charge supplied in a request.
//
// Only needed when selling physical products. Automatically omitted for
// orders containing only digital products.
//
// Example:
//
//	shipping := &order.ShippingLineItemParams{
//	    Fee: money.AmountParams{Currency: money.USD, Value: 500}, // $5.00
//	}
type ShippingLineItemParams struct {
	// Fee is the total shipping charge (required).
	// In minor units. Not multiplied by any quantity.
	Fee money.AmountParams `json:"fee"`

	// TaxCode specifies the tax treatment (optional).
	// Used for tax calculation when tax integration is enabled.
	TaxCode string `json:"tax_code,omitempty"`

	// CustomData holds arbitrary key-value custom data (optional).
	// Both keys and values must be strings.
	// Maximum 25KB when serialized.
	CustomData map[string]string `json:"custom_data,omitempty"`
}

// ShippingLineItem is a delivery charge returned in an order.
type ShippingLineItem struct {
	ID      string       `json:"id"`
	TaxCode string       `json:"tax_code,omitempty"`
	Label   string       `json:"label,omitempty"`
	Fee     money.Amount `json:"fee"`
}

// OrderLineItemParams is a discriminated union supplied in an order request.
//
// Each order line item is one of three types: product, fee, or shipping.
// Set Type and the corresponding field (Product, Fee, or Shipping).
// Leave the other fields nil.
//
// Example (product):
//
//	lineItem := order.LineItemParams{
//	    Type: order.LineItemTypeProduct,
//	    Product: &order.ProductLineItemParams{
//	        Type:     "digital",
//	        Name:     "Premium Subscription",
//	        Quantity: 1,
//	        Price:    price.InlineParams{AmountParams: money.AmountParams{Currency: money.USD, Value: 999}},
//	    },
//	}
//
// Example (fee):
//
//	lineItem := order.LineItemParams{
//	    Type: order.LineItemTypeFee,
//	    Fee: &order.FeeLineItemParams{
//	        Label:  "Platform Fee",
//	        Amount: money.AmountParams{Currency: money.USD, Value: 299},
//	    },
//	}
type LineItemParams struct {
	// Type specifies which variant is active (required).
	Type LineItemType `json:"type"`

	// Product is populated when Type is LineItemTypeProduct.
	// Nil for other types.
	Product *ProductLineItemParams `json:"product,omitempty"`

	// Fee is populated when Type is LineItemTypeFee.
	// Nil for other types.
	Fee *FeeLineItemParams `json:"fee,omitempty"`

	// Shipping is populated when Type is LineItemTypeShipping.
	// Nil for other types.
	Shipping *ShippingLineItemParams `json:"shipping,omitempty"`
}

// OrderLineItem is a discriminated union returned by the API.
type LineItem struct {
	Type     LineItemType      `json:"type"`
	Discount *DiscountLineItem `json:"discount,omitempty"`
	Product  *ProductLineItem  `json:"product,omitempty"`
	Fee      *FeeLineItem      `json:"fee,omitempty"`
	Shipping *ShippingLineItem `json:"shipping,omitempty"`
}

type DiscountLineItem struct{}

// LineItemGroup contains line items grouped by type with totals.
//
// The API returns this grouped structure in order responses to make
// it easier to display cart breakdowns.
type LineItemGroup struct {
	// LineItems is the list of cart items.
	LineItems []LineItem `json:"line_items"`

	// Total is the sum of all line items.
	// This is the order amount before any fees or discounts.
	Total money.Amount `json:"total"`
}
