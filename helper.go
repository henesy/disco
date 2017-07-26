package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
	"fmt"
	"bitbucket.org/henesy/disco/DiscordGo"
	"bufio"
)

//HexColor is a struct gives RGB values
type HexColor struct {
	R     int
	G     int
	B     int
}

//Msg is a composition of Color.New printf functions
func Msg(MsgType, format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

//Header simply prints a header containing state/session information
func Header() {
	Msg(InfoMsg, "Welcome, %s!\n\n", State.Session.User.Username)
	if State.Channel.IsPrivate {
		Msg(InfoMsg, "Channel: %s\n", State.Channel.Recipient.Username)
	} else {
		Msg(InfoMsg, "Guild: %s, Channel: %s\n", State.Guild.Name, State.Channel.Name)
	}
}

//ReceivingMessageParser parses receiving message for mentions, images and MultiLine and returns string array
func ReceivingMessageParser(m *discordgo.Message) []string {
	Message := m.ContentWithMentionsReplaced()

	//Parse images
	for _, Attachment := range m.Attachments {
		Message = Message + " " + Attachment.URL
	}

	// MultiLine comment parsing
	Messages := strings.Split(Message, "\n")

	return Messages
}

//PrintMessages prints amount of Messages to CLI
func PrintMessages(Amount int) {
	for Key, m := range State.Messages {
		if Key >= len(State.Messages)-Amount {
			Messages := ReceivingMessageParser(m)

			for _, Msg := range Messages {
				//log.Printf("> %s > %s\n", UserName(m.Author.Username), Msg)
				MessagePrint(m.Timestamp, m.Author.Username, Msg)

			}
		}
	}
}

//Notify uses Notify-Send from libnotify to send a notification when a mention arrives.
func Notify(m *discordgo.Message) {
	Channel, err := State.Session.DiscordGo.Channel(m.ChannelID)
	if err != nil {
		Msg(ErrorMsg, "(NOT) Channel Error: %s\n", err)
	}
	Guild, err := State.Session.DiscordGo.Guild(Channel.GuildID)
	if err != nil {
		Msg(ErrorMsg, "(NOT) Guild Error: %s\n", err)
	}
	Title := "@" + m.Author.Username + " : " + Guild.Name + "/" + Channel.Name
	cmd := exec.Command("notify-send", Title, m.ContentWithMentionsReplaced())
	err = cmd.Start()
	if err != nil {
		Msg(ErrorMsg, "(NOT) Check if libnotify is installed, or disable notifications.\n")
	}

}

//MessagePrint prints one correctly formatted Message to stdout
func MessagePrint(Time, Username, Content string) {
	//var Color color.Attribute
	TimeStamp, _ := time.Parse(time.RFC3339, Time)
	LocalTime := TimeStamp.Local().Format("2006/01/02 15:04:05")

	log.SetFlags(0)
	log.Printf("%s > %s > %s\n", LocalTime, Username, Content)
	log.SetFlags(log.LstdFlags)
}

func dis(a, b int) float64 {
	return float64((a - b) * (a - b))
}

func Rawon() error {
	consctl, err := os.OpenFile("/dev/consctl", os.O_WRONLY, 0200)
	if err != nil {
		/* not on Plan 9 */
		fmt.Println("\nNot running on Plan 9")
		return err
	}
	
	rawon := []byte("rawon")
	_, err = consctl.Write(rawon)
	if err != nil {
		return err
	}
	
	return nil
}

func RawOff() error {
	consctl, err := os.OpenFile("/dev/consctl", os.O_WRONLY, 0200)
	if err != nil {
		/* not on Plan 9 */
		return err
	}
	
	rawoff := []byte("rawoff")
	_, err = consctl.Write(rawoff)
	if err != nil {
		return err
	}
	
	return nil
}

func GetPass() string {
	cons, err := os.OpenFile("/dev/cons", os.O_RDWR, 0600)
	if err != nil {
		fmt.Println("Failed to open /dev/cons")
	}
	consScan := bufio.NewScanner(cons)
	consScan.Scan()
	return consScan.Text()
}
