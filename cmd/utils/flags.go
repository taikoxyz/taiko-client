package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"unicode"

	"github.com/naoina/toml"
)

// These settings ensure that TOML keys use the same names as Go struct fields.
var tomlSettings = toml.Config{
	NormFieldName: func(_ reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(_ reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		var link string
		if unicode.IsUpper(rune(rt.Name()[0])) && rt.PkgPath() != "main" {
			link = fmt.Sprintf(", see https://godoc.org/%s#%s for available fields", rt.PkgPath(), rt.Name())
		}
		return fmt.Errorf("field '%s' is not defined in %s%s", field, rt.String(), link)
	},
}

// LoadConfigFile loads TOML config file.
func LoadConfigFile(file string, cfg interface{}) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(cfg)
	// Add file name to errors that have a line number.
	var lineError *toml.LineError
	if errors.As(err, &lineError) {
		err = errors.New(file + ", " + err.Error())
	}
	return err
}
