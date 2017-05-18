package main

import (
	. "DiscordBashBot/util"
	. "DiscordBashBot/vote"
	"flag"
	"fmt"
	"io/ioutil"

	. "github.com/bwmarrin/discordgo"
	. "gopkg.in/yaml.v2"
)

// Variables used for command line parameters
var (
	Token string
)

type Configuration struct {
	Token            string
	ValidChannelList []string
}

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	file, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		fmt.Println("Error opening config file: ", err)
	}
	configuration := new(Configuration)
	err = Unmarshal(file, configuration)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
	}

	// Create a new Discord session using the provided bot token.
	dg, err := New("Bot " + configuration.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(IsValidBotResponseChannel(BotMentionedMessageCreate(NonBotMessageCreate(botResponse)), configuration.ValidChannelList))
	dg.AddHandler(IsValidBotResponseChannel(BotMentionedMessageCreate(NonBotMessageCreate(HandleKickVote)), configuration.ValidChannelList))

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func botResponse(s *Session, m *MessageCreate) {

	s.ChannelMessageSend(m.ChannelID, "I heard you, "+m.Author.Username)
}
