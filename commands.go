package main

import (
	"regexp"
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
	case "s/":
		r := regexp.MustCompile(`s\/((\w|\s)+)\/((\w|\s|)+)\/`)
		n := r.FindStringSubmatchIndex(line)
		if len(n) < 1 {
			return ""
		}
		msg := State.Messages
		for i := len(msg)-1; i >= 0; i-- {
			if (msg[i].ChannelID != State.Channel.ID && msg[i].GuildID != State.Guild.ID) {
				break
			}
			if msg[i].Author.ID != Session.User.ID {
				break
			}
			if strings.Contains(msg[i].Content, line[n[2]:n[3]]) {
				cmd, err := regexp.Compile("(" + line[n[2]:n[3]] + ")")
				if err != nil {
					Msg(ErrorMsg, "%s - invalid regex\n", line)
				}
				rep := cmd.ReplaceAllString(msg[i].Content, "")
				if len(line) == 5 {
					rep = cmd.ReplaceAllString(msg[i].Content, line[n[4]:n[5]])
				}
				data := discordgo.NewMessageEdit(msg[i].ChannelID, msg[i].ID)
				data = data.SetContent(rep)
				_, err = State.Session.DiscordGo.ChannelMessageEditComplex(data)
				if err != nil {
					Msg(ErrorMsg, "%s\n", err)
					return ""
				}
				Msg(TextMsg, "%s -> %s\n", msg[i].Content, rep)
				return ""
			}
		}
		return line
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
		return ""
	case ":u":
		session := State.Session
		user := session.User
		newName := strings.TrimPrefix(line, ":u ")
		_, err := State.Session.DiscordGo.UserUpdate(user.Email, session.Password, newName, user.Avatar, "")
		if err != nil {
			Msg(ErrorMsg, "[:u] Argument Error: %s\n", err)
			return ""
		}
		Msg(TextMsg, "name -> %s\n", newName)
		return ""
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
