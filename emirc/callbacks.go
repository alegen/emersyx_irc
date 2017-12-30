package main

import(
    "emersyx.net/emersyx_apis/emircapi"
    "log"
    irc "github.com/fluffle/goirc/client"
)

func newMessage(id string, line *irc.Line) emircapi.Message {
    var m emircapi.Message

    m.Source = id
    m.Raw = line.Raw
    m.Command = line.Cmd
    m.Origin = line.Nick
    m.Parameters = make([]string, len(line.Args))
    copy(m.Parameters, line.Args)

    return m
}

func channelCallback(bot *IRCBot) func(*irc.Conn, *irc.Line) {
    return func(conn *irc.Conn, line *irc.Line) {
        go func() {
            bot.Messages <- newMessage(bot.identifier, line)
        }()
    }
}

func loggingCallback(bot *IRCBot, logger *log.Logger) func(*irc.Conn, *irc.Line) {
    return func(conn *irc.Conn, line *irc.Line) {
        m := newMessage(bot.identifier, line)
        logger.Printf("new message: %s\n", m.Command)
        logger.Printf("raw:         %s\n", m.Raw)
        logger.Printf("origin:      %s\n", m.Origin)
        logger.Printf("parameters:  %s\n", m.Parameters)
    }
}
