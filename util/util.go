package util

import (
	. "github.com/bwmarrin/discordgo"
)

func NonBotMessageCreate(handler func(s *Session, m *MessageCreate)) func(s *Session, m *MessageCreate) {
	return func(s *Session, m *MessageCreate) {
		if !m.Author.Bot {
			handler(s, m)
		}
	}
}

func BotMentionedMessageCreate(handler func(s *Session, m *MessageCreate)) func(s *Session, m *MessageCreate) {
	return func(s *Session, m *MessageCreate) {
		loggedInUser, _ := s.User("@me")

		if IsUserMentionedByUser(m, loggedInUser) {
			handler(s, m)
		}
	}
}

func IsUserMentionedByUser(m *MessageCreate, u *User) bool {
	for _, mentionedUser := range m.Mentions {
		if mentionedUser.ID == u.ID {
			return true
		}
	}
	return false
}

func IsUserMentionedByID(s *Session, m *MessageCreate, id string) bool {
	u, _ := s.User(id)
	return IsUserMentionedByUser(m, u)
}
