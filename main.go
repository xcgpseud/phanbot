package main

import (
	"./commands"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Config struct {
	Token  string
	Prefix string
}

var cfg Config

func init() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Setenv("PREFIX", cfg.Prefix)
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
