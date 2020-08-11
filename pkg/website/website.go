package website

import (
	"github.com/infinytum/go-website/internal/config"
	"github.com/infinytum/go-website/internal/config/encoder/yaml"
	"github.com/infinytum/go-website/internal/config/source"
	"github.com/infinytum/go-website/internal/config/source/file"
	"github.com/infinytum/go-website/pkg/context"
	"github.com/infinytum/go-website/pkg/ui"
	"github.com/sirupsen/logrus"
)

var ApplicationContext *context.Application

func ListenAndServe() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Debug("Loading main configuration file...")
	// Load config.yml from current working directory
	cfg := config.DefaultConfig
	enc := yaml.NewEncoder()
	if err := cfg.Load(file.NewSource(
		file.WithPath("./config/config.yml"),
		source.WithEncoder(enc),
	)); err != nil {
		logrus.Fatal(err)
	}

	logrus.Debug("Main configuration file loaded and validated, updating log level...")
	// Configure the log level according to the configuration file
	// This also validates that the provided log level is a valid one.
	logLevel, err := logrus.ParseLevel(cfg.Get("logging", "level").String("info"))
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Logging level has been set to %s", logLevel.String())
	logrus.SetLevel(logLevel)

	logrus.Debug("Initializing application-wide context...")
	// Initialize application-wide context
	ApplicationContext = context.NewApplication(cfg)

	logrus.Debug("Warming up cache instance to prevent runtime failure")
	// Warm-Up cache to ensure configuration validity
	ApplicationContext.Cache()

	logrus.Debug("Warming up squidex instance to prevent runtime failure")
	// Warm-Up cache to ensure configuration validity
	ApplicationContext.Squidex()

	logrus.Info("Uncover initialization has completed. Launching Uncover UI...")
	ui.ListenAndServe(ApplicationContext)
}
