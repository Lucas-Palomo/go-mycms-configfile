package main

import (
	ConfigFile "github.com/Lucas-Palomo/go-mycms-configfile/v2/configfile"
	Model "github.com/Lucas-Palomo/go-mycms-configfile/v2/examples/internal"
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
