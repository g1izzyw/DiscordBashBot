package group

import {
	"strings"
	"fmt"
	. "github.com/go-redis/redis"
}

func HandleAddToGroup(s *Session, m *MessageCreate) {
	startIndex := strings.Index(m.Content, addToGroupString)

	if startIndex == -1 || len(m.Mentions) != 2 {
		return
	}

	commandContents := Split(m.Content, " ")


	if len(commandContents) != 1 {
		
	}

	botUser, _ := s.User("@me")
}

