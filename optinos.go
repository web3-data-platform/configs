package config

import "io/fs"

var (
	DefaultConfigType             = "json"
	DefaultIndent                 = ""
	DefaultPerm       fs.FileMode = 0644
)

const (
	CTJson = "json"
	CTYaml = "yaml"
)

type Options struct {
	filePath   string
	configType string
	indent     string
	perm       fs.FileMode
}

type Option func(opts *Options)

func newOptions(opts ...Option) *Options {
	options := Options{
		configType: DefaultConfigType,
		perm:       DefaultPerm,
	}

	for _, o := range opts {
		o(&options)
	}
	return &options
}

func ConfigType(t string) Option {
	return func(opts *Options) {
		opts.configType = t
	}
}

func Indent(indent string) Option {
	return func(opts *Options) {
		opts.indent = indent
	}
}

func FilePath(p string) Option {
	return func(opts *Options) {
		opts.filePath = p
	}
}
