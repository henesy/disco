package main

import (
	"strconv"
	"strings"
	"github.com/bwmarrin/discordgo"
)

//ParseForCommands parses input for Commands, returns message if no command specified, else return is empty
func ParseForCommands(line string) string {
	if len(line) < 2 {
		return line
	}
	switch line[:2] {
	case ":?":
		// Show help menu
		Msg(TextMsg, "Commands: ")
		Msg(TextMsg, "[:g] - Select guild")
		Msg(TextMsg, "[:p] - Select private message")
		Msg(TextMsg, "[:c] - Select guild channel")
		Msg(TextMsg, "[:c ?] - List guild channels")
		Msg(TextMsg, "[:c <num>] - Go directly to channel")
		Msg(TextMsg, "[:m <num>] - Display last <num> messages")
		Msg(TextMsg, "[:u <name>] - Change username")
	case ":g":
		SelectGuild()
		return ""
	case ":p":
		SelectPrivate()
		return ""
	case ":c":
		opts := strings.Split(line, " ")
		if len(opts) == 1 {
			SelectChannel()
			return ""
		}
		selectID := 0
		if opts[1] == "?" {
			for _, channel := range State.Channels {
				if channel.Type == 0 {
					Msg(TextMsg, "[%d] %s\n", selectID, channel.Name)
					selectID++
				}
			}
			return ""
		}
		selectMap := make(map[int]*discordgo.Channel)
		for _, channel := range State.Channels {
			if channel.Type == 0 {
				selectMap[selectID] = channel
				selectID++
			}
		}
		selection, err := strconv.Atoi(opts[1])
		if err != nil {
			Msg(ErrorMsg, "[:c] Argument Error: %s\n", err)
			return ""
		}
		if len(State.Channels) < selection || selection < 0 {
			Msg(ErrorMsg, "[:c] Argument Error: Out of bounds\n")
			return ""
		}
		channel := selectMap[selection]
		State.SetChannel(channel.ID)
		ShowContent()
		return ""
	case ":m":
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
		return line
	case ":u":
		session := State.Session
		user := session.User
		newName := strings.TrimPrefix(line, ":u ")
		_, err := State.Session.DiscordGo.UserUpdate(user.Email, session.Password, newName, user.Avatar, "")
		if err != nil {
			Msg(ErrorMsg, "[:u] Argument Error: %s\n", err)
		}
		return line
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
