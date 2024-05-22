package handlers

import (
	"fmt"
	"log"

	"github.com/CyberDefenseEd/ChatProx/util"
	"github.com/bwmarrin/discordgo"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetTelegramUpdates(updates <-chan *tgbotapi.Update, config *util.Config, tgBot *tgbotapi.BotAPI, dg *discordgo.Session) {
	for update := range updates {
		if update.Message == nil || update.Message.From.IsBot {
			continue
		}
		if update.Message.Chat.ID == config.TelegramChatID {
			message := util.FormatTelegramMessage(update.Message.From.UserName, update.Message.Text)

			sendMediaToDiscord := func(fileID string, mediaType string) {
				fileURL, err := util.GetFileURL(tgBot, fileID)
				if err != nil {
					log.Printf("Failed to get %s URL: %v", mediaType, err)
					return
				}
				mediaMessage := fmt.Sprintf("@%s > sent a [%s](%s)", update.Message.From.UserName, mediaType, fileURL)
				_, err = dg.ChannelMessageSend(config.DiscordChannelID, mediaMessage)
				if err != nil {
					log.Printf("Failed to send %s to Discord: %v", mediaType, err)
				}
			}

			if update.Message.Sticker != nil {
				sendMediaToDiscord(update.Message.Sticker.FileID, "sticker")
			} else if update.Message.Animation != nil {
				sendMediaToDiscord(update.Message.Animation.FileID, "GIF")
			} else if update.Message.Photo != nil {
				photo := update.Message.Photo[len(update.Message.Photo)-1]
				sendMediaToDiscord(photo.FileID, "photo")
			} else if update.Message.Video != nil {
				sendMediaToDiscord(update.Message.Video.FileID, "video")
			} else {
				_, err := dg.ChannelMessageSend(config.DiscordChannelID, message)
				if err != nil {
					log.Printf("Failed to send message to Discord: %v", err)
				}
			}
		}
	}
}
