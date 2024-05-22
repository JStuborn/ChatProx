package handlers

import (
	"fmt"
	"log"

	"github.com/CyberDefenseEd/ChatProx/util"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/bwmarrin/discordgo"
)

func ListenToMessageCreation(tgBot *tgbotapi.BotAPI, config *util.Config, s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if m.ChannelID == config.DiscordChannelID {
		message := fmt.Sprintf("@%s > %s", m.Author.Username, m.Content)
		tgMsg := tgbotapi.NewMessage(config.TelegramChatID, message)
		_, err := tgBot.Send(tgMsg)
		if err != nil {
			log.Printf("Failed to send message to Telegram: %v", err)
		}
	}
}
