package config

import (
	"github.com/infinytum/baerenhoehle/internal/config/loader"
	"github.com/infinytum/baerenhoehle/internal/config/reader"
	"github.com/infinytum/baerenhoehle/internal/config/source"
)

// WithLoader sets the loader for manager config
func WithLoader(l loader.Loader) Option {
	return func(o *Options) {
		o.Loader = l
	}
}

// WithSource appends a source to list of sources
func WithSource(s source.Source) Option {
	return func(o *Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) Option {
	return func(o *Options) {
		o.Reader = r
	}
}