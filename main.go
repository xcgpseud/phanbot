package main

import (
	"./commands"
	"./config"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var cfg *config.BotConfig

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	cfg = config.GetBotConfig()
}

func main() {
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("Error creating Discord Session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s == nil || m == nil || m.Author.ID == s.State.User.ID || m.Content == "" {
		return
	}

	prefixlen := len(cfg.Prefix)
	if m.Content[:prefixlen] == cfg.Prefix {
		commands.HandleCommands(s, m.Message, time.Now())
	}
}
