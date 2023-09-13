package config

// NewEnvConfig create a new env file configuration manager instance
func NewEnvConfig(filenames ...string) (Config, error) {
	ec := &envConfig{Files: filenames}
	if err := ec.Load(); err != nil {
		return nil, err
	} else {
		return ec, nil
	}
}

// NewJSONConfig create a new json file configuration manager instance
func NewJSONConfig(filenames ...string) (Config, error) {
	jc := &jsonConfig{Files: filenames}
	if err := jc.Load(); err != nil {
		return nil, err
	} else {
		return jc, nil
	}
}

// NewMemoryConfig create a new in-memory configuration manager instance
func NewMemoryConfig(config map[string]any) (Config, error) {
	mc := new(memoryConfig)
	if config != nil {
		mc.data = config
	}
	if err := mc.Load(); err != nil {
		return nil, err
	} else {
		return mc, nil
	}
}
