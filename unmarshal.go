package config

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type UnmarshalAction func(buf []byte, val interface{}) error

func jsonUnmarshal(buf []byte, val interface{}) error {
	if err := json.Unmarshal(buf, val); err != nil {
		return err
	}
	return nil
}

func yamlUnmarshal(buf []byte, val interface{}) error {
	if err := yaml.Unmarshal(buf, val); err != nil {
		return err
	}
	return nil
}
