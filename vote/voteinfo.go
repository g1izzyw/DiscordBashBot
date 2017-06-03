package vote

import (
	. "DiscordBashBot/util"
	"fmt"
	"strconv"
	"time"

	. "github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis"
)

type voteinfo struct {
	voteId     int
	user       *User
	vote       bool
	timePlaced time.Time
}

func (vote *voteinfo) addToRedis() string {
	connection := GetRedisConnection()
	voteInfoId := strconv.FormatInt(connection.Incr("VoteInfoCounter").Val(), 10)
	err := connection.HSet("VoteInfo:"+voteInfoId, "VoteId", vote)
	if err != nil {
		fmt.Println("failed to set voteinfo")
	}
	connection.HSet("VoteInfo:"+voteInfoId, "UserId", vote.user.ID)
	connection.HSet("VoteInfo:"+voteInfoId, "Vote", vote.vote)
	connection.HSet("VoteInfo:"+voteInfoId, "TimePlaced", vote.timePlaced.Unix())
	connection.LPush("VoteInfos", voteInfoId)
	connection.BgSave()
	connection.Close()
	return voteInfoId
}

func loadVotes(c *redis.Client) []*voteinfo {

	return nil
}

func ConstructVoteInfo(u *User, vote bool) *voteinfo {
	vi := new(voteinfo)
	vi.timePlaced = time.Now()
	vi.user = u
	vi.vote = vote

	return vi
}
