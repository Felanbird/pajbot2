package modules

import (
	"log"
	"strings"

	"github.com/pajlada/pajbot2/bot"
	"github.com/pajlada/pajbot2/command"
	"github.com/pajlada/pajbot2/common"
)

/*
Command xD
*/
type Command struct {
	commands []command.Command
}

/*
some ideas:
not sure how this would work but id like to keep things simple

$(user1) = usersource displayname
$(sender) = sender of msg
$(1) = raw arg from msg
$(bot) = bot
$(channel) = channel

users have methods, for example $(sender.points):
	.low = lowercase
	.points
	.level
	etc ...

bot has:
	bot.uptime
	bot.version
	etc...

channel has:
	channel.uptime
	channel.title
	channel.game
	channel.subs
	channel.name
	etc...

!points would look like this :
	"$(user1) has $(user1.points) points."

!uptime:
	"$(sender), $(channel.name) has been online for $(channel.uptime) PogChamp"

*/

// Ensure the module implements the interface properly
var _ Module = (*Command)(nil)

// Init initializes something
func (module *Command) Init() {
	xdCommand := command.Command{
		Trigger:  "!test",
		Response: ":P",
	}
	module.commands = append(module.commands, xdCommand)
}

// Check xD
func (module *Command) Check(b *bot.Bot, msg *common.Msg, action *bot.Action) error {
	m := strings.Split(msg.Message, " ")
	trigger := strings.ToLower(m[0])
	for _, command := range module.commands {
		log.Println(trigger, command.Trigger)
		if trigger == command.Trigger {
			// TODO: Get response first, and skip if the response is nil or something of that sort
			action.Response = command.GetResponse()
			action.Stop = true
			return nil
		}
	}
	if trigger == "!xd" {
		action.Response = "pajaSWA"
		action.Stop = true
	}
	if trigger == "!quit" && msg.User.Name == "nuuls" {
		b.Quit <- "ayy lmao something bad happened xD"
	}
	return nil
}