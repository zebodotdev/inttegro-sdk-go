package secretkey

type GenerateParams struct {
	Label string `json:"label,omitempty"`
}

type LookupParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
}

type UpdateParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	Label       string `json:"label"`
}

type DestroyParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
}

type PageParams struct {
	Page   int `json:"page,omitempty"`
	Number int `json:"number,omitempty"`
	Size   int `json:"size,omitempty"`
}

type UsageParams struct {
	SecretKeyID string `json:"secret_key_id,omitempty"`
	Page        int    `json:"page,omitempty"`
	Number      int    `json:"number,omitempty"`
	Size        int    `json:"size,omitempty"`
}
