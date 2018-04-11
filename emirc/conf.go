package main

import (
	"crypto/tls"
	"errors"
	"math"
	"strconv"
)

// ircGatewayConfig is the struct for holding IRC gateway configuration values, loaded from the configuration file.
type ircGatewayConfig struct {
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
func (cfg *ircGatewayConfig) validate() error {
	if cfg.Nick == nil || len(*cfg.Nick) == 0 {
		return errors.New("nick cannot have zero length")
	}
	if cfg.Ident == nil || len(*cfg.Ident) == 0 {
		return errors.New("ident cannot have zero length")
	}
	if cfg.Name == nil || len(*cfg.Name) == 0 {
		return errors.New("name cannot have zero length")
	}
	if cfg.ServerAddress == nil || len(*cfg.ServerAddress) == 0 {
		return errors.New("address cannot have zero length")
	}
	if cfg.ServerPort == nil || float64(*cfg.ServerPort) > math.Pow(2, 16)-1 {
		return errors.New("port value is invalid")
	}
	return nil
}

// apply sets the values loaded from the toml configuration file into the ircGateway object received as argument.
func (cfg *ircGatewayConfig) apply(gw *ircGateway) {
	gw.config.Me.Nick = *cfg.Nick
	gw.config.Me.Ident = *cfg.Ident
	gw.config.Me.Name = *cfg.Name
	if cfg.Version != nil {
		gw.config.Version = *cfg.Version
	} else {
		gw.config.Version = ""
	}
	gw.config.Server = *cfg.ServerAddress + ":" + strconv.Itoa(int(*cfg.ServerPort))
	if cfg.ServerUseSSL != nil {
		gw.config.SSL = *cfg.ServerUseSSL
	} else {
		// defaults to secure tls connection
		gw.config.SSL = true
	}
	if gw.config.SSL {
		gw.config.SSLConfig = &tls.Config{ServerName: *cfg.ServerAddress}
	}
	if cfg.QuitMessage != nil {
		gw.config.QuitMessage = *cfg.QuitMessage
	}
}
