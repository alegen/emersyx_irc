package main

import(
    "errors"
    "emersyx.net/emersyx_apis/emcomapi"
    "emersyx.net/emersyx_apis/emircapi"
    irc "github.com/fluffle/goirc/client"
)

// The IRCBot struct defines the implementation of an IRC receptor and resource. The struct implements the
// emircapi.IRCBot and emcomapi.Receptor interfaces.
type IRCBot struct {
    api *irc.Conn
    cfg *irc.Config
    identifier string
    messages chan emcomapi.Event
}

// NewIRCBot creates a new emircapi.IRCBot instance and applies to configuration specified in the arguments.
func NewIRCBot(options ...func (emircapi.IRCBot) error) (emircapi.IRCBot, error) {
    bot := new(IRCBot)

    // create the messages channel
    bot.messages = make(chan emcomapi.Event)

    // create a Config object for the underlying library
    bot.cfg = irc.NewConfig("placeholder")

    // override several default values from the underlying library
	bot.cfg.Me.Ident = "emersyx"
	bot.cfg.Me.Name = "emersyx"
    bot.cfg.Version = "emersyx"
    bot.cfg.SSL = false
	bot.cfg.QuitMessage = "bye"

    // standard function for generating new nicks
    bot.cfg.NewNick = func(n string) string { return n + "^" }

    // apply the configuration options received as arguments
    // this configuration method is inspired from
    // https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
    for _, option := range options {
        err := option(bot)
        if err != nil {
            return nil, err
        }
    }

    // check if the mandatory identifier value has been set
    if len(bot.identifier) == 0 {
        return nil, errors.New("identifier option has not been set")
    }

    // create the underlying Conn object
    bot.api = irc.Client(bot.cfg)

    // initialize callbacks
    bot.initCallbacks()

    return bot, nil
}

func (bot *IRCBot) initCallbacks() {
    bot.api.HandleFunc( irc.PRIVMSG, channelCallback(bot) )
    bot.api.HandleFunc( irc.JOIN,    channelCallback(bot) )
    bot.api.HandleFunc( irc.QUIT,    channelCallback(bot) )
    bot.api.HandleFunc( irc.PART,    channelCallback(bot) )
}
