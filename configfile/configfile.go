package configfile

import (
	"encoding/json"
	"encoding/xml"
	"github.com/BurntSushi/toml"
	"github.com/Lucas-Palomo/go-mycms-configfile/internal/utils"
	"gopkg.in/yaml.v3"
)

type fType int

const (
	XML  fType = iota
	JSON fType = iota
	TOML fType = iota
	YAML fType = iota
)

func (t fType) Unmarshal() func(data []byte, model interface{}) error {
	unmarshal := []func(data []byte, model interface{}) error{
		xml.Unmarshal,
		json.Unmarshal,
		toml.Unmarshal,
		yaml.Unmarshal,
	}

	return unmarshal[t]
}

func ReadFile(filepath string, model interface{}, filetype fType) error {

	data, err := utils.ReadFile(filepath)
	if err != nil {
		return err
	}
	return utils.Unmarshal(data, model, filetype.Unmarshal())
}
