package vote

import (
	"time"

	. "github.com/bwmarrin/discordgo"
)

type voteinfo struct {
	user       *User
	vote       bool
	timePlaced time.Time
}
