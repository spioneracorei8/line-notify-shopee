package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/spioneracorei8/linebot-notification-shopping/data"
)

func main() {
	// Set Thailand time zone
	thailandLocation, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		fmt.Println("Error loading Thailand time zone:", err.Error())
		return
	}

	// Get current time in Thailand
	now := time.Now().In(thailandLocation)

	targetTime1 := time.Date(now.Year(), now.Month(), now.Day(), 11, 58, 0, 0, thailandLocation)
	if now.After(targetTime1) {
		targetTime1 = targetTime1.Add(24 * time.Hour) // Schedule for next day if the target time has passed for today
	}

	targetTime2 := time.Date(now.Year(), now.Month(), now.Day(), 23, 58, 0, 0, thailandLocation)
	if now.After(targetTime2) {
		targetTime2 = targetTime2.Add(24 * time.Hour) // Schedule for next day if the target time has passed for today
	}

	// Calculate duration until the next scheduled time
	duration1 := targetTime1.Sub(now)
	duration2 := targetTime2.Sub(now)
	fmt.Println("Starting and Waiting??!!")
	// Start a goroutine to run the task at the scheduled time
	go func() {
		// Sleep until the scheduled time
		time.Sleep(duration1)

		// Task to send the POST request
		sendLineNotifyMessage("Task at 12:00 PM")
	}()
	go func() {
		// Sleep until the scheduled time
		time.Sleep(duration2)

		// Task to send the POST request
		sendLineNotifyMessage("Task at 00:00 AM")
	}()

	// Keep the main goroutine alive
	select {}
}

func sendLineNotifyMessage(taskName string) {
	// Define Line Notify API endpoint
	fmt.Println("Start Send message!")
	fmt.Printf("%s \n", taskName)
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
			time.Sleep(time.Second * 5)
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
		time.Sleep(7 * time.Second)
		return nil
	}()

}
