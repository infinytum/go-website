package env

import (
	"errors"

	"github.com/infinytum/go-website/internal/config/source"
)

type watcher struct {
	exit chan struct{}
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	<-w.exit

	return nil, errors.New("watcher stopped")
}

func (w *watcher) Stop() error {
	close(w.exit)
	return nil
}

func newWatcher() (source.Watcher, error) {
	return &watcher{exit: make(chan struct{})}, nil
}
