# Go MyCMS ConfigFile

This package is an utility to read config files

The project configuration can be stored on the following file types:
- JSON
- XML
- TOML
- YAML

## Versions

Current version is **v1.0.1-alpha**

- [v1.0.0-alpha](https://github.com/Lucas-Palomo/go-mycms-configfile/releases/tag/v1.0.0-alpha)
- [v1.0.1-alpha](https://github.com/Lucas-Palomo/go-mycms-configfile/releases/tag/v1.0.1-alpha)

## Install

In your go project, run the following command

```shell
go get github.com/Lucas-Palomo/go-mycms-configfile@v1.0.1-alpha
```

## Parse Example

This is toml configuration file.
```toml
[database]
host="localhost"
port=3306
username="root"
password="toor"
```

Calling toml configuration is easy, see the bellow code:

```go

package main

import (
	ConfigFile "github.com/Lucas-Palomo/go-mycms-configfile/pkg"
)

type Conf struct {
	DB DatabaseConf `toml:"database"`
}

type DatabaseConf struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

func main() {
	conf := Conf{}
	ConfigFile.ReadFile("/home/user/config.toml", &conf, ConfigFile.TOML)
	// Now conf has all populated values
}

```