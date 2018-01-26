# disco: [Discord](https://discord.gg) client for 9Front

Hacked up version of theboxmage's [discord-cli](https://github.com/theboxmage/discordcli)

JSON config is in `$home/lib/disco.cfg` for setting password, should be made automatically after first run.

## Install

```
go get github.com/bwmarrin/discordgo
go get github.com/gorilla/websocket
go get golang.org/x/crypto
go get bitbucket.org/henesy/disco
```

## Problems

* You might have to run `mk` or `mk install` twice

* JSON

* PM's are temporarily disabled

* Does not create accounts for you, this still needs to be done in a browser/app

* Does not support 2FA (Discord API explicitly does not allow this)

## Commands
Commands available in chat:

| Command       | Function         |
| ------------- |-------------|
| :q      | Quits disco |
| :g      | Change listening Guild|
| :c      | Change listening Channel inside Guild |
| :m [n]      | Display last [n] messages: ex. `:m 2` displays last two messages |
| :p      | Pulls up the private channel menu |

## Notes

If you can connect to a channel and see messages, but yours aren't sending, check to make sure your e-mail address is verified.

