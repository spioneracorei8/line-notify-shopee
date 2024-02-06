package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spioneracorei8/linebot-notification-shopping/data"
)

func main() {
	app := fiber.New()

	app.Post("/send-notification", func(c *fiber.Ctx) error {
		// Define Line Notify API endpoint
		notifyURL := "https://notify-api.line.me/api/notify"
		data.MessagesData()
		messages := data.Data.Message
		imagesURL := data.Data.ImageURL
		go func() error {
			for i := 0; i < len(messages); i++ {
				msg := messages[i]
				img := imagesURL[i]

				data := url.Values{}
				data.Set("message", msg)
				data.Set("imageThumbnail", img)
				data.Set("imageFullsize", img)
				req, err := http.NewRequest("POST", notifyURL, strings.NewReader(data.Encode()))
				time.Sleep(time.Second * 4)
				if err != nil {
					return err
				}
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				req.Header.Set("Authorization", "Bearer gVrfGt4CQGuBwCCCH3UCxBMBuPZFXpXQzZTqCwsUaQV")

				// Send request to Line Notify API
				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					return err
				}

				// Check response status code
				if resp.StatusCode != http.StatusOK {
					return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
				}
			}
			time.Sleep(6 * time.Second)
			return nil
		}()

		return c.SendString("Notification sent successfully")
	})

	app.Listen(":8999")
}
