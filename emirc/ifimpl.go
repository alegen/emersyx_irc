package main

import(
    "errors"
    "emersyx.net/emersyx_apis/emcomapi"
)

// This file contains implementations of methods that are mandatory to implement the emircapi.IRCBot interfaces.

func (bot *IRCBot) Connect() error {
    err := bot.api.Connect()
    return err
}

func (bot *IRCBot) IsConnected() bool {
    return bot.api.Connected()
}

func (bot *IRCBot) Quit() error {
    err := bot.api.Close()
    return err
}

func (bot *IRCBot) Join(ch string) error {
    if bot.IsConnected() {
        bot.api.Join(ch)
        return nil
    } else {
        return errors.New("The IRCBot instance is not connected to any server.");
    }
}

func (bot *IRCBot) Privmsg(to, msg string) error {
    if bot.IsConnected() {
        bot.api.Privmsg(to, msg)
        return nil
    } else {
        return errors.New("The IRCBot instance is not connected to any server.");
    }
}

func (bot *IRCBot) GetIdentifier() string {
    return bot.identifier;
}

func (bot *IRCBot) GetEventsChannel() chan emcomapi.Event {
    return bot.messages;
}
