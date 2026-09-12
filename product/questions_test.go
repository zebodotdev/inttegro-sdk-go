package product

import (
	"testing"
	"time"
)

func TestProductQuestions(t *testing.T) {
	publishedAt := time.Now()
	product := Product{Active: true, PublishedAt: &publishedAt}
	if product.IsArchived() || !product.IsPublished() || !product.WasEverPublished() {
		t.Fatal("published product reported the wrong publication state")
	}

	product.ArchivedAt = &publishedAt
	if !product.IsArchived() || product.IsPublished() {
		t.Fatal("archived product should not be published")
	}
}
