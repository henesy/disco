# disco: [Discord](https://discord.gg) client for 9front

Fork of theboxmage's [discord-cli](https://github.com/theboxmage/discordcli)

JSON config is in `$home/lib/disco.cfg` for setting password, should be made automatically after first run.

## Install

```
go get github.com/bwmarrin/discordgo
go get github.com/gorilla/websocket
go get golang.org/x/crypto
go get bitbucket.org/henesy/disco
```

## Problems

* JSON

* PM's are temporarily disabled

* Does not create accounts for you, this still needs to be done in a browser/app

* Does not support 2FA (Discord API explicitly does not allow this)

## Commands
Commands available in chat:

| Command       | Function    |
| ------------- |-------------|
| :q        | Quits disco |
| :g        | Change listening Guild|
| :c [n|?]  | Change listening Channel inside Guild, or list channels |
| :m [n]    | Display last [n] messages: ex. `:m 2` displays last two messages |
| :p        | Pulls up the private channel menu |
| :n name   | Change nickname to `name` |

## Notes

If you can connect to a channel and see messages, but yours aren't sending, check to make sure your e-mail address is verified.

## FAQ

Q: What if `go get` doesn't work?

A: If you want to use `go get` on 9front to install disco and its dependencies (recommended) you should use [driusan's dgit](https://github.com/driusan/dgit) as `git`.

Q: What if I can't login because of a captcha error?

A: You'll need to sign in to Discord via the web app (thus solving a captcha) using a browser with html5/js. I recommend an http proxy such as [this](https://bitbucket.org/henesy/http-proxy).

Q: What if I get an error about signing in from a new location?

A: Discord has sent you an e-mail with a location confirmation link, click it, no js should be required.

