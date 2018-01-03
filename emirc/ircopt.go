package main

import(
    "crypto/tls"
    "errors"
    "math"
    "strconv"
    "emersyx.net/emersyx_apis/emircapi"
)

// IRCOptions implements the emircapi.IRCOptions interface. Each method returns a function, which applies a specific
// configuration to an IRCBot object.
type IRCOptions struct {
}

// Identifier sets the receptor identifier value for the IRC bot.
func (o IRCOptions) Identifier(id string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(id) == 0 {
            return errors.New("identifier cannot have zero length")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("unsupported IRCBot implementation")
        }
	    cbot.identifier = id
        return nil
    }
}

// Nick sets the nickname to be used by the IRC bot.
func (o IRCOptions) Nick(nick string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(nick) == 0 {
            return errors.New("nick cannot have zero length")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("unsupported IRCBot implementation")
        }
	    cbot.cfg.Me.Nick = nick
        return nil
    }
}

// Ident sets the ident value to be used by the IRC bot.
func (o IRCOptions) Ident(ident string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(ident) == 0 {
            return errors.New("ident cannot have zero length")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("unsupported IRCBot implementation")
        }
	    cbot.cfg.Me.Ident = ident
        return nil
    }
}

// Name sets the name value to be used by the IRC bot.
func (o IRCOptions) Name(name string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(name) == 0 {
            return errors.New("name cannot have zero length")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("unsupported IRCBot implementation")
        }
	    cbot.cfg.Me.Name = name
        return nil
    }
}

// Version sets the version value to be used by the IRC bot when receiving CTCP version request.
func (o IRCOptions) Version(version string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("unsupported IRCBot implementation")
        }
	    cbot.cfg.Version = version
        return nil
    }
}

// Server sets the server address, port and SSL usage options for the IRC bot.
func (o IRCOptions) Server(address string, port uint, useSSL bool) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(address) == 0 {
            return errors.New("address cannot have zero length")
        }
        if float64(port) > math.Pow(2, 16) - 1 {
            return errors.New("port value is invalid")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("unsupported IRCBot implementation")
        }
        cbot.cfg.Server = address + ":" + strconv.Itoa(int(port))
        cbot.cfg.SSL = useSSL;
        if cbot.cfg.SSL {
            cbot.cfg.SSLConfig = &tls.Config{ ServerName: address }
        }
        return nil
    }
}

// QuitMessage sets the message to be sent by the IRC bot when it disconnects from a server.
func (o IRCOptions) QuitMessage(message string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("unsupported IRCBot implementation")
        }
	    cbot.cfg.QuitMessage = message;
        return nil
    }
}

// NewIRCOptions generates a new IRCOptions object and returns a pointer to it.
func NewIRCOptions() emircapi.IRCOptions {
    return new(IRCOptions)
}
