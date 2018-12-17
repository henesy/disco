package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func removeReaction(s *discordgo.Session, r *discordgo.MessageReactionRemove) {

}

func newReaction(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {

	//State Messages -- don't notify when we're in the channel
	if m.ChannelID == State.Channel.ID && State.Enabled {
		State.AddMessage(m.Message)

		Messages := ReceivingMessageParser(m.Message)

		for _, Msg := range Messages {
			MessagePrint(string(m.Timestamp), m.Author.Username, Msg)
			//log.Printf("> %s > %s\n", UserName(m.Author.Username), Msg)
		}
		return
	}

	//Global Mentions
	Mention := "@" + State.Session.User.Username
	if strings.Contains(m.ContentWithMentionsReplaced(), Mention) {
		go Notify(m.Message)
		return
	}
	DMs, err := Session.DiscordGo.UserChannels()
	if err != nil {
		// No DMs 
		return
	}
	for _, channel := range DMs {
		if m.ChannelID == channel.ID {
			go Notify(m.Message)
			return
		}
	}
}
