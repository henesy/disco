package main

import (
	"fmt"
	"log"
	"strconv"
)

//SelectPrivateMenu is a menu item that changes to a private channel
func SelectPrivateMenu() {

Start:

	Msg(InfoMsg, "Select a Member:\n")

	UserChannels, err := Session.DiscordGo.UserChannels()

	if err != nil {
		Msg(ErrorMsg, "No Private Channels\n")
	}

	UserMap := make(map[int]string)
	SelectID := 0

	for _, user := range UserChannels {
		UserMap[SelectID] = user.ID
		Msg(TextMsg, "[%d] %s\n", SelectID, UserChannels[SelectID].Recipient.Username)
		SelectID++
	}
	Msg(TextMsg, "[b] Extra Options\n")
	var response string
	fmt.Scanf("%s\n", &response)

	ResponseInteger, err := strconv.Atoi(response)

	if response == "b" {
	New:
		Msg(InfoMsg, "Extra Options:\n")
		Msg(TextMsg, "[n] Join New User Channel\n")
		Msg(TextMsg, "[d] Leave User Channel\n")
		Msg(TextMsg, "[b] Go Back\n")

		var response string
		fmt.Scanf("%s\n", &response)

		switch response {
		case "n":
			if State.Channel != nil {
				AddUserChannel()
				ShowEmptyContent()
				goto End
			} else {
				Msg(ErrorMsg, "Join a guild before attempting to join a user channel\n")
				goto New
			}
		case "d":
			SelectDeletePrivate()
			goto Start
		case "b":
			goto Start
		default:
			goto New
		}
	}

	if err != nil {
		Msg(ErrorMsg, "(GU) Conversion Error: %s\n", err)
		goto Start
	}

	if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
		Msg(ErrorMsg, "(GU) Error: ID is out of bounds\n")
		goto Start
	}

	State.Channel = UserChannels[ResponseInteger]
	ShowContent()
End:
}

//SelectDeletePrivateMenu deletes a user channel
func SelectDeletePrivateMenu() {

Start:

	Msg(InfoMsg, "Select a Member:\n")

	UserChannels, err := Session.DiscordGo.UserChannels()

	if err != nil {
		Msg(ErrorMsg, "No Private Channels\n")
	}

	UserMap := make(map[int]string)
	SelectID := 0

	for _, user := range UserChannels {
		UserMap[SelectID] = user.ID
		Msg(TextMsg, "[%d] %s\n", SelectID, UserChannels[SelectID].Recipient.Username)
		SelectID++
	}
	var response string
	fmt.Scanf("%s\n", &response)

	ResponseInteger, err := strconv.Atoi(response)
	if err != nil {
		Msg(ErrorMsg, "(GU) Conversion Error: %s\n", err)
		goto Start
	}

	if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
		Msg(ErrorMsg, "(GU) Error: ID is out of bounds\n")
		goto Start
	}

	Session.DiscordGo.ChannelDelete(UserChannels[ResponseInteger].ID)

}

//SelectGuildMenu is a menu item that creates a new State on basis of Guild selection
func SelectGuildMenu() {
	var err error

Start:

	Msg(InfoMsg, "Select a Guild:\n")

	SelectMap := make(map[int]string)
	SelectID := 0

	for _, guild := range Session.Guilds {
		SelectMap[SelectID] = guild.ID
		Msg(TextMsg, "[%d] %s\n", SelectID, guild.Name)
		SelectID++
	}
	Msg(TextMsg, "[b] Extra Options\n")
	Msg(TextMsg, "[p] Private Channels\n")

	var response string
	fmt.Scanf("%s\n", &response)
	ResponseInteger, err := strconv.Atoi(response)

	if response == "b" {
		ExtraGuildMenuOptions()
		goto Start
	}
	if response == "p" {
		if State != nil {
			SelectPrivate()
		} else {
			State, err = Session.NewState(SelectMap[0], Config.Messages)
			if err != nil {
				log.Fatal(err)
			}
			SelectPrivate()
		}
	} else {
		if err != nil {
			Msg(ErrorMsg, "(GU) Conversion Error: %s\n", err)
			goto Start
		}

		if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
			Msg(ErrorMsg, "(GU) Error: ID is out of bounds\n")
			goto Start
		}

		State, err = Session.NewState(SelectMap[ResponseInteger], Config.Messages)
		if err != nil {
			log.Fatal(err)
		}
	}
}

//SelectChannelMenu is a menu item that sets the current channel
func SelectChannelMenu() {
Start:
	Msg(InfoMsg, "Select a Channel:\n")

	SelectMap := make(map[int]string)
	SelectID := 0

	for _, channel := range State.Channels {
		if channel.Type == "text" {
			SelectMap[SelectID] = channel.ID
			Msg(TextMsg, "[%d] %s\n", SelectID, channel.Name)
			SelectID++
		}
	}
	Msg(TextMsg, "[b] Go Back\n")

	var response string
	fmt.Scanf("%s\n", &response)

	if response == "b" {
		SelectGuildMenu()
		goto Start
	}

	ResponseInteger, err := strconv.Atoi(response)
	if err != nil {
		Msg(ErrorMsg, "(CH) Conversion Error: %s\n", err)
		goto Start
	}

	if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
		Msg(ErrorMsg, "(CH) Error: ID is out of bound\n")
		goto Start
	}

	State.SetChannel(SelectMap[ResponseInteger])
}

