package config

import "github.com/gomig/caster"

// Config is the interface for configuration manager drivers.
type Config interface {
	// Load configurations
	Load() error
	// Set configuration
	//
	// return error if driver not support set or error happend
	//
	// set override local configuration
	Set(key string, value any) error
	// Get configuration
	//
	// return nil if value not exists
	Get(key string) any
	// Exists check if config item exists
	Exists(key string) bool
	// Cast parse config as caster
	Cast(key string) caster.Caster
}
