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

// TextMessageBuilder function to construct text message form
type TextMessageBuilder struct {
	params map[string]string
}

// TextMessageBuilder ..
func (client *TeleConfig) TextMessageBuilder() *TextMessageBuilder {
	return &TextMessageBuilder{params: make(map[string]string)}
}

// ChatID function to set chat id
func (builder *TextMessageBuilder) ChatID(id string) *TextMessageBuilder {
	builder.params["chat_id"] = id
	return builder
}

// ReplyTo function to set chat id
func (builder *TextMessageBuilder) ReplyTo(id string) *TextMessageBuilder {
	builder.params["reply_to_message_id"] = id
	return builder
}

// Content function to set chat id
func (builder *TextMessageBuilder) Content(text string) *TextMessageBuilder {
	builder.params["text"] = text
	return builder
}

// DisableNotification function to disable notification
func (builder *TextMessageBuilder) DisableNotification(disable bool) *TextMessageBuilder {
	builder.params["disable_notification"] = fmt.Sprint(disable)
	return builder
}

// Push function to set chat id
func (builder *TextMessageBuilder) Push(client *TeleConfig) error {
	res, err := HTTPClient.R().
		SetQueryParams(builder.params).
		Get(fmt.Sprintf("%v/sendMessage", client.baseURL))
	fmt.Print(res)
	return err
}
