package uservote

import {
	"fmt"
	"time"
	"logger"
}

type uservote  struct {
	starttime int
	votes []bool
}

func constructUserVote() *uservote {
	uv = new(uservote)
	uv.starttime = time.Now()
	uv.votes = make([5]string)

	return uv
}


func (uv *uservote) yesGreateThanNo() bool {
	yesCount := uv.getYesVoteCount()
	noCount := uv.getNoVoteCount()

	return yesCount > noCount
}

func (uv *uservote) getYesVoteCount() int {
	yesCount := 0
	for i := 0; i < uv.votes.len(); i++ {
		if uv[i] {
			yesCount++
		}
	}
	return yesCount
} 

func (uv *uservote) getNoVoteCount() int {
	noCount := 0
	for i := 0; i < uv.votes.len(); i++ {
		if !uv[i] {
			noCount++
		}
	}
	return noCount
}