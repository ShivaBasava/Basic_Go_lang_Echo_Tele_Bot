package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	//tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// Config represents the structure of the config.json file
type Config struct {
	TelegramTokenB string
}

func readConfig() Config {
	//Opening the JSON file
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)

	if err != nil {
		log.Panic(err)
	}

	return configuration
}

func main() {
	config := readConfig()

	bot, err := tgbotapi.NewBotAPI(config.TelegramTokenB)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		// Ignore any non-Message Updates
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Reply
		text := fmt.Sprintf("~:> %s", update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

		bot.Send(msg)
	}
}
