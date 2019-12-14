package commands

import (
	"../config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

func HandleCommands(s *discordgo.Session, m *discordgo.Message, t time.Time) {
	cfg := config.GetBotConfig()
	msg := strings.TrimPrefix(m.Content, cfg.Prefix)
	cmd := strings.Split(msg, " ")[0]

	switch cmd {
	case "ping":
		HandlePingCommand(s, m)
		break
	case "pong":
		HandlePongCommand(s, m)
		break
	default:
		fmt.Printf("%s is not a command", cmd)
	}
}
