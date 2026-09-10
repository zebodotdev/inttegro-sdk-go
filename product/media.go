package product

// Media contains the known catalog media slots for a product.
type Media struct {
	HeroImage   string   `json:"hero_image,omitempty"`
	Thumbnail   string   `json:"thumbnail,omitempty"`
	WebPageURL  string   `json:"web_page_url,omitempty"`
	BrandLogo   string   `json:"brand_logo,omitempty"`
	Infographic string   `json:"infographic,omitempty"`
	PromoVideo  string   `json:"promo_video,omitempty"`
	DemoVideo   string   `json:"demo_video,omitempty"`
	Gallery     []string `json:"gallery,omitempty"`
	Downloads   []string `json:"downloads,omitempty"`
}
