# ChatProx

ChatProx is a lightweight chat proxy designed to facilitate communication between Discord and Telegram. With ChatProx, you can seamlessly bridge conversations across platforms, allowing users from one platform to interact with users on the other.

## Features

- **Bi-Directional Communication**: ChatProx forwards messages between Discord and Telegram in real-time, ensuring that users on both platforms stay in sync.
  
- **Configurable Settings**: Easily configure ChatProx using the provided `config.json` file. Set up your chat IDs and bot tokens to enable communication between your Discord and Telegram channels.

- **Simple Setup**: ChatProx is designed to be easy to set up and use. Just follow the steps outlined in the [How to Use](#how-to-use) section below to get started.

## How to Use

1. **Set Up Your Configuration**:

   - Copy the provided `config.example.json` to `config.json`.
   - Edit `config.json` to include your chat IDs and bot tokens for Discord and Telegram.

2. **Build the Executable**:
```
go build
```
3. **Run the Proxy**:
```
./ChatProx
```

## Configuration

The `config.json` file contains the following settings:

- `DiscordChannelID`: The ID of the Discord channel you want to proxy messages from.
- `TelegramChatID`: The ID of the Telegram chat where messages will be forwarded to.
- `DiscordBotToken`: The bot token for your Discord bot.
- `TelegramBotToken`: The bot token for your Telegram bot.
- `TelegramTimeout`: Timeout value for Telegram API requests (in seconds).

Ensure that you provide valid IDs and tokens in the configuration file for ChatProx to function correctly.

## Contributing

Contributions are welcome! If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
