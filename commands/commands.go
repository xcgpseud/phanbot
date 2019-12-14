package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
	"time"
)

func HandleCommands(s *discordgo.Session, m *discordgo.Message, t time.Time) {
	msg := strings.TrimPrefix(m.Content, os.Getenv("PREFIX"))
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
