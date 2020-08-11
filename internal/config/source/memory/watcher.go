package memory

import (
	"errors"

	"github.com/infinytum/baerenhoehle/internal/config/source"
)

type memoryWatcher struct {
	u    *memorySource
	exit chan bool
}

func newWatcher(u *memorySource) (*memoryWatcher, error) {
	return &memoryWatcher{
		u:    u,
		exit: make(chan bool),
	}, nil
}

func (u *memoryWatcher) Next() (*source.ChangeSet, error) {
	<-u.exit
	return nil, errors.New("url watcher stopped")
}

func (u *memoryWatcher) Stop() error {
	select {
	case <-u.exit:
	default:
	}
	return nil
}
