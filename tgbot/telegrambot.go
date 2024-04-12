package tgbot

import (
	"bytes"
	"rent-car/api/models"
	"encoding/json"

	"net/http"
)

// -1001981481970
// 7081684820:AAGWSquPDVDYEO8pOwOT95UjhIAOtbNEfhE

func  SendMessageTG(a models.SendMessage)(error) {

	botToken := ""
	chatID := ""

	messageBytes, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	message := string(messageBytes)

	payload := struct {
		ChatID string `json:"chat_id"`
		Text   string `json:"text"`
	}{
		ChatID: chatID,
		Text:   message,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := "https://api.telegram.org/bot" + botToken + "/sendMessage"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
return nil
}

