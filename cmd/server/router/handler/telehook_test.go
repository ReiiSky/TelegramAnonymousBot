package handler_test

import (
	"testing"

	"github.com/go-resty/resty/v2"
)

var httpClient = resty.New()

func TestTelehookNormalMessage(t *testing.T) {
	response, err := httpClient.R().
		SetBody(`{
			"update_id": 678730628,
			"message": {
			  "message_id": 7070,
			  "from": {
				"id": 1272147630,
				"is_bot": false,
				"first_name": "Satsuki",
				"last_name": "Reikaa",
				"username": "Satssuki",
				"language_code": "en"
			  },
			  "chat": {
				"id": 1272147630,
				"first_name": "Satsuki",
				"last_name": "Reikaa",
				"username": "Satssuki",
				"type": "private"
			  },
			  "date": 1600067388,
			  "reply_to_message": {
				"message_id": 7064,
				"from": {
				  "id": 1397462892,
				  "is_bot": true,
				  "first_name": "Anonymous Bot",
				  "username": "anonymouf_bot"
				},
				"chat": {
				  "id": 1272147630,
				  "first_name": "Satsuki",
				  "last_name": "Reikaa",
				  "username": "Satssuki",
				  "type": "private"
				},
				"date": 1599927126,
				"text": "Hello world, Programmed to work and not to feel"
			  },
			  "text": "yawywayaw"
			}
		  }`).Post("http://localhost:6007/telehook")
	if err != nil {
		t.Error("error is not nil")
	}

	if response.StatusCode() != 201 {
		t.Errorf("message status code error")
	}
}

func TestTelehookSearchCommand(t *testing.T) {
	response, err := httpClient.R().
		SetBody(`{
			"update_id": 678730628,
			"message": {
			  "message_id": 7070,
			  "from": {
				"id": 1272147630,
				"is_bot": false,
				"first_name": "Satsuki",
				"last_name": "Reikaa",
				"username": "Satssuki",
				"language_code": "en"
			  },
			  "chat": {
				"id": 1272147630,
				"first_name": "Satsuki",
				"last_name": "Reikaa",
				"username": "Satssuki",
				"type": "private"
			  },
			  "date": 1600067388,
			  "text": "/search"
			}
		  }`).Post("http://localhost:6007/telehook")
	if err != nil {
		t.Error("error is not nil")
	}

	if response.StatusCode() != 204 {
		t.Errorf("message status code error")
	}
}