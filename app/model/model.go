package model

type SMSDataModel struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

type MMSDataModel struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type VoiceDataModel struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

type EmailDataModel struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type BillingDataModel struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

type SupportDataModel struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type IncidentDataModel struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]SMSDataModel              `json:"sms"`
	MMS       [][]MMSDataModel              `json:"mms"`
	VoiceCall []VoiceDataModel              `json:"voice_call"`
	Email     map[string][][]EmailDataModel `json:"email"`
	Billing   BillingDataModel              `json:"billing"`
	Support   []int                         `json:"support"`
	Incident  []IncidentDataModel           `json:"incident"`
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
