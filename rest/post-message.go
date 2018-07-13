package rest

import (
	"net/http"
	"encoding/json"

	"github.com/Silverhammertech/sms-svc/model"
	"github.com/Silverhammertech/sms-svc/sms"
	glog "github.com/Silverhammertech/sms-svc/log"
)

func HandlePostMessage(w http.ResponseWriter, r *http.Request) (err error){
	defer func() {
		defErr := r.Body.Close()
		if defErr != nil {
			glog.Fatal(defErr.Error())
		}
	}()

	var smsparams model.SendableSMSRequest
	err = json.NewDecoder(r.Body).Decode(&smsparams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Use Twilio, Make request
	smsSender := sms.Twilio{}
	smsResponseList, err := smsparams.SendSMS(smsSender)
	if err != nil {
		glog.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}


	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if smsResponseList != nil && len(smsResponseList) > 0 {
		bodyBytes, err := json.Marshal(smsResponseList)
		_, err = w.Write(bodyBytes)
		if err != nil {
			glog.Error(err.Error())
			http.Error(w, "sms: " + err.Error(), http.StatusInternalServerError)
		}
	}

	return err
}


