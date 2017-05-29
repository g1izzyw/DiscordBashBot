package group

import {
	. "github.com/bwmarrin/discordgo"
}

type groupinfo struct {
	groupName string
	groupMembers []*User
}

func ConstructGroupInfo() *groupinfo {
	gi := new(groupinfo)
	gi.groupMembers = []*User{}

	return gi
}