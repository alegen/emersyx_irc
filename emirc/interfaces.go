package emirc

import(
    "errors"
    "plugin"
)

type IRCBot interface {
    GetEventsChannel() chan interface{}
    Connect() error
    Quit()
    Join(string)
    Privmsg(string, string)
}

func NewIRCBot(irc_plugin *plugin.Plugin, nick string, server string, port int, useSSL bool) (IRCBot, error) {
    if irc_plugin == nil {
        return nil, errors.New("Invalid IRC plugin handle")
    }

    f, err := irc_plugin.Lookup("NewIRCBot")
    if err != nil {
        return nil, errors.New("IRC plugin does not have the NewIRCBot symbol.")
    }

    var fc func(string, string, int, bool) IRCBot
    fc = f.(func(string, string, int, bool) IRCBot)
    if err != nil {
        return nil, errors.New("The NewIRCBot function does not have the correct signature.")
    }

    bot := fc(nick, server, port, useSSL)
    return bot, nil
}
