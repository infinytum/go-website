package reader

import (
	"github.com/infinytum/baerenhoehle/internal/config/encoder"
	"github.com/infinytum/baerenhoehle/internal/config/encoder/json"
	"github.com/infinytum/baerenhoehle/internal/config/encoder/toml"
	"github.com/infinytum/baerenhoehle/internal/config/encoder/xml"
	"github.com/infinytum/baerenhoehle/internal/config/encoder/yaml"
)

type Options struct {
	Encoding map[string]encoder.Encoder
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Encoding: map[string]encoder.Encoder{
			"json": json.NewEncoder(),
			"yaml": yaml.NewEncoder(),
			"toml": toml.NewEncoder(),
			"xml":  xml.NewEncoder(),
			"yml":  yaml.NewEncoder(),
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		if o.Encoding == nil {
			o.Encoding = make(map[string]encoder.Encoder)
		}
		o.Encoding[e.String()] = e
	}
}