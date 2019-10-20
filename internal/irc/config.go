package main

import (
	"crypto/tls"
	"errors"
	"math"
	"strconv"
)

// config is the struct for holding IRC gateway configuration values, loaded from the configuration file.
type config struct {
	Nick          *string
	Ident         *string
	Name          *string
	Version       *string
	ServerAddress *string `toml:"server_address"`
	ServerPort    *uint   `toml:"server_port"`
	ServerUseSSL  *bool   `toml:"server_use_ssl"`
	QuitMessage   *string `toml:"quit_message"`
	PluginPath    *string `toml:"plugin_path"`
}

// validate checks the values loaded from the toml configuration file. If any value is found to be invalid, then an
// error is returned.
func (c *config) validate() error {
	if c.Nick == nil || len(*c.Nick) == 0 {
		return errors.New("nick cannot have zero length")
	}
	if c.Ident == nil || len(*c.Ident) == 0 {
		return errors.New("ident cannot have zero length")
	}
	if c.Name == nil || len(*c.Name) == 0 {
		return errors.New("name cannot have zero length")
	}
	if c.ServerAddress == nil || len(*c.ServerAddress) == 0 {
		return errors.New("address cannot have zero length")
	}
	if c.ServerPort == nil || float64(*c.ServerPort) > math.Pow(2, 16)-1 {
		return errors.New("port value is invalid")
	}
	return nil
}

// apply sets the values loaded from the toml configuration file into the gateway object received as argument.
func (c *config) apply(gw *gateway) {
	gw.config.Me.Nick = *c.Nick
	gw.config.Me.Ident = *c.Ident
	gw.config.Me.Name = *c.Name
	if c.Version != nil {
		gw.config.Version = *c.Version
	} else {
		gw.config.Version = ""
	}
	gw.config.Server = *c.ServerAddress + ":" + strconv.Itoa(int(*c.ServerPort))
	if c.ServerUseSSL != nil {
		gw.config.SSL = *c.ServerUseSSL
	} else {
		// defaults to secure tls connection
		gw.config.SSL = true
	}
	if gw.config.SSL {
		gw.config.SSLConfig = &tls.Config{ServerName: *c.ServerAddress}
	}
	if c.QuitMessage != nil {
		gw.config.QuitMessage = *c.QuitMessage
	}
}
