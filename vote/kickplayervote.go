package vote

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	. "DiscordBashBot/util"

	. "github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis"
)

var (
	kickPlayerVoteMap map[string]*kickplayervote
)

//const kickString string = "KICK_PLAYER"

func init() {
	kickPlayerVoteMap = make(map[string]*kickplayervote)
}

type kickplayervote struct {
	playerToKick *User
	userVotes    *uservote
	channelID    string
}

func ConstructKickPlayer(u *User, cId string) *kickplayervote {
	kpv := new(kickplayervote)

	kpv.userVotes = ConstructUserVote()
	kpv.playerToKick = u
	kpv.channelID = cId

	return kpv
}

func (newVote *kickplayervote) addToRedis() string {
	fmt.Println("Adding kick player vote to redis")
	connection := GetRedisConnection()
	voteId := strconv.FormatInt(connection.Incr("KickPlayerVoteCounter").Val(), 10)
	connection.LPush("KickPlayerVotes", voteId)
	connection.HSet("KickPlayerVote:"+voteId, "Id", voteId)
	connection.HSet("KickPlayerVote:"+voteId, "UserId", newVote.playerToKick.ID)
	connection.HSet("KickPlayerVote:"+voteId, "ChannelId", newVote.channelID)
	userVotesId := newVote.userVotes.addToRedis()
	connection.HSet("KickPlayerVote:"+voteId, "UserVoteId", userVotesId)
	connection.BgSave()
	connection.Close()
	return voteId
}

func (voteToStart *kickplayervote) startVote(s *Session) {
	//TODO: put into redis
	s.ChannelMessageSend(voteToStart.channelID, fmt.Sprintf("You have 2 min to kick player %v", voteToStart.playerToKick.Username))

	WarningOutputByBot(time.Minute+(30*time.Second), fmt.Sprintf("You have 30 second remaining for your vote to kick %v", voteToStart.playerToKick.Username), s, voteToStart.channelID)

	voteToStart.addToRedis()

	go func() {
		time.Sleep(2 * time.Minute)
		if voteToStart, ok := kickPlayerVoteMap[voteToStart.playerToKick.ID]; ok {
			delete(kickPlayerVoteMap, voteToStart.playerToKick.ID)
			s.ChannelMessageSend(voteToStart.channelID, fmt.Sprintf("Vote has expired for kick player %v", voteToStart.playerToKick.Username))
		}
	}()
}

func (voteToUpdate *kickplayervote) addVote(u *User, vote bool) {
	voteToUpdate.userVotes.AddVoteToList(u, vote)
}

func (voteToCheck *kickplayervote) votePassed() bool {
	fmt.Printf("%d uservote length", len(voteToCheck.userVotes.votes))
	return (len(voteToCheck.userVotes.votes) >= 3)
}

func HandleKickVote(s *Session, m *MessageCreate) {
	startIndex := strings.Index(m.Content, KICK_STRING)

	if startIndex == -1 || len(m.Mentions) != 2 {
		return
	}

	botUser, _ := s.User("@me")

	for _, mentionedUser := range m.Mentions {
		if mentionedUser.ID != botUser.ID {
			var kv *kickplayervote
			var ok bool

			if kv, ok = kickPlayerVoteMap[mentionedUser.ID]; ok {
				for _, info := range kv.userVotes.votes {
					if m.Author.ID == info.user.ID {
						s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("<@%v> can't vote to kick again for player %v", m.Author.ID, kv.playerToKick.Username))
						return
					}
				}

			} else {
				kv = ConstructKickPlayer(mentionedUser, m.ChannelID)
				kv.startVote(s)
				kickPlayerVoteMap[mentionedUser.ID] = kv
			}
			fmt.Println("here2")
			kv.addVote(m.Author, true)

			if kv.votePassed() {
				fmt.Println("here1")

				//var guildToBanFrom *Guild
				channel, _ := s.Channel(m.ChannelID)
				guildToBanFrom, _ := s.Guild(channel.GuildID)

				if guildToBanFrom == nil {
					fmt.Println("Guild to ban from does not exist")
					return
				}
				fmt.Println("here6")

				err := s.GuildBanCreate(guildToBanFrom.ID, kv.playerToKick.ID, 1)
				if err != nil {
					fmt.Print("user: ")
					fmt.Println(kv.playerToKick.Username)
					fmt.Print("guild: ")
					fmt.Println(guildToBanFrom.Name)
					fmt.Print("Failed to ban: ")
					fmt.Println(err)
				}
				delete(kickPlayerVoteMap, kv.playerToKick.ID) // remove vote from the map, when player is kicked

				time.Sleep(time.Second * 60)

				s.GuildBanDelete(guildToBanFrom.ID, kv.playerToKick.ID)
			}

			return
		}
	}
	// can't kick the bot
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("..|.. <@%v>", m.Author.ID))
}

//TODO: read and load all on going votes
func LoadOngoingKickPlayerVotes() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
