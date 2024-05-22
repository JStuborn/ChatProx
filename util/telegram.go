package util

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetFileURL(bot *tgbotapi.BotAPI, fileID string) (string, error) {
	fileConfig := tgbotapi.FileConfig{FileID: fileID}
	file, err := bot.GetFile(fileConfig)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath), nil
}
