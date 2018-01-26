// This file provides a basic "quick start" example of using the Discordgo
// package to connect to Discord using the New() helper function.
package main

import (
	"log"
	"regexp"
	"bitbucket.org/henesy/disco/DiscordState"
	"fmt"
	"bufio"
	"os"
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

func main() {
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
			_ ,err := State.Session.DiscordGo.ChannelMessageSend(State.Channel.ID, line)
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

//ShowContent shows defaulth Channel content
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
	r, err := regexp.Compile("\\@\\w+")
	if err != nil {
		Msg(ErrorMsg, "Regex Error: ", err)
	}

	lineByte := r.ReplaceAllFunc([]byte(line), ReplaceMentions)

	return string(lineByte[:])
}

//ReplaceMentions replaces mentions to ID
func ReplaceMentions(input []byte) []byte {
	var OutputString string

	SizeByte := len(input)
	InputString := string(input[1:SizeByte])

	if Member, ok := State.Members[InputString]; ok {
		OutputString = "<@" + Member.User.ID + ">"
	} else {
		OutputString = "@" + InputString
	}
	return []byte(OutputString)
}
