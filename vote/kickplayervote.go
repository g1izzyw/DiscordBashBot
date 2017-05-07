package vote

import (
	"strings"

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

	return kpv
}

func (voteToStart *kickplayervote) startVote() {

}

func (voteToStart *kickplayervote) addVote(u *User, vote bool) {
	kickplayervote.userVotes.AddVoteToList(u, vote)
}

func (voteToStart *kickplayervote) votePassed() bool {
	return (len(kickplayervote.userVotes.votes) > 3)
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
				kv := ConstructKickPlayer(mentionedUser)
				kv.startVote() //TODO implement
				kickPlayerVoteMap[mentionedUser.ID] = kv
			}

			kv.addVote(m.Author, true)

			if kv.votePassed() {
				//TODO actually kick the player
			}

			return
		}
	}

	//TODO: can't kick bot ..|..
}
