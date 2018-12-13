// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"bitbucket.org/henesy/disco/DiscordState"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

//Global Message Types
const (
	ErrorMsg  = "Error"
	InfoMsg   = "Info"
	HeaderMsg = "Head"
	TextMsg   = "Text"
)

//Version is current version const
const Version = "1.0"

//Session is global Session
var Session *DiscordState.Session

//State is global State
var State *DiscordState.State

//UserChannels is global User Channels

//MsgType is a string containing global message type
type MsgType string

var hideTimeStamp = flag.Bool("t", false, "Hide timestamps in channel log")
var enableNotify = flag.Bool("n", false, "Enable notifications")
var notifyFlag = flag.String("w", "10,10,260,90", "Dimensions to pass through to statusmsg")

func main() {

	flag.Parse()
	if flag.Lookup("h") != nil {
		flag.Usage()
		os.Exit(1)
	}
	if flag.Lookup("w") != nil {
		*notifyFlag = fmt.Sprintf("-w %s", *notifyFlag)
	}
	//Initialize Config
	GetConfig()
	CheckState()
	Msg(HeaderMsg, "disco version: %s\n\n", Version)

	//NewSession
	Session = DiscordState.NewSession(Config.Username, Config.Password) //Please don't abuse
	err := Session.Start()
	if err != nil {
		log.Println("Session Failed")
		log.Fatalln(err)
	}
	//Attach New Window
	InitWindow()

	//Attach Even Handlers
	State.Session.DiscordGo.AddHandler(newMessage)
	//State.Session.DiscordGo.AddHandler(newReaction)

	//defer rl.Close()
	log.SetOutput(os.Stderr) // let "log" write to l.Stderr instead of os.Stderr
	State.Session.DiscordGo.UpdateStatus(0, "Plan 9")

	//Start Listening
	reader := bufio.NewReader(os.Stdin)
	for {
		//fmt.Print("> ")
		//line, _ := rl.Readline()
		line, _ := reader.ReadString('\n')
		line = line[:len(line)-1]

		//QUIT
		if line == ":q" {
			break
		}

		//Parse Commands
		line = ParseForCommands(line)

		line = ParseForMentions(line)

		if line != "" {
			_, err := State.Session.DiscordGo.ChannelMessageSend(State.Channel.ID, line)
			if err != nil {
				fmt.Print("Error: ", err, "\n")
			}
		}
	}

	return
}

//InitWindow creates a New CLI Window
func InitWindow() {
	SelectGuildMenu()
	if State.Channel == nil {
		SelectChannelMenu()
	}
	State.Enabled = true
	ShowContent()
}

//ShowContent shows default Channel content
func ShowContent() {
	Header()
	if Config.MessageDefault {
		State.RetrieveMessages(Config.Messages)
		PrintMessages(Config.Messages)
	}
}

//ShowEmptyContent shows an empty channel
func ShowEmptyContent() {
	Header()

}

//ParseForMentions parses input string for mentions
func ParseForMentions(line string) string {
	r, err := regexp.Compile("@\\w+")
	if err != nil {
		Msg(ErrorMsg, "Regex Error: ", err)
	}

	lineByte := r.ReplaceAllStringFunc(line, ReplaceMentions)

	return lineByte
}

//ReplaceMentions replaces mentions to ID 
func ReplaceMentions(input string) string {
	// Check for guild members that match
	channel := State.Guild.Members
	for _, member := range channel {
		if member.Nick == input[1:] {
			return member.User.Mention()
		}
		if strings.HasPrefix(member.User.Username, input[1:]) {
			return member.User.Mention()
		}
	}
	// Walk all PM channels
	userChannels, err := Session.DiscordGo.UserChannels()
	if err != nil {
		return input
	}
	for _, channel := range userChannels {
		for _, recipient := range channel.Recipients {
			if strings.HasPrefix(input[1:], recipient.Username) {
				fmt.Println("usermatch")
				return recipient.Mention()
			}
		}
	}
	return input
}

//Parse for guild-specific emoji
func ParseForEmoji(line string) string {
	r, err := regexp.Compile("<(:\\w+:)[0-9]+>")
	if err != nil {
		Msg(ErrorMsg, "Regex Error: ", err)
	}
	return r.ReplaceAllString(line, "$1")
}
