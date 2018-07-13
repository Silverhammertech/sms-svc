package model

type TwilioResponse struct {
	Sid string `json:"sid"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	DateSent string `json:"date_sent"`
	AccountSid string `json:"account_sid"`
	To string `json:"to"`
	From string `json:"from"`
	MessagingServiceSid string `json:"messaging_service_sid"`
	Body string `json:"body"`
	Status string `json:"status"`
	NumSegments string `json:"num_segments"`
	NumMedia string `json:"num_media"`
	Direction string `json:"direction"`
	APIVersion string `json:"api_version"`
	Price string `json:"price"`
	PriceUnit string `json:"price_unit"`
	ErrorCode string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	URI string `json:"uri"`
	SubresourceUris struct {
		    Media string `json:"media"`
	    } `json:"subresource_uris"`
}

type TwilioErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status int `json:"status"`
}
