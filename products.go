package inttegro

import "context"

// ProductsService manages catalog products.
type ProductsService struct {
	client *Client
}

// Create creates a product.
func (s *ProductsService) Create(ctx context.Context, params CreateProductParams) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.do(ctx, "POST", "/products/create", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// AddPrice adds a price to an existing product.
func (s *ProductsService) AddPrice(ctx context.Context, params AddProductPriceParams) (*CatalogPrice, error) {
	var resp struct {
		Price CatalogPrice `json:"price"`
	}
	if err := s.client.do(ctx, "POST", "/products/add_price", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Price, nil
}

// SetDefaultUnitPrice sets an existing product price as the product's default unit price.
func (s *ProductsService) SetDefaultUnitPrice(ctx context.Context, params SetDefaultUnitPriceParams) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.do(ctx, "POST", "/products/set_default_unit_price", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Lookup retrieves a product by ID.
func (s *ProductsService) Lookup(ctx context.Context, productID string) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.do(ctx, "POST", "/products/lookup", LookupProductParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Page retrieves a page of products.
func (s *ProductsService) Page(ctx context.Context, params PageProductsParams) (*ProductsPage, error) {
	var resp struct {
		Page ProductsPage `json:"page"`
	}
	if err := s.client.do(ctx, "POST", "/products/page", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Page, nil
}

// Update updates a product.
func (s *ProductsService) Update(ctx context.Context, params UpdateProductParams) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.do(ctx, "POST", "/products/update", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Publish publishes a product.
func (s *ProductsService) Publish(ctx context.Context, productID string) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.do(ctx, "POST", "/products/publish", ProductActionParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Unpublish unpublishes a product.
func (s *ProductsService) Unpublish(ctx context.Context, productID string) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.do(ctx, "POST", "/products/unpublish", ProductActionParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}

// Archive archives a product.
func (s *ProductsService) Archive(ctx context.Context, productID string) (*Product, error) {
	var resp struct {
		Product Product `json:"product"`
	}
	if err := s.client.do(ctx, "POST", "/products/archive", ProductActionParams{ProductID: productID}, &resp); err != nil {
		return nil, err
	}
	return &resp.Product, nil
}
