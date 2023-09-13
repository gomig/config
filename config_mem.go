package config

import (
	"github.com/gomig/caster"
)

type memoryConfig struct {
	data map[string]any
}

func (mc memoryConfig) Load() error {
	// Load do nothing for memory config
	if mc.data == nil {
		mc.data = make(map[string]any)
	}
	return nil
}

func (mc memoryConfig) Set(key string, value any) error {
	mc.data[key] = value
	return nil
}

func (mc memoryConfig) Get(key string) any {
	if v, ok := mc.data[key]; ok {
		return v
	}
	return nil
}

func (mc memoryConfig) Exists(key string) bool {
	if _, ok := mc.data[key]; ok {
		return true
	}
	return false
}

func (mc memoryConfig) Cast(key string) caster.Caster {
	return caster.NewCaster(mc.data[key])
}
