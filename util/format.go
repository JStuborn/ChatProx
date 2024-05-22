package util

import "fmt"

func FormatDiscordMessage(username, content string) string {
	return fmt.Sprintf("@%s > %s", username, content)
}

func FormatTelegramMessage(username, content string) string {
	return fmt.Sprintf("[@%s](<https://t.me/%s>) > %s", username, username, content)
}
