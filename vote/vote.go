package vote

import (
	"time"

	. "github.com/bwmarrin/discordgo"
	. "github.com/go-redis/redis"
)

type voteinfo struct {
	voteId     int
	user       *User
	vote       bool
	timePlaced time.Time
	isOngoing  bool
}

func loadVotes(c *Client) []*voteinfo {

}
