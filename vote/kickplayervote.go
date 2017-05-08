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
	fmt.Printf("%d uservote length", len(voteToCheck.userVotes.votes))
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
			var kv *kickplayervote
			var ok bool

			if kv, ok = kickPlayerVoteMap[mentionedUser.ID]; ok {
				for _, info := range kv.userVotes.votes {
					if m.Author.ID == info.user.ID {
						// TODO: bot message to print -> can't vote twice
						return
					}
				}

			} else {
				kv = ConstructKickPlayer(mentionedUser)
				kv.startVote() //TODO implement
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
					fmt.Println("here5")
					//TODO: print failure message of some kind
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
				time.Sleep(time.Second * 60)
				s.GuildBanDelete(guildToBanFrom.ID, kv.playerToKick.ID)
			}

			return
		}
	}

	//TODO: can't kick bot ..|..
}
