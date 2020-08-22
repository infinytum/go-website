package language

import (
	"github.com/infinytum/go-website/internal/config"
)

type Language struct {
	Cache  map[string]config.Config
	Config config.Config
}
