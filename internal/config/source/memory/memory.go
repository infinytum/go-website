// Package memory loads changesets from a memory
package memory

import (
	"time"

	"github.com/infinytum/go-website/internal/config/source"
)

type memorySource struct {
	raw  string
	opts source.Options
}

var (
	DefaultString = "{}"
)

func (u *memorySource) Read() (*source.ChangeSet, error) {
	cs := &source.ChangeSet{
		Data:      []byte(u.raw),
		Format:    u.opts.Encoder.String(),
		Timestamp: time.Now(),
		Source:    "memory",
	}
	cs.Checksum = cs.Sum()

	return cs, nil
}

func (u *memorySource) Watch() (source.Watcher, error) {
	return newWatcher(u)
}

func (u *memorySource) String() string {
	return "memory"
}

func NewSource(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)

	raw, ok := options.Context.Value(memoryKey{}).(string)
	if !ok {
		raw = DefaultString
	}

	return &memorySource{raw: raw, opts: options}
}
