package main

import (
	"crypto/tls"
	"emersyx.net/emersyx_apis/emircapi"
	"errors"
	"io"
	"math"
	"strconv"
)

// IRCOptions implements the emircapi.IRCOptions interface. Each method returns a function, which applies a specific
// configuration to an IRCGateway object.
type IRCOptions struct {
}

// Logging sets the io.Writer instance to write logging messages to and the verbosity level.
func (o IRCOptions) Logging(writer io.Writer, level uint) func(emircapi.IRCGateway) error {
	return func(gw emircapi.IRCGateway) error {
		if writer == nil {
			return errors.New("writer argument cannot be nil")
		}
		cgw, ok := gw.(*IRCGateway)
		if ok == false {
			return errors.New("unsupported IRCGateway implementation")
		}
		cgw.log.SetOutput(writer)
		cgw.log.SetLevel(level)
		return nil
	}
}

// Identifier sets the receptor identifier value for the IRC gateway.
func (o IRCOptions) Identifier(id string) func(emircapi.IRCGateway) error {
	return func(gw emircapi.IRCGateway) error {
		if len(id) == 0 {
			return errors.New("identifier cannot have zero length")
		}
		cgw, ok := gw.(*IRCGateway)
		if ok == false {
			return errors.New("unsupported IRCGateway implementation")
		}
		cgw.identifier = id
		cgw.log.SetComponentID(id)
		return nil
	}
}

// Nick sets the nickname to be used by the IRC gateway.
func (o IRCOptions) Nick(nick string) func(emircapi.IRCGateway) error {
	return func(gw emircapi.IRCGateway) error {
		if len(nick) == 0 {
			return errors.New("nick cannot have zero length")
		}
		cgw, ok := gw.(*IRCGateway)
		if ok == false {
			return errors.New("unsupported IRCGateway implementation")
		}
		cgw.cfg.Me.Nick = nick
		return nil
	}
}

// Ident sets the ident value to be used by the IRC gateway.
func (o IRCOptions) Ident(ident string) func(emircapi.IRCGateway) error {
	return func(gw emircapi.IRCGateway) error {
		if len(ident) == 0 {
			return errors.New("ident cannot have zero length")
		}
		cgw, ok := gw.(*IRCGateway)
		if ok == false {
			return errors.New("unsupported IRCGateway implementation")
		}
		cgw.cfg.Me.Ident = ident
		return nil
	}
}

// Name sets the name value to be used by the IRC gateway.
func (o IRCOptions) Name(name string) func(emircapi.IRCGateway) error {
	return func(gw emircapi.IRCGateway) error {
		if len(name) == 0 {
			return errors.New("name cannot have zero length")
		}
		cgw, ok := gw.(*IRCGateway)
		if ok == false {
			return errors.New("unsupported IRCGateway implementation")
		}
		cgw.cfg.Me.Name = name
		return nil
	}
}

// Version sets the version value to be used by the IRC gateway when receiving CTCP version request.
func (o IRCOptions) Version(version string) func(emircapi.IRCGateway) error {
	return func(gw emircapi.IRCGateway) error {
		cgw, ok := gw.(*IRCGateway)
		if ok == false {
			return errors.New("unsupported IRCGateway implementation")
		}
		cgw.cfg.Version = version
		return nil
	}
}

// Server sets the server address, port and SSL usage options for the IRC gateway.
func (o IRCOptions) Server(address string, port uint, useSSL bool) func(emircapi.IRCGateway) error {
	return func(gw emircapi.IRCGateway) error {
		if len(address) == 0 {
			return errors.New("address cannot have zero length")
		}
		if float64(port) > math.Pow(2, 16)-1 {
			return errors.New("port value is invalid")
		}
		cgw, ok := gw.(*IRCGateway)
		if ok == false {
			return errors.New("unsupported IRCGateway implementation")
		}
		cgw.cfg.Server = address + ":" + strconv.Itoa(int(port))
		cgw.cfg.SSL = useSSL
		if cgw.cfg.SSL {
			cgw.cfg.SSLConfig = &tls.Config{ServerName: address}
		}
		return nil
	}
}

// QuitMessage sets the message to be sent by the IRC gateway when it disconnects from a server.
func (o IRCOptions) QuitMessage(message string) func(emircapi.IRCGateway) error {
	return func(gw emircapi.IRCGateway) error {
		cgw, ok := gw.(*IRCGateway)
		if ok == false {
			return errors.New("unsupported IRCGateway implementation")
		}
		cgw.cfg.QuitMessage = message
		return nil
	}
}

// NewIRCOptions generates a new IRCOptions object and returns a pointer to it.
func NewIRCOptions() emircapi.IRCOptions {
	return new(IRCOptions)
}
