package pkg

// TeleConfig struct that store your telegram config
type TeleConfig struct {
	apiToken string
}

// TeleBaseURL storing Telegram 
const TeleBaseURL = `https://api.telegram.org/bot`

// CreateTeleClient constructor function to create Tele Object client
func CreateTeleClient(apiToken string) *TeleConfig {
	return &TeleConfig{apiToken: apiToken}
}

// SetWebhook set webhook to telegram server
func (client *TeleConfig) SetWebhook(URL string) error {

}
