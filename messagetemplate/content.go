package messagetemplate

type SMSContent struct {
	MessageTemplate string `json:"message_template"`
}

type Mailbox struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type EmailContent struct {
	From    *Mailbox          `json:"from,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	HTML    string            `json:"html"`
	ReplyTo *Mailbox          `json:"reply_to,omitempty"`
	Subject string            `json:"subject"`
}

type RenderedContent struct {
	Channel Channel        `json:"channel"`
	Email   map[string]any `json:"email,omitempty"`
	SMS     map[string]any `json:"sms,omitempty"`
}
