package config

import (
	"io/ioutil"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

var (
	unmarshals map[string]UnmarshalAction = map[string]UnmarshalAction{
		CTJson: jsonUnmarshal,
		CTYaml: yamlUnmarshal,
	}

	marshals map[string]MarshalAction = map[string]MarshalAction{
		CTJson: jsonMarshal,
		CTYaml: yamlMarshal,
	}
)

func Read(val interface{}, opts ...Option) error {
	options := newOptions(opts...)

	filePath, err := homedir.Expand(options.filePath)
	if err != nil {
		return xerrors.Errorf("expanding failed: %w", err)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	unmarshal, ok := unmarshals[options.configType]
	if !ok {
		return xerrors.Errorf("unmarshals:%s is not exist", options.configType)
	}

	return unmarshal(data, val)
}

func Write(val interface{}, opts ...Option) error {
	options := newOptions(opts...)

	filePath, err := homedir.Expand(options.filePath)
	if err != nil {
		return xerrors.Errorf("expanding failed: %w", err)
	}

	marshal, ok := marshals[options.configType]
	if !ok {
		return xerrors.Errorf("marshals:%s is not exist", options.configType)
	}

	data, err := marshal(val, opts...)
	if err != nil {
		return nil
	}

	return ioutil.WriteFile(filePath, data, options.perm)
}
