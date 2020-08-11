package memory

import (
	"context"
	"encoding/json"

	"github.com/infinytum/baerenhoehle/internal/config/source"
)

type memoryKey struct{}

func WithString(u string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, memoryKey{}, u)
	}
}

func WithBytes(u []byte) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, memoryKey{}, string(u))
	}
}

func WithInterface(u interface{}) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		data, _ := json.Marshal(u)
		o.Context = context.WithValue(o.Context, memoryKey{}, string(data))
	}
}
