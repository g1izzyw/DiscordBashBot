package group

import {
	. "github.com/go-redis/redis"
	. "github.com/bwmarrin/discordgo"
}


func HandleAddToGroup(s *Session, m *MessageCreate) {
	startIndex := strings.Index(m.Content, messageGroupString)

	
}
