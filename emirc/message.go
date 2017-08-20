package main

import(
    irc "github.com/fluffle/goirc/client"
)

const(
    PRIVMSG = "PRIVMSG"
)

// This is the basic structure for an IRC message received by the client when an event occurs.
// Names of the struct members have been taken from RFC-1459 and RFC-2812.
type Message struct {
    Raw         string
    Command     string
    Origin      string
    Parameters  []string
}

func NewMessage(raw, command, origin string, parameters []string) *Message {
    m := new(Message)

    m.Raw = raw
    m.Command = command
    m.Origin = origin
    m.Parameters = make([]string, len(parameters))
    copy(m.Parameters, parameters)

    return m
}

func newMessage(line *irc.Line) *Message {
    m := new(Message)

    m.Raw = line.Raw
    m.Command = line.Cmd
    m.Origin = line.Nick
    m.Parameters = make([]string, len(line.Args))
    copy(m.Parameters, line.Args)

    return m
}
