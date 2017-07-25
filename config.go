package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"

//	"golang.org/x/crypto/ssh/terminal"
)

//Configuration is a struct that contains all configuration fields
type Configuration struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	MessageDefault bool   `json:"messagedefault"`
	Messages       int    `json:"messages"`
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
	fmt.Println("Input your email")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()

	EmptyStruct.Username = scan.Text()
	fmt.Println("Input your password")
	//password, err := terminal.ReadPassword(0)
	
	plan9 := true
	/* Plan 9 raw mode for rio */
	consctl, err := os.OpenFile("/dev/consctl", os.O_WRONLY, 0200)
	if err != nil {
		/* not on Plan 9 */
		plan9 = false
	}
	
	password := "";
	
	if plan9 {
		rawon := make([]byte, 5)
		rawon = []byte("rawon")
		_, err = consctl.Write(rawon)
		if err != nil {
			fmt.Println("Failed to set rawon")
		} else {
			//cons, err := os.OpenFile("/dev/cons", os.O_RDWR, 0600)
			//if err != nil {
			//	fmt.Println("Failed to open /dev/cons")
			//}
			//scan := bufio.NewScanner(cons)
			scan.Scan()
			password = scan.Text()
			rawoff := make([]byte, 6)
			rawoff = []byte("rawoff")
			_, err = consctl.Write(rawoff)
			if err != nil {
				fmt.Println("Failed to set rawoff")
			}
		}
	}
	
	EmptyStruct.Password = string(password)
	EmptyStruct.Messages = 10
	EmptyStruct.MessageDefault = true

	//Create File
	file, err := os.Create(usr.HomeDir + "/lib/disco.cfg")
	if err != nil {
		log.Fatalln(err)
	}

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

	if Config.Password == "" {
		log.Fatalln("No Password Specified, please edit " + usr.HomeDir + "/lib/disco.cfg")
	}

}
