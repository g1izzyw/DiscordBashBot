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

func ConstructVoteInfo(u *User, vote bool) *voteinfo {
	vi := new(voteinfo)
	vi.timePlaced = time.Now()
	vi.user = u
	vi.vote = vote

	return vi
}