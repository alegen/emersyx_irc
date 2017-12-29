package main

import(
    "emersyx.net/emersyx_apis/emircapi"
    irc "github.com/fluffle/goirc/client"
)

// The IRCBot struct defines the implementation of an IRC receptor and resource. The struct implements the
// emircapi.IRCBot and emcomapi.Receptor interfaces.
type IRCBot struct {
    api *irc.Conn
    cfg *irc.Config
    Messages chan emircapi.Message
}

// This function creates a new emircapi.IRCBot instance and applies to configuration specified in the arguments.
func NewIRCBot(options ...func (emircapi.IRCBot) error) (emircapi.IRCBot, error) {
    bot := new(IRCBot)
    bot.Messages = make(chan emircapi.Message)

    // create a Config object for the underlying library
    bot.cfg = irc.NewConfig("placeholder")

    // override several default values from the underlying library
	bot.cfg.Me.Ident = "emersyx"
	bot.cfg.Me.Name = "emersyx"
    bot.cfg.Version = "emersyx"
    bot.cfg.SSL = true
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

    // create the underlying Conn object
    bot.api = irc.Client(bot.cfg)

    // initialize callbacks
    bot.initCallbacks()

    return bot, nil
}

func (bot *IRCBot) initCallbacks() {
    bot.api.HandleFunc(irc.PRIVMSG, makeSendToChannelCallback(bot)  )
    bot.api.HandleFunc(irc.JOIN,    makeSendToChannelCallback(bot)  )
    bot.api.HandleFunc(irc.QUIT,    makeSendToChannelCallback(bot)  )
    bot.api.HandleFunc(irc.PART,    makeSendToChannelCallback(bot)  )
}
