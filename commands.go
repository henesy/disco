package main

import (
	"strconv"
	"strings"
)

//ParseForCommands parses input for Commands, returns message if no command specified, else return is empty
func ParseForCommands(line string) string {
	//One Key Commands
	switch line {
	case ":g":
		SelectGuild()
		line = ""
	case ":c":
		SelectChannel()
		line = ""
	case ":p":
		SelectPrivate()
		line = ""
	default:
		// Nothing
	}

	//Argument Commands
	if strings.HasPrefix(line, ":m") {
		AmountStr := strings.Split(line, " ")
		if len(AmountStr) < 2 {
			Msg(ErrorMsg, "[:m] No Arguments \n")
			return ""
		}

		Amount, err := strconv.Atoi(AmountStr[1])
		if err != nil {
			Msg(ErrorMsg, "[:m] Argument Error: %s \n", err)
			return ""
		}

		Msg(InfoMsg, "Printing last %d messages!\n", Amount)
		State.RetrieveMessages(Amount)
		PrintMessages(Amount)
		line = ""
	}
	if strings.HasPrefix(line, ":u") {
		session := State.Session
		user := session.User
		newName := strings.TrimPrefix(line, ":u ")
		_, err := State.Session.DiscordGo.UserUpdate(user.Email, session.Password, newName, user.Avatar, "")
		if err != nil {
			Msg(ErrorMsg, "[:u] Argument Error: %s\n", err)
		}
		line = ""
	}
		
	return line
}

//SelectGuild selects a new Guild
func SelectGuild() {
	State.Enabled = false
	SelectGuildMenu()
	// Segfaults would happen here
	SelectChannelMenu()
	State.Enabled = true
	ShowContent()
}

//AddUserChannel moves a user to a private channel with another user.
func AddUserChannel() {
	State.Enabled = false
	AddUserChannelMenu()
	State.Enabled = true
	ShowContent()
}

//SelectChannel selects a new Channel
func SelectChannel() {
	State.Enabled = false
	SelectChannelMenu()
	State.Enabled = true
	ShowContent()
}

//SelectPrivate a private channel
func SelectPrivate() {
	State.Enabled = false
	SelectPrivateMenu()
	State.Enabled = true
	ShowContent()
}

//SelectDeletePrivate a private channel
func SelectDeletePrivate() {
	State.Enabled = false
	SelectDeletePrivateMenu()
	State.Enabled = true
	ShowContent()
}
