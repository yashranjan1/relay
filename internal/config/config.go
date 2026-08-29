package config

import (
	"errors"
	"os"
	"path"

	"github.com/maniac-en/req/internal/log"
	"github.com/maniac-en/req/internal/tui/styles"
	"github.com/pelletier/go-toml/v2"
)

func LoadConfig() error {
	configDir, err := os.UserConfigDir()

	if err != nil {
		log.Error("Could not load user config dir")
		return errors.New("Could not load user config dir")
	}
	path := path.Join(configDir, "req", "config.toml")

	byteData, err := os.ReadFile(path)
	if err != nil {
		log.Error("Could not load config")
		initDefaults()
		return nil
	}

	var config Config

	toml.Unmarshal(byteData, &config)

	theme, ok := styles.ThemeMap[config.Theme.Name]

	if !ok {
		theme = styles.ThemeMap["default"]
	}

	styles.SetTheme(theme)
	// styles.SetTheme()

	return nil
}

func initDefaults() {
	theme := styles.ThemeMap["default"]
	log.Info("setting default config")
	styles.SetTheme(theme)
}
