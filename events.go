package main

import (
	"strings"

	"bitbucket.org/henesy/disco/DiscordGo"
)

func removeReaction(s *discordgo.Session, r *discordgo.ReactionRemove) {

}

func newReaction(s *discordgo.Session, m *discordgo.ReactionAdd) {
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated user has access to.
func newMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Global Mentions
	Mention := "@" + State.Session.User.Username
	if strings.Contains(m.ContentWithMentionsReplaced(), Mention) {
		go Notify(m.Message)
	}

	// Do nothing when State is disabled
	if !State.Enabled {
		return
	}

	//State Messages
	if m.ChannelID == State.Channel.ID {
		State.AddMessage(m.Message)

		Messages := ReceivingMessageParser(m.Message)

		for _, Msg := range Messages {
			MessagePrint(m.Timestamp, m.Author.Username, Msg)
			//log.Printf("> %s > %s\n", UserName(m.Author.Username), Msg)
		}
	}
}
