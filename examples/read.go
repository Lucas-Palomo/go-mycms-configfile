package main

import (
	"fmt"
	ConfigFile "github.com/Lucas-Palomo/go-mycms-configfile/v2/configfile"
	Model "github.com/Lucas-Palomo/go-mycms-configfile/v2/examples/internal"
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