//ExtraGuildMenuOptions prints and handles extra options for SelectGuildMenu
func ExtraGuildMenuOptions() {
Start:
	Msg(InfoMsg, "Extra Options:\n")
	Msg(TextMsg, "[n] Join New Server\n")
	Msg(TextMsg, "[d] Leave Server\n")
	Msg(TextMsg, "[o] Join Official discord-cli Server\n")
	Msg(TextMsg, "[b] Go Back\n")

	var response string
	fmt.Scanf("%s\n", &response)

	switch response {
	case "n":
	New:
		Msg(TextMsg, "Please input invite number ([b] back):\n")
		fmt.Scanf("%s\n", &response)
		if response == "b" {
			goto Start
		}
		Invite, err := Session.DiscordGo.Invite(response)
		if err != nil {
			Msg(ErrorMsg, "Invalid Invite\n")
			goto New
		}
		Msg(TextMsg, "Join %s ? [y/n]:\n", Invite.Guild.Name)
		fmt.Scanf("%s\n", &response)
		if response == "y" {
			Session.DiscordGo.InviteAccept(Invite.Code)
			err := Session.Update()
			if err != nil {
				Msg(ErrorMsg, "Session Update Failed: %s\n", err)
			}
		} else {
			goto Start
		}
	case "o":
		_, err := Session.DiscordGo.InviteAccept("0pXWCo5RQbVuFHDM")
		if err != nil {
			Msg(ErrorMsg, "Joining Official discord-cli Server failed\n")
			goto Start
		}
		Msg(InfoMsg, "Joined Official discord-cli Server!\n")
	case "d":
		LeaveServerMenu()
		goto Start
	default:
		return
	}

	return
}

//ExtraPrivateMenuOptions adds functionality to UserChannels.
func ExtraPrivateMenuOptions() {

	return
}

//AddUserChannelMenu takes a user from the current guild and adds them to a private message. WILL RETURN ERROR IF IN USER CHANNEL.
func AddUserChannelMenu() {
	if State.Channel.IsPrivate {
		Msg(ErrorMsg, "Currently in a user channel, move to a guild with :g\n")
	} else {
		SelectMap := make(map[int]string)
	Start:
		SelectID := 0
		for _, Member := range State.Members {
			SelectMap[SelectID] = Member.User.ID
			Msg(TextMsg, "[%d] %s\n", SelectID, Member.User.Username)
			SelectID++
		}
		var response string
		fmt.Scanf("%s\n", &response)

		if response == "b" {
			return
		}

		ResponseInteger, err := strconv.Atoi(response)
		if err != nil {
			Msg(ErrorMsg, "(CH) Conversion Error: %s\n", err)
			goto Start
		}

		if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
			Msg(ErrorMsg, "(CH) Error: ID is out of bound\n")
			goto Start
		}
		Chan, err := Session.DiscordGo.UserChannelCreate(SelectMap[ResponseInteger])

		if Chan.LastMessageID == "" {
			var firstMessage string
			fmt.Scanf("%s\n", &firstMessage)
			Session.DiscordGo.ChannelMessageSend(Chan.ID, "Test")
		}
		State.Channel = Chan
	}
}

//LeaveServerMenu is a copy of SelectGuildMenu that leaves instead of selects
func LeaveServerMenu() {
	var err error

Start:

	Msg(InfoMsg, "Leave a Guild:\n")

	SelectMap := make(map[int]string)
	SelectID := 0

	for _, guild := range Session.Guilds {
		SelectMap[SelectID] = guild.ID
		Msg(TextMsg, "[%d] %s\n", SelectID, guild.Name)
		SelectID++
	}
	Msg(TextMsg, "[b] Go Back\n")

	var response string
	fmt.Scanf("%s\n", &response)

	if response == "b" {
		return
	}

	ResponseInteger, err := strconv.Atoi(response)
	if err != nil {
		Msg(ErrorMsg, "(GUD) Conversion Error: %s\n", err)
		goto Start
	}

	if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
		Msg(ErrorMsg, "(GUD) Error: ID is out of bounds\n")
		goto Start
	}

	Guild, err := Session.DiscordGo.Guild(SelectMap[ResponseInteger])
	if err != nil {
		Msg(ErrorMsg, "(GUD) Unknown Error: %s\n", err)
		goto Start
	}

	Msg(TextMsg, "Leave %s ? [y/n]:\n", Guild.Name)
	fmt.Scanf("%s\n", &response)
	if response == "y" {
		Session.DiscordGo.GuildLeave(Guild.ID)
		err := Session.Update()
		if err != nil {
			Msg(ErrorMsg, "Session Update Failed: %s\n", err)
		}
	} else {
		goto Start
	}

}
