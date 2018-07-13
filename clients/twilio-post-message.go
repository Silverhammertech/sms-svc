package clients

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"

	"github.com/Silverhammertech/sms-svc/model"
)


func PostTwilioMessage(twilioPosturlFormat string,smsParams model.SendableSMSRequest, recepient string) (*http.Response, error){
	urlStr := fmt.Sprintf(twilioPosturlFormat, smsParams.SID)

	// Build out the data for our message
	v := url.Values{}
	v.Set("To",recepient)
	v.Set("From",smsParams.From)
	v.Set("Body",smsParams.Message)

	rb := *strings.NewReader(v.Encode())

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(smsParams.SID, smsParams.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	// Make request
	return client.Do(req)

}
