package main

import (
	"emersyx.net/emersyx/api"
	"errors"
	"github.com/BurntSushi/toml"
	"io"
)

// IRCOptions implements the api.IRCOptions interface. Each method returns a function, which applies a specific
// configuration to an IRCGateway object.
// TODO make this type private
type IRCOptions struct {
}

// Logging sets the io.Writer instance to write logging messages to and the verbosity level.
func (o IRCOptions) Logging(writer io.Writer, level uint) func(api.Peripheral) error {
	return func(peripheral api.Peripheral) error {
		if writer == nil {
			return errors.New("writer argument cannot be nil")
		}
		gw := assertIRCGateway(peripheral)
		gw.log.SetOutput(writer)
		gw.log.SetLevel(level)
		return nil
	}
}

// Identifier sets the receptor identifier value for the IRC gateway.
func (o IRCOptions) Identifier(id string) func(api.Peripheral) error {
	return func(peripheral api.Peripheral) error {
		if len(id) == 0 {
			return errors.New("identifier cannot have zero length")
		}
		gw := assertIRCGateway(peripheral)
		gw.identifier = id
		gw.log.SetComponentID(id)
		return nil
	}
}

// ConfigPath loads the toml configuration file and validates the contents. If valid, the contents are applied to the
// IRC gateway.
func (o IRCOptions) ConfigPath(path string) func(api.Peripheral) error {
	return func(peripheral api.Peripheral) error {
		config := new(ircGatewayConfig)
		_, err := toml.DecodeFile(path, config)
		if err != nil {
			return err
		}
		if err := config.validate(); err != nil {
			return err
		}
		gw := assertIRCGateway(peripheral)
		config.apply(gw)
		return nil
	}
}

// Core sets the api.Core instance which provides services to the IRC gateway.
func (o IRCOptions) Core(core api.Core) func(api.Peripheral) error {
	return func(peripheral api.Peripheral) error {
		if core == nil {
			return errors.New("core argument cannot be nil")
		}
		gw := assertIRCGateway(peripheral)
		gw.core = core
		return nil
	}
}

// assertIRCGateway tries to make a type assertion on the peripheral argument, to the *IRCGateway type. If the type
// assertion fails, then panic() is called. A call to recover() is in the applyOptions function.
func assertIRCGateway(peripheral api.Peripheral) *IRCGateway {
	gw, ok := peripheral.(*IRCGateway)
	if ok == false {
		panic("unsupported IRCGateway implementation")
	}
	return gw
}

// applyOptions executes the functions provided as the options argument with IRC gateway as argument. The implementation
// calls recover() in order to stop panicking, which may be caused by the call to panic() within the assertProcessor
// function. assertProcessor is used by functions returned by IRCOptions.
func applyOptions(peripheral api.Peripheral, options ...func(api.Peripheral) error) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	for _, option := range options {
		if e = option(peripheral); e != nil {
			return
		}
	}

	return
}

// NewPeripheralOptions generates a new IRCOptions object and returns a pointer to it.
func NewPeripheralOptions() api.PeripheralOptions {
	return new(IRCOptions)
}
