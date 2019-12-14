package config

import (
	"os"
	"strconv"
	"sync"
)

type BotConfig struct {
	Token        string
	Prefix       string
	LogChannelId int64
}

var botInstance *BotConfig
var once sync.Once

func GetBotConfig() *BotConfig {
	once.Do(func() {
		botInstance = &BotConfig{
			Token:        getEnv("BOT_TOKEN", ""),
			Prefix:       getEnv("BOT_PREFIX", "!"),
			LogChannelId: getEnvAsInt("BOT_LOG_CHANNEL_ID", -1),
		}
	})
	return botInstance
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return int64(value)
	}
	return defaultVal
}
