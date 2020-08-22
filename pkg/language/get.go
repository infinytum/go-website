package language

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/infinytum/go-website/internal/config"
	"github.com/infinytum/go-website/internal/config/encoder/yaml"
	"github.com/infinytum/go-website/internal/config/source"
	"github.com/infinytum/go-website/internal/config/source/file"
	"github.com/sirupsen/logrus"
)

func (l *Language) GetOrLoad(language string) config.Config {
	if cfg, ok := l.Cache[language]; ok {
		return cfg
	}
	defaultLang := l.Config.Get("language").String("en")
	cwd, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return config.NewConfig()
	}

	finalPath, _ := filepath.Abs(cwd + "/static/lang/" + language + ".yaml")
	if !strings.HasPrefix(finalPath, cwd) {
		return config.NewConfig()
	}

	if _, err := os.Stat(finalPath); os.IsNotExist(err) {
		logrus.Error(err)
		finalPath = cwd + "/static/lang/" + defaultLang + ".yaml"
	}

	cfg := config.NewConfig()
	enc := yaml.NewEncoder()
	if err := cfg.Load(file.NewSource(
		file.WithPath(finalPath),
		source.WithEncoder(enc),
	)); err != nil {
		logrus.Error(err)
		return config.NewConfig()
	}
	return cfg
}
