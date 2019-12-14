package commands

import "github.com/bwmarrin/discordgo"

func HandlePongCommand(s *discordgo.Session, m *discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "I'm not gonna say Ping, fuck off.")
}

func HandlePingCommand(s *discordgo.Session, m *discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "FuCkInG pOnG1!1!")
}
