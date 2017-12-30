package main

import(
    "crypto/tls"
    "errors"
    "math"
    "strconv"
    "emersyx.net/emersyx_apis/emircapi"
)

// This struct implements the emircapi.IRCOptions interface. Each method returns a function, which applies a specific
// configuration to an IRCBot object.
type IRCOptions struct {
}

func (o IRCOptions) Identifier(id string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(id) == 0 {
            return errors.New("Identifier cannot have zero length.")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.identifier = id
        return nil
    }
}

func (o IRCOptions) Nick(nick string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(nick) == 0 {
            return errors.New("Nick cannot have zero length.")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.Me.Nick = nick
        return nil
    }
}

func (o IRCOptions) Ident(ident string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(ident) == 0 {
            return errors.New("Ident cannot have zero length.")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.Me.Ident = ident
        return nil
    }
}

func (o IRCOptions) Name(name string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(name) == 0 {
            return errors.New("Name cannot have zero length.")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.Me.Name = name
        return nil
    }
}

func (o IRCOptions) Version(version string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.Version = version
        return nil
    }
}

func (o IRCOptions) Server(address string, port uint, useSSL bool) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        if len(address) == 0 {
            return errors.New("Address cannot have zero length.")
        }
        if float64(port) > math.Pow(2, 16) - 1 {
            return errors.New("Port value is invalid.")
        }
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
        cbot.cfg.Server = address + ":" + strconv.Itoa(int(port))
        cbot.cfg.SSL = useSSL;
        if cbot.cfg.SSL {
            cbot.cfg.SSLConfig = &tls.Config{ ServerName: address }
        }
        return nil
    }
}

func (o IRCOptions) UseSSL(useSSL bool) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
        cbot.cfg.SSL = useSSL;
        return nil
    }
}

func (o IRCOptions) QuitMessage(message string) func(emircapi.IRCBot) error {
    return func(bot emircapi.IRCBot) error {
        cbot, ok := bot.(*IRCBot)
        if ok == false {
            return errors.New("Unsupported IRCBot implementation")
        }
	    cbot.cfg.QuitMessage = message;
        return nil
    }
}

func NewIRCOptions() emircapi.IRCOptions {
    return new(IRCOptions)
}
