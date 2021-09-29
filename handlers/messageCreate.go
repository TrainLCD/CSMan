package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (h *Handlers) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	sliced := strings.Split(m.Content, " ")

	if sliced[0] == "!close" {
		if len(sliced) != 3 {
			return
		}

		tid := sliced[1]
		resolveReason := sliced[2]

		err := h.resolver.Close(tid, resolveReason)

		if err != nil {
			handleError(err, s, m)
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":tada:Ticket `%s` successfully resolved:tada:", tid))
	}

	if sliced[0] == "!reopen" {
		if len(sliced) != 2 {
			return
		}

		tid := sliced[1]

		err := h.resolver.Reopen(tid)

		if err != nil {
			handleError(err, s, m)
			return
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(":muscle:Ticket `%s` reopened!:muscle:", tid))
	}
}
