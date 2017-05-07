package vote

import (
	"fmt"
	"strings"
	"time"

	. "github.com/bwmarrin/discordgo"
)

var (
	kickPlayerVoteMap map[string]*kickplayervote
)

const kickString string = "KICK_PLAYER"

func init() {
	kickPlayerVoteMap = make(map[string]*kickplayervote)
}

type kickplayervote struct {
	playerToKick *User
	userVotes    *uservote
}

func ConstructKickPlayer(u *User) *kickplayervote {
	kpv := new(kickplayervote)

	kpv.userVotes = ConstructUserVote()
	kpv.playerToKick = u
	fmt.Println("here3")
	return kpv
}

func (voteToStart *kickplayervote) startVote() {

}

func (voteToUpdate *kickplayervote) addVote(u *User, vote bool) {
	voteToUpdate.userVotes.AddVoteToList(u, vote)
}

func (voteToCheck *kickplayervote) votePassed() bool {
	return (len(voteToCheck.userVotes.votes) > 0)
}

func HandleKickVote(s *Session, m *MessageCreate) {
	startIndex := strings.Index(m.Content, kickString)

	if startIndex == -1 || len(m.Mentions) != 2 {
		return
	}

	botUser, _ := s.User("@me")

	for _, mentionedUser := range m.Mentions {
		if mentionedUser.ID != botUser.ID {
			var kv kickplayervote

			if kv, ok := kickPlayerVoteMap[mentionedUser.ID]; ok {
				for _, info := range kv.userVotes.votes {
					if m.Author.ID == info.user.ID {
						// TODO: bot message to print -> can't vote twice
						return
					}
				}

			} else {
				fmt.Println("here1")
				kv := ConstructKickPlayer(mentionedUser)
				kv.startVote() //TODO implement
				kickPlayerVoteMap[mentionedUser.ID] = kv
			}
			fmt.Println("here4")
			kv.addVote(m.Author, true)

			if kv.votePassed() {

				var guildToBanFrom *Guild
				guilds, _ := s.UserGuilds()
				foundGuild := false
				for _, userGuild := range guilds {
					guild, _ := s.Guild(userGuild.ID)
					for _, member := range guild.Members {
						if member.User.ID == kv.playerToKick.ID {
							foundGuild = true
							guildToBanFrom = guild
							break
						}
					}
					if foundGuild {
						break
					}
				}
				if guildToBanFrom == nil {
					//TODO: print failure message of some kind
					return
				}
				s.GuildBanCreate(guildToBanFrom.ID, kv.playerToKick.ID, 1)
				time.Sleep(time.Second * 10)
				s.GuildBanDelete(guildToBanFrom.ID, kv.playerToKick.ID)
			}

			return
		}
	}

	//TODO: can't kick bot ..|..
}
