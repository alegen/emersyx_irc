package main

import(
    "emersyx.net/emersyx_irc/emirc"
    irc "github.com/fluffle/goirc/client"
)

func newMessage(line *irc.Line) *emirc.Message {
    m := new(emirc.Message)

    m.Raw = line.Raw
    m.Command = line.Cmd
    m.Origin = line.Nick
    m.Parameters = make([]string, len(line.Args))
    copy(m.Parameters, line.Args)

    return m
}

func makeSendToChannelCallback(bot *IRCBot) func(*irc.Conn, *irc.Line) {
    return func(conn *irc.Conn, line *irc.Line) {
        go func(){
            bot.Messages <- newMessage(line)
        }()
    }
}

func makeLoggingCallback(bot *IRCBot) func(*irc.Conn, *irc.Line) {
    return func(conn *irc.Conn, line *irc.Line) {
        m := newMessage(line)
        log().Printf("new message: %s\n", m.Command)
        log().Printf("raw:         %s\n", m.Raw)
        log().Printf("origin:      %s\n", m.Origin)
        log().Printf("parameters:  %s\n", m.Parameters)
    }
}
