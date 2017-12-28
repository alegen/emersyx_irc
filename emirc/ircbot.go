package main

import(
    "strconv"
    "crypto/tls"
    "emersyx.net/emersyx_irc/emirc"
    irc "github.com/fluffle/goirc/client"

)

type IRCBot struct {
    api *irc.Conn
    Messages chan interface{}
}

func NewIRCBot(nick string, server string, port int, useSSL bool) emirc.IRCBot {
    bot := new(IRCBot)
    bot.Messages = make(chan interface{})

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

func (bot *IRCBot) GetEventsChannel() chan interface{} {
    return bot.Messages
}

func (bot *IRCBot) initCallbacks() {
    bot.api.HandleFunc(irc.PRIVMSG, makeSendToChannelCallback(bot)  )
    bot.api.HandleFunc(irc.JOIN,    makeSendToChannelCallback(bot)  )
    bot.api.HandleFunc(irc.QUIT,    makeSendToChannelCallback(bot)  )
    bot.api.HandleFunc(irc.PART,    makeSendToChannelCallback(bot)  )
}

func (bot *IRCBot) Connect() error {
    err := bot.api.Connect()
    if err == nil {
        bot.waitToConnect()
        return nil
    } else {
        return err
    }
}

func (bot *IRCBot) waitToConnect() {
    for bot.api.Connected() == false {
    }
}

func (bot *IRCBot) IsConnected() bool {
    return bot.api.Connected()
}


func (bot *IRCBot) Quit() {
    go bot.api.Close()
    bot.waitToQuit()
}

func (bot *IRCBot) waitToQuit() {
    for bot.IsConnected() {
    }
}

func (bot *IRCBot) Join(ch string) {
    bot.api.Join(ch)
}

func (bot *IRCBot) Privmsg(to, msg string) {
    bot.api.Privmsg(to, msg)
}
