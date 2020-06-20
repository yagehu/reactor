package config

import (
	"os"
	"path/filepath"

	"go.uber.org/config"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type Config struct {
	Context struct {
		RuntimeEnvironment RuntimeEnvironment `json:"-"`
	} `json:"-"`
}

type RuntimeEnvironment int

const (
	RuntimeEnvironmentDevelopment RuntimeEnvironment = iota
	RuntimeEnvironmentProduction

	_envKeyPrefix             = "REACTOR_"
	_envKeyConfigDir          = "CONFIG_DIR"
	_envKeyRuntimeEnvironment = "RUNTIME_ENVIRONMENT"
)

func New() (Config, error) {
	var (
		c                  Config
		runtimeEnvironment RuntimeEnvironment
		configDir          = os.Getenv(_envKeyPrefix + _envKeyConfigDir)
	)

	switch os.Getenv(_envKeyPrefix + _envKeyRuntimeEnvironment) {
	case "production":
		runtimeEnvironment = RuntimeEnvironmentProduction
	default:
		runtimeEnvironment = RuntimeEnvironmentDevelopment
	}

	opts := []config.YAMLOption{
		config.File(filepath.Join(configDir, "base.yaml")),
	}

	switch runtimeEnvironment {
	case RuntimeEnvironmentProduction:
		opts = append(
			opts, config.File(filepath.Join(configDir, "production.yaml")),
		)
	default:
		opts = append(
			opts, config.File(filepath.Join(configDir, "development.yaml")),
		)
	}

	opts = append(opts, config.File(filepath.Join(configDir, "secrets.yaml")))

	provider, err := config.NewYAML(opts...)
	if err != nil {
		return Config{}, err
	}

	if err := provider.Get(config.Root).Populate(&c); err != nil {
		return Config{}, err
	}

	c.Context.RuntimeEnvironment = runtimeEnvironment

	return c, nil
}
