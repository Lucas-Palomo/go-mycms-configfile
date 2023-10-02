package test

import (
	"encoding/json"
	"encoding/xml"
	"github.com/BurntSushi/toml"
	"github.com/Lucas-Palomo/go-mycms-configfile/internal/utils"
	"github.com/Lucas-Palomo/go-mycms-configfile/pkg"
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
	}

	t.Run("open file error", func(t *testing.T) {
		setup()
		object := Object{}

		utils.Open = func(name string) (*os.File, error) {
			return nil, os.ErrNotExist
		}

		err := pkg.ReadFile("./testdata/not.found", &object, pkg.XML)
		assert.Error(t, err)
	})
	t.Run("read file error", func(t *testing.T) {
		setup()
		object := Object{}

		utils.ReadAll = func(r io.Reader) ([]byte, error) {
			return []byte{}, io.ErrUnexpectedEOF
		}

		err := pkg.ReadFile("./testdata/test.xml", &object, pkg.XML)
		assert.Error(t, err)
	})
	t.Run("parse error", func(t *testing.T) {
		setup()
		object := Object{}

		err := pkg.ReadFile("./testdata/test.yaml", &object, pkg.XML)
		assert.Error(t, err)
	})
	t.Run("read xml", func(t *testing.T) {
		setup()
		object := Object{}

		err := pkg.ReadFile("./testdata/test.xml", &object, pkg.XML)
		assert.Nil(t, err)
		assert.Equal(t, object.Model.Message, "xml")
	})
	t.Run("read json", func(t *testing.T) {
		setup()
		object := Object{}

		err := pkg.ReadFile("./testdata/test.json", &object, pkg.JSON)
		assert.Nil(t, err)
		assert.Equal(t, object.Model.Message, "json")
	})
	t.Run("read  toml", func(t *testing.T) {
		setup()
		object := Object{}

		err := pkg.ReadFile("./testdata/test.toml", &object, pkg.TOML)
		assert.Nil(t, err)
		assert.Equal(t, object.Model.Message, "toml")
	})
	t.Run("read yaml", func(t *testing.T) {
		setup()
		object := Object{}

		err := pkg.ReadFile("./testdata/test.yaml", &object, pkg.YAML)
		assert.Nil(t, err)
		assert.Equal(t, object.Model.Message, "yaml")
	})
}

func TestFileType(t *testing.T) {
	t.Run("unmarshal reference", func(t *testing.T) {
		assert.Equal(t, reflect.ValueOf(pkg.XML.Unmarshal()), reflect.ValueOf(xml.Unmarshal))
		assert.Equal(t, reflect.ValueOf(pkg.JSON.Unmarshal()), reflect.ValueOf(json.Unmarshal))
		assert.Equal(t, reflect.ValueOf(pkg.TOML.Unmarshal()), reflect.ValueOf(toml.Unmarshal))
		assert.Equal(t, reflect.ValueOf(pkg.YAML.Unmarshal()), reflect.ValueOf(yaml.Unmarshal))
	})
}
