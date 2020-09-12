package pkg_test

import (
	"testing"

	. "github.com/Satssuki/tele-anon-bot-queue/pkg"
)

const token = `1397462892:AAEcpmYz4dy9duqfB7EGsJlULGyBcXd5saQ`

func TestTeleConfig_SetWebhook(t *testing.T) {
	client := CreateTeleClient(token)
	status, err := client.SetWebhook("https://07924b3bd3de.ngrok.io/telehook")
	if err != nil {
		t.Errorf("error exist with message: %v", err.Error())
	}
	if !status.Ok {
		t.Error("OK is not true")
	}

	if !status.Result {
		t.Error("Result is not true")
	}

	if len(status.Description) < 4 {
		t.Errorf("Description is not 'Webhook is already set' but '%v'", status.Description)
	}
}
