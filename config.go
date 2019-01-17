package main

import (
	"bitbucket.org/mischief/libauth"
	"bufio"
	"fmt"
	"github.com/mischief/ndb"
	"log"
	"os"
	"os/user"
	sc "strconv"
	"strings"
)

// Types of auth we can do ­ add handling for new modes to atoam()
type AuthModes int

const (
	Pass AuthModes = iota
	Factotum
	Unknown // Placeholder
)

//Configuration is a struct that contains all configuration fields
type Configuration struct {
	AuthMode      AuthModes
	Username      string
	LoadBacklog   bool
	Messages      int
	PromptChar    string
	TimestampChar string
	password      string
}

// Config is the global configuration of discord-cli
var Config Configuration

var ConfigPath string = "/lib/disco.ndb"

// Convert a string such as "true" into true
func atob(s string) bool {
	s = strings.ToLower(s)

	if s == "true" {
		return true
	}

	return false
}

// Convert a string such as "factotum" into factotum
func atoam(s string) AuthModes {
	s = strings.ToLower(s)

	if s == "pass" {
		return Pass
	}

	if s == "factotum" {
		return Factotum
	}

	return Unknown
}

//GetConfig retrieves configuration file from $home/lib/disco.ndb, if it doesn't exist it calls CreateConfig()
func GetConfig() {
	//Get User
Start:
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Get File
	file, err := os.Open(usr.HomeDir + ConfigPath)

	if err != nil {
		log.Println("Creating new config file")
		CreateConfig()
		goto Start
	}

	file.Close()

	// Decode File
	ndb, err := ndb.Open(usr.HomeDir + ConfigPath)
	if err != nil {
		log.Fatal("error: Could not ")
	}

	Config.Username = ndb.Search("username", "").Search("username")
	Config.AuthMode = atoam(ndb.Search("auth", "").Search("auth"))

	if Config.AuthMode == Pass {
		Config.password = ndb.Search("username", "").Search("password")
	}

	Config.LoadBacklog = atob(ndb.Search("loadbacklog", "").Search("loadbacklog"))
	Config.Messages, _ = sc.Atoi(ndb.Search("messages", "").Search("messages"))
	Config.PromptChar = ndb.Search("promptchar", "").Search("promptchar")
	Config.TimestampChar = ndb.Search("timestampchar", "").Search("timestampchar")

}

//CreateConfig creates folder inside $home and makes a new empty configuration file
func CreateConfig() {
	// Get User
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// Set Default values
	fmt.Print("Input your email: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	raw := fmt.Sprintf("auth=pass\nloadbacklog=true\nmessages=10\npromptchar=→\ntimestampchar=>\n\nusername=%s	password=\n", scan.Text())

	// Create File
	os.Mkdir(usr.HomeDir+"/lib", 0775)

	file, err := os.Create(usr.HomeDir + ConfigPath)

	if err != nil {
		log.Fatalln(err)
	}

	file.Chmod(0600)

	// PrintToFile
	_, err = file.Write([]byte(raw))

	if err != nil {
		log.Fatalln(err)
	}

	file.Close()
}

// CheckState checks the current state for essential missing information, errors will fail the program
func CheckState() {
	//Get User
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	if Config.Username == "" {
		log.Fatalln("Error: No Username Specified, please edit " + usr.HomeDir + ConfigPath)
	}

	// Check and handle password loading
	switch Config.AuthMode {
	case Factotum:
		// Acquire password from factotum
		userPwd, err := libauth.Getuserpasswd("proto=pass service=discord user=%s server=discordapp.com", Config.Username)

		if err != nil {
			// Factotum didn't get anything
			fmt.Fprintln(os.Stderr, "Warning: No success getting key from factotum, consider disabling it in ndb.")
			fmt.Fprintln(os.Stderr, "Libauth gave: ", err)
		} else {
			Config.password = userPwd.Password
		}

	case Unknown:
		log.Fatalln("Error: incorrect, or no, authmode specified via auth= config tuple. Consider: auth=(pass factotum).")

	default:
		// Password is already loaded in auth=pass mode
	}

	if Config.password == "" {
		log.Fatalln("Error: No password loaded, cannot auth. Are you missing a factotum key or password= tuple?")
	}
}
