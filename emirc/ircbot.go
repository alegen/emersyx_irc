package main

import(
    "strconv"
    "crypto/tls"
    "emersyx.net/emersyx_apis/emircapi"
    irc "github.com/fluffle/goirc/client"
)

// The IRCBot struct defines the implementation of an IRC receptor and resource. The struct implements the
// emircapi.IRCBot and emcomapi.Receptor interfaces.
type IRCBot struct {
    api *irc.Conn
    Messages chan emircapi.Message
}

// This function creates a new emircapi.IRCBot instance and applies to configuration specified in the arguments.
func NewIRCBot(nick string, server string, port int, useSSL bool) emircapi.IRCBot {
    bot := new(IRCBot)
    bot.Messages = make(chan emircapi.Message)

    cfg := irc.NewConfig(nick)
	cfg.Me.Ident = "emersyx"
	cfg.Me.Name = "emersyx"
	cfg.QuitMessage = "buh`bye"
    cfg.SSL = useSSL
    cfg.Server = server + ":" + strconv.Itoa(port)
    cfg.NewNick = func(n string) string { return n + "^" }

    if useSSL {
        cfg.SSLConfig = &tls.Config{ ServerName: server }
    }

    bot.api = irc.Client(cfg)
    bot.initCallbacks()

    return bot
}

func (bot *IRCBot) initCallbacks() {
    bot.api.HandleFunc(irc.PRIVMSG, makeSendToChannelCallback(bot)  )
    bot.api.HandleFunc(irc.JOIN,    makeSendToChannelCallback(bot)  )
    bot.api.HandleFunc(irc.QUIT,    makeSendToChannelCallback(bot)  )
    bot.api.HandleFunc(irc.PART,    makeSendToChannelCallback(bot)  )
}
