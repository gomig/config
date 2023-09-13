package config_test

import (
	"testing"

	"github.com/gomig/config"
)

func TestEnvConfig(t *testing.T) {
	env, err := config.NewEnvConfig("config_test.env")
	if err != nil {
		t.Fatal(err)
	}
	v, err := env.Cast("APP_TITLE").String()
	if err != nil {
		t.Fatal(err)
	}
	if v != "My App" {
		t.Errorf(`Failed check APP_TITLE == "My App"`)
	}
}

func TestJsonConfig(t *testing.T) {
	env, err := config.NewJSONConfig("config_test.json")
	if err != nil {
		t.Fatal(err)
	}
	v, err := env.Cast("app.title").String()
	if err != nil {
		t.Fatal(err)
	}
	if v != "My App" {
		t.Errorf(`Failed check app.title == "My App"`)
	}
}

func TestMemConfig(t *testing.T) {
	env, err := config.NewMemoryConfig(map[string]any{
		"title": "My App",
	})
	if err != nil {
		t.Fatal(err)
	}
	v, err := env.Cast("title").String()
	if err != nil {
		t.Fatal(err)
	}
	if v != "My App" {
		t.Errorf(`Failed check title == "My App"`)
	}
}
