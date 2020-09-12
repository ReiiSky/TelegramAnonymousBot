package pkg

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// TeleConfig struct that store your telegram config
type TeleConfig struct {
	apiToken string
	baseURL  string
}

// TeleBaseURL storing Telegram base url
const TeleBaseURL = `https://api.telegram.org/bot`

// HTTPClient resty http client
var HTTPClient = resty.New()

// CreateTeleClient constructor function to create Tele Object client
func CreateTeleClient(apiToken string) *TeleConfig {
	return &TeleConfig{apiToken: apiToken, baseURL: fmt.Sprintf("%v%v", TeleBaseURL, apiToken)}
}

// SetWebhookResponse struct to storing webhook response
type SetWebhookResponse struct {
	Ok          bool   `json:"ok"`
	Result      bool   `json:"result"`
	Description string `json:"description"`
}

// SetWebhook set webhook to telegram server
func (client *TeleConfig) SetWebhook(URL string) (*SetWebhookResponse, error) {
	res, err := HTTPClient.R().
		SetQueryParam("url", URL).
		Post(fmt.Sprintf("%v/setWebHook", client.baseURL))

	response := SetWebhookResponse{}
	if err == nil {
		err = JSONUnmarshal(res.Body(), &response)
	}
	return &response, err
}
