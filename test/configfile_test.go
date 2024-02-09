package test

import (
	"encoding/json"
	"encoding/xml"
	"github.com/Lucas-Palomo/go-mycms-configfile/v2/configfile"
	"github.com/Lucas-Palomo/go-mycms-configfile/v2/internal/utils"
	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"reflect"
	"testing"
)

type Object struct {
	Model Model `xml:"test" json:"test" yaml:"test" toml:"test"`
}

type Model struct {
	Message string `xml:"message" json:"message" yaml:"message" toml:"message"`
}

func TestReadFile(t *testing.T) {

	setup := func() {
		utils.Open = os.Open
		utils.ReadAll = io.ReadAll
		utils.OpenFile = os.OpenFile
	}

	createObject := func(message string) Object {
		return Object{
			Model: Model{
				Message: message,
			},
		}
	}

	writeFile := func(filepath string, object interface{}, fileType configfile.FileType) error {
		return configfile.WriteFile(filepath, object, fileType)
	}

	readFile := func(filepath string, object interface{}, fileType configfile.FileType) error {
		return configfile.ReadFile(filepath, object, fileType)
	}

	t.Run("error to open file", func(t *testing.T) {
		setup()
		object := Object{}

		utils.Open = func(name string) (*os.File, error) {
			return nil, os.ErrNotExist
		}

		err := configfile.ReadFile("./testdata/not.found", &object, configfile.XML)
		assert.Error(t, err)
	})
	t.Run("error to open file and create", func(t *testing.T) {
		setup()
		object := createObject("it's a xml")

		utils.OpenFile = func(name string, flag int, perm os.FileMode) (*os.File, error) {
			return nil, os.ErrPermission
		}

		err := writeFile("./testdata/test.xml", &object, configfile.XML)
		assert.Error(t, err)
	})
	t.Run("error to read file", func(t *testing.T) {
		setup()
		object := Object{}

		utils.ReadAll = func(r io.Reader) ([]byte, error) {
			return []byte{}, io.ErrUnexpectedEOF
		}

		err := configfile.ReadFile("./testdata/test.xml", &object, configfile.XML)
		assert.Error(t, err)
	})
	t.Run("error to Marshal", func(t *testing.T) {
		setup()
		err := writeFile("./testdata/test.xml", make([]byte, 1), configfile.XML)
		assert.Error(t, err)
	})
	t.Run("error to Unmarshal", func(t *testing.T) {
		setup()
		path := "./testdata/test.yaml"
		object := createObject("it's a yaml")

		err := writeFile(path, &object, configfile.YAML)
		assert.Nil(t, err)

		object = Object{}

		err = readFile(path, &object, configfile.XML)
		assert.Error(t, err)
	})
	t.Run("write and read xml", func(t *testing.T) {
		setup()
		path := "./testdata/test.xml"
		object := createObject("it's a xml")

		err := writeFile(path, &object, configfile.XML)
		assert.Nil(t, err)

		object = Object{}

		err = readFile(path, &object, configfile.XML)
		assert.Nil(t, err)
		assert.Equal(t, object.Model.Message, "it's a xml")
	})
	t.Run("write and read json", func(t *testing.T) {
		setup()
		path := "./testdata/test.json"
		object := createObject("it's a json")

		err := writeFile(path, &object, configfile.JSON)
		assert.Nil(t, err)

		object = Object{}

		err = readFile(path, &object, configfile.JSON)
		assert.Nil(t, err)
		assert.Equal(t, object.Model.Message, "it's a json")
	})
	t.Run("write and read toml", func(t *testing.T) {
		setup()
		path := "./testdata/test.toml"
		object := createObject("it's a toml")

		err := writeFile(path, &object, configfile.TOML)
		assert.Nil(t, err)

		object = Object{}

		err = readFile(path, &object, configfile.TOML)
		assert.Nil(t, err)
		assert.Equal(t, object.Model.Message, "it's a toml")
	})
	t.Run("write and read yaml", func(t *testing.T) {
		setup()
		path := "./testdata/test.yaml"
		object := createObject("it's a yaml")

		err := writeFile(path, &object, configfile.YAML)
		assert.Nil(t, err)

		object = Object{}

		err = readFile(path, &object, configfile.YAML)
		assert.Nil(t, err)
		assert.Equal(t, object.Model.Message, "it's a yaml")
	})
}

func TestFileType(t *testing.T) {
	t.Run("unmarshal reference", func(t *testing.T) {
		assert.Equal(t, reflect.ValueOf(configfile.XML.Unmarshal()), reflect.ValueOf(xml.Unmarshal))
		assert.Equal(t, reflect.ValueOf(configfile.JSON.Unmarshal()), reflect.ValueOf(json.Unmarshal))
		assert.Equal(t, reflect.ValueOf(configfile.TOML.Unmarshal()), reflect.ValueOf(toml.Unmarshal))
		assert.Equal(t, reflect.ValueOf(configfile.YAML.Unmarshal()), reflect.ValueOf(yaml.Unmarshal))
	})
	t.Run("marshal reference", func(t *testing.T) {
		assert.Equal(t, reflect.ValueOf(configfile.XML.Marshal()), reflect.ValueOf(xml.Marshal))
		assert.Equal(t, reflect.ValueOf(configfile.JSON.Marshal()), reflect.ValueOf(json.Marshal))
		assert.Equal(t, reflect.ValueOf(configfile.TOML.Marshal()), reflect.ValueOf(toml.Marshal))
		assert.Equal(t, reflect.ValueOf(configfile.YAML.Marshal()), reflect.ValueOf(yaml.Marshal))
	})
}
