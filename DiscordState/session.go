//Package DiscordState is an abstraction layer that gives proper structs and functions to get and set the current state of the cli server
package DiscordState

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

//!----- Session -----!//

//NewSession Creates a new Session
func NewSession(Username, Password string) *Session {
	Session := new(Session)
	Session.Username = Username
	Session.Password = Password

	return Session
}

//Start attaches a discordgo listener to the Sessions and fills it.
func (Session *Session) Start() error {

	fmt.Printf("Connecting...")

	dg, err := discordgo.New(Session.Username, Session.Password)
	if err != nil {
		return err
	}

	// Open the websocket and begin listening.
	dg.Open()

	//Retrieve GuildID's from current User
	//need index of Guilds[] rather than UserGuilds[] (maybe)
	Guilds, err := dg.UserGuilds(0, "", "")
	if err != nil {
		return err
	}

	Session.Guilds = Guilds

	Session.DiscordGo = dg

	Session.User, _ = Session.DiscordGo.User("@me")

	fmt.Printf(" PASSED!\n")

	return nil
}

//NewState (constructor) attaches a new state to the Guild inside a Session, and fills it.
func (Session *Session) NewState(GuildID string, MessageAmount int) (*State, error) {
	State := new(State)

	//Disable Event Handling
	State.Enabled = false

	//Set Session
	State.Session = Session

	//Set Guild
	for _, guildID := range Session.Guilds {
		if guildID.ID == GuildID {
			Guild, err := State.Session.DiscordGo.Guild(guildID.ID)
			if err != nil {
				return nil, err
			}

			State.Guild = Guild
		}
	}

	//Retrieve Members

	State.Members = make(map[string]*discordgo.Member)

	for _, Member := range State.Guild.Members {
		State.Members[Member.User.Username] = Member
	}

	//RetrieveMemberRoles
	State.MemberRole = make(map[string]*discordgo.Role)

	for _, Member := range State.Guild.Members {
		var MemberRole string

		if len(Member.Roles) > 0 {
			MemberRole = Member.Roles[0]
		} else {
			break
		}

		for _, Role := range State.Guild.Roles {
			if Role.ID == MemberRole {
				State.MemberRole[Member.User.Username] = Role
				break
			}
		}
	}

	//Set MessageAmount
	State.MessageAmount = MessageAmount

	//Init Messages
	State.Messages = []*discordgo.Message{}

	//Retrieve Channels

	State.Channels = State.Guild.Channels

	//Set User Channels
	//State.Chan = Session.DiscordGo.UserChannels()

	return State, nil
}

//Update updates the session, this reloads the Guild list
func (Session *Session) Update() error {
	UserGuilds, err := Session.DiscordGo.UserGuilds(0, "", "")
	if err != nil {
		return err
	}

	// this doesn't seem to actually work.
	Session.Guilds = UserGuilds
	return nil
}
