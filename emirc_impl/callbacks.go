package main

import(
    irc "github.com/fluffle/goirc/client"
)

func messageToChannel(bot *IRCBot, m *Message) {
    bot.Messages <- m
}

func makeSendToChannelCallback(bot *IRCBot) func(*irc.Conn, *irc.Line) {
    return func(conn *irc.Conn, line *irc.Line) {
        go messageToChannel(bot, newMessage(line))
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
