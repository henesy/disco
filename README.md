# disco: [Discord](https://discord.gg) client for 9front

Fork of theboxmage's [discord-cli](https://github.com/theboxmage/discordcli).

Ndb config is in `$home/lib/disco.ndb` for setting password, should be made automatically after first run. Alternatively, you may use factotum. 

## Install

### Dependencies

- 9fans.net/go/plan9
- github.com/Plan9-Archive/libauth
- github.com/mischief/ndb
- golang.org/x/crypto
- github.com/gorilla/websocket
- github.com/bwmarrin/discordgo

### Installation

```
% go get github.com/henesy/disco
% go install github.com/henesy/disco
```

## Usage

```
% disco -h
Usage of disco:
  -n	Enable notifications
  -t	Hide timestamps in channel log
  -w string
    	Dimensions to pass through to statusmsg (default "10,10,260,90")
```

## Commands

Commands available in chat:

| Command       | Function    |
| ------------- |-------------|
| :q        | Quits disco |
| :g        | Change listening Guild |
| :c [n ?]  | Change listening Channel inside Guild, or list channels |
| :m [n]    | Display last [n] messages: ex. `:m 2` displays last two messages |
| :p        | Pulls up the private channel menu |
| :n name   | Change nickname to `name` |
| :!        | Print current server information |
| :?        | List the available commands |

You can regex the last message sent using a format such as:

	s/forsynth/forsyth/

## Config

A basic `$home/lib/disco.ndb` looks something like:

```
auth=pass
loadbacklog=true
messages=10
promptchar=→
timestampchar=>

username=coolperson@mycooldomain.com	password=somepassword1
```

Note that the auth= tuple accepts

	auth=factotum

for authentication using a factotum key and will ignore the password= tuple.

If used, the factotum key should resemble something to the effect of:

	proto=pass server=discordapp.com service=discord user=youremail@domain.com !password=hunter2

## Notes

If you can connect to a channel and see messages, but yours aren't sending, check to make sure your e-mail address is verified.

## FAQ

Q: What if `go get` doesn't work?

A: If you want to use `go get` on 9front to install disco and its dependencies (recommended) you should use [driusan's dgit](https://github.com/driusan/dgit) as `git`.

Q: What if I can't login because of a captcha error?

A: You'll need to sign in to Discord via the web app (thus solving a captcha) using a browser with html5/js. I recommend an http proxy such as [this](https://github.com/henesy/http-proxy) in conjunction with a system with such a browser..

Q: What if I get an error about signing in from a new location?

A: Discord has sent you an e-mail with a location confirmation link, click it, no js should be required.

## Problems

* Does not create accounts for you, this still needs to be done in a browser/app
* Does not support 2FA (Discord API explicitly does not allow this)

