package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/CyberDefenseEd/ChatProx/handlers"
	"github.com/CyberDefenseEd/ChatProx/util"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type HandlerArgs struct {
	TGbot   *tgbotapi.BotAPI
	Config  *util.Config
	Session *discordgo.Session
}

func loadConfig() (*util.Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &util.Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	tgBot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}
	tgBot.Debug = true

	dg, err := discordgo.New("Bot " + config.DiscordBotToken)
	if err != nil {
		log.Panic(err)
	}

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		handlers.ListenToMessageCreation(tgBot, config, s, m)
	})

	u := tgbotapi.NewUpdate(0)
	u.Timeout = config.TelegramTimeout

	updates := make(chan *tgbotapi.Update)
	go func() {
		for update := range tgBot.GetUpdatesChan(u) {
			updates <- &update
		}
	}()

	go handlers.GetTelegramUpdates(updates, config, tgBot, dg)

	// Open Discord websocket connection
	err = dg.Open()
	if err != nil {
		log.Panic(err)
	}
	defer dg.Close()

	tgMsg := tgbotapi.NewMessage(config.TelegramChatID, "Discord -> Telegram proxy is now online!")
	_, err = tgBot.Send(tgMsg)
	if err != nil {
		log.Printf("Failed to send message to Telegram: %v", err)
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	select {}
}
