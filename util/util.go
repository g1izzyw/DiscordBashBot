package util

import (
	"fmt"
	"time"

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

func IsValidBotResponseChannel(handler func(s *Session, m *MessageCreate), channelIdList []string) func(s *Session, m *MessageCreate) {
	return func(s *Session, m *MessageCreate) {

		messageChannel, _ := s.Channel(m.ChannelID)

		for _, channelId := range channelIdList {
			channel, _ := s.Channel(channelId)
			if messageChannel.Name == channel.Name {
				handler(s, m)
				//return
			}
		}

		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v this channel does not process command, try another channel", messageChannel))
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

func WarningOutputByBot(t time.Duration, message string, s *Session, channelId string) {
	go func() {
		time.Sleep(t)
		s.ChannelMessageSend(channelId, message)
	}()
}
