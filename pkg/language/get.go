package language

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/infinytum/go-website/internal/config"
	"github.com/infinytum/go-website/internal/config/encoder/yaml"
	"github.com/infinytum/go-website/internal/config/source"
	"github.com/infinytum/go-website/internal/config/source/file"
)

func (l *Language) GetOrLoad(language string) config.Config {
	if cfg, ok := l.Cache[language]; ok {
		return cfg
	}
	defaultLang := l.Config.Get("language").String("en")
	cwd, err := os.Getwd()
	if err != nil {
		return config.NewConfig()
	}

	finalPath, _ := filepath.Abs(cwd + "/" + language + ".yaml")
	if !strings.HasPrefix(finalPath, cwd) {
		return config.NewConfig()
	}

	if _, err := os.Stat(cwd + "/static/lang/"); os.IsNotExist(err) {
		finalPath = cwd + "/" + defaultLang + ".yaml"
	}

	cfg := config.NewConfig()
	enc := yaml.NewEncoder()
	if err := cfg.Load(file.NewSource(
		file.WithPath(finalPath),
		source.WithEncoder(enc),
	)); err != nil {
		return config.NewConfig()
	}
	return cfg
}
