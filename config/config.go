package config

var (
	DEBUG bool
	DEFAULT_PORT string
	BASE_PATH string

	TWILIO_POST_URL_FORMAT  string
	TWILIO_NUMBER_DEFAULT string
	TWILIO_AUTH_DEFAULT string
	TWILIO_SID_DEFAULT string
)

func init() {

	// TODO get config from somewhere

	// Server config
	DEBUG = true
	DEFAULT_PORT = "8080"
	BASE_PATH = "/api/v1/"

	TWILIO_POST_URL_FORMAT  = "https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json"
	TWILIO_NUMBER_DEFAULT = "+15005550006" // default
	TWILIO_AUTH_DEFAULT = "b1aa25ec776f34aca841857906c7c739" // default
	TWILIO_SID_DEFAULT = "ACda58eb5f25fd588f09cff21f2d0bedde" // default
}
