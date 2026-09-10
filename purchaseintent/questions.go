package purchaseintent

// IsActive reports whether the purchase intent can currently be presented as active.
func (p PurchaseIntent) IsActive() bool { return p.Status == StatusActive }

// IsSingleUse reports whether the purchase intent can create at most one order.
func (p PurchaseIntent) IsSingleUse() bool {
	return p.Usage.SingleUse != nil && *p.Usage.SingleUse
}

// UsedOrderID returns the order that consumed a single-use purchase intent.
func (p PurchaseIntent) UsedOrderID() (string, bool) {
	if !p.IsSingleUse() || p.Usage.Order == nil || p.Usage.Order.ID == "" {
		return "", false
	}
	return p.Usage.Order.ID, true
}
