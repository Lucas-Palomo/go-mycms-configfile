package configfile

import (
	"encoding/json"
	"encoding/xml"
	"github.com/Lucas-Palomo/go-mycms-configfile/internal/utils"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"os"
)

type FileType int

const (
	XML  FileType = iota
	JSON FileType = iota
	TOML FileType = iota
	YAML FileType = iota
)

func (t FileType) Unmarshal() func(data []byte, model interface{}) error {
	unmarshal := []func(data []byte, model interface{}) error{
		xml.Unmarshal,
		json.Unmarshal,
		toml.Unmarshal,
		yaml.Unmarshal,
	}

	return unmarshal[t]
}

func (t FileType) Marshal() func(model interface{}) ([]byte, error) {
	marshal := []func(model interface{}) ([]byte, error){
		xml.Marshal,
		json.Marshal,
		toml.Marshal,
		yaml.Marshal,
	}

	return marshal[t]
}

func ReadFile(filepath string, model interface{}, filetype FileType) error {

	data, err := utils.ReadFile(filepath)
	if err != nil {
		return err
	}
	return utils.Unmarshal(data, model, filetype.Unmarshal())
}

func WriteFile(filepath string, model interface{}, filetype FileType) error {
	content, err := utils.Marshal(model, filetype.Marshal())

	if err != nil {
		return err
	}

	file, err := utils.OpenFile(filepath, os.O_SYNC|os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	_, err = file.Write(content)

	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
