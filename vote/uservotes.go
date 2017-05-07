package vote

import (
	"time"

	. "github.com/bwmarrin/discordgo"
)

type IUserVote interface {
	votePassed() bool
	startVote()
	addVote(*User, bool)
}

type uservote struct {
	starttime time.Time
	votes     []*voteinfo
}

func ConstructUserVote() *uservote {
	uv := new(uservote)
	uv.starttime = time.Now()
	uv.votes = []*voteinfo{}
	return uv
}

func (uv *uservote) YesGreaterThanNo() bool {
	yesCount := uv.GetYesVoteCount()
	noCount := uv.GetNoVoteCount()

	return yesCount > noCount
}

func (uv *uservote) GetYesVoteCount() int {
	yesCount := 0
	for _, vote := range uv.votes {
		if vote.vote {
			yesCount++
		}
	}
	return yesCount
}

func (uv *uservote) GetNoVoteCount() int {
	noCount := 0
	for _, vote := range uv.votes {
		if !vote.vote {
			noCount++
		}
	}
	return noCount
}

func (uv *uservote) AddVoteToList(u *User, vote bool) {
	newUserVote := new(voteinfo)
	newUserVote.timePlaced = time.Now()
	newUserVote.user = u
	newUserVote.vote = vote
	uv.votes = append(uv.votes, newUserVote)
}
