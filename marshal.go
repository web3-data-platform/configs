package config

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type MarshalAction func(val interface{}, opts ...Option) ([]byte, error)

func jsonMarshal(val interface{}, opts ...Option) ([]byte, error) {
	options := newOptions(opts...)
	if options.indent == DefaultIndent {
		return json.Marshal(val)
	}
	return json.MarshalIndent(val, "", options.indent)
}

func yamlMarshal(val interface{}, opts ...Option) ([]byte, error) {
	return yaml.Marshal(val)
}
