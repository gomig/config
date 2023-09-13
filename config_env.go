package config

import (
	"fmt"
	"os"

	"github.com/gomig/caster"
	"github.com/gomig/utils"
	"github.com/joho/godotenv"
)

type envConfig struct {
	Files []string
	data  map[string]any
}

func (envConfig) err(format string, args ...any) error {
	return utils.TaggedError([]string{"EnvConfig"}, format, args...)
}

func (ec *envConfig) Load() error {
	if ec.data == nil {
		ec.data = make(map[string]any)
	}

	if err := godotenv.Overload(ec.Files...); err != nil {
		return ec.err(err.Error())
	}

	for k, v := range ec.data {
		if err := ec.Set(k, v); err != nil {
			return err
		}
	}

	return nil
}

func (ec *envConfig) Set(key string, value any) error {
	ec.data[key] = value
	if err := os.Setenv(key, fmt.Sprintf("%v", value)); err != nil {
		return ec.err(err.Error())
	}
	return nil
}

func (envConfig) Get(key string) any {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return nil
}

func (envConfig) Exists(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}

func (ec envConfig) Cast(key string) caster.Caster {
	return caster.NewCaster(ec.Get(key))
}
