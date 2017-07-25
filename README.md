# disco: [Discord](https://discord.gg) client for 9Front

Hacked up version of theboxmage's [discord-cli](https://github.com/theboxmage/discordcli)

JSON config is in `$home/lib//lib/disco.cfg` for setting password, should be made automatically after first run.

## Install

`mk deps`

`mk install`

## Problems

* You might have to run `mk` or `mk install` twice

* JSON

* Does not create accounts for you, this still needs to be done in a browser/app

* Does not support 2FA (Discord API explicitly does not allow this)

* Cutting and pasting text might break things

## Commands
Commands available in chat:

| Command       | Function         |
| ------------- |-------------|
| :q      | Quits discord-cli |
| :g      | Change listening Guild|
| :c      | Change listening Channel inside Guild |
| :m [n]      | Display last [n] messages: ex. `:m 2` displays last two messages |
| :p      | Pulls up the private channel menu |

## Notes

If you can connect to a channel and see messages, but yours aren't sending, check to make sure your e-mail address is verified.
