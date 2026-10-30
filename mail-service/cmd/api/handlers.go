package main

import (
	"fmt"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type MailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload MailMessage

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		fmt.Println(err)
		fmt.Println(requestPayload)
		app.errorJSON(w, err)
		return
	}

	var msg Message = Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		fmt.Println(err)
		fmt.Println(requestPayload)
		fmt.Println(msg)
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Message sent to " + msg.To,
	}

	app.writeJSON(w, payload, http.StatusAccepted)

}
