package secretkey

type UsageRow struct {
	SecretKeyID string     `json:"secret_key_id"`
	OccurredAt  string     `json:"occurred_at"`
	AuthResult  AuthResult `json:"auth_result"`
}

type UsagePage struct {
	Number  int        `json:"number"`
	Size    int        `json:"size"`
	Count   int        `json:"count"`
	Total   int        `json:"total"`
	HasMore bool       `json:"has_more"`
	Rows    []UsageRow `json:"rows"`
}

// SecretKeyUsage contains key metadata and its recent authentication outcomes.
type Usage struct {
	Key   SecretKey `json:"key"`
	Usage UsagePage `json:"usage"`
}
