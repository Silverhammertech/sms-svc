package model

type SMSRequest struct{
	SID string     // A 34 character string that uniquely identifies resource
	Token string   // Authtoken
	From string    // The phone number in E164 format or alphanumeric sender ID that initiated the message. For incoming messages, this will be the remote phone. For outgoing messages, this will be one of your Twilio phone numbers or the alphanumeric sender ID used.
	To []string     `json:"to"`// The phone number that received the message in E614 format. For incoming messages, this will be one of your Twilio phone numbers. For outcoming messages, this will be the remote phone.
	Message string 	`json:"message"`
}


type SMSResponse struct {
	Recepient string `json:"recepient"`
	Status string `json:"status"`
	StatusDescription string `json:"statusdescription"`
}

type SendableSMSRequest SMSRequest

type ISMSSender interface {
	Send(params SendableSMSRequest) ([]SMSResponse, error)
}

func (params SendableSMSRequest)SendSMS(s ISMSSender) ([]SMSResponse, error){
	return s.Send(params)
}