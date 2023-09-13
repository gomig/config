package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gomig/caster"
	"github.com/gomig/utils"
	"github.com/tidwall/gjson"
)

type jsonConfig struct {
	Files []string
	json  string
	data  map[string]any
}

func (jsonConfig) err(format string, args ...any) error {
	return utils.TaggedError([]string{"JSONConfig"}, format, args...)
}

func (jc *jsonConfig) fetch(key string) (any, bool, bool) {
	if v, ok := jc.data[key]; ok {
		return v, false, true
	}

	if val := gjson.Get(jc.json, key); val.Exists() {
		return val, true, true
	}

	return nil, false, false
}

func (jc *jsonConfig) Load() error {
	if jc.data == nil {
		jc.data = make(map[string]any)
	}

	contents := make([]string, 0)
	for _, f := range jc.Files {
		bytes, err := ioutil.ReadFile(f)
		if err != nil {
			return jc.err(err.Error())
		}
		content := string(bytes)
		if !gjson.Valid(content) {
			return jc.err(fmt.Sprintf("%s content is invalid!", f))
		}

		fileName := filepath.Base(f)
		fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

		if len(jc.Files) > 1 {
			contents = append(contents, `"`+fileName+`":`+content)
		} else {
			contents = append(contents, content)
		}

	}
	if len(jc.Files) > 1 {
		jc.json = "{" + strings.Join(contents, ",") + "}"
	} else {
		if !strings.HasPrefix(contents[0], "{") {
			contents[0] = "{" + contents[0] + "}"
		}
		jc.json = contents[0]
	}
	return nil
}

func (jc *jsonConfig) Set(key string, value any) error {
	jc.data[key] = value
	return nil
}

func (jc jsonConfig) Get(key string) any {
	if v, isJSON, exists := jc.fetch(key); !exists {
		return nil
	} else if isJSON {
		return v.(gjson.Result).Value()
	} else {
		return v
	}
}

func (jc jsonConfig) Exists(key string) bool {
	_, _, exists := jc.fetch(key)
	return exists
}

func (jc jsonConfig) Cast(key string) caster.Caster {
	return caster.NewCaster(jc.Get(key))
}
