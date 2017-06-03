package util

import (
	. "DiscordBashBot/configuration"
	"fmt"
	"time"

	. "github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis"
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

func IsValidBotResponseChannel(s *Session, m *MessageCreate, channelIdList []string) bool {
	messageChannel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel: %v", err)
	}
	for _, channelId := range channelIdList {
		if messageChannel.Name == channelId {
			return true
		}
	}

	return false
}

func HandleIfValidBotResponseChannel(handler func(s *Session, m *MessageCreate), channelIdList []string) func(s *Session, m *MessageCreate) {
	return func(s *Session, m *MessageCreate) {
		if IsValidBotResponseChannel(s, m, channelIdList) {
			handler(s, m)
		}
	}
}

func NotifyInvalidChannel(channelIdList []string) func(s *Session, m *MessageCreate) {
	return func(s *Session, m *MessageCreate) {
		messageChannel, _ := s.Channel(m.ChannelID)
		if !IsValidBotResponseChannel(s, m, channelIdList) {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%v this channel does not process command, try another channel", messageChannel.Name))
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

func WarningOutputByBot(t time.Duration, message string, s *Session, channelId string) {
	go func() {
		time.Sleep(t)
		s.ChannelMessageSend(channelId, message)
	}()
}

func GetRedisConnection() *redis.Client {
	redisOptions := new(redis.Options)

	if BotConfiguration.RedisPassword != nil {
		redisOptions.Password = *BotConfiguration.RedisPassword
	} else {
		redisOptions.Password = ""
	}
	if BotConfiguration.RedisUnixSocket != nil {
		redisOptions.Addr = *BotConfiguration.RedisUnixSocket
	} else {
		redisOptions.Addr = *BotConfiguration.RedisHost + ":" + *BotConfiguration.RedisPort
	}

	redisOptions.DB = 0

	connection := redis.NewClient(redisOptions)
	return connection
}
