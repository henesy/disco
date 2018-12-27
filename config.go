package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	//	"golang.org/x/crypto/ssh/terminal"
	"bitbucket.org/mischief/libauth"
)

//Configuration is a struct that contains all configuration fields
type Configuration struct {
	Username       string `json:"username"`
	MessageDefault bool   `json:"messagedefault"`
	Messages       int    `json:"messages"`
	CompletionChar string `json:"completionchar"`
	TimeCompChar   string `json:"timecompchar"`
	password       string
}

// Config is the global configuration of discord-cli
var Config Configuration

//GetConfig retrieves configuration file from $home/lib/disco.cfg, if it doesn't exist it calls CreateConfig()
func GetConfig() {
	//Get User
Start:
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	//Get File
	file, err := os.Open(usr.HomeDir + "/lib/disco.cfg")
	if err != nil {
		log.Println("Creating new config file")
		CreateConfig()
		goto Start
	}

	//Decode File
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		log.Println("Failed to decode configuration file")
		log.Fatalf("Error: %s", err)
	}
}

//CreateConfig creates folder inside $home and makes a new empty configuration file
func CreateConfig() {
	//Get User
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	var EmptyStruct Configuration
	//Set Default values
	fmt.Print("Input your email: ")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	EmptyStruct.Username = scan.Text()
	EmptyStruct.Messages = 10
	EmptyStruct.MessageDefault = true
	EmptyStruct.CompletionChar = ">"
	EmptyStruct.TimeCompChar = ">"

	//Create File
	os.Mkdir(usr.HomeDir+"/lib", 0775)
	file, err := os.Create(usr.HomeDir + "/lib/disco.cfg")
	if err != nil {
		log.Fatalln(err)
	}
	file.Chmod(0600)

	//Marshall EmptyStruct
	raw, err := json.Marshal(EmptyStruct)
	if err != nil {
		log.Fatalln(err)
	}

	//PrintToFile
	_, err = file.Write(raw)
	if err != nil {
		log.Fatalln(err)
	}

	file.Close()
}

//CheckState checks the current state for essential missing information, errors will fail the program
func CheckState() {
	//Get User
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	if Config.Username == "" {
		log.Fatalln("No Username Specified, please edit " + usr.HomeDir + "/lib/disco.cfg")
	}
	userPwd, err := libauth.Getuserpasswd("proto=pass service=discord user=%s server=discordapp.com", Config.Username)
	if err != nil {
		log.Fatal(err)
	}
	Config.password = userPwd.Password
}
