// utils/redis.go
package utils

import (
	"fmt"

	"github/batuhanzorbeyzengin/insider-messaging-system/configs"

	"github.com/go-redis/redis"
)

const messageCacheKey = "sent_messages"

var RedisClient *redis.Client

func InitRedis(config configs.RedisConfig) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

// CacheMessageDetails caches the message details in a Redis hash.
func CacheMessageDetails(messageID, sentAt string) error {
	err := RedisClient.HSet(messageCacheKey, messageID, sentAt).Err()
	if err != nil {
		return err
	}

	return nil
}

type CachedMessageDetails struct {
	MessageID string `json:"messageId"`
	SendTime  string `json:"sendTime"`
}

// GetCachedMessageDetails retrieves the cached message details from the Redis hash.
func GetCachedMessageDetails() ([]CachedMessageDetails, error) {
	cachedMessages, err := RedisClient.HGetAll(messageCacheKey).Result()
	if err != nil {
		return nil, err
	}

	var messageDetails []CachedMessageDetails
	for messageID, sendTime := range cachedMessages {
		messageDetails = append(messageDetails, CachedMessageDetails{
			MessageID: messageID,
			SendTime:  sendTime,
		})
	}

	return messageDetails, nil
}
