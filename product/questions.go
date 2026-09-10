package product

// IsArchived reports whether the product is archived.
func (p Product) IsArchived() bool { return p.ArchivedAt != nil }

// IsPublished reports whether the product is currently published and available.
func (p Product) IsPublished() bool { return p.Active && !p.IsArchived() }

// WasEverPublished reports whether the product has a recorded first publication.
func (p Product) WasEverPublished() bool { return p.PublishedAt != nil }
