// services/messaging_service.go
package services

import (
	"log"
	"time"

	"github/batuhanzorbeyzengin/insider-messaging-system/database"
	"github/batuhanzorbeyzengin/insider-messaging-system/database/models"
	"github/batuhanzorbeyzengin/insider-messaging-system/utils"
	"sync"
)

const (
	messageBatchSize     = 2
	messagingInterval    = 2 * time.Minute
	messagingLoopTimeout = 5 * time.Second
)

var (
	messagingLoop sync.WaitGroup
	messagingStop = make(chan struct{})
)

// StartMessagingService starts the messaging service and sends the first batch of messages immediately.
func StartMessagingService() {
	messagingLoop.Add(1)
	go handleMessagingLoop()
	sendMessages()
}

// StopMessagingService stops the messaging service and waits for the messaging loop to finish.
func StopMessagingService() {
	close(messagingStop)
	messagingLoop.Wait()
}

func handleMessagingLoop() {
	defer messagingLoop.Done()
	ticker := time.NewTicker(messagingInterval)
	defer ticker.Stop()
	for {
		select {
		case <-messagingStop:
			return
		case <-ticker.C:
			sendMessages()
		}
	}
}

func sendMessages() {
	var messages []models.Message
	database.DB.Where("sent_status = ?", false).Limit(messageBatchSize).Find(&messages)

	for _, message := range messages {
		if !message.SentStatus {
			webhookResponse, err := utils.SendMessageToWebhook(message)
			if err != nil {
				log.Printf("Error sending message to webhook: %v", err)
				continue
			}

			sentAt := time.Now().Format("2006-01-02 15:04:05")

			err = utils.CacheMessageDetails(webhookResponse.MessageID, sentAt)
			if err != nil {
				log.Printf("error caching message details: %v", err)
				continue
			}

			database.DB.Model(&message).Update("sent_status", true)
		}
	}
}
