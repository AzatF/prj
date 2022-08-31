package model

type SMSDataModel struct {
	Country      string `json:"country,omitempty"`
	Bandwidth    string `json:"bandwidth,omitempty"`
	ResponseTime string `json:"response_time,omitempty"`
	Provider     string `json:"provider,omitempty"`
}

type MMSDataModel struct {
	Country      string `json:"country,omitempty"`
	Bandwidth    string `json:"bandwidth,omitempty"`
	ResponseTime string `json:"response_time,omitempty"`
	Provider     string `json:"provider,omitempty"`
}

type VoiceDataModel struct {
	Country             string  `json:"country,omitempty"`
	Bandwidth           string  `json:"bandwidth,omitempty"`
	ResponseTime        string  `json:"response_time,omitempty"`
	Provider            string  `json:"provider,omitempty"`
	ConnectionStability float32 `json:"connection_stability,omitempty"`
	TTFB                int     `json:"ttfb,omitempty"`
	VoicePurity         int     `json:"voice_purity,omitempty"`
	MedianOfCallsTime   int     `json:"median_of_calls_time,omitempty"`
}

type EmailDataModel struct {
	Country      string `json:"country,omitempty"`
	Provider     string `json:"provider,omitempty"`
	DeliveryTime int    `json:"delivery_time,omitempty"`
}

type BillingDataModel struct {
	CreateCustomer bool `json:"create_customer,omitempty"`
	Purchase       bool `json:"purchase,omitempty"`
	Payout         bool `json:"payout,omitempty"`
	Recurring      bool `json:"recurring,omitempty"`
	FraudControl   bool `json:"fraud_control,omitempty"`
	CheckoutPage   bool `json:"checkout_page,omitempty"`
}

type SupportDataModel struct {
	Topic         string `json:"topic,omitempty"`
	ActiveTickets int    `json:"active_tickets,omitempty"`
}

type IncidentDataModel struct {
	Topic  string `json:"topic,omitempty"`
	Status string `json:"status,omitempty"`
}

type ResultT struct {
	Status bool       `json:"status,omitempty"`
	Data   ResultSetT `json:"data,omitempty"`
	Error  string     `json:"error,omitempty"`
}

type ResultSetT struct {
	SMS       [][]SMSDataModel              `json:"sms,omitempty"`
	MMS       [][]MMSDataModel              `json:"mms,omitempty"`
	VoiceCall []VoiceDataModel              `json:"voice_call,omitempty"`
	Email     map[string][][]EmailDataModel `json:"email,omitempty"`
	Billing   BillingDataModel              `json:"billing,omitempty"`
	Support   []int                         `json:"support,omitempty"`
	Incident  []IncidentDataModel           `json:"incident,omitempty"`
}

type ISO3166 struct {
	Country string
	Alpha2  string
	Alpha3  string
	Code    string
}

type Providers struct {
	Provider []string
}
