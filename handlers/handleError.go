package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func handleError(err error, s *discordgo.Session, m *discordgo.MessageCreate) {
	if err != nil {
		log.Println(err)
		s.ChannelMessageSend(m.ChannelID, "An error occrred:weary:")
	}
}
