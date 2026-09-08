package order

import (
	"context"
)

// Page returns a paginated list of recent orders.
//
// Use this to display order history, search for orders, or sync order state
// with your system. Results are sorted by creation date (newest first).
//
// Parameters:
//   - params.PageNumber: Page to retrieve (optional, default: 1)
//   - params.PageSize: Orders per page (optional, default: 20, max: 100)
//
// Returns a slice of orders for the requested page.
//
// Example:
//
//	orders, err := client.Orders.Page(ctx, order.PageParams{
//	    PageNumber: 1,
//	    PageSize:   50,
//	})
//	for _, order := range orders {
//	    fmt.Printf("Order %s: %s\n", order.ID, order.Status)
//	}
func (s *Service) Page(ctx context.Context, params PageParams) ([]Resource, error) {
	var resp struct {
		Page struct {
			Orders []Resource `json:"orders"`
		} `json:"page"`
	}
	if err := s.client.Do(ctx, "POST", "/orders/page", params, &resp); err != nil {
		return nil, err
	}
	return resp.Page.Orders, nil
}

// OrderPageParams specifies pagination for listing orders.
//
// Use this to fetch recent orders with pagination support.
//
// Example:
//
//	params := order.PageParams{
//	    PageNumber: 1,
//	    PageSize:   50,
//	}
//	orders, err := client.Orders.Page(ctx, params)
type PageParams struct {
	// PageNumber is the page to retrieve (optional, default: 1).
	// Pages are 1-indexed. First page is 1, not 0.
	PageNumber int `json:"page_number,omitempty"`

	// PageSize is the number of orders per page (optional, default: 20).
	// Maximum 100. Minimum 1.
	PageSize int `json:"page_size,omitempty"`
}
