package sms

import (
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"errors"

	"github.com/Silverhammertech/sms-svc/model"
	"github.com/Silverhammertech/sms-svc/clients"
	glog "github.com/Silverhammertech/sms-svc/log"
	"github.com/Silverhammertech/sms-svc/config"
)

const (
	PENDING = "Pending"
	NOTIFIED = "Notified"
	FAILED = "Failed"
)
type Twilio struct {}

func (twilioSMS Twilio) Send(params model.SendableSMSRequest) (twilioResp []model.SMSResponse,twilioErr  error) {

	// set twilio params
	params.SID = config.TWILIO_SID_DEFAULT
	params.Token = config.TWILIO_AUTH_DEFAULT
	params.From = config.TWILIO_NUMBER_DEFAULT

	for _,recipient := range params.To {
		responseOut, err := processRecepient(recipient, params)
		if err != nil {
			glog.Error(err.Error())
			twilioErr = err
			return
		}
		twilioResp = append(twilioResp, responseOut)
	}

	return
}

func processRecepient(recipient string , twilioParams model.SendableSMSRequest) (retResponse model.SMSResponse, err error) {
	retResponse = model.SMSResponse{
		Recepient: recipient,
		Status: PENDING,
	}
	var httpResp *http.Response
	httpResp, err = clients.PostTwilioMessage(config.TWILIO_POST_URL_FORMAT, twilioParams, recipient)

	if err != nil {
		return
	}

	if httpResp == nil {
		err = errors.New("No http response from twilio request")
		return
	}

	switch  httpResp.StatusCode {
	case http.StatusCreated: // we will treat this as 200
		var goodResponse model.TwilioResponse
		err = json.Unmarshal(GetBytes(httpResp), &goodResponse)
		if err != nil {
			glog.Error(err.Error())
			fmt.Printf("error unmarshaling 200 response: ", err.Error())
		}
		retResponse.Status = NOTIFIED
		retResponse.StatusDescription = goodResponse.Status
		break;
	case http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized: // 401 occurs when your account expired.
		retResponse.Status = FAILED
		var badResponse model.TwilioErrorResponse
		err = json.Unmarshal(GetBytes(httpResp), &badResponse)
		if err != nil {
			glog.Error(err.Error())
			fmt.Printf("error unmarshaling 500 response: ", err.Error())
		}
		retResponse.StatusDescription = badResponse.Message
		err = errors.New(badResponse.Message)
		break;
	default:
		retResponse.Status = FAILED
		errMsg := "unknown status code response"
		retResponse.StatusDescription = errMsg
		err = errors.New(errMsg)
		glog.Error(string(GetBytes(httpResp)))
	}

	return
}

func GetBytes(httpResp *http.Response) (bytes []byte){
	if httpResp.Body != nil {
		defer func() {
			defErr := httpResp.Body.Close()
			if defErr != nil {
				glog.Fatal(defErr.Error())
			}
		}()
		var err error
		bytes, err = ioutil.ReadAll(httpResp.Body)

		if err != nil {
			glog.Fatal(err.Error())
			return
		}
	}

	return
}
