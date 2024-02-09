# Go MyCMS ConfigFile

This package is a tool to read and write config files

The project configuration can be stored on the following file types:
- JSON
- XML
- TOML
- YAML

## Versions

Current version is **v2.0.0**

- [v1.0.0-alpha](https://github.com/Lucas-Palomo/go-mycms-configfile/releases/tag/v1.0.0-alpha)
- [v1.0.1-alpha](https://github.com/Lucas-Palomo/go-mycms-configfile/releases/tag/v1.0.1-alpha)
- [v1.0.2-alpha](https://github.com/Lucas-Palomo/go-mycms-configfile/releases/tag/v1.0.2-alpha)
- [v2.0.0](https://github.com/Lucas-Palomo/go-mycms-configfile/releases/tag/v2.0.0)

## Install

In your go project, run the following command

```shell
go get github.com/Lucas-Palomo/go-mycms-configfile@latest
```
## Examples

All examples can be found in "examples" directory

### Read Example

This is toml configuration file.
```toml
[database]
host="localhost"
port=3306
username="root"
password="toor"
```

I created a struct to represent this configuration

```go
package examples_internal

type Conf struct {
	DB DatabaseConf `toml:"database" json:"database"`
}

type DatabaseConf struct {
	Host     string `toml:"host" json:"host"`
	Port     int    `toml:"port" json:"port"`
	Username string `toml:"username" json:"username"`
	Password string `toml:"password" json:"password"`
}
```

Calling toml configuration is easy, see the bellow code:

```go
package main

import (
	"fmt"
	ConfigFile "github.com/Lucas-Palomo/go-mycms-configfile/configfile"
	Model "github.com/Lucas-Palomo/go-mycms-configfile/examples/internal"
	"log"
)

func main() {
	conf := Model.Conf{}
	err := ConfigFile.ReadFile("./data/config.toml", &conf, ConfigFile.TOML)
	if err != nil {
		log.Fatal(err)
	}
	// Now conf has all populated values
	fmt.Printf("%#v\n", conf)
}
```

The output is

```text
examples_internal.Conf{DB:examples_internal.DatabaseConf{Host:"localhost", Port:3306, Username:"root", Password:"toor"}}
```

### Write Example
Another usage is to create configuration files, in this example, we're creating a json to store the project configuration:

```go
package main

import (
	ConfigFile "github.com/Lucas-Palomo/go-mycms-configfile/configfile"
	Model "github.com/Lucas-Palomo/go-mycms-configfile/examples/internal"
	"log"
)

func main() {
	conf := Model.Conf{
		DB: Model.DatabaseConf{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "toor",
		},
	}
	err := ConfigFile.WriteFile("./data/config.json", &conf, ConfigFile.JSON)
	if err != nil {
		log.Fatal(err)
	}
	// Now conf content is in the config.json
}
```