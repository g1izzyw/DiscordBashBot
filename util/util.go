package util

import (
	"github.com/bwmarrin/discordgo"
)

func NonBotMessageCreate(handler func(s *discordgo.Session, m *discordgo.MessageCreate)) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate){
		if (!m.Author.Bot) {
			handler(s, m)
		}
	}
}

func BotMentionedMessageCreate(handler func(s *discordgo.Session, m *discordgo.MessageCreate)) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		loggedInUser, _ := s.User("@me")
		for _, mentionedUser := range m.Mentions {
			if (mentionedUser.ID == loggedInUser.ID) {
				handler(s, m)
				break
			}
		}
	}


}