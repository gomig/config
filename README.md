# Config

Configuration manager with default `Env`, `Json` and `Memory` driver.

**Mote:** Config use [Caster](https://github.com/gomig/caster) for get dependencies by type.

## Create New Config Driver

Config library contains three different driver by default.

### Env Driver

Env driver use environment file (.env) for managing configuration.

```go
import "github.com/gomig/config"
envConf, err := config.NewEnvConfig("app.env", "db.env", ".env")
```

### JSON Driver

JSON driver use json file for managing configuration.

**Caution:** When you pass multiple file, for accessing config you must pass file name as first part of config path!

```go
import "github.com/gomig/config"
jsonConf, err := config.NewJSONConfig("app.json", "db.json", "global.json")
```

### Memory Driver

Use in-memory array for keeping and managing configuration.

```go
import "github.com/gomig/config"
memConf, err := config.NewMemoryConfig(map[string]any{
    "name": "My First App",
    "key": "My Secret Key",
})
```

## Usage

Config interface contains following methods:

### Load

Load/Reload configurations.

```go
// Signature:
Load() error

// Example
err := envConf.Load()
```

### Set

Set configuration item. this function override preloaded config.

```go
// Signature:
Set(key string, value any) error

// Example
err := memConf.Set("name", "My App")
err = envConf.Set("APP_NAME", "My App")
err = jsonConf.Set("app_name", "My App")
```

**Cation:** For setting/overriding config item in `JSON` driver with multiple files pass filename as first part of config path.

```go
import "github.com/gomig/config"
jsonConf, err := config.NewJSONConfig("file1.json", "file2.json")
err = jsonConf.Set("file1.app.title", "Some")
```

### Get

Get configuration. Get function return config item as `any`. if you need get config with type use helper get functions described later.

```go
// Signature:
Get(key string) any
```

**Caution:** For `JSON` driver with multiple file you must pass filename as first part of config path!

```go
item := jsonConf.Get("file1.app.title")
```

### Exists

Check if config item exists.

```go
// Signature:
Exists(key string) bool
```

### Cast

Parse config as caster.

```go
// Signature:
Cast(name string) caster.Caster

// Example:
v, err := conf.Cast("timeout").Int()
```
